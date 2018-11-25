
package main

import (
    "fmt"

    "./hyper"
)

func main() {
    fmt.Println("Starting server. Waiting for messages")

    // Constantly listen for messages
    for {
        data := hyper.WaitFor()

        fmt.Println(string(data))

        pack := hyper.Decode(data)

        fmt.Println(string(pack.GetData()))

        i, _  := pack.UnpackInt()
        i2, _ := pack.UnpackInt()

        fmt.Println(i)
        fmt.Println(i2)
    }
}
