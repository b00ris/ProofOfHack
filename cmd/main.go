package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/b00ris/proof_of_hack"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"github.com/ethereum/go-ethereum/params"
)

const maxPeers = 80

func main() {
	listenAddr := flag.String("listenaddr", "127.0.0.1:8080", "")
	flag.Parse()
	p := proofofhack.ProofOfHack{}

	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Public key: %s\n", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
	server := &p2p.Server{
		Config: p2p.Config{
			PrivateKey:     key,
			MaxPeers:       maxPeers,
			Name:           "proof of hack test node",
			Protocols:      p.Protocol(),
			ListenAddr:     *listenAddr,
			NAT:            nat.Any(),
			BootstrapNodes: getBootnodes(params.TestnetBootnodes),
		},
	}
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln()
}

func getBootnodes(nodes []string) []*discover.Node {
	BootstrapNodes := make([]*discover.Node, 0, len(nodes))
	for _, url := range nodes {
		node, err := discover.ParseNode(url)
		if err != nil {
			log.Printf("Bootstrap URL invalid: %v, err: %v", url, err)
			continue
		}
		BootstrapNodes = append(BootstrapNodes, node)
	}
	return BootstrapNodes

}
