package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/buraksenn/expense-tracker/internal/bot"
	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/internal/store"
	"github.com/buraksenn/expense-tracker/internal/worker"
	"github.com/buraksenn/expense-tracker/pkg/aws/dynamo"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
	"github.com/buraksenn/expense-tracker/pkg/aws/textract"
	"github.com/buraksenn/expense-tracker/pkg/logger"
	"github.com/buraksenn/expense-tracker/pkg/telegram"
)

var (
	logPath   = flag.String("log_path", "", "Log file path")
	bot_token = flag.String("bot_token", "", "Telegram bot token")
)

func main() {
	flag.Parse()
	prepareBotToken()
	var file *os.File
	if *logPath != "" {
		var err error
		file, err = os.OpenFile(*logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file, err: " + err.Error())
		}
		defer func() {
			if err := file.Close(); err != nil {
				panic("Failed to close log file, err: " + err.Error())
			}
		}()
	}
	logger.Init(os.Stdout)
	telegramClient, err := telegram.New(*bot_token)
	if err != nil {
		logger.Fatal("Failed to create telegram client, err: %v", err)
	}
	incomingMessageChan := make(common.IncomingMessageChan)
	outGoingMessageChan := make(common.OutgoingMessageChan)
	telegramBot := bot.New(telegramClient, incomingMessageChan, outGoingMessageChan)

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

	worker := worker.New(repo, s3Client, textractClient, incomingMessageChan, outGoingMessageChan)
	go telegramBot.Start()
	go worker.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)

	// Block until we receive our signal.
	sig := <-c
	logger.Info("Got signal: %v", sig)
}

func prepareBotToken() {
	if b := os.Getenv("BOT_TOKEN"); b != "" {
		*bot_token = b
	}
}
