
package main

import (
    "fmt"
    "./hyper"
)

func main() {
    // Declare our info
    usn := "Park"
    psk := "hunter2"

    packet := hyper.NewPacket()

    packet.Pack("usn", usn)
    packet.Pack("psk", psk)

    hyper.SendPacket(packet, "localhost")

    // Receive the packet
    p := hyper.ReceivePacket()

    ans := p.Unpack("answer")
    fmt.Println(ans)
}
