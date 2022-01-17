package main

import "fmt"

func main() {
	name := "César O. Araújo"

	template := `
	<!DOCTYPE html>
	<html lang="pt-br">
	<head>
	<meta charset="UTF-8" />
	<title>Hello, world!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(template)
}
