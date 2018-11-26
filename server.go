
package main

import (
    "fmt"

    "./hyper"

    _ "reflect"

    "bytes"
)

func main() {
    fmt.Println("Starting server. Waiting for messages")

    // Constantly listen for messages
    for {
        data := hyper.WaitFor()

        pack := hyper.Decode(data)
        info := pack.Unpack()

        usn := info[0].(string)
        psk := string(bytes.Trim([]byte(info[1].(string)), "\x00"))

        result := "No"

        sendp := hyper.Packet{}

        fmt.Println([]byte(psk))
        fmt.Println([]byte(usn))

        if (psk == "hunter2") {
            // Logon successful
            result = "Yes"
        }

        sendp.PackString(result)

        // Send the result
        hyper.SendPacket(sendp, "localhost")

        /*
        data := hyper.WaitFor()

        fmt.Println(string(data))

        pack := hyper.Decode(data)

        fmt.Println(string(pack.GetData()))

        info := pack.Unpack()

        for i := range info {
            fmt.Print(info[i])
            fmt.Print("   ")
            fmt.Println(reflect.TypeOf(info[i]))
        }
        */
    }
}
