package main

import "fmt"

func average(ar []float64) float64 {
//    panic("Not Implemented")

    total := 0.0
    for _, v := range ar {
        total += v
    }
    return total / float64(len(ar))
}

func return_bool(x int) (r bool) {
    r = x != 0
    return
}

func return_two_vars(a int, b int) (int, int) {
    return a, b
}

func many_args(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    fmt.Println("Sum args =", total)
    return total
}

func lambda_function()  {
    fmt.Println("Lambda func");
    x := 0
    increment := func() int {
        x++
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
}


func evenGenerator()  {
    fmt.Println("Generator");
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}

func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}


func recursion()  {
    fmt.Println("Recursion")
    fmt.Println(factorial(5))
}

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}


func defer_function()  {

    fmt.Println("Defer")
    defer second()
    first()
}

func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2nd")
}


func panic_handler()  {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
    fmt.Println(return_bool(0))
    fmt.Println(return_two_vars(1,2))
    x, y := return_two_vars(3,4)
    fmt.Println(x,y)
    many_args(1,2)
    lambda_function()
    evenGenerator()
    recursion()
    defer_function()
    panic_handler()
}

