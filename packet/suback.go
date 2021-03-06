package packet

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// A SubackPacket is sent by the server to the client to confirm receipt and
// processing of a SubscribePacket. The SubackPacket contains a list of return
// codes, that specify the maximum QOS levels that have been granted.
type SubackPacket struct {
	// The granted QOS levels for the requested subscriptions.
	ReturnCodes []uint8

	// The packet identifier.
	ID ID
}

// NewSubackPacket creates a new SubackPacket.
func NewSubackPacket() *SubackPacket {
	return &SubackPacket{}
}

// Type returns the packets type.
func (sp *SubackPacket) Type() Type {
	return SUBACK
}

// String returns a string representation of the packet.
func (sp *SubackPacket) String() string {
	var codes []string

	for _, c := range sp.ReturnCodes {
		codes = append(codes, fmt.Sprintf("%d", c))
	}

	return fmt.Sprintf("<SubackPacket ID=%d ReturnCodes=[%s]>",
		sp.ID, strings.Join(codes, ", "))
}

// Len returns the byte length of the encoded packet.
func (sp *SubackPacket) Len() int {
	ml := sp.len()
	return headerLen(ml) + ml
}

// Decode reads from the byte slice argument. It returns the total number of
// bytes decoded, and whether there have been any errors during the process.
func (sp *SubackPacket) Decode(src []byte) (int, error) {
	total := 0

	// decode header
	hl, _, rl, err := headerDecode(src[total:], SUBACK)
	total += hl
	if err != nil {
		return total, err
	}

	// check buffer length
	if len(src) < total+2 {
		return total, fmt.Errorf("[%s] insufficient buffer size, expected %d, got %d", sp.Type(), total+2, len(src))
	}

	// check remaining length
	if rl <= 2 {
		return total, fmt.Errorf("[%s] expected remaining length to be greater than 2, got %d", sp.Type(), rl)
	}

	// read packet id
	sp.ID = ID(binary.BigEndian.Uint16(src[total:]))
	total += 2

	// check packet id
	if sp.ID == 0 {
		return total, fmt.Errorf("[%s] packet id must be grater than zero", sp.Type())
	}

	// calculate number of return codes
	rcl := int(rl) - 2

	// read return codes
	sp.ReturnCodes = make([]uint8, rcl)
	copy(sp.ReturnCodes, src[total:total+rcl])
	total += len(sp.ReturnCodes)

	// validate return codes
	for i, code := range sp.ReturnCodes {
		if !validQOS(code) && code != QOSFailure {
			return total, fmt.Errorf("[%s] invalid return code %d for topic %d", sp.Type(), code, i)
		}
	}

	return total, nil
}

// Encode writes the packet bytes into the byte slice from the argument. It
// returns the number of bytes encoded and whether there's any errors along
// the way. If there is an error, the byte slice should be considered invalid.
func (sp *SubackPacket) Encode(dst []byte) (int, error) {
	total := 0

	// check return codes
	for i, code := range sp.ReturnCodes {
		if !validQOS(code) && code != QOSFailure {
			return total, fmt.Errorf("[%s] invalid return code %d for topic %d", sp.Type(), code, i)
		}
	}

	// check packet id
	if sp.ID == 0 {
		return total, fmt.Errorf("[%s] packet id must be grater than zero", sp.Type())
	}

	// encode header
	n, err := headerEncode(dst[total:], 0, sp.len(), sp.Len(), SUBACK)
	total += n
	if err != nil {
		return total, err
	}

	// write packet id
	binary.BigEndian.PutUint16(dst[total:], uint16(sp.ID))
	total += 2

	// write return codes
	copy(dst[total:], sp.ReturnCodes)
	total += len(sp.ReturnCodes)

	return total, nil
}

// Returns the payload length.
func (sp *SubackPacket) len() int {
	return 2 + len(sp.ReturnCodes)
}
