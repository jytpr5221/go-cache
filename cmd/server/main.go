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
	cacheServer := cache.NewCache()

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
            exists, value := cacheServer.Get(key)
			if exists{
				fmt.Println(value)
			}else{
				fmt.Println("No such key found!")
			}

        case "set":
            fmt.Println("Insert your key")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)

            fmt.Println("Insert your value")
            valReader := bufio.NewReader(os.Stdin)
            val, _ := valReader.ReadString('\n')
            val = strings.TrimSpace(val)

            cacheServer.Set(key, val, 10)

        case "del":
            fmt.Println("Insert your key")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)
            cacheServer.Del(key)

        default:
            return
        }
    }
}