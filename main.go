package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	msg := githubactions.GetInput("msg")
	if msg == "" {
		githubactions.Fatalf("Missing input 'msg'")
	}
	webhook := githubactions.GetInput("webhook")
	if webhook == "" {
		githubactions.Fatalf("Missing input 'webshook'")
	}

	fmt.Println("URL:> ", webhook)

	var jsonStr = []byte(fmt.Sprintf("{'text' : '%s'}", msg))
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
