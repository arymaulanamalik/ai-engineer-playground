package claude

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"

	"github.com/arymaulanamalik/ai-engineer-playground/ai-playground/configs"
	"github.com/arymaulanamalik/ai-engineer-playground/ai-playground/pkg/logger"
	"go.uber.org/zap"
)

type Client interface {
	GenerateText(
		ctx context.Context,
		prompt string,
	) (string, error)
}

type client struct {
	config *configs.Config
	sdk    anthropic.Client
}

func NewClient(
	config *configs.Config,
) Client {
	sdk := anthropic.NewClient(
		option.WithAPIKey(config.LLM.ClaudeAPIKey),
	)

	return &client{
		config: config,
		sdk:    sdk,
	}
}

func (c *client) GenerateText(
	ctx context.Context,
	prompt string,
) (string, error) {
	logger.Info(
		ctx,
		"calling claude api",
		zap.String("model", c.config.LLM.ClaudeModel),
	)

	resp, err := c.sdk.Messages.New(
		ctx,
		anthropic.MessageNewParams{
			MaxTokens: int64(c.config.LLM.ClaudeMaxTokens),
			Model:     anthropic.Model(c.config.LLM.ClaudeModel),
			Messages: []anthropic.MessageParam{
				anthropic.NewUserMessage(
					anthropic.NewTextBlock(prompt),
				),
			},
		},
	)

	if err != nil {
		logger.Error(
			ctx,
			err,
			"claude api failed",
		)
		return "", err
	}

	if len(resp.Content) == 0 {
		logger.Warn(
			ctx,
			"claude returned empty response",
		)
		return "", nil
	}

	text := resp.Content[0].Text

	logger.Debug(
		ctx,
		"claude response received",
		zap.String("response", text),
	)

	return text, nil
}
