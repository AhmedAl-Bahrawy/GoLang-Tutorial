
package main

import "fmt"

func getFunc(str string) func(string) func(string) string {
    return func(str2 string) func(string) string {
        return func(str3 string) string {
            return str + " " + str2 + " " + str3
        }
    }
}

func main() {
    f1 := getFunc("Hello")
    f2 := f1("World")
    value := f2("Ahmed")

    fmt.Println(value)
}
