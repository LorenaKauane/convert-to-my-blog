# Why?

Automate content creation for my blog with GO, [Notion API](https://developers.notion.com/), [OPENAI](https://platform.openai.com/docs/introduction)

# Functionalities

- Get a Notion Page
- Convert to markdown
- Resume with OpenAI
- Translates into multiple languages ​​(Portuguese and English)
- Move file to desired folder

# How to run

1. You need the notion API keys and give access to the folder
2. You need the OPENAI keys to resume'
3. Configurate `.env` file with the keys `NOTION_SECRET` and `OPEN_AI`

## Execute:

`go build .`

`./convert-to-my-blog.exe`
