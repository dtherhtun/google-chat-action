package main

import (
	"bytes"
	"fmt"
	"net/http"

	gh "github.com/sethvargo/go-githubactions"
)

// Generates the data to send to the webhook
func generateTemplate(ghc *gh.GitHubContext) []byte {

	data := `{
        "cards": [
            {
                "header": {
                    "title": "Code review requested on",
                    "subtitle": "",
                    "imageUrl": "https://github.githubassets.com/images/modules/logos_page/Octocat.png",
                    "imageStyle": "IMAGE"
                },
                "sections": [
                    {
                        "widgets": [
                            {
                                "textParagraph": {
                                    "text": "<b>Project:</b> %s<br><b>Commit-id:</b>  <font color=\"#FF0000\">%s</font><br><b>Branch:</b> <font color=\"#00FF00f\">%s</font><br><b>Build Status:</b> <font color=\"#0000ff\">%s</font>"
                                },
                                "buttons": [
                                    {
                                        "textButton": {
                                            "text": "Job Details",
                                            "onClick": {
                                                "openLink": {
                                                    "url": "https://github.com/%s"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    }`

	// project := ghc.Repository
	// branch := ghc.RefName

	return []byte(fmt.Sprintf(data, "1", "2", "3", "4", "5"))

}

// Send data to the webhook
func sendMessage(webhook string, data []byte) error {
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	return nil
}

func main() {

	githubContext, err := gh.Context()
	if err != nil {
		panic(err)
	}

	webhook := gh.GetInput("webhook")

	if webhook == "" {
		gh.Fatalf("Missing input 'webhook'")
	}
	fmt.Println("URL:> ", webhook)

	fmt.Printf("%#v", githubContext)

	data := generateTemplate(githubContext)
	err = sendMessage(webhook, data)

	if err != nil {
		panic(err)
	}

}
