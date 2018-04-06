package proofofhack

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/p2p"
)

type ProofOfHack struct{}
type Message struct {
	Text string
}

func (p *ProofOfHack) Protocol() []p2p.Protocol {
	return []p2p.Protocol{{
		Name:    "ProofOfHack",
		Version: 1,
		Length:  1,
		Run:     p.Handler,
	}}
}

func (p *ProofOfHack) Handler(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
	for {
		p2p.SendItems(ws, 0, "foo")
		msg, err := ws.ReadMsg()
		if err != nil {
			return err
		}

		var testMessage Message
		err = msg.Decode(&testMessage)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch testMessage.Text {
		case "foo":
			log.Printf("foo: %v => %v", peer.ID().String(), testMessage)

			err := p2p.SendItems(ws, 0, peer.Name())
			if err != nil {
				fmt.Println(err)
				return err
			}
		default:
			log.Printf("default: %v => %v", peer.ID().String(), testMessage)
			err := p2p.SendItems(ws, 0, peer.Name())
			if err != nil {
				return err
			}
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}
