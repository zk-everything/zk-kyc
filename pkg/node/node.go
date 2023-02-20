package node

import "fmt"

func node() {
	ipfs := ipfshandler.NewIPFS("localhost:5001")
	files := make(map[string]string)
	ipfs.Files = files
	ipfs.UploadFile("path/to/file.txt", "file.txt")
	ipfs.DownloadFile("file.txt")
	fmt.Print("LOADED")

	// Consensus Algorithm

	pbft := &PBFT{Blockchain: blockchain, Nodes: nodes, F: 1}
	block := &Block{ID: 1, NodeList: nodes, Hash: "abc"}
	pbft.ProposeBlock(block)
	pbft.AddTransaction("example transaction")
}
