package api_middleware

import (
	"log"
	"net/http"
	"strings"
)

var Counter = make(map[string]int)

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("\n In middleware, URI-Method:", r.RequestURI, r.Method)
		if (r.Method == "GET" && strings.Contains(r.RequestURI, "student")) {
			_, ok := Counter["get_stud"]
			if ok {
				Counter["get_stud"] += 1
			} else {
				Counter["get_stud"] = 1
			}
		} else if (r.Method == "GET" && strings.Contains(r.RequestURI, "attendance")) {
			_, ok := Counter["get_attn"]
			if ok {
				Counter["get_attn"] += 1
			} else {
				Counter["get_attn"] = 1
			}
		} else if (r.Method == "POST" && strings.Contains(r.RequestURI, "student")) {
			_, ok := Counter["post_stud"]
			if ok {
				Counter["post_stud"] += 1
			} else {
				Counter["post_stud"] = 1
			}
		} else if (r.Method == "POST" && strings.Contains(r.RequestURI, "attendance")) {
			_, ok := Counter["post_attn"]
			if ok {
				Counter["post_attn"] += 1
			} else {
				Counter["post_attn"] = 1
			}
		}
		next.ServeHTTP(w, r)
	})
}
