package main

import (
	"bufio"
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

type dsaPrivateKey struct {
	Version       int
	P, Q, G, Y, X *big.Int
}

func main() {
	fmt.Println(" run loadpivatepermfile.go  ping openssl ")

	//load DSA PEM encoded public key file
	pemFile, err := os.Open("/home/huangping/temp/d9/p1.pem")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// need to convert PEM file to []byte for decoding

	pemFileInfo, _ := pemFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	// read pemFile content into pemBytes
	buffer := bufio.NewReader(pemFile)
	_, err = buffer.Read(pemBytes)

	// proper decoding now
	block, _ := pem.Decode([]byte(pemBytes))

	pemFile.Close()

	var rawPriv dsaPrivateKey
	rest, err := asn1.Unmarshal(block.Bytes, &rawPriv)
	if len(rest) != 0 {
		panic("trailing garbage")
	}
	if err != nil {
		panic(err)
	}

	priv := &dsa.PrivateKey{
		PublicKey: dsa.PublicKey{
			Parameters: dsa.Parameters{
				P: rawPriv.P,
				Q: rawPriv.Q,
				G: rawPriv.G,
			},
			Y: rawPriv.Y,
		},
		X: rawPriv.X,
	}

	fmt.Printf("%#v\n", priv)

	checkY := new(big.Int).Exp(priv.G, priv.X, priv.P)
	if checkY.Cmp(priv.Y) != 0 {
		panic("invalid key")
	}
	singverify(priv)

}

func singverify(privatekey *dsa.PrivateKey) {
	// Sign
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	str := "/com/jingoal/token/signatureImpl/DSA"
	io.WriteString(h, str)
	signhash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("src message:  %v\n", str)
	fmt.Printf("token : %x\n", signature)

	var pubkey dsa.PublicKey
	pubkey = privatekey.PublicKey

	// Verify
	verifystatus := dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be true

	// we add additional data to change the signhash
	io.WriteString(h, "This message is NOT to be signed and verified!")
	signhash = h.Sum(nil)

	verifystatus = dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be false
}
