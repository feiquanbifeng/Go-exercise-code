package main

import (
    "io"
    // "bufio"
    "bytes"
    "fmt"
    "net"
    "net/url"
    "strings"
)

var version = "HTTP/1.1"

type Protoer interface {
    Conn(url string)
    Get()
    Post()
    Close()
}

var Http *HTTP

func init() {
    url := "http://pan.baidu.com/disk/home"
    Http = &HTTP{
        Header: make([]string, 0),
    }
    Http.Conn(url)
    Http.SetHeader("Host: " + Http.Url.Host)
}

type HTTP struct {
    Line   string
    Url    *url.URL
    Header []string
    Con    net.Conn
    Body   string
}

func (t *HTTP) Conn(rawurl string) error {
    u, err := url.Parse(rawurl)
    if err != nil {
        return err
    }
    t.Url = u
    conn, err := net.Dial("tcp", u.Host+":80")
    if err != nil {
        return err
    }
    t.Con = conn
    return nil
}

func (t *HTTP) SetHeader(header string) {
    t.Header = append(t.Header, header)
}

func (t *HTTP) Get(method string) {
    t.Line = fmt.Sprint(method, " ", t.Url.Path, " ", version)
}

func (t *HTTP) Request() ([]byte, error) {
    defer t.Con.Close()
    result := fmt.Sprint(t.Line+"\n"+strings.Join(t.Header, "\r\t"), "\r\n\r\n")
    fmt.Println(result)
    fmt.Fprintf(t.Con, result)
    // status, err := bufio.NewReader(t.Con).ReadString('\n')
    res := bytes.NewBuffer(nil)
    var buf [512]byte
    for {
        n, err := t.Con.Read(buf[0:])
        res.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }
    return res.Bytes(), nil
}

func main() {
    Http.Get("GET")
    fmt.Println(Http)
    data, _ := Http.Request()
    fmt.Println(string(data))
}
