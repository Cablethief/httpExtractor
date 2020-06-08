package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	// "github.com/gorilla/mux"
)

// func parseFlags() {
// 	flag.BoolVar(&flgProduction, "production", false, "if true, we start HTTPS server")
// 	flag.Parse()
// }

func base(w http.ResponseWriter, r *http.Request, l *log.Logger, re *regexp.Regexp) {
	requestDump, err := httputil.DumpRequest(r, true)

	if err != nil {
		fmt.Println(err)
	}

	matched := re.Match(requestDump)
	if matched {
		l.Printf("%s - %s - %s\n", r.Method, r.Host, r.RequestURI)
		l.Printf("%q\n", re.FindAll(requestDump, -1))
	}
	fmt.Fprintf(w, "OK")
}

func main() {
	// parseFlags()
	println("Starting up :D")

	logger := log.New(os.Stdout, "", log.LstdFlags)

	if os.Getenv("MATCH_REGEX") == "" {
		logger.Fatal("Please Provide a MATCH_REGEX environment variable")
	}

	logger.Println("MATCH_REGEX - '" + os.Getenv("MATCH_REGEX") + "'")

	regexMatch, err := regexp.Compile(os.Getenv("MATCH_REGEX"))
	if err != nil {
		logger.Fatal(err)
	}

	// go emailer()

	// router := mux.NewRouter()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		base(w, r, logger, regexMatch)
	})

	logger.Println("Listening on port 8080")
	logger.Println("------------------------")
	logger.Fatal(http.ListenAndServe(":8080", nil))

}
