package notion

import (
	"context"
	"convert-to-my-blog/configs"
	"convert-to-my-blog/internal/entity"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
)

func GetPageNotion(idPage string) (string, error) {
	if idPage == "" {
		log.Fatal("ERROR: I need idPage")
	}
	client := notion.NewClient(os.Getenv(configs.Env.NOTION_SECRET))

	children, err := client.FindBlockChildrenByID(context.Background(), idPage, nil)
	if err != nil {
		log.Fatal("(GetPageNotion) client.FindBlockChildrenByID: %v", err)
	}

	return  mountTexts(children), nil
}

func GetPropPage(idPage string) (entity.PropertiesNotion) {

	client := notion.NewClient(os.Getenv(configs.Env.NOTION_SECRET))
	page, err := client.FindPageByID(context.TODO(), idPage)
	if err != nil {
		log.Fatal(err)
	}
	
	props := page.Properties.(notion.DatabasePageProperties)
	
	title := props["title"].RichText[0].Text.Content
	thumbnail := props["thumbnail"].RichText[0].Text.Content

	tag := props["tag"].Select.Name
	date := props["date"].Date.Start.Format("2006-01-02")

	properties := entity.PropertiesNotion{
		Title:       title,
		Description: title,
		Tag:         tag,
		Thumbnail:   thumbnail,
		Date:        date,
	}

	return properties
}

func mountTexts(children notion.BlockChildrenResponse) (string) {
	body := ""
	for _, el := range children.Results {
		switch block := el.(type) {
		case *notion.ParagraphBlock:
			bstr := ParseRichText(block.RichText)
			if *bstr != "" {
				body += *bstr + "\n\n"
			}
		case *notion.Heading1Block:
			bstr := ParseRichText(block.RichText)
			body += *bstr + "\n\n"
		case *notion.Heading2Block:
			bstr := ParseRichText(block.RichText)
			body += *bstr
		case *notion.Heading3Block:
			bstr:= ParseRichText(block.RichText)
			body += *bstr
		}
	}

	return body
}

func ParseRichText(richText []notion.RichText) (*string) {
	str := ""
	for _, el := range richText {
		str += el.Text.Content
	}
	return &str
}

