
package hyper

/* Packet */
// Hyper's custom packet class
// Packets allow data packing into a single <[]byte> which can later be sent off

// In comments, a % is used to represent the escape character(s)
// Data information:
// %s - string
// %i - int
// %B - []byte

// The escape string is used to inform the receiver that there is data to be
// unpacked after it, with the immediate character determining what the data
// type of the information is.
var escape string = "[*^}"

// var escapeb []byte{0, 32, 41, 255, 210}

// Here is the declaration of all of the characters mentioned above.
var (
    string_char = "s"
    bytes_char = "B"
    bool_char = "b"
    int_char = "i"
)

type Packet struct {
    Data map[string]interface{}
}

func NewPacket() Packet {
    pack := Packet{}

    pack.Data = make(map[string]interface{})

    return pack
}

func (this *Packet) Pack(name string, data interface{}) Packet {
    this.Data[name] = data

    return *this
}

func (this *Packet) Unpack(name string) interface{} {
    return this.Data[name]
}
