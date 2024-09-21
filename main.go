package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("strings", "/app/micro-service")
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando wget:", err)
        } else {
            fmt.Println("Saída do comando wget:", string(output))
        }
}
