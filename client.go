
package main

import (
    _ "fmt"

    "./hyper"
)

func main() {
    pack := hyper.NewPacket()

    pack.Pack("value", 80)
    pack.Pack("thing", "teststr")

    hyper.SendPacket(pack, "localhost")
}
