package parser

import (
	"strings"
)

/*
curl --request POST \
  --url 'http://127.0.0.1:8080/api/v1/atm/terminal/new?=' \
  --header 'content-type: application/json' \
  --data '{"name":"igor", "password":"1323", "invite_token":"we23e"}'

curl -i -X GET -H 'Authorization: bearer <your_token>' 'https://appliance.com/api/v1.0/discovery'


*/

type Curl struct {
	Url       string
	Headers   map[string]string
	GetParams []string
	Method    string
	Body      []byte

	InString string
}

func (c *Curl) Fill() {
	parsed := c.ParseAll()
	for i, v := range parsed {
		if v == "--header" || v == "-H" {
			if c.Headers == nil {
				c.Headers = make(map[string]string)
			}
			c.Headers[parsed[i+1]] = parsed[i+2]
		}
		if v == "--url" {
			c.Url = parsed[i+1]
		}

		if v == "--request" {
			c.Method = parsed[i+1]
		}
	}
}

func (c *Curl) ParseAll() []string {
	end := strings.Replace(strings.Replace(c.InString, "\n", "", -1), `\`, "", -1)
	end = strings.Replace(end, "'", "", -1)
	split := strings.Split(end, " ")
	res := make([]string, 0, len(split))
	for _, v := range split {
		v := strings.Trim(v, " ")
		if v == "" {
			continue
		}
		res = append(res, v)
	}
	return res
}

func (c *Curl) parseMethod() string {
	methods := make(map[string]string)
	methods["Get"] = "GET"
	methods["Head"] = "HEAD"
	methods["Post"] = "POST"
	methods["Put"] = "PUT"
	methods["Patch"] = "PATCH"
	methods["Delete"] = "DELETE"
	methods["Connect"] = "CONNECT"
	methods["Options"] = "OPTIONS"
	methods["Trace"] = "TRACE"
	for _, v := range methods {
		if strings.Contains(c.InString, v) {
			return v
		}
	}
	return ""
}
