// kill process
// refer to douban python CaoE
package main

import (
    "fmt"
    "syscall"
    "os"
    "os/signal"
)

const (
    SYS_FORK = 57
)

func main() {
    pid, _, sysErr := syscall.RawSyscall(SYS_FORK, 0, 0, 0)
    if sysErr != 0 {
        panic(sysErr.Error())
    }
    fmt.Println(pid)
    //install(true, syscall.SIGTERM)
}

type signalHandler func(gid int, sig os.Signal) error

type signalSet struct {
    m map[os.Signal]signalHandler
}

func signalSetNew() *signalSet {
    return &signalSet{
        m: make(map[os.Signal]signalHandler),
    }
}

func (s *signalSet) register(sig os.Signal, handler signalHandler) {
    if _, ok := s.m[sig]; !ok {
        s.m[sig] = handler
    }
}

func install(fork bool, sig syscall.Signal) {
    var reg = func(gid int, sig os.Signal) {
        var handler signalHandler
        handler = makeQuitSignalHandler(gid, sig)
        ss := signalSetNew()
        ss.register(syscall.SIGINT, handler)
        ss.register(syscall.SIGQUIT, handler)
        ss.register(syscall.Signal(0xf), handler)
        ss.register(syscall.SIGCHLD, makeChildDieSignalHandler(gid, sig))

        for {
            c := make(chan os.Signal, 1)
            var sigs []os.Signal
            for sig := range ss.m {
                sigs = append(sigs, sig)
            }
            signal.Notify(c, sigs...)
            sig := <- c
            for _, h := range ss.m {
                err := h(gid, sig)
                if err != nil {
                    fmt.Printf("unknown signal received: %v\n", sig)
                    os.Exit(1)
                }
            }
        }
    }
    if !fork {
        reg(os.Getpid(), sig)
        return
    }
    pid, _, sysErr := syscall.RawSyscall(SYS_FORK, 0, 0, 0)
    if sysErr != 0 {
        panic(sysErr.Error())
    }
    if pid == 0 {
        //syscall.Setpgid(0, 0)
    } else {
        gid := pid
        reg(int(gid), sig)
        for {

        }
    }
}

func makeQuitSignalHandler(gid int, sig os.Signal) signalHandler {
    return func(signum int, frame os.Signal) error {
        return nil
    }
}

func makeChildDieSignalHandler(gid int, sig os.Signal) signalHandler {
    return func(signum int, frame os.Signal) error {

        // defer os.Exit((status & 0xff00) >> 8)
        return nil
    }
}

func exitWhenParentOrChildDies(sig os.Signal) error {
    pid := os.Getppid()
    fmt.Println(pid)
    return nil
}
