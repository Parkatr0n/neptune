
package main

import (
    "fmt"

    "bufio"
    "os"

    "strings"

    "github.com/fatih/color"
)

var (
    c_white  = color.New(color.FgWhite)
    c_blue   = color.New(color.FgHiBlue)
    c_green  = color.New(color.FgHiGreen, color.Bold)
    c_red    = color.New(color.FgRed)
    c_yellow = color.New(color.FgYellow)
)

func main() {
    fmt.Println("Starting")

    // Initialization

    // Start networking stuff
    c_white.Print("Hello! Please enter the computer which you want to login to: ")
    box := input()
    c_white.Print("Which user? ")
    user := input()

    // Bash shell
    c_green.Print(user + "@" + box)
    c_white.Print(":")
    c_blue.Print("/")
    c_white.Print("$ ")

    cmd := input()

    fmt.Println(cmd)
}

func parseCommand(s string) {
    split := strings.Fields(s)

    program := split[0]

    args := split[1:]
    fmt.Println(program, args)
}

func input() string {
    reader := bufio.NewReader(os.Stdin)

    string, _ := reader.ReadString('\n')

    return strings.Replace(string, "\n", "", -1)
}
