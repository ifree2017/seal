package seal_rtmp_conn

import (
	"UtilsTools/identify_panic"
	"log"
)

func (rtmpSession *RtmpConn) HandleRtmpSession() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, "-", identify_panic.IdentifyPanic())
		}

		rtmpSession.Conn.Close()
		MapPublishingStreams.Delete(rtmpSession.StreamInfo.stream)

		log.Println("One RtmpConn finished.remote=", rtmpSession.Conn.RemoteAddr())
	}()

	log.Println("One RtmpConn come in. remote=", rtmpSession.Conn.RemoteAddr())

	err := rtmpSession.HandShake()
	if err != nil {
		log.Println("rtmp handshake failed, err=", err)
		return
	}

	log.Println("rtmp handshake success.remote=", rtmpSession.Conn.RemoteAddr())

	err = rtmpSession.RtmpMsgLoop()

	log.Println("rtmp msg loop quit.err=", err, ",remote=", rtmpSession.Conn.RemoteAddr())
}