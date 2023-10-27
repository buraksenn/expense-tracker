package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambda/messages"
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

	telegramBot = bot.New(telegramClient, outGoingMessageChan)
	worker = pkgWorker.New(repo, s3Client, textractClient, outGoingMessageChan)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req messages.InvokeRequest) {
	logger.Debug("Starting handler")

	msg, err := parseInvokeRequest(req)
	if err != nil {
		logger.Error("Failed to parse invoke request, err: %v", err)
		return
	}

	if err := worker.HandleCommand(ctx, msg); err != nil {
		logger.Error("Failed to handle command, err: %v", err)
		return
	}
	telegramBot.Stop()

	logger.Info("Handler finished")
}

func parseInvokeRequest(req messages.InvokeRequest) (*common.IncomingMessage, error) {
	logger.Debug("Parsing invoke request: %v", req)
	var update tgbotapi.Update
	if err := json.Unmarshal(req.Payload, &update); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload, err: %v", err)
	}

	return telegramBot.PrepareMessage(&update)
}
