package templates

import (
	"bytes"
	"html/template"
	"regexp"

	"golang.org/x/net/html"
)

var parameter string = `{{.+}}`

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
			// check that token data doesn't start with \n
			tokenData := tokenizer.Token().Data
			if match, _ := regexp.MatchString(parameter, tokenData); match {
				parameters = append(parameters, tokenData)
			}
		case html.ErrorToken:
			return parameters, nil
		default:
			continue
		}
	}
}
