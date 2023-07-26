package application

import (
	"convert-to-my-blog/internal/entity"
	"convert-to-my-blog/internal/utils"
	"convert-to-my-blog/pkg/notion"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func GetDataUser() *entity.ResultUser {
	notionIdPromptContent := entity.PromptContent{
		ErrorMsg: "Please provide an ID for the Notion page.",
		Label:    "❓ Qual e o ID da pagina do notion?",
	}

	idPageNotion := *PromptGetInput(notionIdPromptContent)

	objPropPT := notion.GetPropPage(idPageNotion)

	pathPromptContent := entity.PromptContent{
		ErrorMsg: "Informe a pasta",
		Label:    "❓ Qual pasta eu devo salvar o arquivo?",
	}

	path := *PromptGetInput(pathPromptContent)

	languagePromptContent := entity.PromptContent {
		Label: "Quais linguagens você quer gerar?",
		Items:append(utils.LANGUAGES, "Todas"),
	}

	valueLanguage := PromptSelect(languagePromptContent)
	language :=  utils.LANGUAGES
	if(valueLanguage != "Todas") {
		language =[]string{valueLanguage} 
	}

	videoPromptContent := entity.PromptContent {
		Label: "Tem video no youtube?",
		Items: []string{"Sim", "Não"},
	}

	IsVideoYoutube := PromptBoolean(videoPromptContent)
	idYoutubePage := ""
	if(IsVideoYoutube) {
		idYoutubePromptContent:= entity.PromptContent{
			ErrorMsg: "Please provide an ID for the video.",
			Label:    "ID Video:",
		}
	
		idYoutubePage = *PromptGetInput(idYoutubePromptContent)
	}

	return &entity.ResultUser {
		Language: language,
		IdPage: idPageNotion,
		IsVideoYoutube: IsVideoYoutube,
		IdVideoYoutube: idYoutubePage,
		Properties: objPropPT,
		Path: path,
	}
}


func PromptGetInput(pc entity.PromptContent) *string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := &promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return &result
}

func PromptSelect(pc entity.PromptContent) string {
    prompt := &promptui.Select{
        Label: pc.Label,
        Items: pc.Items,
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }
    return result
}

func PromptBoolean(pc entity.PromptContent) bool {
    prompt := &promptui.Select{
        Label: pc.Label,
        Items: pc.Items,
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }
    return result == "Sim"
}