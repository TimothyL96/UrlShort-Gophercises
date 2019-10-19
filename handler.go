package urlshort_gophercises

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	handler := func(writer http.ResponseWriter, request *http.Request) {
		if pathsToUrls[request.URL.RequestURI()] != "" {
			http.Redirect(writer, request, pathsToUrls[request.URL.RequestURI()], 302)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}

	return handler
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	parsedYAMLs, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYAMLs)

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) (ymlMaps []map[string]string, err error) {
	err = yaml.Unmarshal(yml, &ymlMaps)
	return
}

func buildMap(parsedYAMLs []map[string]string) map[string]string {
	pathMap := make(map[string]string)

	for _, yml := range parsedYAMLs {
		pathMap[yml["path"]] = yml["url"]
	}

	return pathMap
}
