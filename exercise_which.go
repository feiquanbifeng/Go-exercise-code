// which just to find command in you path
// refer to the nodejs which
package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
    "runtime"
    strs "strings"
)

var (
    isAbsolute = absUnix
    r          *regexp.Regexp
    colon      = ":"
    isWin      = false
    root       string
)

func init() {
    if runtime.GOOS == "windows" {
        isAbsolute = absWin
        colon = ";"
        isWin = true

        wd, _ := os.Getwd()
        root = filepath.VolumeName(wd)
    }

    r = regexp.MustCompile("^([a-zA-Z]:|[\\\\/]{2}[^\\\\/]+[\\\\/][^\\\\/]+)?([\\\\/])?")
}

// to check the file is executable
func isExe(p string) bool {
    _, err := exec.LookPath(p)
    return err == nil
}

// resolve the path
// you can pass more than one parameter
// just like nodejs path.resolve
func resolve(arg ...string) string {

    var (
        l   int    = len(arg) - 1
        p   string = arg[l]
        ret string
    )

    if l == 0 {
        cw, _ := os.Getwd()
        ret, _ = filepath.Abs(cw + "/" + p)
        return filepath.FromSlash(ret)
    }
       
    for ; l >= 0; l-- {
        // if the last elem is absolute path then return
        if filepath.IsAbs(p) {
            ret = filepath.Clean(p)
            break
        }
        
        if l == 0 {
            ret = p
            break
        }
        
        p = arg[l-1] + "/" + p
    }
    
    return filepath.FromSlash(ret)
}

// it is similar to which in linux environment
// you can pass a func when call
func Which(cmd string, cb func(args ...interface{}) interface{}) interface{} {
    if isAbsolute(cmd) {
        return cb(nil, cmd)
    }

    var (
        pathEnv = strs.Split(os.Getenv("PATH"), colon)
        pathExt = []string{""}
    )

    if isWin {
        cwd, err := os.Getwd()

        if err != nil {
            log.Fatalf("获取路径失败！%v", err)
            return nil
        }

        pathEnv = append(pathEnv, cwd)
        ext := os.Getenv("PATHEXT")

        if len(ext) != 0 {
            pathExt = strs.Split(ext, colon)
        } else {
            pathExt[0] = ".EXE"
        }

        if strs.Index(cmd, ".") != -1 {
            pathExt = append([]string{""}, pathExt...)
        }
    }

    i, l := 0, len(pathEnv)

    for ; i <= l; i++ {
        if i == l {
            return cb(errors.New("not found:" + cmd))
        }

        var p = resolve(pathEnv[i], cmd)
        ii, ll := 0, len(pathExt)

        for ; ii < ll; ii++ {
            cur := p + strs.ToLower(pathExt[ii])
            fi, err := os.Stat(cur)
            if err == nil && !fi.IsDir() && isExe(cur) {
                return cb(cur)
            }
        }
    }
    return nil
}

// when cmd is in you path home
// will return the complete path
// in terms of os platform to loop the path
func WhichSync(cmd string) string {
    if isAbsolute(cmd) {
        return cmd
    }

    var (
        pathEnv = strs.Split(os.Getenv("PATH"), colon)
        pathExt = []string{""}
    )

    if isWin {

        cwd, err := os.Getwd()

        if err != nil {
            log.Fatalf("获取路径失败！%v", err)
            return ""
        }

        pathEnv = append(pathEnv, cwd)
        ext := os.Getenv("PATHEXT")

        if len(ext) != 0 {
            pathExt = strs.Split(ext, colon)
        } else {
            pathExt[0] = ".EXE"
        }

        if strs.Index(cmd, ".") != -1 {
            pathExt = append([]string{""}, pathExt...)
        }
    }

    for i, l := 0, len(pathEnv); i < l; i++ {
        var p = filepath.Join(pathEnv[i], cmd)

        for j, ll := 0, len(pathExt); j < ll; j++ {
            cur := p + strs.ToLower(pathExt[j])
            fi, err := os.Stat(cur)

            if err != nil {
                continue
            }

            if !fi.IsDir() && isExe(cur) {
                return cur
            }
        }
    }

    fmt.Fprintf(os.Stderr, "not found: %s\n", cmd)
    return ""
}

// absolute when platform is windows
func absWin(p string) bool {
    if absUnix(p) {
        return true
    }

    // match the regexp
    result := r.FindStringSubmatch(p)

    // pull off the device/UNC bit from a windows path.
    var (
        device     string
        isUnc      = false
        isAbsolute = false
    )

    if len(result[1]) != 0 {
        device = result[1]
        isUnc = device[1] != ':'
    }

    if len(result[2]) != 0 || isUnc {
        isAbsolute = true
    }

    return isAbsolute
}

// absolute when platform is other os
func absUnix(p string) bool {
    return p == "" || p[0] == '/'
}

func main() {
    r := Which("java.exe", func(argv ...interface{}) interface{} {
        switch len(argv) {
        case 1:
            switch argv[0].(type) {
            case string:
                fmt.Println(argv[0].(string))
            case error:
                fmt.Println(argv[0].(error))
            default:
            }
        default:
            fmt.Println(argv[1].(string))
        }
        return nil
    })
    fmt.Println(r)
}
