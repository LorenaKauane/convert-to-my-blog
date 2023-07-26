package application

import (
	"convert-to-my-blog/internal/utils"
	"convert-to-my-blog/pkg/notion"
	openai "convert-to-my-blog/pkg/openAi"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func CreateNewResume() {
	resultUser := GetDataUser()
	loading := spinner.New(spinner.CharSets[6], 100*time.Millisecond)
	loading.Prefix = "Loading...."
	loading.Start()  
	content, err := notion.GetPageNotion(resultUser.IdPage) 
	loading.Stop()
	if(err != nil) {
		log.Fatal("Error get page notion")
	}

	fmt.Println("\n🔥 Ja estou com seu conteúdo do notion, agora vamos para a parte divertida 🔥 ")

	for _, language := range resultUser.Language {
		fmt.Println("===============================================================================")
		fmt.Println("Traduzindo para:", language)
		fmt.Println("===============================================================================")
		
		isDifferentePT := language != utils.PT
		if(isDifferentePT) {
			fmt.Print(fmt.Sprintf("\nChatgpt está traduzindo o Header para: %s", language))

			promptValue := fmt.Sprintf("Traduza os valores para %s e me retorne APENAS o json", language)
			promptHeader := utils.StructToJSONWithPrompt(promptValue, resultUser.Properties)
			loading.Start()  
			contentHeaderAI := openai.CallOpenAi(promptHeader)
			loading.Stop()  
			resultUser.Properties = utils.ConverterStringToStruct(contentHeaderAI)
			fmt.Println("\n✔️  Pronto")
		}
	
		header := utils.MountHeaderMD(resultUser.Properties)

		fmt.Println("\n🚀 Chatgpt está resumindo seu conteudo...")
		loading.Start()  
		markdownText := openai.CallOpenAi(fmt.Sprintf("Pegue esse texto formate para post em blog com secoes, de titulos, lista, se necessario explique alguns items na sequencia me retorne em formato markdown %s", content))
		markdownText = openai.CallOpenAi(fmt.Sprintf("Traduza esse conteudo para %s: %s", language, markdownText))
		loading.Stop()  
		fmt.Println(fmt.Sprintf("\n✔️  Conteúdo resumido, formatado em markdown e traduzido para: %s", language))

		if(resultUser.IsVideoYoutube && !isDifferentePT) {
			markdownText = fmt.Sprintf("Caso você goste de conteúdo em vídeo, aproveite para se inscrever no meu canal e ativar as notificações para não perder nenhum conteúdo novo!\n\n[![VIDEO IMAGE](https://img.youtube.com/vi/%s/0.jpg)](https://www.youtube.com/watch?v=%s)\n\n%s", resultUser.IdVideoYoutube,  resultUser.IdVideoYoutube, markdownText)
		}

		fmt.Println("\n🙌  Etapas finais por aqui... Estou montando o header para seguir o template do Markdown")
		markdownText = fmt.Sprintf("%s\n%s", header, markdownText)
		clearNameFile:= utils.RemoveCaracteresEspeciais(resultUser.Properties.Title)
		fileName := fmt.Sprintf(`%s-%s.md`, resultUser.Properties.Date, strings.ReplaceAll(clearNameFile, " ", "-") )
		
		creteFile(fileName, markdownText)
		moveFile(resultUser.Path, utils.NAMEFOLDER[language], fileName)

		fmt.Println(fmt.Sprintf("\n🌟  Pronto! Arquivo criado e movido para pasta desejada: %s", language))
	}
	fmt.Println("\n🌟  Antes de subir em produção teste local, adicione uns gifs, mude algum texto que achar necessario e seja feliz 🌟")
}

func creteFile(fileName string, markdownText string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(markdownText))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
}

func moveFile(path string, lang string, fileName string) {
	currentDir, _ := os.Getwd()
	newPath  := fmt.Sprintf("%s\\%s\\%s",path, lang, fileName)
	
	oldPath := filepath.Join(currentDir, fileName)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("Erro ao mover o arquivo:", err)
		return
	}
}
