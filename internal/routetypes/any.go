package routetypes

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Format
// struct any
//
//	{
//	    __u16 proto;
//	    __u16 port;
//	    __u32 PAD;
//	};
//
// Written out in the serialise function
type Any struct {
	Proto uint16
	Port  uint16
}

func (a Any) Bytes() []byte {
	output := make([]byte, 8)
	binary.LittleEndian.PutUint16(output[0:2], a.Proto)

	binary.LittleEndian.PutUint16(output[2:], a.Port)

	return output
}

func (a *Any) Unpack(b []byte) error {
	if len(b) < 8 {
		return errors.New("too short")
	}

	a.Proto = binary.LittleEndian.Uint16(b[0:2])
	a.Port = binary.LittleEndian.Uint16(b[2:])

	return nil
}

func (a Any) String() string {
	return fmt.Sprintf("proto %d, %d", a.Proto, a.Port)
}
