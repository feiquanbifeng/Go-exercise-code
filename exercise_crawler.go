// simple crawler use go
// just demo
package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "regexp"
    "strings"
)

func main() {
    client := &http.Client{}
    request, _ := http.NewRequest("GET", "http://www.36kr.com/", nil)
    request.Header.Set("Accept", "text/html")
    response, _ := client.Do(request)
    if response.StatusCode == 200 {
        body, _ := ioutil.ReadAll(response.Body)
        bodystr := string(body)
        rr := strings.NewReplacer("&nbsp;", "", "\t", "", "\n", "", "\r", "", "\f", "", " ", "")
        bodystr = rr.Replace(strings.TrimSpace(bodystr))
        r := regexp.MustCompile(`(?U)\"right-col\"><h1><a.*>(.*)</a>`)
        var matchs [][]string
        matchs = r.FindAllStringSubmatch(bodystr, -1)
        for _, v := range matchs {
            fmt.Println(v[1])
        }
    }
}
