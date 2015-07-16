// simple interpreter
// see more detail on oschina.net
// also you can read this article http://ruslanspivak.com/lsbasi-part1/

package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "unicode"
)

type Elem string

// 定义加号以及整数和结束符
const (
    INTEGER Elem = "INTEGER"
    PLUS    Elem = "PLUS"
    EOF     Elem = "EOF"
)

type Token struct {
    kind  Elem
    value interface{}
}

func (t *Token) String() string {
    return fmt.Sprintf("Token(%s, %v)", t.kind, t.value)
}

type Interpreter struct {
    text          string
    pos           int
    current_token *Token
}

func (i *Interpreter) Error() {
    panic("Error parsing input")
}

func (i *Interpreter) GetNextToken() (*Token, error) {
    text := i.text
    if i.pos > len(text)-1 {
        return &Token{
            kind:  EOF,
            value: nil,
        }, nil
    }
    current_char := text[i.pos]
    if unicode.IsDigit(rune(current_char)) {
        v, err := strconv.Atoi(string(current_char))
        if err != nil {
            return nil, err
        }
        token := &Token{
            kind:  INTEGER,
            value: v,
        }
        i.pos += 1
        return token, nil
    }

    if current_char == '+' {
        token := &Token{
            kind:  PLUS,
            value: current_char,
        }
        i.pos += 1
        return token, nil
    }
    i.Error()
    return nil, nil
}

func (i *Interpreter) Eat(kind Elem) error {
    if i.current_token.kind == kind {
        current_token, err := i.GetNextToken()
        if err != nil {
            return err
        }
        i.current_token = current_token
    }
    return nil
}

func (i *Interpreter) Expr() (int, error) {
    current_token, err := i.GetNextToken()
    if err != nil {
        return 0, err
    }
    i.current_token = current_token
    left := current_token
    i.Eat(INTEGER)

    // _ = i.current_token
    i.Eat(PLUS)

    right := i.current_token
    i.Eat(INTEGER)

    result := 0
    if leftValue, ok := left.value.(int); ok {
        result += leftValue
    }

    if rightValue, ok := right.value.(int); ok {
        result += rightValue
    }

    return result, nil
}

func main() {
    b := make([]byte, 100)
    f := os.Stdin
    w := os.Stdout
    defer f.Close()
    defer w.Close()

    for {
        w.WriteString("Input> ")
        c, _ := f.Read(b)
        bb := b[:c-2]
        if len(bb) == 0 {
            continue
        }
        if strings.Compare(string(bb), "exit") == 0 {
            break
        }
        i := &Interpreter{
            text:          string(bb),
            pos:           0,
            current_token: nil,
        }
        r, err := i.Expr()
        if err != nil {
            panic(err)
        }
        fmt.Println(r)
    }
}
