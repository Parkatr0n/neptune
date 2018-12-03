
package hyper

import (
    _ "fmt"
    "net"

    _ "reflect"

    "strconv"
    "strings"

    "bytes"
)

const (
    port = "8080"
)

func Send(data []byte, location string) {
    conn, _ := net.Dial("udp", location + ":" + port)
    defer conn.Close()

    // Because they aren't sending a packet, we have to (as hyper) make sure that
    //  when we receive it (as hyper) that we know it's not a packet
    // We use "notp" to designate that what we are sending is NOT a Packet.
    // If we were to send a packet, we would use "pack" to designate that all of
    //  the data following is unpackable.
    data = append([]byte("notp"), data...)
    conn.Write(data)

}

func ReceivePacket() Packet {
    pc, _ := net.ListenPacket("udp", "localhost:" + port)

    buffer := make([]byte, 1024)
    pc.ReadFrom(buffer)

    // Alright, now that we have the data, we need to turn it into a Packet
    packet := NewPacket()

    // Split the packet by the escape character
    escape := "/"
    data := string(bytes.Trim(buffer, "\x00")) // convert the []byte buffer to a string

    split := strings.Split(data, escape)

    for i, val := range split {
        if i == 0 { continue }
        valsplit := strings.Split(val, ":")
        typeof := valsplit[0]
        name := valsplit[1]

        switch (typeof) {
        case "int":
            num, _ := strconv.Atoi(valsplit[2])
            packet.Pack(name, num)
        case "string":
            packet.Pack(name, valsplit[2])
        default:
            packet.Pack(name, "ERR" + valsplit[2])
        }
    }

    pc.Close()

    return packet
}

func SendPacket(pack Packet, location string) {
    data := pack.Data
    datastr := "pack"

    escape := "/"

    // Convert all of the Packet's data into a byte slice to send
    for key, value := range data {
        typeof := ""

        topack := ""

        switch v := value.(type) {
        case int:
            topack = strconv.Itoa(v)
            typeof = "int"
            break
        case string:
            topack = v
            typeof = "string"
            break
        }

        datastr += escape + typeof + ":" + key + ":" + topack
    }

    rawdata := []byte(datastr)

    conn, _ := net.Dial("udp", location + ":" + port)
    conn.Write(rawdata)

    conn.Close()
}

func WaitFor() []byte {
    pc, _ := net.ListenPacket("udp", "localhost:" + port)

    defer pc.Close()

    // Read
    buffer := make([]byte, 1024)
    pc.ReadFrom(buffer)

    return buffer
}
