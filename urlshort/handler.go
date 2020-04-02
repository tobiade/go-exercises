package urlshort

import (
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trimmedURL := trimURL(r.URL.Path)
		if url, ok := pathsToUrls[trimmedURL]; ok {
			http.Redirect(w, r, url, 301)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
}

//DBHandler will read from DB and map paths to corresponding URL
func DBHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trimmedURL := trimURL(r.URL.Path)
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket(bytes(bucketName))
			redirectURL := b.Get(bytes(trimmedURL))
			if redirectURL != nil {
				http.Redirect(w, r, string(redirectURL), 301)
			} else {
				fallback.ServeHTTP(w, r)
			}
			return nil
		})
	}
}

func trimURL(url string) string {
	index := strings.LastIndex(url, "/")
	trimmedURL := url[index:]
	return trimmedURL
}
