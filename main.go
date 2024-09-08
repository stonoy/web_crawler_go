package main

import "fmt"

func main() {
	fmt.Println(getURLsFromHTML(`
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
	`, "https://blog.boot.dev"))

}
