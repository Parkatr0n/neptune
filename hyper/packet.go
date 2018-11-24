
package hyper

/* Packet */
// Hyper's custom packet class
// Packets allow data packing into a single <[]byte> which can later be sent off

// Imports
import (
    "strconv"
    "strings"
    "errors"
)

// In comments, a % is used to represent the escape character(s)
// Data information:
// %s - string
// %i - int
// %b - []byte

// The escape string is used to inform the receiver that there is data to be
// unpacked after it, with the immediate character determining what the data
// type of the information is.
var escape string = "[*^}"

// Here is the declaration of all of the characters mentioned above.
var (
    string_char = 's'
    bytes_char = 'B'
    bool_char = 'b'
    int_char = 'i'
)

type Packet struct {
    data []byte
}

/* Packing Functions */
// The Packing functions pack the data into the class in a way that can be
// extracted easily.
func (this *Packet) PackString(data string) {
    // Pack some data into the packet
    information := escape + string_char + data
    convert := []byte(information)
    this.data = append(this.data, convert)
}

func (this *Packet) PackInt(data int) {
    // Pack some data into the packet
    // Convert the int to string
    data_string := strconv.Itoa(data)

    information := escape + int_char + data_string

    convert := []byte(information)
    this.data = append(this.data, information)
}

/* Unpacking Functions */
// Separate reads all of the information and returns it as a <[][]byte>
//  (a slice of byte slices)
func (this *Packet) Separate() []string {
    data := string(this.data)

    // Split it into a []string with the delimiter being the escape string
    split := strings.Split(data, escape)

    return split
}

// The problem with unpacking is that we don't know what the next thing is going
// to be. Because of this, the user, when unpacking, will have to unpack in the
// exact same order that the data was packed in.

// The unpacking starts at the beginning of the packet.
// If the packing order was string > []byte > int, the unpacking order will
// be string > []byte > int
func (this *Packet) UnpackInt() (int, error) {
    // First we have to separate all of the data
    split := this.Separate()

    // Then, extract the data from the first index
    to_unpack := split[0]

    // Just make sure it's the type that we want to unpack
    if (to_unpack[0] != int_char) {
        // Uh oh. They just tried to unpack something that wasn't an int.
        return -1, errors.New("E: Tried to unpack non-int: item is a '" + to_unpack[0] "'")
    }

    unpack_data := to_unpack[1:]

    data_int, _ := strconv.Atoi(unpack_data)

    return data_int, nil
}

// GetData() returns the raw data as a <[]byte>
func (this *Packet) GetData() []byte {
    return this.data
}
