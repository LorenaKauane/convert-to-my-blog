package openai

import (
	"context"
	"convert-to-my-blog/configs"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func CallOpenAi(prompt string) string {
	client := openai.NewClient(os.Getenv(configs.Env.OPEN_AI))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo16K,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}
