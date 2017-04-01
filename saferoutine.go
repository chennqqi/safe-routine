package safe_routine

import (
    "fmt"
    "runtime"
    "io"
    "os"
)

var RecoverReporter io.Writer = os.Stdout

type mustPanic struct {
    Err interface{}
}

func Panic(v interface{}) {
    panic(mustPanic{Err:v})
}

func catchPanic() {
    if r := recover(); r != nil {
        p, ok := r.(mustPanic)
        if ok {
            panic(p.Err)
        } else {
            buf := make([]byte, 10000)
            runtime.Stack(buf, false)
            detail := fmt.Sprintf("Recoverd from safe routine %v\n%s", r, string(buf))
            RecoverReporter.Write([]byte(detail))
        }
    }
}

type Routine func()

type RoutineExt func(...interface{})

func NewRoutine(rt RoutineExt) RoutineExt {
    return func(args ...interface{}) {
        defer catchPanic()
        rt(args)
    }
}

func New(rt Routine) {
    defer catchPanic()
    rt()
}