package main

import (
    "bytes"
    "log"
    "sync"
)

type Waiter struct {
    expected      []byte
    expectedMutex sync.Mutex
    actual        []byte
    done          bool
    cond          *sync.Cond
}

func (w *Waiter) Write(p []byte) (int, error) {
    w.expectedMutex.Lock()
    defer w.expectedMutex.Unlock()
    w.actual = append(w.actual, p...)
    if bytes.Contains(w.actual, w.expected) {
        w.cond.L.Lock()
        defer w.cond.L.Unlock()
        w.done = true
        w.cond.Broadcast()
    }
    return len(p), nil
}

func (w *Waiter) Wait() {
    w.cond.L.Lock()
    defer w.cond.L.Unlock()
    for !w.done {
        w.cond.Wait()
    }
}

func New(expected []byte) *Waiter {
    return &Waiter{
        expected: expected,
        cond:     sync.NewCond(&sync.Mutex{}),
    }
}

func main() {
    const count = 10
    expected := []byte("42")
    w := New(expected)

    event := make(chan bool)
    for i := 0; i < count; i++ {
        go func() {
            if r := recover(); r != nil {
                log.Fatal(r)
            }
            w.Write([]byte("4"))
            w.Write([]byte("2"))
            event <- true
        }()
    }

    for i := 0; i < count; i++ {
        go func() {
            if r := recover(); r != nil {
                log.Fatal(r)
            }
            w.Wait()
            event <- true
        }()
    }

    done := make(chan struct{})
    go func() {
        defer close(done)
        for i := 0; i < count*2; i++ {
            <-event
        }
    }()

    <-done
    w.Wait()
}
