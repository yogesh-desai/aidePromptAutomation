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
		"What is Angular and how does it differ from AngularJS?",
		"Explain the key features of Angular.",
		"What is data binding in Angular and explain the different types of data binding.",
		"Describe the Angular component architecture.",
		"What is Angular CLI and how is it used in Angular development?",
		"Explain the concept of directives in Angular and provide examples.",
		"What are Angular services and why are they used?",
		"Describe the difference between template-driven forms and reactive forms in Angular.",
		"What is Angular routing and how does it work?",
		"Explain the concept of dependency injection in Angular.",
		"What is AWS Cloud and how does it differ from traditional hosting?",
		"What are microservices and why are they used in modern application development?",
		"How can AWS Cloud support the deployment of microservices architecture?",
		"What are the key AWS services used for deploying and managing microservices?",
		"How can you ensure security and compliance when deploying microservices on AWS Cloud?",
		"What is AWS Elastic Beanstalk and how can it simplify the deployment of microservices?",
		"How does AWS Lambda fit into the microservices architecture on AWS Cloud?",
		"What are the best practices for monitoring and scaling microservices on AWS Cloud?",
		"How can you automate the deployment and management of microservices using AWS CloudFormation?",
		"What are the cost implications of deploying microservices on AWS Cloud?",
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
