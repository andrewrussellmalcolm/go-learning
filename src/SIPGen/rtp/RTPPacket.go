package rtp

type RTPPacket struct {
	packet []byte
}

/** factory*/
func NewRPTPacket(marker bool, packetType, sequenceNumber, timestamp int, payload []byte) *RTPPacket {
	rtpPacket := new(RTPPacket)
	rtpPacket.init(marker, packetType, sequenceNumber, timestamp, payload)
	return rtpPacket
}

/** private initialiser */
func (rtpPacket *RTPPacket) init(marker bool, packetType, sequenceNumber, timestamp int, payload []byte) {

	const HEADERLEN = 12
	const VERSION = 2
	const PADDING = 0
	const EXTENSION = 0
	const CC = 0
	const SSRC = 1337

	rtpPacket.packet = make([]byte, HEADERLEN)

	rtpPacket.packet[0] = VERSION<<6 | PADDING<<5 | EXTENSION<<4 | CC

	if marker {
		rtpPacket.packet[1] = 0x80 | byte(packetType&0xff)
	} else {
		rtpPacket.packet[1] = 0x00 | byte(packetType&0xff)
	}
	rtpPacket.packet[2] = byte(sequenceNumber & 0x0000ff00 >> 8)
	rtpPacket.packet[3] = byte(sequenceNumber & 0x000000ff >> 0)

	rtpPacket.packet[4] = byte(timestamp & 0xff000000 >> 24)
	rtpPacket.packet[5] = byte(timestamp & 0x00ff0000 >> 16)
	rtpPacket.packet[6] = byte(timestamp & 0x0000ff00 >> 8)
	rtpPacket.packet[7] = byte(timestamp & 0x000000ff >> 0)

	rtpPacket.packet[8] = byte(SSRC & 0xff000000 >> 24)
	rtpPacket.packet[9] = byte(SSRC & 0x00ff0000 >> 16)
	rtpPacket.packet[10] = byte(SSRC & 0x0000ff00 >> 8)
	rtpPacket.packet[11] = byte(SSRC & 0x000000ff >> 0)

	rtpPacket.packet = append(rtpPacket.packet, payload...)
}

/** */
func (rtpPacket *RTPPacket) AsByteArray() []byte {
	return rtpPacket.packet

}
