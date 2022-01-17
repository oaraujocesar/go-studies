package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "César O. Araújo"

	template := fmt.Sprint(`
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
	`)

	newfile, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creating new file", err)
	}

	defer newfile.Close()

	io.Copy(newfile, strings.NewReader(template))
}
