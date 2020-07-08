package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"runtime"
)

func printStats() {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	envVars := os.Environ()

	osname := runtime.GOOS
	arch := runtime.GOARCH

	fmt.Println("Runtime Info:")
	fmt.Println("    pwd: ", workingDir)
	fmt.Println("    os: ", osname)
	fmt.Println("    arch: ", arch)
	fmt.Println("    hostname: ", hostname)
	fmt.Println("    username: ", user.Username)

	fmt.Println("    Env Vars:")
	for _, eVar := range envVars {
		fmt.Println("        ", eVar)
	}
	fmt.Println("")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Saying Hello")

	page := `
<html>
<head>
<title>Make Way For Docklings</title
</head>
<body>
<h1>Hello Dockling!</h1>
<p>Why not change the output of this page and then build and run it again in docker?</p>
</body>`

	fmt.Fprintf(w, page)
}

func serve(port string) {
	http.HandleFunc("/", helloHandler)
	fmt.Printf("Serving HTTP on :%s\n", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	fmt.Println("Make Way For Docklings\n")
	printStats()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8042"
	}

	serve(port)

}
