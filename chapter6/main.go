package main

import "fmt"

func main() {
//    array_process()
//    slice_process()
    map_process()
}

func array_process()  {

//    var x [5]float64
//    x[0] = 98

    x := [4]float64{
        98,
        93,
        77,
        82,
        // 83,
    }


    var total float64 = 0
/*    for i := 0; i < len(x); i++ {
        total += x[i]
    }*/

    for _, value := range x{
        total += value
    }
    fmt.Println(total / float64(len(x)))
}

func slice_process()  {

    arr := [5]float64{1,2,3,4,5}
    x := arr[1:4]
    fmt.Println(x)

    copy1 := []int{1,2,3}
    copy2 := make([]int, 2)
    copy(copy2, copy1)
    fmt.Println(copy1, copy2)

    append1 := []int{1,2,3}
    append2 := append(append1, 4, 5)
    fmt.Println(append1, append2)
}

func map_process()  {
    elements := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }

    if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }

    name, ok := elements["Un"]
    fmt.Println(name, ok)

    fmt.Println(elements["Li"])
    fmt.Println(elements["Un"])
    fmt.Println(elements)


}