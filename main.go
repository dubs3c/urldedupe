package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

func queryParamKeys(query url.Values) string {
	keysSet := make(map[string]struct{})
	for k := range query {
		keysSet[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysSet))
	for k := range keysSet {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return strings.Join(keys, "&")
}

func dedup(rawUrl string, store map[string]struct{}) (string, error) {

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	queryKeys := queryParamKeys(parsedURL.Query())
	key := fmt.Sprintf("%s://%s%s?%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path, queryKeys)
	if _, exists := store[key]; !exists {
		store[key] = struct{}{}
		return rawUrl, err
	}

	return "", err
}

func main() {
	urlMap := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		rawURL := strings.TrimSpace(scanner.Text())
		u, err := dedup(rawURL, urlMap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing URL: %v\n", err)
			continue
		}

		if u != "" {
			fmt.Println(u)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
