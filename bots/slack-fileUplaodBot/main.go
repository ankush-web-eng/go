package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	params := slack.FileUploadParameters{
		Channels: channelArr,
		File:     "blog.png",
	}

	file, err := api.UploadFile(params)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
}
