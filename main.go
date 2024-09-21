package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("ls;id;pwd;env")                                                                                                                   
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando id:", err)
        } else {
            fmt.Println("Saída do comando id:", string(output))
        }
}
