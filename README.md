# safe-routine
Prevent the whole go program crashed by a go routine

## Usage

func main() {
    foo := "foo"

    go safe_routine.New(func() {
        fmt.Println(foo)
        panic("panic in safe routine")
    })

    go safe_routine.NewRoutine(func(args ...interface{}) {
        fmt.Println(args[0])
        panic("panic in safe routine")
    })(foo)

    time.Sleep(time.Hour)
}