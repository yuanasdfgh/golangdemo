package initialize

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func initChatgpt() {
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
	fmt.Println("---===")
}
