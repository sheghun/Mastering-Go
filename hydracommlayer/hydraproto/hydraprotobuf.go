package hydraproto

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
	"time"
)

type ProtoHandler struct{}

func NewProtoHandler() *ProtoHandler {
	return new(ProtoHandler)
}

func (pSender *ProtoHandler) EncodeAndSend(obj interface{}, destination string) error {
	v, ok := obj.(*Ship)
	if !ok {
		return errors.New("proto: Unknown message type")
	}

	data, err := proto.Marshal(v)
	if err != nil {
		return err
	}

	return sendMessage(data, destination)
}

func (pSender *ProtoHandler) DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)

	return pb, proto.Unmarshal(buffer, pb)
}

func (pSender *ProtoHandler) ListenAndDecode(listenAddress string) (chan interface{}, error) {
	outChan := make(chan interface{})
	li, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return outChan, err
	}
	log.Println("Listening to ", listenAddress)

	go func() {
		defer li.Close()
		for {
			c, err := li.Accept()
			if err != nil {
				break
			}
			log.Println("Accepted Connection from ", c.RemoteAddr())
			go func(c net.Conn) {
				for {
					buffer, err := ioutil.ReadAll(c)
					if err != nil {
						break
					}
					if len(buffer) == 0 {
						continue
					}
					obj, err := pSender.DecodeProto(buffer)
					if err != nil {
						continue
					}
					select {
					case outChan <- obj:
					default:
					}
				}
			}(c)
		}
	}()
	return outChan, nil
}

func sendMessage(buffer []byte, destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("Sending %d bytes to %s \n", len(buffer), destination)
	_, err = conn.Write(buffer)

	return err
}
