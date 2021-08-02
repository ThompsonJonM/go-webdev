package main

import "fmt"

func main() {
	n := "James Bond"

	tpl := `
	<!DOCTYPE html>
	<hmtl lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + n + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)
}
