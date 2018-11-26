
package hyper

/* Packet */
// Hyper's custom packet class
// Packets allow data packing into a single <[]byte> which can later be sent off

// Imports
import (
    "strings"
    "strconv"

    "errors"
    "bytes"

    "fmt"
)

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
    data []byte
    unpacked int
    unpackdata []interface{}
}

/* Packing Functions */
// The Packing functions pack the data into the class in a way that can be
// extracted easily.
func (this *Packet) PackString(data string) {
    // Pack some data into the packet
    information := escape + string_char + data
    convert := []byte(information)

    for i := range convert {
        this.data = append(this.data, convert[i])
    }
    // this.data = append(this.data, convert)
}

func (this *Packet) PackInt(data int) {
    // Pack some data into the packet
    // Convert the int to string
    data_string := strconv.Itoa(data)

    information := escape + int_char + data_string

    convert := []byte(information)

    for i := range convert {
        this.data = append(this.data, convert[i])
    }
    // this.data = append(this.data, convert)
}

/* Unpacking Functions */
// Separate reads all of the information and returns it as a <[][]byte>
//  (a slice of byte slices)
func (this *Packet) Separate() [][]byte {
    // Split it into a []string with the delimiter being the escape string
    split := bytes.Split(this.data, []byte(escape))

    return split
}

func (this *Packet) SSep() []string {
    return strings.Split(string(this.data), escape)
}

func TrimBytes(data []byte) []byte {
    return bytes.Trim(data, "\x00")
}

func Trim(data string) string {
    return string(TrimBytes([]byte(data)))
}

/*

The Unpack.

*/
func (this *Packet) Unpack() []interface{} {
    // unpack
    split := this.SSep() // bytes

    result := make([]interface{}, 0)

    for i := 1; i < len(split); i++ {
        s := string(split[i])

        char := string(s[0])

        rest := s[1:]

        if char == int_char {
            // take out an int
            trim := Trim(rest)
            conv, err := strconv.Atoi(trim)
            fmt.Println(err)
            result = append(result, conv)
        } else if char == string_char {
            result = append(result, rest)
        }
    }

    return result
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
    to_unpack := split[this.unpacked + 1]

    // Just make sure it's the type that we want to unpack
    if (string(to_unpack[0]) != int_char) {
        // Uh oh. They just tried to unpack something that wasn't an int.
        return -1, errors.New("E: Tried to unpack non-int: item is a '" + string(to_unpack[0]) + "'")
    }

    unpack_data := string(bytes.Trim([]byte(to_unpack[1:]), "\x00"))

    data_int, err := strconv.Atoi(unpack_data)

    // Now we have to let ourselves know that we've
    this.unpacked += 1

    return data_int, err
}

// GetData() returns the raw data as a <[]byte>
func (this *Packet) GetData() []byte {
    return this.data
}
