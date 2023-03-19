package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"server/model/chat"
)

type converssation struct{}

const key = "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"

var Converssation = new(converssation)

type ClientConfig struct {
	authToken          string
	HTTPClient         *http.Client
	BaseURL            string
	OrgID              string
	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		HTTPClient: &http.Client{},
		BaseURL:    "",
		OrgID:      "",
		authToken:  authToken,

		EmptyMessagesLimit: 300,
	}
}
func (chatSessionService *chatSessionService) GetGpt(converssation chat.Converssation) (resStream *openai.ChatCompletionStream, err error) {
	//var list []chat.Converssation
	//global.DB.Order("id desc").Limit(10).Find(&list)
	//prompt := q
	//
	//for _, c := range list {
	//	prompt = fmt.Sprintf("%q%s\n%s", c.Question, c.Answer, prompt)
	//}
	//fmt.Println("%q\nprompt数据库+input拼接", prompt)

	//config := openai.ClientConfig{
	//	authToken:          key,
	//	BaseURL:            "http://13.229.248.26",
	//	HTTPClient:         &http.Client{},
	//	OrgID:              "",
	//	EmptyMessagesLimit: 300,
	//}

	config := openai.DefaultConfig(key)
	config.BaseURL = "http://13.229.248.26"
	client := openai.NewClientWithConfig(config)
	fmt.Println("%+v", config)
	message := converssation.Messages
	outputMessage := make([]openai.ChatCompletionMessage, len(message))

	for i, obj := range message {
		outputMessage[i] = openai.ChatCompletionMessage{obj.Role, obj.Content}
	}
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			Messages:  outputMessage,
			Stream:    true,
			MaxTokens: 2049,
		},
	)
	defer stream.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//for {
	//	response, err := stream.Recv()
	//	if errors.Is(err, io.EOF) {
	//		fmt.Println("Stream finished")
	//		break
	//	}
	//
	//	if err != nil {
	//		fmt.Printf("Stream error: %v\n", err)
	//		break
	//	}
	//
	//	fmt.Printf("Stream response: %v\n", response)
	//
	//}

	fmt.Println("gpt返回数据", stream)
	return stream, err
	//conversation := chat.Converssation{
	//	Question: q,
	//	Answer:   text,
	//}
	//global.DB.Create(&conversation)
	//params := Params{
	//	Model:    "gpt-3.5-turbo",
	//	Messages: converssation.Messages,
	//}
	//fmt.Println(fmt.Sprintf("%+v", params))
	//dataByte, err := json.Marshal(params)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//bodyReader := bytes.NewReader(dataByte)
	//client := http.Client{}
	//
	//req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://13.229.248.26/chat/completions", bodyReader)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Accept", "text/event-stream")
	//req.Header.Set("Cache-Control", "no-cache")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "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"))
	//res, err = client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("res", res)
}
