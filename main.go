package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	project := githubactions.GetInput("project")
	if msg == "" {
		githubactions.Fatalf("Missing input 'project'")
	}
	commit := githubactions.GetInput("commit")
        if msg == "" {
                githubactions.Fatalf("Missing input 'commit'")
        }
	branch := githubactions.GetInput("branch")
        if msg == "" {
                githubactions.Fatalf("Missing input 'branch'")
        }
	status := githubactions.GetInput("status")
        if msg == "" {
                githubactions.Fatalf("Missing input 'status'")
        }
	actionid := githubactions.GetInput("actionid")
        if msg == "" {
                githubactions.Fatalf("Missing input 'actionid'")
        }
	webhook := githubactions.GetInput("webhook")
	if webhook == "" {
		githubactions.Fatalf("Missing input 'webshook'")
	}

	fmt.Println("URL:> ", webhook)
	data = `{
    "cards": [
        {
            "header": {
                "title": "GitHub Action",
                "subtitle": "Build Job",
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
                                                "url": "https://github.com/DTherHtun/web-app-flux/runs/%s"
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
	var jsonStr = []byte(fmt.Sprintf(data, project, commit, branch, status, actionid ))
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
}
