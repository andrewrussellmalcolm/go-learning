package sip

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type SIPConn struct {
	sipConn   net.Conn
	rtpConn   net.Conn
	debugging bool
}

/** factory*/
func NewSIPConn(localAddress, localPort, remoteAddress, remotePort, localMediaPort, remoteMediaPort string) *SIPConn {
	sipConn := new(SIPConn)
	sipConn.init(localAddress, localPort, remoteAddress, remotePort, localMediaPort, remoteMediaPort)
	return sipConn
}

/** private initialiser */
func (conn *SIPConn) init(localAddress, localPort, remoteAddress, remotePort, localMediaPort, remoteMediaPort string) {

	localAddr, err := net.ResolveUDPAddr("udp", localAddress+":"+localPort)
	conn.checkError("ResolveUDPAddr", err)

	remoteAddr, err := net.ResolveUDPAddr("udp", remoteAddress+":"+remotePort)
	conn.checkError("ResolveUDPAddr", err)

	localMediaAddr, err := net.ResolveUDPAddr("udp", localAddress+":"+localMediaPort)
	conn.checkError("ResolveUDPAddr", err)

	remoteMediaAddr, err := net.ResolveUDPAddr("udp", remoteAddress+":"+remoteMediaPort)
	conn.checkError("ResolveUDPAddr", err)

	conn.sipConn, err = net.DialUDP("udp", localAddr, remoteAddr)
	conn.checkError("DialUDP", err)

	conn.rtpConn, err = net.DialUDP("udp", localMediaAddr, remoteMediaAddr)
	conn.checkError("DialUDP", err)
}

/** */
func (conn *SIPConn) ReadSIP() {

	packet := make([]byte, 1024)

	conn.sipConn.SetDeadline(time.Now().Add(time.Millisecond * 500))

	n, err := conn.sipConn.Read(packet)
	conn.checkError("ReadSIP", err)

	if conn.debugging {
		fmt.Printf("READ SIP %d bytes: %.*s\n", n, strings.Index(string(packet), "\r\n"), packet)
	}
}

/** */
func (conn *SIPConn) ReadRTP() {

	packet := make([]byte, 1024)

	conn.rtpConn.SetDeadline(time.Now().Add(time.Millisecond * 500))

	n, err := conn.rtpConn.Read(packet)
	conn.checkError("ReadRTP", err)

	if conn.debugging {
		fmt.Printf("\tREAD RTP %d bytes\n", n)
	}
}

/** */
func (conn *SIPConn) WriteSIP(message string) {

	if conn.debugging {
		fmt.Printf("WRITE SIP %d bytes: %.*s\n", len(message), strings.Index(message, "\r\n"), message)
	}

	conn.sipConn.SetDeadline(time.Now().Add(time.Millisecond * 500))
	_, err := conn.sipConn.Write([]byte(message))
	conn.checkError("WriteSIP", err)
}

func (conn *SIPConn) WriteRTP(message []byte) {

	if conn.debugging {
		fmt.Printf("\tWRITE RTP %d bytes\n", len(message))
	}

	conn.rtpConn.SetDeadline(time.Now().Add(time.Millisecond * 500))
	_, err := conn.rtpConn.Write(message)
	conn.checkError("WriteRTP", err)
}

func (conn *SIPConn) SetDebugging(debugging bool) {
	conn.debugging = debugging
}

/* A Simple function to verify error */
func (conn *SIPConn) checkError(message string, err error) {
	if err != nil {
		fmt.Printf("%s : %s", message, err)
		os.Exit(0)
	}
}
