package main

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    r, err := zip.OpenReader("readme.zip")
    if err != nil {
        log.Fatal(err)
    }
    defer r.Close()
    for _, f := range r.File {
        fmt.Printf("Contents of %s:\n", f.Name)
        rc, err := f.Open()

        if err != nil {
            log.Fatal(err)
        }
        _, err = io.CopyN(os.Stdout, rc, 68)
        if err == io.EOF {
            rc.Close()
            fmt.Println()
            continue
        }
        if err != nil {
            log.Fatal(err)
        }
        rc.Close()
        fmt.Println()
    }
}
