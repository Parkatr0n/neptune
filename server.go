
package main

import (
    "fmt"

    "./hyper"

    _ "reflect"

    _ "bytes"
)

func main() {
    fmt.Println("Starting server. Waiting for messages")

    // Constantly listen for messages
    for {
        pack := hyper.ReceivePacket()

        fmt.Println(pack.Unpack("value"))
        fmt.Println(pack.Unpack("thing"))
    }
}
