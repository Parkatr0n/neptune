
package main

import (
    "fmt"

    "./hyper"
)

func main() {
    usn := "Park"
    psk := "hunter2"

    pack := hyper.Packet{}

    pack.PackString(usn)
    pack.PackString(psk)

    // Send it
    hyper.SendPacket(pack, "localhost")

    // Wait for information to come back
    data := hyper.WaitFor()

    rpack := hyper.Decode(data)
    info := rpack.Unpack()

    answer := info[0]

    // Answer will be "Yes" or "No" to whether we can logon
    fmt.Println(answer)
}
