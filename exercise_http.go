package main

import (
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

const (
    routeObject = "/object"
)

var routes map[string]Controller

func init() {
    routes = make(map[string]Controller)
    RESTRouter(routeObject, &MainController{})
    start()
}

func RESTRouter(patter string, c Controller) {
    routes[patter] = c
}

type Server struct{}

func (t *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    url := r.URL.String()
    m := r.Method
    r.ParseForm()
    fmt.Println(url, m, r.Form["name"])
    if strings.HasPrefix(url, routeObject) {
        if c, ok := routes[routeObject]; ok {
            c.Init(w, r)
            switch m {
            case "POST":
                c.Post()
            case "GET":
                c.Get()
            case "PUT":
                c.Put()
            case "DELETE":
                c.Delete()
            }
        }
    }
}

func start() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: &Server{},
    }
    server.ListenAndServe()
}

type Controller interface {
    Post()
    Get()
    Put()
    Delete()
    Input() url.Values
    Init(w http.ResponseWriter, r *http.Request)
}

type MainController struct {
    ResponseWrite http.ResponseWriter
    Request       *http.Request
}

func (s *MainController) Post() {
    fmt.Println("Post ...")
}

func (s *MainController) Get() {
    fmt.Println("call get method")
    s.ResponseWrite.Write([]byte("Hello"))
}

func (s *MainController) Put() {

}

func (s *MainController) Delete() {

}

func (s *MainController) Init(w http.ResponseWriter, r *http.Request) {
    s.ResponseWrite = w
    s.Request = r
}

func (s *MainController) Input() url.Values {
    if s.Request.Form != nil {
        s.Request.ParseForm()
    }
    return s.Request.Form
}

func main() {
}
