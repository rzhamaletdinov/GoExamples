package main

import "fmt"

// this is a comment

const Q_NAME string = "My Fucking Constant"

var y = "Global var"

func main() {

    x := "Hellw world"
    fmt.Println(x)
    fmt.Println(y)
    f()
    show_number()
    show_list()
}

func f() {
    fmt.Println("Function F")
}

func show_number(){
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println(output)
}

func show_list()  {

    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }
}