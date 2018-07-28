package main

import (
	"encoding/hex"
	"math/rand"
	"testing"
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipfs-blockstore"
)

func getRandomSmor(t uint64) *Smor {
	buf := make([]byte, 16)
	rand.Read(buf)
	data := hex.EncodeToString(buf)

	return &Smor{
		CreatedAt: t,
		Data:      data,
	}
}

func TestBasicPostlist(t *testing.T) {
	// make a memory backed blockstore for testing
	memds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(memds)

	// 'construct' our merklelist
	ml := MerkleList{bs: bs}

	// setup a few random posts to use for test data
	var posts []*Smor
	for i := 1; i < 18; i++ {
		posts = append(posts, getRandomSmor(uint64(i)))
	}

	for _, p := range posts {
		if err := ml.InsertPost(p); err != nil {
			t.Fatal("Failed to split node", err)
		}
		fmt.Println(ml.root)
	}
}