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
		"What is a regular expression?",
		"What are the benefits of using regular expressions?",
		"How do you define a regular expression in JavaScript?",
		"Explain the difference between greedy and lazy quantifiers in regular expressions.",
		"What is the purpose of the test() method in JavaScript regular expressions?",
		"How do you match a specific word in a string using regular expressions?",
		"Explain the difference between exec() and test() methods in JavaScript regular expressions.",
		"How do you match any character using a dot . in regular expressions?",
		"What is the purpose of the ^ and $ anchors in regular expressions?",
		"How do you match a digit in a string using regular expressions?",
		"Explain the difference between match() and test() methods in JavaScript regular expressions.",
		"How do you match a word boundary in a string using regular expressions?",
		"What is the purpose of character classes in regular expressions?",
		"How do you match a specific number of characters using quantifiers in regular expressions?",
		"Explain the difference between match() and exec() methods in JavaScript regular expressions.",
		"How do you match a non-word character in a string using regular expressions?",
		"What is the purpose of the i flag in regular expressions?",
		"How do you match a whitespace character in a string using regular expressions?",
		"Explain the difference between the test() method and the search() method in JavaScript regular expressions.",
		"How do you match a specific range of characters using character classes in regular expressions?",
		"What is the purpose of the g flag in regular expressions?",
		"How do you match a specific word with case insensitivity using regular expressions?",
		"Explain the difference between the exec() method and the search() method in JavaScript regular expressions.",
		"How do you match a specific character multiple times using quantifiers in regular expressions?",
		"What is the purpose of the m flag in regular expressions?",
		"How do you match a specific word at the beginning of a string using regular expressions?",
		"Explain the difference between the replace() method and the test() method in JavaScript regular expressions.",
		"How do you match a specific word at the end of a string using regular expressions?",
		"What is the purpose of the s flag in regular expressions?",
		"How do you match a specific word anywhere in a string using regular expressions?",
		"Explain the difference between the match() method and the replace() method in JavaScript regular expressions.",
		"How do you match a word that starts with a specific letter using regular expressions?",
		"What is the purpose of the y flag in regular expressions?",
		"How do you match a word that ends with a specific letter using regular expressions?",
		"Explain the difference between the split() method and the match() method in JavaScript regular expressions.",
		"How do you match a word that contains a specific letter using regular expressions?",
		"What is the purpose of the u flag in regular expressions?",
		"How do you match a word that does not contain a specific letter using regular expressions?",
		"Explain the difference between the exec() method and the split() method in JavaScript regular expressions.",
		"How do you match a word that starts and ends with the same letter using regular expressions?",
		"What is the purpose of the sticky property in regular expressions?",
		"How do you match a word that starts with a vowel using regular expressions?",
		"Explain the difference between the source property and the flags property in JavaScript regular expressions.",
		"How do you match a word that contains only vowels using regular expressions?",
		"What is the purpose of the ignoreCase property in regular expressions?",
		"How do you match a word that contains only consonants using regular expressions?",
		"Explain the difference between the lastIndex property and the global property in JavaScript regular expressions.",
		"How do you match a word that has at least three characters using regular expressions?",
		"What is the purpose of the dotAll property in regular expressions?",
		"How do you match a word that has exactly five characters using regular expressions?",
		"Explain the difference between the exec() method and the matchAll() method in JavaScript regular expressions.",
		"How do you match a word that has an even number of characters using regular expressions?",
		"What is the purpose of the unicode property in regular expressions?",
		"How do you match a word that has an odd number of characters using regular expressions?",
		"Explain the difference between the exec() method and the replaceAll() method in JavaScript regular expressions.",
		"How do you match a word that starts and ends with the same character using regular expressions?",
		"What is the purpose of the flags property in regular expressions?",
		"How do you match a word that does not start with a specific letter using regular expressions?",
		"Explain the difference between the source property and the sticky property in JavaScript regular expressions.",
		"How do you match a word that does not end with a specific letter using regular expressions?",
		"What is the purpose of the matchAll() method in regular expressions?",
		"How do you match a word that does not contain a specific letter using regular expressions?",
		"Explain the difference between the global property and the ignoreCase property in JavaScript regular expressions.",
		"How do you match a word that does not start and end with the same letter using regular expressions?",
		"What is the purpose of the replaceAll() method in regular expressions?",
		"How do you match a word that does not start with a vowel using regular expressions?",
		"Explain the difference between the unicode property and the dotAll property in JavaScript regular expressions.",
		"How do you match a word that does not contain only vowels using regular expressions?",
		"What is the purpose of the hasIndices property in regular expressions?",
		"How do you match a word that does not contain only consonants using regular expressions?",
		"Explain the difference between the hasIndices property and the flags property in JavaScript regular expressions.",
		"How do you match a word that does not have at least three characters using regular expressions?",
		"What is the purpose of the toString() method in regular expressions?",
		"How do you match a word that does not have exactly five characters using regular expressions?",
		"Explain the difference between the toString() method and the test() method in JavaScript regular expressions.",
		"How do you match a word that does not have an even number of characters using regular expressions?",
		"What is the purpose of the Symbol.match property in regular expressions?",
		"How do you match a word that does not have an odd number of characters using regular expressions?",
		"Explain the difference between the Symbol.match property and the Symbol.replace property in JavaScript regular expressions.",
		"How do you match a word that does not start and end with the same character using regular expressions?",
		"What is the purpose of the Symbol.search property in regular expressions?",
		"How do you match a word that does not start with a specific letter using regular expressions?",
		"Explain the difference between the Symbol.search property and the Symbol.split property in JavaScript regular expressions.",
		"How do you match a word that does not end with a specific letter using regular expressions?",
		"What is the purpose of the Symbol.split property in regular expressions?",
		"How do you match a word that does not contain a specific letter using regular expressions?",
		"Explain the difference between the Symbol.split property and the Symbol.matchAll property in JavaScript regular expressions.",
		"How do you match a word that does not start and end with the same letter using regular expressions?",
		"What is the purpose of the Symbol.matchAll property in regular expressions?",
		"How do you match a word that does not start with a vowel using regular expressions?",
		"Explain the difference between the Symbol.replace property and the Symbol.replaceAll property in JavaScript regular expressions.",
		"How do you match a word that does not contain only vowels using regular expressions?",
		"What is the purpose of the Symbol.replace property in regular expressions?",
		"How do you match a word that does not contain only consonants using regular expressions?",
		"Explain the difference between the Symbol.replaceAll property and the Symbol.match property in JavaScript regular expressions.",
		"How do you match a word that does not have at least three characters using regular expressions?",
		"What is the purpose of the Symbol.replace property in regular expressions?",
		"How do you match a word that does not have exactly five characters using regular expressions?",
		"Explain the difference between the Symbol.replace property and the Symbol.search property in JavaScript regular expressions.",
		"How do you match a word that does not have an even number of characters using regular expressions?",
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
