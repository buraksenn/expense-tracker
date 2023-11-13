package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/buraksenn/expense-tracker/config"
	"github.com/buraksenn/expense-tracker/internal/bot"
	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/internal/store"
	"github.com/buraksenn/expense-tracker/pkg/aws/dynamo"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
	"github.com/buraksenn/expense-tracker/pkg/aws/textract"
	"github.com/buraksenn/expense-tracker/pkg/logger"

	"github.com/buraksenn/expense-tracker/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	pkgWorker "github.com/buraksenn/expense-tracker/internal/worker"
)

var (
	cfg         *config.Config
	telegramBot *bot.Bot
	worker      *pkgWorker.Worker
)

func init() {
	logger.Debug("Initializing...")
	var err error
	cfg, err = config.GetConfig()
	if err != nil {
		logger.Fatal("Failed to get config, err: %v", err)
	}

	telegramClient, err := telegram.New(cfg.TelegramBotToken)
	if err != nil {
		logger.Fatal("Failed to create telegram client, err: %v", err)
	}
	outGoingMessageChan := make(common.OutgoingMessageChan)
	doneChan := make(chan struct{})
	ctx := context.Background()
	dynamoClient, err := dynamo.NewClient(ctx)
	if err != nil {
		logger.Fatal("Failed to create dynamo client, err: %v", err)
	}
	repo := store.NewRepo(dynamoClient)
	s3Client, err := s3.NewClient(ctx)
	if err != nil {
		logger.Fatal("Failed to create s3 client, err: %v", err)
	}
	textractClient, err := textract.NewDefaultClient(ctx)
	if err != nil {
		logger.Fatal("Failed to create textract client, err: %v", err)
	}

	telegramBot = bot.New(telegramClient, outGoingMessageChan, doneChan)
	worker = pkgWorker.New(repo, s3Client, textractClient, outGoingMessageChan)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (int, error) {
	logger.Debug("Starting handler")

	msg, err := parseRequestBody(request)
	if err != nil {
		logger.Error("Failed to parse invoke request, err: %v", err)
		return 200, nil
	}

	telegramBot.Start()
	if err := worker.HandleCommand(ctx, msg); err != nil {
		logger.Error("Failed to handle command, err: %v", err)
		return 200, nil
	}
	telegramBot.Stop()
	<-telegramBot.DoneChan

	logger.Info("Handler finished")
	return 200, nil
}

func parseRequestBody(req events.APIGatewayProxyRequest) (*common.IncomingMessage, error) {
	logger.Debug("Parsing request body")

	var update tgbotapi.Update
	if err := json.Unmarshal([]byte(req.Body), &update); err != nil {
		return nil, fmt.Errorf("unmarshalling request body: %w", err)
	}

	return telegramBot.PrepareMessage(&update)
}
