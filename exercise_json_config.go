// json config
// isolation code

package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
)

type Config struct {
    Store string
    StoreConfig json.RawMessage
}

type RedisConfig struct {
    Addr string `json:"addr"`
    DB int `json:"db"`
}

func NewConfig(m json.RawMessage) *RedisConfig {
    c := new(RedisConfig)
    json.Unmarshal(m, c)
    return c
}

func main() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Fatalln("Read file config.json error:", err)
    }
    c := new(Config)
    // notice that slice 3 maybe the bow code
    json.Unmarshal(data[3:], c)
    r := NewConfig(c.StoreConfig)
    fmt.Println(r.Addr)
}
