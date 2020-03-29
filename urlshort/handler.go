package urlshort

import (
	"net/http"
	"strings"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index := strings.LastIndex(r.URL.Path, "/")
		trimmedURL := r.URL.Path[index:]
		if url, ok := pathsToUrls[trimmedURL]; ok {
			http.Redirect(w, r, url, 301)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
}
