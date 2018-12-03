
package main

import (
    "fmt"

    "./hyper"
)

func main() {
    fmt.Println("starting")

    for {
        packet := hyper.ReceivePacket()

        usn := packet.Unpack("usn")
        psk := packet.Unpack("psk")

        fmt.Println(usn, psk)

        answer := "No"

        if usn == "Park" && psk == "hunter2" {
            answer = "Yes"
        }

        packet.Pack("answer", answer)

        // send the answer
        hyper.SendPacket(packet, "localhost")
    }
}
