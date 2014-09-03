// create temporary director
// refer to nodejs temp
package main

import (
    "errors"
    "math/rand"
    "os"
    "fmt"
    "os/exec"
    "path/filepath"
    "regexp"
    "runtime"
    "strconv"
    "strings"
    "time"
)

var (
    removeObjects     = []func(){}
    TmpDir            = os.TempDir()
    gracefulCleanup   = false
)

type removeOrCreateTmp func(name string) error

type Tmp struct {
    Mode          os.FileMode
    Prefix        string
    Postfix       string
    Template      string
    Dir           string
    Tries         int
    Keep          bool
    UnsafeCleanup bool
}

func init() {
    if TmpDir == "" {
        TmpDir = getTMPDir()
    }
}

// create temporary name
func (tmp *Tmp) GetTmpName() (error, string) {

    var (
        randomChars        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz"
        randomCharsLength  = len(randomChars)
        errInvalidTemplate = errors.New("Invalid template provided")
        errNotUnique       = errors.New("could not get a unique tmp filename, max tries reached")
    )

    templateDefined := false
    template := tmp.Template

    if template != "" {
        templateDefined = true
    }

    var tries int
    if tmp.Tries == 0 {
        tries = 3
    }

    if matched, _ := regexp.MatchString("XXXXXX", template); templateDefined && !matched {
        return errInvalidTemplate, ""
    }

    var getName = func() string {
        rand.Seed(time.Now().UnixNano())
        
        if !templateDefined {

            var name = ""

            if prefix := tmp.Prefix; prefix == "" {
                name += "tmp-"
            } else {
                name += prefix
            }
            name += strconv.Itoa(os.Getpid())
            randHex := uint64(rand.Float64() * 0x1000000000)
            name += strconv.FormatUint(randHex, 36)
            name += tmp.Postfix

            var tmpDir string
            if tmpDir = tmp.Dir; tmpDir == "" {
                tmpDir = TmpDir
            }

            return createDirName(filepath.Join(tmpDir, name))
        }

        var chars = []byte{}
        
        for i := 0; i < 6; i++ {
            chars = append(chars, randomChars[rand.Intn(randomCharsLength)])
        }

        return createDirName(strings.Replace(template, "XXXXXX", string(chars), -1))
    }

    var name string
    for {
        name = getName()
        if isExist(name) {

            if tries -= 1; tries > 0 {
                continue
            }
            return errNotUnique, ""
        }
        break
    }

    return nil, name
}

// create temporary file
// pass callbak function
func (tmp *Tmp) CreateTmpFile() (err error, name string, fd *os.File, removeCallBack func()) {

    if tmp.Postfix == "" {
        tmp.Postfix = ".tmp"
    }

    err, name = tmp.GetTmpName()

    var mode os.FileMode = tmp.Mode

    if uint32(mode) == 0 {
        mode = 0600
    }

    flag := os.O_CREATE | os.O_EXCL | os.O_RDWR

    var dirName = string(name[:strings.LastIndex(name, string(filepath.Separator))])

    _, err = exec.LookPath(dirName)

    if err != nil {
        er := os.MkdirAll(dirName, mode)

        if er != nil {
            return
        }
    }

    fd, err = os.OpenFile(name, flag, mode)

    if err != nil {
        return
    }

    removeCallBack = prepareRemoveCallback(removeFile, name)

    if !tmp.Keep {
        removeObjects = append([]func(){removeCallBack}, removeObjects...)
    }

    defer fd.Close()

    return
}

// create temporary directory
// return multiple value
func (tmp *Tmp) CreateTmpDir() (err error, name string, removeCallBack func()) {
    err, name = tmp.GetTmpName()

    var mode os.FileMode = tmp.Mode
    if uint32(mode) == 0 {
        mode = 0700
    }

    err = os.MkdirAll(name, mode)

    if err != nil {
        return
    }

    var removeFunc removeOrCreateTmp = removeDirectory

    if tmp.UnsafeCleanup {
        removeFunc = os.RemoveAll
    }

    removeCallBack = prepareRemoveCallback(removeFunc, name)

    if !tmp.Keep {
        removeObjects = append([]func(){removeCallBack}, removeObjects...)
    }

    return
}

// set clean up flag
func (tmp *Tmp) SetGracefulCleanup() {
    gracefulCleanup = true
}

// gc
func (tmp *Tmp) GarbageCollector() {
    if !gracefulCleanup {
        return
    }

    for _, v := range removeObjects {
        v()
    }
}

// create dir name in case of platform is windows
// according to the os return the absolute path
func createDirName(path string) string {
    if runtime.GOOS == "windows" {
        if !filepath.IsAbs(path) {
            wd, _ := os.Getwd()
            path = filepath.VolumeName(wd) + `/` + path
        }
    }

    return filepath.FromSlash(path)
}

// check the path is exist or not
func isExist(path string) bool {
    _, err := os.Stat(path)
    return err == nil || os.IsExist(err)
}

// delete directory
func removeDirectory(path string) error {
    return removeFileOrDirectory(path, true)
}

// delete file
func removeFile(path string) error {
    return removeFileOrDirectory(path, false)
}

// common operator
func removeFileOrDirectory(path string, flag bool) error {
    fd, err := os.Stat(path)

    if err != nil {
        return err
    }

    if fd.IsDir() {
        if flag {
            os.Remove(path)
        }
    }

    if !flag {
        os.Remove(path)
    }

    return nil
}

// set closure
func prepareRemoveCallback(removeFunction removeOrCreateTmp, path string) func() {
    called := false
    return func() {
        if called {
            return
        }
        removeFunction(path)
        called = true
    }
}

// get temporary directory when os.TempDir() return ""
func getTMPDir() string {
    tempNames := []string{"TMPDIR", "TMP", "TEMP"}

    for _, v := range tempNames {
        if env := os.Getenv(v); env != "" {
            return env
        }
    }

    return "/tmp"
}

func main() {
    var p = &Tmp{
        Mode:   0644,
        Prefix: "myTmpDir_",
        // prefix: "prefix-",
        Template: "/tmp/tmp-XXXXXX/good.txt",
        Postfix:  ".txt",
    }

    err, tmpName, _, _ := p.CreateTmpFile()
    fmt.Println("return tmpname: ", tmpName)

    if err != nil {
        panic(err)
    }
    // opt, cb := parseArguments(p, r)
    p.SetGracefulCleanup()
    // cb("abd", 1235)
    // fmt.Println(opt, cb)

    time.Sleep(time.Second * 6)
    p.GarbageCollector()
}
