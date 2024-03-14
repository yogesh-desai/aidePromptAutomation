package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {

	url := "https://aide.imagine.tech/api/v1/user-story"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	questions := []string{
		"What is GoLang and what are its key features?",
		"Explain Goroutines and how they are different from threads.",
		"What is a channel in GoLang and how is it used for communication between Goroutines?",
		"How does GoLang handle errors and what are the best practices for error handling?",
		"What is a defer statement in GoLang and how is it used?",
		"Explain the concept of interfaces in GoLang and provide an example.",
		"How does GoLang manage memory compared to other programming languages?",
		"What is the difference between a map and a slice in GoLang?",
		"How does GoLang support concurrent programming?",
		"Explain the concept of pointers in GoLang and provide an example of their usage.",
		"What is the purpose of the init function in GoLang?",
		"How does GoLang handle garbage collection?",
		"What is the difference between a method and a function in GoLang?",
		"Explain the concept of packages in GoLang and how they are used.",
		"What is the role of the main function in a GoLang program?",
		"How does error handling work in GoLang compared to other languages?",
		"What is the difference between a struct and an interface in GoLang?",
		"Explain the concept of type assertion in GoLang and provide an example.",
		"How can you create a new instance of a struct in GoLang?",
		"What is the purpose of the blank identifier (_) in GoLang?",
		"How does GoLang support polymorphism?",
		"What are the different types of loops available in GoLang?",
		"Explain the concept of method receivers in GoLang.",
		"How can you work with files in GoLang?",
		"What is the purpose of the sync package in GoLang?",
		"How can you perform unit testing in GoLang?",
		"What is the difference between the new keyword and make function in GoLang?",
		"How can you work with JSON data in GoLang?",
		"Explain the concept of function closures in GoLang.",
		"What is the purpose of the context package in GoLang?",
		"How can you handle HTTP requests and responses in GoLang?",
		"What is the purpose of the recover function in GoLang?",
		"How can you work with time and dates in GoLang?",
		"Explain the concept of method chaining in GoLang.",
		"What is the purpose of the sync.Mutex type in GoLang?",
		"How can you create and use custom error types in GoLang?",
		"What is the purpose of the sync.WaitGroup type in GoLang?",
		"How can you work with command-line arguments in GoLang?",
		"Explain the concept of deferred function calls in GoLang.",
		"What is the purpose of the context.Context type in GoLang?",
		"How can you work with SQL databases in GoLang?",
		"What is the purpose of the flag package in GoLang?",
		"Explain the concept of method overloading in GoLang.",
		"How can you work with environment variables in GoLang?",
		"What is the purpose of the log package in GoLang?",
		"How can you implement a web server in GoLang?",
		"Explain the concept of function literals in GoLang.",
		"What is the purpose of the http package in GoLang?",
		"How can you work with regular expressions in GoLang?",
		"What is the purpose of the sync.Cond type in GoLang?",
		"Explain the concept of method sets in GoLang.",
		"How can you work with TCP and UDP connections in GoLang?",
		"What is the purpose of the context.WithCancel function in GoLang?",
		"How can you work with goroutine pools in GoLang?",
		"Explain the concept of method expressions in GoLang.",
		"What is the purpose of the context.WithTimeout function in GoLang?",
		"How can you work with channels for synchronization in GoLang?",
		"What is the purpose of the context.WithValue function in GoLang?",
		"How can you work with context deadlines in GoLang?",
		"Explain the concept of method values in GoLang.",
		"What is the purpose of the context.Background function in GoLang?",
		"How can you work with context cancellation in GoLang?",
		"What is the purpose of the context.WithDeadline function in GoLang?",
		"Explain the concept of method sets in GoLang.",
		"How can you work with context propagation in GoLang?",
		"What is the purpose of the context.TODO function in GoLang?",
		"How can you work with context values in GoLang?",
		"Explain the concept of method expressions in GoLang.",
		"What is the purpose of the context.WithValue function in GoLang?",
		"How can you work with context timeouts in GoLang?",
		"Explain the concept of method values in GoLang.",
		"What is the purpose of the context.WithCancel function in GoLang?",
		"How can you work with context deadlines in GoLang?",
		"Explain the concept of method values in GoLang.",
		"What is the purpose of the context.WithTimeout function in GoLang?",
		"How can you work with context cancellation in GoLang?",
		"What is the purpose of the context.WithDeadline function in GoLang?",
		"Explain the concept of method sets in GoLang.",
		"How can you work with context propagation in GoLang?",
		"What is the purpose of the context.TODO function in GoLang?",
		"How can you work with context values in GoLang?",
		"Explain the concept of method expressions in GoLang.",
		"What is the purpose of the context.WithValue function in GoLang?",
		"How can you work with context timeouts in GoLang?",
		"Explain the concept of method values in GoLang.",
		"What is the purpose of the context.WithCancel function in GoLang?",
		"How can you work with context deadlines in GoLang?",
		"Explain the concept of method values in GoLang.",
		"What is the purpose of the context.WithTimeout function in GoLang?",
		"How can you work with context cancellation in GoLang?",
		"What is the purpose of the context.WithDeadline function in GoLang?",
		"Explain the concept of method sets in GoLang.",
		"How can you work with context propagation in GoLang?",
		"What is the purpose of the context.TODO function in GoLang?",
		"How can you work with context values in GoLang?",
		"Explain the concept of method expressions in GoLang.",
		"What is the purpose of the context.WithValue function in GoLang?",
		"How can you work with context timeouts in GoLang?",
		"Explain the concept of method values in GoLang.",
		"What is the purpose of the context.WithCancel function in GoLang?",
	}

	for _, question := range questions {
		fmt.Println("Processing for: ", question)

		valueStory := "{\"project\":\"9de5e1e6-97fd-4f6e-a186-4d13eae39288\",\"title\":\"" + question + "\",\"storyType\":\"Back end\",\"domain\":\"Development\",\"expectedTime\":0}"
		_ = writer.WriteField("userStory", valueStory)
		err := writer.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Accept", "application/json, text/plain, */*")
		req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("Cookie", "AzureAppProxyAnalyticCookie_48b159c3-9e1e-47a4-bb66-1197e604f374_https_1.3=4|6Oksvk0FVx+0iJeVFGLqxDJsBuuzzr7LdRrWvLqzvuIiL25/HBaQRGxPf4IMKf07rfJpkpCjRo287gd/lxmPoWjINhETi8ZgGlwVFDaokAWpJJEfEpa3ad4LyX07mJEam47d56OvM7UkEKlvDa03Pg==; _ga=GA1.2.841991317.1710239967; _gid=GA1.2.1114055079.1710239967; _token=FF8ycp_448oaVHICwlIUHE-66GnHSgBEjrjitffOHrI; user-info=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjJjMzAxOGMwLWNhNTMtNDRhMC04N2NjLTdlOGY3ZDRhZTFkMSIsImlzUmVwb3J0QWRtaW4iOmZhbHNlLCJpc0FkbWluIjpmYWxzZSwiaXNTdXBlcmFkbWluIjpmYWxzZSwiZXhwIjoxNzIwNjA3OTc1fQ; AzureAppProxyAnalyticCookie_48b159c3-9e1e-47a4-bb66-1197e604f374_https_1.3=4|qfADs9QDsuvpPfK7rY0ZHd7UxzkxUqpYmTPcGX103svBicFdaEkuVQnU6L4B+l0b/TZ4/cVczBU1G7ehKeowueozjW2CLlmVlgBkpbY8ndVCz0vko4c2xXJ10nowQOvpdZCb8rRrHz34MgxPwOFANg==")
		req.Header.Add("Origin", "https://aide.imagine.tech")
		req.Header.Add("Referer", "https://aide.imagine.tech/app/project/9de5e1e6-97fd-4f6e-a186-4d13eae39288")
		req.Header.Add("Sec-Fetch-Dest", "empty")
		req.Header.Add("Sec-Fetch-Mode", "cors")
		req.Header.Add("Sec-Fetch-Site", "same-origin")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
		req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", "\"macOS\"")

		req.Header.Set("Content-Type", writer.FormDataContentType())
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("error while making request: ", err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("error while reading the body: ", err)
			return
		}

		fmt.Println("Done creating story.", string(body))
	}
}
