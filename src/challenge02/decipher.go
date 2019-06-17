package main

import (
    "fmt"
)

func main() {
    cipherText := "CSOITEUIWUIZNSROCNKFD"
    keyword := "GOLANG"

    for i, c := range cipherText {
        key := rune(keyword[i % len(keyword)])
        c -= key - 'A'
        if c < 'A' {
            c += 26
        }
        fmt.Printf("%c", c)
    }
    fmt.Println()
}
