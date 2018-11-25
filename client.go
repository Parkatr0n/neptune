
package main

import (
    "fmt"

    "./hyper"
)

func main() {
    fmt.Println("Starting: sending 'hello world' (in a packet)")

    pack := hyper.Packet{}

    pack.PackInt(35)
    pack.PackInt(64)

    // Send it
    location := "localhost"
    hyper.SendPacket(pack, location)
}
