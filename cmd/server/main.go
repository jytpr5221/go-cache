package main

import (
	"bufio"
	"fmt"
    "strconv"
	"os"
	"predictive-cache/internal/cache"
	"strings"
)

func main() {


    fmt.Println("Welcome to Cache Server.")
    fmt.Println("Enter your eviction policy")

    policyReader := bufio.NewReader(os.Stdin)
    policy, _ := policyReader.ReadString('\n')
    intPolicy, _ := strconv.Atoi(strings.TrimSpace(policy))

    c := cache.NewCache(cache.EvictionPolicy(intPolicy))
    for {
        fmt.Println("Select your option: ")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))

        switch input {
        case "get":
            fmt.Println("Insert your key:")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)

            value, ok := c.Get(key)
            if ok {
                fmt.Println(value)
            } else {
                fmt.Println("not found")
            }

        case "set":
            fmt.Println("Insert your key:")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)

            fmt.Println("Insert your value:")
            valReader := bufio.NewReader(os.Stdin)
            val, _ := valReader.ReadString('\n')
            val = strings.TrimSpace(val)

            c.Set(key, val, 60)

        case "del":
            fmt.Println("Insert your key:")
            keyReader := bufio.NewReader(os.Stdin)
            key, _ := keyReader.ReadString('\n')
            key = strings.TrimSpace(key)

            c.Del(key)

        case "exit":
            return

        default:
            fmt.Println("valid commands: get, set, del, exit")
        }
    }
}