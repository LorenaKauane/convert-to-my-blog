package entity

type ResultUser struct {
	Language       []string
	IdPage         string
	Properties     PropertiesNotion
	PropertiesEN   PropertiesNotion
	IsVideoYoutube bool
	IdVideoYoutube string
	Path           string
}