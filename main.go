package main

import(
"fmt"
"os"
"github.com/slack-go/slack"
)


func main(){

	 os.Setenv("SLACK_BOT_TOKEN", "xoxb-4434836363509-4440617340421-ijWODOh8gqMuT6eOSZ4NRPdT")
	 os.Setenv("CHANNEL_ID", "C04CVP0DT6W")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Reading_Material_1.pdf"} 

	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err !=nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}