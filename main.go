package main

import (
	"fmt"

	gh "github.com/sethvargo/go-githubactions"
)

func main() {
	githubContext, _ := gh.Context()
	repositoryOwner, repositoryName := githubContext.Repo()
	project := fmt.Sprintf("%s/%s", repositoryOwner, repositoryName)

	fmt.Printf("%#v", project)
	fmt.Printf("%#v", githubContext)
	fmt.Printf("%#v", githubContext)

	// 	branch := gh.GetInput("branch")
	// 	if branch == "" {
	// 		gh.Fatalf("Missing input 'branch'")
	// 	}
	// 	status := gh.GetInput("status")
	// 	if status == "" {
	// 		gh.Fatalf("Missing input 'status'")
	// 	}
	// 	actionid := gh.GetInput("actionid")
	// 	if actionid == "" {
	// 		gh.Fatalf("Missing input 'actionid'")
	// 	}
	// 	webhook := gh.GetInput("webhook")
	// 	if webhook == "" {
	// 		gh.Fatalf("Missing input 'webshook'")
	// 	}

	//		fmt.Println("URL:> ", webhook)
	//		data := `{
	//	    "cards": [
	//	        {
	//	            "header": {
	//	                "title": "GitHub Action",
	//	                "subtitle": "Build Job",
	//	                "imageUrl": "https://github.githubassets.com/images/modules/logos_page/Octocat.png",
	//	                "imageStyle": "IMAGE"
	//	            },
	//	            "sections": [
	//	                {
	//	                    "widgets": [
	//	                        {
	//	                            "textParagraph": {
	//	                                "text": "<b>Project:</b> %s<br><b>Commit-id:</b>  <font color=\"#FF0000\">%s</font><br><b>Branch:</b> <font color=\"#00FF00f\">%s</font><br><b>Build Status:</b> <font color=\"#0000ff\">%s</font>"
	//	                            },
	//	                            "buttons": [
	//	                                {
	//	                                    "textButton": {
	//	                                        "text": "Job Details",
	//	                                        "onClick": {
	//	                                            "openLink": {
	//	                                                "url": "https://github.com/%s"
	//	                                            }
	//	                                        }
	//	                                    }
	//	                                }
	//	                            ]
	//	                        }
	//	                    ]
	//	                }
	//	            ]
	//	        }
	//	    ]
	//	}`
	//
	//	var jsonStr = []byte(fmt.Sprintf(data, project, commit, branch, status, actionid))
	//	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
	//	if err != nil {
	//		panic(err)
	//	}
	//	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//	client := &http.Client{}
	//	resp, err := client.Do(req)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer resp.Body.Close()
	//	fmt.Println("response Status:", resp.Status)
}
