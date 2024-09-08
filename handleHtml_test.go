package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs with query params and http",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<div>
					<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one/search?q=any">
					<span>Boot.dev</span>
				</a>
				</div>
				<h1>Welcome</h1>
				<ul>
					<li><a href="/any/one"></a></li>
					<li><a href="http://google.com"></a></li>
				</ul>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one/search?q=any", "https://blog.boot.dev/any/one", "http://google.com"},
		},
		{
			name:     "absolute and relative URLs with hashs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one#thearea">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one/search?q=any#requiredarea">
					<span>Boot.dev</span>
				</a>
				<section>
					<div>
						<img src="any/where/in/web" />
					</div>
				</section>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one#thearea", "https://other.com/path/one/search?q=any#requiredarea"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			urls, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("%v : expected -> %v, got -> %v", i, tc.expected, urls)
				return
			}

			if !reflect.DeepEqual(urls, tc.expected) {
				t.Errorf("%v : expected -> %v, got -> %v", i, tc.expected, urls)
			}
		})
	}
}
