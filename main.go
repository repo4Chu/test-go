package main

import (
        "fmt"
        "io/ioutil"
)

func main() {
        fmt.Println("Lendo o arquivo /etc/passwd...")

        // Abre o arquivo /etc/passwd para leitura
        data, err := ioutil.ReadFile("/go/app.go")
        if err != nil {
                fmt.Println("Erro ao ler o arquivo:", err)
                return
        }

        // Converte o conteúdo para uma string e imprime
        fmt.Println(string(data))
}
