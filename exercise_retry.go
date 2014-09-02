package main

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "sort"
    "time"
)

type RetryOperation struct {
    timeouts           []int
    fn                 func(attempts int)
    errors             []error
    attempts           int
    operationTimeout   int
    operationTimeoutCb func(attempts ...int)
}

type operation struct {
    retries    int
    factor     int
    minTimeout float64
    maxTimeout float64
    randomize  bool
    *RetryOperation
}

func New() *operation {
    return &operation{
        retries:    10,
        factor:     2,
        minTimeout: 1 * 1000,
        maxTimeout: math.Inf(0),
        RetryOperation: &RetryOperation{
            operationTimeoutCb: nil,
        },
    }
}

// count timeouts
func timeouts() (timeouts []int, err error) {
    o := New()
    if o.minTimeout > o.maxTimeout {
        err = errors.New("minTimeout is greater than maxTimeout")
        return
    }
    for i := 0; i < o.retries; i++ {
        timeouts = append(timeouts, o.createTimeout(i))
    }
    sort.Ints(timeouts)
    return
}

// according opts count the timeout
func (o *operation) createTimeout(attempt int) int {
    var (
        random  float64
        timeout float64
    )
    rand.Seed(time.Now().UnixNano())
    if o.randomize {
        random = rand.Float64() + 1
    } else {
        random = 1
    }
    timeout = math.Floor(random*o.minTimeout*math.Pow(float64(o.factor), float64(attempt)) + 0.5)
    timeout = math.Min(timeout, o.maxTimeout)
    return int(timeout)
}

func (r *RetryOperation) retry(err error) bool {
    if err != nil {
        return false
    }
    if r.errors == nil {
        r.errors = []error{}
    }
    r.errors = append(r.errors, err)
    var timeout int
    if len(r.timeouts) > 0 {
        timeout = r.timeouts[0]
        r.timeouts = r.timeouts[1:]
    } else {
        return false
    }

    r.attempts++
    timer1 := time.NewTimer(time.Second * time.Duration(timeout))
    <-timer1.C
    r.fn(r.attempts)
    if r.operationTimeoutCb != nil {
        fmt.Println("===")
        timer2 := time.NewTimer(time.Second * time.Duration(r.operationTimeout))
        <-timer2.C
        r.operationTimeoutCb(r.attempts)
    }
    return true
}

func (r *RetryOperation) attempt(fn func(int), args ...interface{}) {
    r.fn = fn
    for i, v := range args {
        if i == 0 {
            r.operationTimeout = v.(int)
        }
        if i == 1 {
            r.operationTimeoutCb = v.(func(...int))
        }
    }
    r.fn(r.attempts)

    if r.operationTimeoutCb != nil {
        timer1 := time.NewTimer(time.Second * time.Duration(r.operationTimeout))
        <-timer1.C
        r.operationTimeoutCb()
    }
}

func (r *RetryOperation) try(fn func(int)) {
    fmt.Println("Using RetryOperation.try() is deprecated")
    r.attempt(fn)
}

func (r *RetryOperation) start(fn func(int)) {
    fmt.Println("sing RetryOperation.start() is deprecated")
    r.attempt(fn)
}

// find main error
func (r *RetryOperation) mainError() error {
    lenErrors := len(r.errors)
    if lenErrors == 0 {
        return nil
    }
    var (
        mainError      error
        mainErrorCount = 0
        counts         = map[error]int{}
    )
    for _, err := range r.errors {
        c, ok := counts[err]
        count := 1
        if ok {
            count = c + 1
        }
        counts[err] = count
        if count >= mainErrorCount {
            mainError = err
            mainErrorCount = count
        }
    }
    return mainError
}

func main() {
    operation := &operation{
        retries:    5,
        factor:     3,
        minTimeout: 1 * 1000,
        maxTimeout: 60 * 1000,
        randomize:  true,
    }
    fmt.Println(operation)

    // test timeout  begin //
    // timeouts, _ := timeouts()
    // fmt.Println(len(timeouts), timeouts[0], timeouts[1], timeouts[2])
    // test timeout  end //

    retry := New()
    err := errors.New("some error")
    err2 := errors.New("some other error")
    retry.errors = append([]error{}, err)
    retry.errors = append(retry.errors, err2)
    retry.errors = append(retry.errors, err)
    fmt.Println(retry.mainError())
    retry.attempt(func(int) {}, 1, func(...int) {})
    fmt.Println(retry.operationTimeout)
}
