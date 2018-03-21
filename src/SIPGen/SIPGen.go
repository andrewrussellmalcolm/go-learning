package main

import (
	"SIPGen/rtp"
	"SIPGen/sip"
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

/** program entry point */
func main() {

	/** check we have all the args we  need */
	if len(os.Args) < 7 {
		fmt.Print("\n\tUsage: SIPGen LOCAL_HOST LOCAL_SIP_PORT REMOTE_HOST REMOTE_SIP_PORT LOCAL_RTP_PORT REMOTE_RTP_PORT\n")
		fmt.Print("\te.g: SIPGen.exe 192.168.33.224 5062 10.5.1.127 5060 12800 12802\n\n")
		os.Exit(0)
	}

	// get the configuration from the args: TODO - how about some validation
	localHost := os.Args[1]
	localSIPPort := os.Args[2]
	remoteHost := os.Args[3]
	remoteSIPPort := os.Args[4]
	localRTPPort := os.Args[5]
	remoteRTPPort := os.Args[6]

	// generate the SIP massages for this conversation
	invite := sip.PrepareMessage(sip.INVITE, localHost, localSIPPort, remoteHost, remoteSIPPort, remoteHost, remoteRTPPort, sip.GenerateCallID(localHost))
	ack := sip.PrepareMessage(sip.ACK, localHost, localSIPPort, remoteHost, remoteSIPPort, remoteHost, remoteRTPPort, sip.GenerateCallID(localHost))
	bye := sip.PrepareMessage(sip.BYE, localHost, localSIPPort, remoteHost, remoteSIPPort, remoteHost, remoteRTPPort, sip.GenerateCallID(localHost))

	// open connections to remote server
	sipConn := sip.NewSIPConn(localHost, localSIPPort, remoteHost, remoteSIPPort, localRTPPort, remoteRTPPort)

	// send invite
	sipConn.WriteSIP(invite)

	// wait fro treying, ringing, ok
	sipConn.ReadSIP()
	sipConn.ReadSIP()
	sipConn.ReadSIP()

	// send acknowledge
	sipConn.WriteSIP(ack)

	// send some audio
	// 160 samples at 8kHz = 20ms
	// 250 20ms packets = 5 seconds
	timestamp := 0
	payload := make([]byte, 160)
	ticker := time.NewTicker(20 * time.Millisecond)
	seq := 0
	stop := make(chan bool)

	sendAndReceive := func() {
		// generate packet, send and wait for response
		rand.Read(payload)
		outPacket := rtp.NewRPTPacket(seq == 0, 0x08, seq, timestamp, payload)
		sipConn.WriteRTP(outPacket.AsByteArray())
		sipConn.ReadRTP()
		timestamp += len(payload)
		seq++

		// stop when we have sent enough packets (250 20ms packets)
		if seq >= 250 {
			stop <- true
		}
	}

	// send first packet
	sendAndReceive()

	// wakes up every 20ms to send/recieve
	go func() {
		for range ticker.C {
			sendAndReceive()
		}
	}()

	// wait for signal to say all rtp sent
	<-stop
	ticker.Stop()

	// send bye and wait for ack
	sipConn.WriteSIP(bye)
	sipConn.ReadSIP()
}
