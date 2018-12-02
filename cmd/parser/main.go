package main

import (
	"fmt"

	"github.com/khigor777/http-swagger/parser"
)

func main() {
	str := `curl --request DELETE \
  --url 'https://hybrid.humaniq.com/tapatybe/api/v5/service/delete_user/375295665333?X-Tester=1eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMiLCJuYW1lIjoidGVzdGVyIiwiaWF0IjoyMTMxMjMxMzIzNDI0fQ.01OaFxlraNJo4T4pMyUGujPhX6zPNbIQVvR8UvLXQoNnOsVWynvxWdXJfWH5AMLtq_OMmni2gp1mVJ-ka2gwGw&=' \
  --header 'content-type: application/json' \
  --header 'x-tester: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMiLCJuYW1lIjoidGVzdGVyIiwiaWF0IjoyMTMxMjMxMzIzNDI0fQ.01OaFxlraNJo4T4pMyUGujPhX6zPNbIQVvR8UvLXQoNnOsVWynvxWdXJfWH5AMLtq_OMmni2gp1mVJ-ka2gwGw'`
	res := parser.Curl{InString: str}
	res.Fill()
	fmt.Println(res.Headers)
}
