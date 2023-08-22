package main

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/zohaibsoomro/slack-bot/config"
)

func main() {
	client := slack.New(config.SlackAppToken)
	channels := []string{config.ChannelId}

	fileUploadParams := slack.FileUploadParameters{
		File:     "./file.txt",
		Channels: channels,
	}

	file, err := client.UploadFile(fileUploadParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name:", file.Name, "\ntype:", file.Filetype, "\nsize:", file.Size, "\ndownload url:", file.URLPrivateDownload)
}
