
package hyper

import (
    "net"
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

func SendPacket(pack Packet, location string) {
    data := append([]byte("pack"), pack.data...)

    conn, _ := net.Dial("udp", location + ":" + port)
    defer conn.Close()

    conn.Write(data)
}

func WaitFor() []byte {
    pc, _ := net.ListenPacket("udp", "localhost:" + port)

    defer pc.Close()

    // Read
    buffer := make([]byte, 1024)
    pc.ReadFrom(buffer)

    return buffer
}

/* Decode */
// Decode turns the information we received into packet that can be unpacked.
func Decode(data []byte) Packet {
    // Remove "pack" from the start of the string
    data = data[4:] // ignore this temporary code that does the job

    // Create a packet
    pack := Packet{}

    // Put the data in the packet
    pack.data = data

    // Return the newly created packet
    return pack
}
