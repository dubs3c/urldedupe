package main

import (
	"net/url"
	"testing"
)

func TestQueryParamKeys(t *testing.T) {
	testCases := []struct {
		rawQuery     string
		expectedKeys string
	}{
		{"a=1&b=2&c=3", "a&b&c"},
		{"c=3&a=1&b=2", "a&b&c"},
		{"b=2", "b"},
		{"", ""},
	}

	for _, tc := range testCases {
		parsedQuery, _ := url.ParseQuery(tc.rawQuery)
		keys := queryParamKeys(parsedQuery)
		if keys != tc.expectedKeys {
			t.Errorf("Expected keys '%s', got '%s' for query '%s'", tc.expectedKeys, keys, tc.rawQuery)
		}
	}
}

func TestDeduplicateURLs(t *testing.T) {

}
