package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`
<html>
	<head>
		<meta charset="UTF-8">
		<title>Hello, World</title>
		<script src="static/test.js"></script>
	</head>
	<body>
		<h1 id="title">Hello World 😊</h1>
		<a href="https://www.google.com/">Let's Google :D</a>
	</body>
</html>`,
	)
}

func main() {
	const (
		ConnHost = "http://localhost"
		ConnPort = 9000
	)

	http.HandleFunc("/hello", hello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	go http.ListenAndServe(fmt.Sprintf(":%d", ConnPort), nil)

	fmt.Printf("Access %s:%d/hello\n", ConnHost, 9000)

	var input string
	fmt.Scanln(&input)
}
