package hydracommlayer

import "Mastering-Go/hydracommlayer/hydraproto"

const (
	Protobuf uint8 = iota
)

type HydraConnection interface {
	EncodeAndSend(obj interface{}, destination string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}

func NewConnection(connType uint8) HydraConnection {
	switch connType {
	case Protobuf:
		return hydraproto.NewProtoHandler()
	}
	return nil
}
