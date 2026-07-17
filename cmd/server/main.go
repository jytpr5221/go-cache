package main

import (
    "bufio"
    "fmt"
    "os"
    "predictive-cache/internal/cache"
    "strings"
)

func main() {
    fmt.Println("Welcome to Cache Server.")

    for {
        fmt.Println("Select ur Option: ")

        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "get":
            fmt.Println("Insert your key")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)
            cache.GET(key)

        case "set":
            fmt.Println("Insert your key")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)

            fmt.Println("Insert your value")
            valReader := bufio.NewReader(os.Stdin)
            val, _ := valReader.ReadString('\n')
            val = strings.TrimSpace(val)

            cache.SET(key, val)

        case "del":
            fmt.Println("Insert your key")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)
            cache.DEL(key)

        default:
            return
        }
    }
}