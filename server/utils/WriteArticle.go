package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"
	Prompt "github.com/flipped-aurora/gin-vue-admin/server/model/Promt"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func TestChatMain(Prompt Prompt.Promt, title chan string, filename string) {
	for {
		title1, ok := <-title
		if !ok {
			fmt.Println("通道已关闭，退出循环")
			break
		}
		fmt.Println("title:", title1)
		fmt.Println("filename:", filename)
		fmt.Println("Prompt:", Prompt.PromtData)
	}
}

func Chatmain(Prompt Prompt.Promt, title chan string, filename string) {
	// azureOpenAIKey := os.Getenv("AZURE_OPENAI_API_KEY")
	// modelDeploymentID := os.Getenv("YOUR_MODEL_DEPLOYMENT_NAME")
	azureOpenAIKey := "75189fda8719425dbbc5947d59d661e7"
	modelDeploymentID := "gpt-4o-mini"
	maxTokens := int32(4000)

	// azureOpenAIEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	azureOpenAIEndpoint := "https://jmheathyswedencentral.openai.azure.com"
	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "跳过示例，环境变量缺失\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)
	if err != nil {
		log.Printf("错误: %s", err)
		return
	}

	prompts, err := PromptJsonArrayToStringSlice(Prompt.PromtData)
	// title1 := <-title
	// article := generateArticle(title1, client, modelDeploymentID, maxTokens, prompts)
	// var articles []string
	// for _, arti := range article {
	// 	articles = append(articles, ReplaceString(arti))

	// }
	// // 保存文章到本地markdown文件
	// saveArticle(title1, articles)
	for {
		title1, ok := <-title
		if !ok {
			fmt.Println("通道已关闭，退出循环")
			break
		}
		article := generateArticle(title1, client, modelDeploymentID, maxTokens, prompts)
		var articles []string
		for _, arti := range article {
			articles = append(articles, ReplaceString(arti))
		}
		// 保存文章到本地markdown文件
		saveArticle(filename+"/"+title1, articles)

	}
}

func generateArticle(title string, client *azopenai.Client, modelDeploymentID string, maxTokens int32, prompts []string) []string {
	var article []string
	var historyMessages []azopenai.ChatRequestMessageClassification // 用于保存历史消息

	for _, prompt := range prompts {
		flag := strings.Split(prompt, ":")[0]
		data := strings.Split(prompt, ":")[1]
		if flag != "part" {
			data = "\n" + data + "\n"
			article = append(article, data)
		} else {
			part := generatePart(title, client, modelDeploymentID, maxTokens, data, historyMessages)
			article = append(article, part)
			fmt.Println("生成的部分为：", part)
			// 记录生成的部分作为历史消息
			historyMessages = append(historyMessages, &azopenai.ChatRequestAssistantMessage{Content: to.Ptr(part)})
		}
	}

	return article
}

func generatePart(title string, client *azopenai.Client, modelDeploymentID string, maxTokens int32, prompt string, historyMessages []azopenai.ChatRequestMessageClassification) string {
	// fmt.Printf("generatePart函数接受的提示词为：%s\n", prompt)

	// 合并上下文和当前提示
	answer, _ := getCompletion(title, client, modelDeploymentID, prompt, maxTokens, historyMessages)
	return answer
}

func getCompletion(title string, client *azopenai.Client, modelDeploymentID string, prompt string, maxTokens int32, historyMessages []azopenai.ChatRequestMessageClassification) (string, []azopenai.ChatRequestMessageClassification) {

	origionalPrompt := fmt.Sprintf("你是一位擅长写文章的健康科普专栏医生，你会以医生的视角看待问题。并且会用简明易懂的语言生成引人入胜的科普小文章，要求第三人称，以患者为主体，标题为：%s,不要在文中生成标题", title)

	// 将历史消息加入到消息列表中
	messages := append(historyMessages, &azopenai.ChatRequestSystemMessage{Content: to.Ptr(origionalPrompt)},
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent(prompt)})

	temperature := float32(0.7)

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		Messages:       messages,
		DeploymentName: &modelDeploymentID,
		MaxTokens:      &maxTokens,
		Temperature:    &temperature,
	}, nil)

	if err != nil {
		log.Printf("错误: %s", err)
		return "", messages
	}

	for _, choice := range resp.Choices {
		if choice.Message != nil && choice.Message.Content != nil {
			return *choice.Message.Content, messages
		}
	}

	return "", messages
}

func saveArticle(title string, articles []string) {
	// 保存文章到本地markdown文件
	filename := fmt.Sprintf("%s.md", title)
	fileutil.CreateFile(filename + "/" + title)
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("创建文件时出错: %s", err)
		return
	}
	defer f.Close()

	for _, article := range articles {
		// _, err = f.WriteString(fmt.Sprintf("%s\n", article))
		// fmt.Println(article)

		fileutil.WriteStringToFile(filename+"/"+title, article, true)
		if err != nil {
			log.Printf("写入文件时出错: %s", err)
			return
		}
	}

	fmt.Println("文章已成功保存为", filename)
}
