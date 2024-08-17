package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    for {
        ip, err := net.LookupIP("www.baidu.com")
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        f, err := os.OpenFile("record.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        defer f.Close()
        _, err = f.WriteString(fmt.Sprintf("%s: %s\n", time.Now().Format(time.RFC3339), ip[0].String()))
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        time.Sleep(10 * time.Second)
    }
}
