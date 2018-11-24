
package hyper

import (
    "net"
)

const (
    port = "8080"
)

func Send(data []byte, location string) {
    conn, _ := net.Dial("udp", location + port)
    defer conn.Close()

    buffer := make([]byte, 1024)

    conn.Write(data)
}

func SendPacket(pack *Packet, location string) {
    Send(pack.data, location)
}

func WaitFor() []byte {
    pc, _ := net.ListenPacket("udp", "localhost:" + port)

    defer pc.Close()

    // Read
    buffer := make([]byte, 1024)
    pc.ReadFrom(buffer)

    return buffer
}
