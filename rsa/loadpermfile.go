package main

import (
	"bufio"
	"crypto/dsa"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	//load DSA PEM encoded public key file
	//pemFile, err := os.Open("/home/huangping/temp/dsa/dsakeytest_pub.pem")
	pemFile, err := os.Open("DSApublickey.pem")

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
	pemBlock, _ := pem.Decode([]byte(pemBytes))

	pemFile.Close()

	// convert PEM block to dsa.PublicKey
	// we use ASN1 because the PEM block was encoded with ASN1
	// see https://www.socketloop.com/tutorials/golang-generate-dsa-private-public-key-and-pem-files-example

	// if the PEM block is not encoded by ASN1, the unmarshal will fail
	// you need to find out the Public Key algorithm used to encode the PEM bytes

	//var publicKey dsa.PublicKey
	var publicKey dsa.PublicKey

	_, err = asn1.Unmarshal(pemBlock.Bytes, &publicKey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Printf("Public key : \n%x\n", publicKey)

	fmt.Printf("Public key parameter P: %v\n", publicKey.Parameters.P)
	fmt.Printf("Public key parameter Q: %v\n", publicKey.Parameters.Q)
	fmt.Printf("Public key parameter G: %v\n", publicKey.Parameters.G)
	fmt.Printf("Public key Y: %v\n", publicKey.Y)

}
