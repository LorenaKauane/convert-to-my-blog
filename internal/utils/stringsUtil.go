package utils

import (
	"convert-to-my-blog/internal/entity"
	"fmt"
)

func MountHeaderMD(obj entity.PropertiesNotion) string {
	header := fmt.Sprintf(`---
description: %s
title: %s
date: '%s'
thumbnail: %s
tag: %s
---`, obj.Description, obj.Title, obj.Date, obj.Thumbnail, obj.Tag)

	return header
}

func StructToJSONWithPrompt(prompt string , objPropPT entity.PropertiesNotion) (string) {
	return fmt.Sprintf(`%s {
		"title": "%s",
		"description": "%s",
		"tag": "%s",
		"thumbnail": "%s",
		"image": "%s",
		"date": "%s"
	}`, prompt, objPropPT.Title, objPropPT.Description, objPropPT.Tag, objPropPT.Thumbnail, objPropPT.Image, objPropPT.Date)
}
