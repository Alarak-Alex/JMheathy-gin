package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/duke-git/lancet/strutil"
	"github.com/duke-git/lancet/system"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/slice"
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

func Chatmain(Prompt Prompt.Promt, title chan string, filePath string) {
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
			// fmt.Println("通道已关闭，退出循环")
			break
		}
		article := generateArticle(title1, client, modelDeploymentID, maxTokens, prompts)
		var articles []string
		for _, arti := range article {
			articles = append(articles, ReplaceString(arti))
		}
		// 保存文章到本地markdown文件
		saveArticle(filePath, articles, title1)

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
			// fmt.Println("生成的部分为：", part)
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

func saveArticle(Path string, articles []string, title string) {
	// 保存文章到本地markdown文件

	// 生成文档名称
	part1 := title + "_part1"
	part2 := title + "_part2"
	filepart1 := fmt.Sprintf("%s.md", part1)
	filepart2 := fmt.Sprintf("%s.md", part2)
	fileutil.CreateFile(Path + "/" + filepart1)
	fileutil.CreateFile(Path + "/" + filepart2)

	f1, err := os.OpenFile(Path+"/"+filepart1, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("错误: %s", err)
		return
	}
	f2, err := os.OpenFile(Path+"/"+filepart2, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("错误: %s", err)
		return
	}

	var cutindex int

	for index, str := range articles {
		if strutil.ContainsAll(str, []string{"---文章分割线---"}) {
			cutindex = index + 1
			break
		}
	}

	deleterange := len(articles) - cutindex + 1

	var part1Content []string
	var part2Content []string

	part1Content = slice.DropRight(articles, deleterange)
	part2Content = slice.DeleteRange(articles, 0, cutindex)
	// 写入part1内容
	for _, part1 := range part1Content {
		_, err = f1.WriteString(part1 + "\n")
		if err != nil {
			log.Printf("错误: %s", err)
			return
		}
	}

	// 写入part2内容
	for _, part2 := range part2Content {
		_, err = f2.WriteString(part2 + "\n")
		if err != nil {
			log.Printf("错误: %s", err)
			return
		}
	}

	f1.Close()
	f2.Close()

	base1 := Path + "/" + filepart1
	base2 := Path + "/" + filepart2
	docx1 := strings.Split(base1, ".")[0] + ".docx"
	docx2 := strings.Split(base2, ".")[0] + ".docx"

	// 打印Pandoc版本信息
	command1 := "pandoc " + base1 + " -f markdown -t docx -s -o " + docx1
	command2 := "pandoc " + base2 + " -f markdown -t docx -s -o " + docx2
	_, _, err = system.ExecCommand(command1)
	if err != nil {
		fmt.Errorf("执行命令失败: %w", err)
	}
	// fmt.Println("std out: ", stdout)
	// fmt.Println("std err: ", stderr)
	fileutil.RemoveFile(base1)
	_, _, err = system.ExecCommand(command2)
	if err != nil {
		fmt.Errorf("执行命令失败: %w", err)
	}
	// fmt.Println("std out: ", stdout, fileutil.CurrentPath())
	// fmt.Println("std err: ", stderr)
	fileutil.RemoveFile(base2)
}
