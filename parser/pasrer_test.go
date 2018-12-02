package parser

import (
	"fmt"
	"testing"
)

func TestCurlStrParse(t *testing.T) {
	str := `curl --request POST \
  --url 'http://127.0.0.1:8080/api/v1/atm/terminal/new?=' \
  --header 'content-type: application/json' \
  --data '{"name":"igor", "password":"1323", "invite_token":"we23e"}`
	res := Curl{InString: str}
	res.Fill()
	fmt.Println(res.Method)

}
