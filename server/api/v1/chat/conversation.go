package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sashabaranov/go-openai"
	"io"
	"server/model/chat"
	"server/response"
)

const key = "xxxxxxxxx"

func GetConversation(c *gin.Context) {

	response.Success("成功", "对话列表", c)
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// websocket
//func GetGpt(c *gin.Context) {
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	for {
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		err = conn.WriteMessage(messageType, p)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//}

func GetGpt(c *gin.Context) {
	var converssation chat.Converssation
	err := c.ShouldBindJSON(&converssation)
	config := openai.DefaultConfig(key)
	config.BaseURL = "http://xx proxy xx/v1"
	client := openai.NewClientWithConfig(config)
	message := converssation.Messages
	outputMessage := make([]openai.ChatCompletionMessage, len(message))

	for i, obj := range message {
		outputMessage[i] = openai.ChatCompletionMessage{obj.Role, obj.Content}
	}
	fmt.Println("%+v", "outputMessage21211", outputMessage)
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

	fmt.Println("gpt返回数据", stream)
	for k, v := range stream.GetResponse().Header {
		c.Writer.Header().Set(k, v[0])
	}
	//c.Writer.Header().Set("Content-Type", "text/event-stream")
	//c.Writer.Header().Set("Connection", "keep-alive")
	//c.Writer.Header().Set("Cache-Control", "must-revalidate,no-cache,")
	//向客户端返回流数据
	c.Stream(func(w io.Writer) bool {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finish %v", stream)
			return false
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return false
		}

		//fmt.Printf("Stream response: %q\n", response.Choices[0].Delta.Content)

		jsonString, err := json.Marshal(response.Choices[0].Delta)
		if err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Fprintf(w, "data:%s\n\n", jsonString)
		// 写入流数据
		//_, err = w.Write(jsonString)
		//if err != nil {
		//	return false
		//}
		//c.Writer.Flush()
		// 继续读取流数据
		return true
	})

	//response.Success("成功", 12, c)
}
