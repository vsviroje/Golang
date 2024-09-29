package util

import (
	"encoding/json"
	"net/url"
	"strings"
)

// RegexPattern List of available regular expressions to be used with function DoesStringMatch below
var RegexPattern = map[string]string{
	"ExceptAlphaAndDashes":        `[^a-zA-Z_\-]`,    // Match: Any other characters besides alphabets, underscore and hypen
	"ExceptAlphaNumericAndDashes": `[^a-zA-Z0-9_\-]`, // Match: Any other characters besides alphabets, numbers, underscore and hypen
}

func JSONDecoder(s string) *json.Decoder {
	d := json.NewDecoder(strings.NewReader(s))
	d.UseNumber()
	return d
}

func CreateUrlString(baseUrl string, subUrl string) (*string, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(subUrl)
	if err != nil {
		return nil, err
	}
	uri := base.ResolveReference(endpoint).String()
	return &uri, nil
}
