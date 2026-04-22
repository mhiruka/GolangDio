package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        resultado := ""

        if i%3 == 0 {
            resultado += "Pin"
        }
        if i%5 == 0 {
            resultado += "Pan"
        }

        if resultado == "" {
            fmt.Println(i)
        } else {
            fmt.Println(resultado)
        }
    }
}