package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/gomodule/redigo/redis"
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

// namesHandler has a form for saving names and displays a list of names saved
func namesHandler(redis_addr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// attempt to connect to redis
		c, err := redis.Dial("tcp", redis_addr)
		if err != nil {
			if strings.HasSuffix(err.Error(), "connect: connection refused") ||
				strings.HasSuffix(err.Error(), "no such host") {
				fmt.Println("Redis Not Available on ", redis_addr)

				// If redis is not connected, we just print a static page.
				page := `
<html>
<head>
<title>Make Way For Docklings</title
</head>
<body>
<h1>Redis Dockling!</h1>
<p>redis is not available on %s</p>
</body>`

				fmt.Fprintf(w, fmt.Sprintf(page, redis_addr))
				return

			} else {
				// Error out.
				fmt.Println("Unexpected Error Connecting to Redis: ", err)
				http.Error(w, "Unexpected Error Connecting To Redis", 500)
				return
			}
		}

		// From here on, we're connected to redis

		var body string
		// get a list of things from redis.
		names, err := redis.Strings(c.Do("LRANGE", "names", 0, -1))
		if err != nil {
			fmt.Println("ERR FROM REDIS", err)
			http.Error(w, "Unexpected Error Reading From Redis", 500)
			return
		}

		// a simple form and a simple <ul>
		bodyTemplate :=
			`
<form action="/name_saver/save_name">
    <label for="name">Save a name:</label>
    <input name="name" id="name" value="">
    <button>Save Name</button>
</form>
<p>you are connected to redis on %s<p>
<p>Here is the list of names you've saved:<p>
<ul>
%s
</ul>
`
		nameList := ""
		for _, name := range names {
			nameList = nameList + fmt.Sprintf("<li>%s</li>\n", name)
		}

		body = fmt.Sprintf(bodyTemplate, redis_addr, nameList)

		page := `
<html>
<head>
<title>Make Way For Docklings</title
</head>
<body>
<h1>Redis Dockling!</h1>
%s
</body>`

		renderedPage := fmt.Sprintf(page, body)

		fmt.Fprintf(w, renderedPage)

	}

}

func addNameHandler(redis_addr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get parameter from get params,
		fmt.Println("Adding Name")

		addedNames, ok := r.URL.Query()["name"]

		if !ok {
			fmt.Println("NO KEYS")
		}

		// attempt to connect to redis
		c, err := redis.Dial("tcp", redis_addr)
		if err != nil {
			// Error out.
			fmt.Println("Error Connecting to Redis: ", err)
			http.Error(w, "Error Connecting To Redis", 500)
			return
		}

		// add a key to redis
		for _, addedName := range addedNames {
			fmt.Println("Adding: ", addedName)
			_, err := c.Do("RPUSH", "names", addedName)
			if err != nil {
				fmt.Println("BAD ADD ", err)
				http.Error(w, "Error Adding To Redis", 500)
				return
			}
		}

		// redirect
		http.Redirect(w, r, "/name_saver", 302)
	}
}

func serve(port string, redis_addr string) {
	http.HandleFunc("/name_saver", namesHandler(redis_addr))
	http.HandleFunc("/name_saver/save_name", addNameHandler(redis_addr))
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

	redis_host := os.Getenv("REDIS_HOST")
	if redis_host == "" {
		redis_host = "localhost"
	}

	redis_port := os.Getenv("REDIS_PORT")
	if redis_port == "" {
		redis_port = "6379"
	}

	redis_addr := fmt.Sprintf("%s:%s", redis_host, redis_port)

	serve(port, redis_addr)

}
