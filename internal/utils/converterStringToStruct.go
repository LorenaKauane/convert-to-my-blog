package utils

import (
	"convert-to-my-blog/internal/entity"
	"encoding/json"
	"log"
)

func ConverterStringToStruct(text string) ( entity.PropertiesNotion) {
	englishProperties := entity.PropertiesNotion{}
	err := json.Unmarshal([]byte(text), &englishProperties)
	if err != nil {
		log.Fatal("Erro ao converter resposta em inglês para struct:", err)
	}
	return englishProperties
}