package main

import (
        "os/exec"
)

func main() {
        cmd := exec.Command("find", "/","-name","flag.txt")
        output, err := cmd.Output()
        if err != nil {
            fmt.Println("Erro:", err)
        } else {
            fmt.Println("output:", string(output))
        }
}
