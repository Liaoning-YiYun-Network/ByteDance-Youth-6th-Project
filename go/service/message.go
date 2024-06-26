package service

import (
	"SkyLine/data"
	"SkyLine/entity"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net"
	"sync"
)

var chatConnMap = sync.Map{}

func RunMessageServer() {
	port := viper.GetString("server.port")
	if port == "" {
		data.Logger.Warn("No port specified, use default port 9090")
		port = "9090"
	}
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			data.Logger.Errorf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			data.Logger.Errorf("Close conn failed: %v\n", err)
		}
	}(conn)

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			data.Logger.Errorf("Read message failed: %v\n", err)
			continue
		}

		var event = entity.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		data.Logger.Infof("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			data.Logger.Warnf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := entity.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			data.Logger.Errorf("Push message failed: %v\n", err)
		}
	}
}
