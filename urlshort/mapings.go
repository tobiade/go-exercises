package urlshort

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type pathToURLEntry struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

//GetMappingsFromJSON will parse JSON and get URL mappings
func GetMappingsFromJSON() (map[string]string, error) {
	sample := []byte(`[{"path":"/fantasy", "url":"https://fantasy.premierleague.com/"}]`)
	f := func(v interface{}) error {
		return json.Unmarshal(sample, v)
	}
	return unmarshalToMap(f)
}

//GetMappingsFromYAML will parse YAML and get URL mappings
func GetMappingsFromYAML(filePath string) (map[string]string, error) {
	mapping, err := parseYAML([]byte(defaultYAML()))
	if err != nil {
		return nil, err
	}

	//this is so annoying - how can we make this less verbose/cleaner???
	if filePath != "" {
		if d, err := ioutil.ReadFile(filePath); err != nil {
			return nil, err
		} else {
			f := func(v interface{}) error {
				return yaml.Unmarshal(d, v)
			}
			mapping, err = unmarshalToMap(f)
			if err != nil {
				return nil, err
			}

		}
	}
	return mapping, nil
}

func defaultYAML() string {
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	return yaml
}

func parseYAML(yml []byte) (map[string]string, error) {
	var entries []pathToURLEntry
	err := yaml.Unmarshal(yml, &entries)
	m := make(map[string]string)
	for _, entry := range entries {
		m[entry.Path] = entry.URL
	}
	return m, err
}

func unmarshalToMap(f func(v interface{}) error) (map[string]string, error) {
	var entries []pathToURLEntry
	err := f(&entries)
	m := make(map[string]string)
	for _, entry := range entries {
		m[entry.Path] = entry.URL
	}
	return m, err
}
