package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("No user ID provided.\nUsage: `better-emote-grabber 60255659ba4b112dc952075c`")
	}

	url := "https://api.betterttv.net/3/users/"+os.Args[1]+"?limited=true&personal=false"

	// Get request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // response body is []byte

	var result BTTVResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		log.Println("Can not unmarshal JSON")
	}

	// Loop through the data node for the FirstName
	for _, emote := range result.ChannelEmotes {
		if err = DownloadEmote(emote); err != nil {
			log.Println("Couldn't download", emote.Code)
		}
	}

}

func DownloadEmote(emote BTTVChannelEmotes) error {

	log.Println("Downloading", emote.Code+"."+emote.ImageType)

	out, err := os.Create(emote.Code + "." + emote.ImageType)
	if err != nil {
		return err
	}
	defer out.Close()

	// https://cdn.betterttv.net/emote/60b14a6df8b3f62601c34c01/3x.webp
	resp, err := http.Get("https://cdn.betterttv.net/emote/" + emote.ID + "/3x." + emote.ImageType)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
