// simple crawler use go
// just demo
package main

import (
    "net/http"
    "io/ioutil"
    "time"
    "fmt"
    "regexp"
    "strings"
    _ "database/sql"

    "log"
    "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

const (
    URL = "127.0.0.1:27017"
)

type Article struct {
    Title string
    CreateTime time.Time
}

func save(title string) {
    session, err := mgo.Dial(URL)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("crawler").C("article")
        err = c.Insert(&Article{title, time.Now()})
        if err != nil {
                fmt.Println(err)
                log.Fatal(err)
        }

        /*result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)*/
}

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
            save(v[1])
        }
    }
}
