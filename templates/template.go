package templates

import (
	"bytes"
	"fmt"
	"html/template"

	"golang.org/x/net/html"
)

// TODO: check for regexp {{.*}}
// ParseTemplate reads the data and return recognized parameters
func ParseTemplate(data []byte) ([]string, error) {
	var parameters []string
	// parsing html template
	_, err := template.New("RequestTemplate").Parse(string(data))
	if err != nil {
		return nil, err
	}

	// parsing html and fetching all parameters
	tokenizer := html.NewTokenizer(bytes.NewReader(data))
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.TextToken:
			fmt.Println(string(tokenizer.Text()))
		case html.ErrorToken:
			return parameters, nil
		default:
			continue
		}
	}
}
