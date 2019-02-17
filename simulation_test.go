package skeleton

import (
	"crypto/ecdsa"
	"crypto/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/p2p"
)

func TestSimulation(t *testing.T) {
	key, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	key2, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)

	srv1 := SubProtocol{}
	server1P2P := &p2p.Server{
		Config: p2p.Config{
			Name:       "test1",
			MaxPeers:   10,
			ListenAddr: ":8080",
			PrivateKey: key,
			Protocols:  srv1.Protocol(),
		},
	}

	if err := server1P2P.Start(); err != nil {
		t.Fatalf("Could not start server: %v", err)
	}

	srv2 := SubProtocol{}
	server2P2P := &p2p.Server{
		Config: p2p.Config{
			Name:       "test2",
			MaxPeers:   10,
			ListenAddr: ":8081",
			PrivateKey: key2,
			Protocols:  srv2.Protocol(),
		},
	}

	if err := server2P2P.Start(); err != nil {
		t.Fatalf("Could not start server: %v", err)
	}

	server1P2P.AddPeer(server2P2P.Self())

	time.Sleep(time.Second * 10)
}
