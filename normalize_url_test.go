package main

import (
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name     string
		rawUrl   string
		expected string
	}{
		{
			name:     "https_/",
			rawUrl:   "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "https",
			rawUrl:   "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http_/",
			rawUrl:   "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http",
			rawUrl:   "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			normalizedUrl, err := normalizeUrl(tc.rawUrl)
			if err != nil {
				t.Errorf("%vth test - %v failed, error -> %v", i, tc.name, err)
				return
			}

			if tc.expected != normalizedUrl {
				t.Errorf("raw_url -> %v, expected -> %v, got -> %v", tc.rawUrl, tc.expected, normalizedUrl)
			}
		})
	}
}
