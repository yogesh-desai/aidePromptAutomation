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
		"What is GraphQL?",
		"How does GraphQL differ from REST?",
		"Explain the advantages of using GraphQL over REST.",
		"Can you define a schema in GraphQL?",
		"What is a resolver in GraphQL?",
		"Explain the role of a resolver in GraphQL.",
		"How do you define a type in GraphQL?",
		"What is a scalar type in GraphQL?",
		"List some scalar types in GraphQL.",
		"How do you define custom scalar types in GraphQL?",
		"What is a query in GraphQL?",
		"Differentiate between query and mutation in GraphQL.",
		"What is a mutation in GraphQL?",
		"Explain the concept of a subscription in GraphQL.",
		"How does GraphQL handle versioning?",
		"What are GraphQL directives?",
		"Provide examples of GraphQL directives.",
		"How do you define arguments in a GraphQL query?",
		"Can you have nested queries in GraphQL?",
		"Explain the concept of fragments in GraphQL.",
		"What is an operation type in GraphQL?",
		"How does GraphQL handle authentication and authorization?",
		"Explain the concept of introspection in GraphQL.",
		"Can you disable introspection in GraphQL?",
		"What is the difference between a fragment and a query in GraphQL?",
		"How does GraphQL handle errors?",
		"Explain the concept of pagination in GraphQL.",
		"What is batching in GraphQL?",
		"How do you handle file uploads in GraphQL?",
		"Explain the concept of union types in GraphQL.",
		"What is a schema stitching in GraphQL?",
		"How do you perform error handling in GraphQL?",
		"Explain the concept of schema delegation in GraphQL.",
		"What is a federation in GraphQL?",
		"How does GraphQL handle caching?",
		"What are the differences between GraphQL subscriptions and WebSockets?",
		"How do you handle long-running queries in GraphQL?",
		"What are some best practices for designing a GraphQL schema?",
		"How do you optimize GraphQL queries?",
		"Explain the concept of DataLoader in GraphQL.",
		"Can you nest mutations in GraphQL?",
		"What is a schema directive in GraphQL?",
		"How does GraphQL handle partial updates?",
		"Explain the concept of schema-first development in GraphQL.",
		"What are some popular GraphQL clients?",
		"How does GraphQL handle circular dependencies in a schema?",
		"Can you provide an example of a GraphQL query?",
		"How do you handle errors in a GraphQL resolver?",
		"Explain the concept of schema stitching in GraphQL.",
		"How do you handle versioning in GraphQL APIs?",
		"What is the role of Apollo Server in GraphQL?",
		"Explain the concept of subscriptions in GraphQL.",
		"How does GraphQL handle data validation?",
		"What are some common security concerns in GraphQL?",
		"How does GraphQL handle caching?",
		"What are some performance optimization techniques for GraphQL?",
		"How do you implement authentication and authorization in GraphQL?",
		"Explain the concept of schema delegation in GraphQL.",
		"What is the difference between a schema and a resolver in GraphQL?",
		"How do you handle nested mutations in GraphQL?",
		"Can you provide an example of a GraphQL mutation?",
		"What are some tools for testing GraphQL APIs?",
		"How do you handle file uploads in GraphQL mutations?",
		"Explain the concept of schema stitching in GraphQL.",
		"What are some best practices for writing efficient GraphQL queries?",
		"How do you handle batch processing in GraphQL?",
		"What are some common anti-patterns in GraphQL?",
		"How does GraphQL handle error reporting?",
		"Can you provide an example of a GraphQL subscription?",
		"What are some strategies for caching GraphQL responses?",
		"How do you handle schema evolution in GraphQL?",
		"Explain the concept of schema delegation in GraphQL.",
		"What are some tools for generating GraphQL schemas?",
		"How do you handle pagination in GraphQL queries?",
		"Can you provide an example of a GraphQL schema?",
		"What are some limitations of GraphQL?",
		"How do you handle authorization in GraphQL subscriptions?",
		"What is the role of GraphQL introspection?",
		"How do you handle concurrency in GraphQL?",
		"Explain the concept of federated GraphQL schemas.",
		"What are some popular GraphQL frameworks?",
		"How do you handle versioning in GraphQL subscriptions?",
		"Can you provide an example of using custom scalar types in GraphQL?",
		"What are some strategies for optimizing GraphQL performance?",
		"How do you handle database transactions in GraphQL mutations?",
		"Explain the concept of data loaders in GraphQL.",
		"What are some common patterns for structuring GraphQL schemas?",
		"How do you handle schema validation in GraphQL?",
		"What is the role of the GraphQL schema language?",
		"How do you handle errors in GraphQL subscriptions?",
		"What are some considerations for deploying GraphQL APIs?",
		"How do you handle schema migrations in GraphQL?",
		"What are some popular GraphQL server implementations?",
		"How do you handle error handling in GraphQL subscriptions?",
		"What are some differences between GraphQL and gRPC?",
		"How do you handle schema validation in GraphQL mutations?",
		"What are some strategies for optimizing GraphQL queries?",
		"How do you handle data consistency in GraphQL mutations?",
		"What are some considerations for scaling GraphQL APIs?",
		"Can you provide an example of using directives in GraphQL schemas?",
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
