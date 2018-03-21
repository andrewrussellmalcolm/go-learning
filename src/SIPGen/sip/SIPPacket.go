package sip

import (
	"fmt"
	"os"
	"strings"
)

/** */
const OPTIONS = "OPTIONS sip:[service]@[remote_ip]:[remote_port] SIP/2.0\r\n" +
	"Call-ID: [call_id]\r\n\r\n"

/** */
const INVITE = "INVITE sip:[service]@[remote_ip]:[remote_port] SIP/2.0\r\n" +
	"Via: SIP/2.0/[transport] [local_ip]:[local_port];branch=[branch]\r\n" +
	"From: sipp <sip:sipgen@[local_ip]:[local_port]>;tag=[call_number]\r\n" +
	"To: sut <sip:[service]@[remote_ip]:[remote_port]>\r\n" +
	"Call-ID: [call_id]\r\n" +
	"CSeq: 1 INVITE\r\n" +
	"Contact: sip:sipp@[local_ip]:[local_port]\r\n" +
	"Max-Forwards: 70\r\n" +
	"Subject: Performance Test\r\n" +
	"Content-Type: application/sdp\r\n" +
	"Content-Length: [len]\r\n\r\n" +
	"v=0\r\n" +
	"o=user1 53655765 2353687637 IN IP[local_ip_type] [local_ip]\r\n" +
	"s=-\r\n" +
	"c=IN IP[media_ip_type] [media_ip]\r\n" +
	"t=0 0\r\n" +
	"m=audio [media_port] RTP/AVP 0\r\n" +
	"a=rtpmap:0 PCMU/8000\r\n"

/** */
const ACK = "ACK sip:[service]@[remote_ip]:[remote_port] SIP/2.0\r\n" +
	"Via: SIP/2.0/[transport] [local_ip]:[local_port];branch=[branch]\r\n" +
	"From: sipp <sip:sipp@[local_ip]:[local_port]>;tag=[call_number]\r\n" +
	"o: sut <sip:[service]@[remote_ip]:[remote_port]>[peer_tag_param]\r\n" +
	"Call-ID: [call_id]\r\n" +
	"CSeq: 1 ACK\r\n" +
	"Contact: sip:sipp@[local_ip]:[local_port]\r\n" +
	"Max-Forwards: 70\r\n" +
	"Subject: Performance Test\r\n" +
	"Content-Length: 0\r\n\r\n"

/** */
const BYE = "BYE sip:[service]@[remote_ip]:[remote_port] SIP/2.0\r\n" +
	"Via: SIP/2.0/[transport] [local_ip]:[local_port];branch=[branch]\r\n" +
	"From: sipp <sip:sipp@[local_ip]:[local_port]>;tag=[call_number]\r\n" +
	"To: sut <sip:[service]@[remote_ip]:[remote_port]>[peer_tag_param]\r\n" +
	"Call-ID: [call_id]\r\n" +
	"CSeq: 2 BYE\r\n" +
	"Contact: sip:sipp@[local_ip]:[local_port]\r\n" +
	"Max-Forwards: 70\r\n" +
	"Subject: Performance Test\r\n" +
	"Content-Length: 0\r\n\r\n"

/** */
func PrepareMessage(message, localAddress, localPort, remoteAddress, remotePort, mediaAddress, mediaPort, callID string) string {

	message = strings.Replace(message, "[service]", "sipecho", -1)
	message = strings.Replace(message, "[local_ip]", localAddress, -1)
	message = strings.Replace(message, "[local_port]", string(localPort), -1)
	message = strings.Replace(message, "[remote_ip]", remoteAddress, -1)
	message = strings.Replace(message, "[remote_port]", string(remotePort), -1)
	message = strings.Replace(message, "[branch]", "", -1)
	message = strings.Replace(message, "[transport]", "UDP", -1)
	message = strings.Replace(message, "[call_number]", "1", -1)
	message = strings.Replace(message, "[call_id]", callID, -1)
	message = strings.Replace(message, "[media_ip]", mediaAddress, -1)
	message = strings.Replace(message, "[media_ip_type]", "4", -1)
	message = strings.Replace(message, "[local_ip_type]", "4", -1)
	message = strings.Replace(message, "[media_port]", string(mediaPort), -1)
	len := len(message) - strings.LastIndex(message, "v=0")
	message = strings.Replace(message, "[len]", fmt.Sprintf("%d", len), -1)

	return message
}

/** */
func GenerateCallID(localAddr string) string {
	pid := os.Getpid()
	return fmt.Sprintf("%d-%d@%s", 100, pid, localAddr)
}
