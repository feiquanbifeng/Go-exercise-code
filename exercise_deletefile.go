package main

import (
    "path/filepath"
    "fmt"
    "os"
    "strings"
)

func main() {
    path := "F:\\plugins"
    absPath, _ := filepath.Abs(path)

    // names := []string{"aa.js", "zh.js", "zh-cn.js"}

    err := filepath.Walk(absPath, func(path string, fi os.FileInfo, err error) error {
        if nil == fi {
            return err
        }

        if fi.IsDir() {
            return nil
        }

        name := fi.Name()
        p := filepath.Dir(path)

        if strings.HasSuffix(p, "lang") {
            switch name {
            case "en.js":
            case "zh.js":
            case "zh-cn.js":
            default:
                os.Remove(p + "\\" + name)
            }
        }

        // fmt.Println(name, filepath.Dir(path))
        return nil
    })

    if err != nil {
        panic(err)
    }

    fmt.Println("删除指定文件夹下的文件成功！")
}
