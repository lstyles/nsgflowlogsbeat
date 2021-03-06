package packets

import (
	"fmt"
	"io"
)

//PubcompPacket is an internal representation of the fields of the
//Pubcomp MQTT packet
type PubcompPacket struct {
	FixedHeader
	MessageID uint16
}

func (pc *PubcompPacket) String() string {
	return fmt.Sprintf("%s MessageID: %d", pc.FixedHeader, pc.MessageID)
}

func (pc *PubcompPacket) Write(w io.Writer) error {
	var err error
	pc.FixedHeader.RemainingLength = 2
	packet := pc.FixedHeader.pack()
	packet.Write(encodeUint16(pc.MessageID))
	_, err = packet.WriteTo(w)

	return err
}

//Unpack decodes the details of a ControlPacket after the fixed
//header has been read
func (pc *PubcompPacket) Unpack(b io.Reader) error {
	var err error
	pc.MessageID, err = decodeUint16(b)

	return err
}

//Details returns a Details struct containing the Qos and
//MessageID of this ControlPacket
func (pc *PubcompPacket) Details() Details {
	return Details{Qos: pc.Qos, MessageID: pc.MessageID}
}
