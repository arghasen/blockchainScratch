package blockchain

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/big"
)

const Difficulty = 10

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) CreateData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash),
			[]byte(pow.Block.Data),
			make([]byte, 8),
			make([]byte, 8),
		},
		[]byte{},
	)
	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))
	return data
}

func (pow *ProofOfWork) MineBlock() (int, []byte) {
	var hashInt big.Int
	var hash [16]byte
	nonce := 0

	for {
		data := pow.CreateData(nonce)
		hash = md5.Sum(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.CreateData(pow.Block.Nonce)
	hash := md5.Sum(data)
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.Target) == -1 {
		return true
	} else {
		return false
	}
}
