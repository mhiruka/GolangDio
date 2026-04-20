package main

import "fmt"

func main() {
    var kelvin float64
    var celsius float64

    // Entrada do usuário
    fmt.Print("Digite a temperatura em Kelvin: ")
    fmt.Scanln(&kelvin)

    // Conversão
    celsius = kelvin - 273

    // Saída
    fmt.Printf("Temperatura em Celsius: %.2f°C\n", celsius)
}