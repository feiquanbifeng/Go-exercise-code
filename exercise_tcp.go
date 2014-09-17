package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
)

func main() {
    service := "my.oschina.net:80"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    defer conn.Close()
    checkError(err)
    _, err = conn.Write([]byte("HEAD /u/593413/blog/305014 HTTP/1.0\r\nHost: my.oschina.net\r\n\r\n"))
    checkError(err)
    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
