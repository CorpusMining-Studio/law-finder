package router

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Router struct {
	laws embed.FS
}

func New(lawResources embed.FS) *Router {
	return &Router{
		laws: lawResources,
	}
}

func (r *Router) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("/law-finder", r.findLaw)
	return mux
}

type findLawRequest struct {
	Article string `json:"article"`
}

func (rr *Router) findLaw(w http.ResponseWriter, r *http.Request) {
	// Check target law in query parameter
	target := r.URL.Query().Get("law")
	if target == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	// Read the file from the embedded filesystem
	targetFile := "static/" + target + ".md"
	law, err := rr.laws.ReadFile(targetFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(law))
		return
	case http.MethodPost:
		// Check if the request body is empty
		if r.Body == nil {
			http.Error(w, "Bad Request: Body is nil", http.StatusBadRequest)
			return
		}
		// Parse the request body
		var req findLawRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: Can't parse request body", http.StatusBadRequest)
			return
		}
		if req.Article == "" {
			http.Error(w, "Bad Request: Missing required field", http.StatusBadRequest)
			return
		}
		articleInt, err := normalizeArticle(req.Article)
		if err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(findArticle(law, articleInt)))
		return
	}
	w.Write([]byte(law))
}

func findArticle(law []byte, article int) string {
	text := string(law)
	// Construct the regex pattern for finding the specified article and the next article.
	pattern := fmt.Sprintf(`(?s)(### \*\*第 %d 條\*\*)\n(.*?)\n### \*\*第 %d 條\*\*`, article, article+1)

	// Compile the regex.
	re := regexp.MustCompile(pattern)

	// Find the match
	match := re.FindStringSubmatch(text)

	if len(match) > 2 {
		return match[1] + match[2] // Return the title and the captured content between the articles
	} else {
		// Handle the case for the last article if "第 n+1 條" doesn't exist
		endPattern := fmt.Sprintf(`(?s)(### \*\*第 %d 條\*\*)\n(.*)`, article)
		endRe := regexp.MustCompile(endPattern)
		match = endRe.FindStringSubmatch(text)
		if len(match) > 2 {
			return match[1] + match[2]
		}
	}
	return "Article not found"
}
