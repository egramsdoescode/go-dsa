package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
    str1 := encodeNaive([]string{"we", "say",":", "yes"})
    fmt.Println(decodeNaive(str1))
    str2 := encodeEfficient([]string{"we", "say",":", "yes"})
    fmt.Println(decodeEfficient(str2))
}

// O(n)
func encodeEfficient(strs []string) string {
    var sb strings.Builder

    for _, str := range strs {
        sb.WriteString(strconv.Itoa(len(str)))
        sb.WriteString("#")
        sb.WriteString(str)
    }

    return sb.String() 
}

// O(n)
func decodeEfficient(str string) []string {
    result := make([]string, 0)
    i := 0

    for i < len(str) {
        j := i
        for str[j] != '#' {
            j++
        }

        length, _ := strconv.Atoi(str[i:j])
        i = j + 1

        result = append(result, str[i:i + length])
        i += length
    }
    return result 
}

// O(n^2)
func encodeNaive(strs []string) string {
    result := ""
    for _, str := range strs {
        result += str + "#"
    } 
    return result 
}

// O(n)
func decodeNaive(str string) []string {
    result := make([]string, 0)
    word := ""
    for _, char := range str {
        if char == '#' {
            result = append(result, word)
            word = ""
            continue
        }
        word += string(char)
    } 
    return result 
}
