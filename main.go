package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	usage :=
		`
Usage: ./emote-grabber <bettertv | 7tv> <id>

Get ID from url at:
[7tv]      https://7tv.app/emote-sets/<ID>
[BetterTV] https://betterttv.com/users/<ID>

For Example:
emote-grabber 7tv 61f0db5c1300d0c637b96133
emote-grabber bettertv 60255659ba4b112dc952075c
		`

	if len(os.Args) != 3 {
		log.Fatal(usage)
	}

	emotePackID := os.Args[2]

	switch os.Args[1] {
	case "7tv":
		url := "https://7tv.io/v3/emote-sets/" + emotePackID
		result, err := GetEmotePackData(url, STVResponse{})
		if err != nil {
			log.Fatal("Failed to fetch Emote pack metadata")
		}

		for _, emote := range result.Emotes {
			if err = DownloadSTVEmote(emote); err != nil {
				log.Println("Couldn't download", emote.Name)
			}
		}
	case "bettertv":
		url := "https://api.betterttv.net/3/users/" + emotePackID + "?limited=true&personal=false"
		result, err := GetEmotePackData(url, BTTVResponse{})
		if err != nil {
			log.Fatal("Failed to fetch Emote pack metadata")
		}

		for _, emote := range result.Emotes {
			if err = DownloadBTTVEmote(emote); err != nil {
				log.Println("Couldn't download", emote.Code)
			}
		}
	default:
		log.Fatal(usage)
	}
}

func GetEmotePackData[K BTTVResponse | STVResponse](url string, result K) (K, error) {
	log.Println("Fetching", url)
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	return result, nil
}

func downloadToFile(url, filename string) error {
	log.Println("Downloading", url, "to", filename)
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadSTVEmote(emote STVEmote) error {
	url := "https:" + emote.Data.Host.URL + "/4x.webp"
	filename := emote.Name + ".webp"
	return downloadToFile(url, filename)
}

func DownloadBTTVEmote(emote BTTVEmote) error {
	url := "https://cdn.betterttv.net/emote/" + emote.ID + "/3x." + emote.ImageType
	filename := emote.Code + "." + emote.ImageType
	return downloadToFile(url, filename)
}
