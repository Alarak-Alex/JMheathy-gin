package article

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

type ModelDeployment struct {
	AzureOpenAIKey      string
	AzureOpenAIEndpoint string
	DeploymentName      string
}

type Prompt struct {
	InsidePrompt    string
	JsonPromptArrry []string
}

func WriteWordArticle(AiConfig ModelDeployment, prompt Prompt) {
	azureOpenAIKey := AiConfig.AzureOpenAIKey
	modelDeploymentID := AiConfig.DeploymentName
	maxTokens := int32(4000)

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := AiConfig.AzureOpenAIEndpoint

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "写文启动失败，请检查配置信息是否正确。\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	// This is a conversation in progress.
	// NOTE: all messages, regardless of role, count against token usage for this API.
	messages := []azopenai.ChatRequestMessageClassification{
		// 设置系统内置角色
		&azopenai.ChatRequestSystemMessage{Content: to.Ptr("你好，欢迎使用 Azure OpenAI 聊天机器人。")},

		// The user asks a question
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("Does Azure OpenAI support customer managed keys?")},

		// The reply would come back from the model. You'd add it to the conversation so we can maintain context.
		&azopenai.ChatRequestAssistantMessage{Content: to.Ptr("Yes, customer managed keys are supported by Azure OpenAI")},

		// The user answers the question based on the latest reply.
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("What other Azure Services support customer managed keys?")},

		// from here you'd keep iterating, sending responses back from ChatGPT
	}

	gotReply := false

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		// This is a conversation in progress.
		// NOTE: all messages count against token usage for this API.
		Messages:       messages,
		DeploymentName: &modelDeploymentID,
		MaxTokens:      &maxTokens,
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	for _, choice := range resp.Choices {
		gotReply = true

		if choice.ContentFilterResults != nil {
			fmt.Fprintf(os.Stderr, "Content filter results\n")

			if choice.ContentFilterResults.Error != nil {
				fmt.Fprintf(os.Stderr, "  Error:%v\n", choice.ContentFilterResults.Error)
			}

			fmt.Fprintf(os.Stderr, "  Hate: sev: %v, filtered: %v\n", *choice.ContentFilterResults.Hate.Severity, *choice.ContentFilterResults.Hate.Filtered)
			fmt.Fprintf(os.Stderr, "  SelfHarm: sev: %v, filtered: %v\n", *choice.ContentFilterResults.SelfHarm.Severity, *choice.ContentFilterResults.SelfHarm.Filtered)
			fmt.Fprintf(os.Stderr, "  Sexual: sev: %v, filtered: %v\n", *choice.ContentFilterResults.Sexual.Severity, *choice.ContentFilterResults.Sexual.Filtered)
			fmt.Fprintf(os.Stderr, "  Violence: sev: %v, filtered: %v\n", *choice.ContentFilterResults.Violence.Severity, *choice.ContentFilterResults.Violence.Filtered)
		}

		if choice.Message != nil && choice.Message.Content != nil {
			fmt.Fprintf(os.Stderr, "Content[%d]: %s\n", *choice.Index, *choice.Message.Content)
		}

		if choice.FinishReason != nil {
			// this choice's conversation is complete.
			fmt.Fprintf(os.Stderr, "Finish reason[%d]: %s\n", *choice.Index, *choice.FinishReason)
		}
	}

	if gotReply {
		fmt.Fprintf(os.Stderr, "Received chat completions reply\n")
	}

}
