package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	str := "eNp9UTFOxDAQ7E+6P0yZSFEqREMHAokGJBoKROGz14mlxA5rh7vwEP7Cd3gJu9GdjgJRuPDuzuzMbOVyqVs8lp54HzI1CB5SQxingUaKJUN6eCLjiO84jQixEHtjZVY7Nk0LQj4DyGG3wJphCLFTrlbBCq0y27rdbrab788v3B6MQvTr52hxI0TXs/fEf5Qq1fTMQVY3EJqjoAa72ePldbcUqlHtZaBQVImXFw2IWV/iers5U61inQgNohElrVtAB0tTEUdGvBfkYjpS65zmrl+NTpzeBed0pdJUklSKpHRMb3NgcjXYaJJKEyEJJGuKpmBQaJwSG14U0+Ler8oFG8PQnHiOCHJXSKeL6EFEUG8yPogTBopd6Ztf2WAyMdj8b7APP+t2qzM="

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("data: %d\n", len(data))
	fmt.Printf("str: %d\n", len(str))

	deCompressedData := new(bytes.Buffer)
	decompressor := gzip.NewReader(bytes.NewReader(data))
	w, err := io.Copy(deCompressedData, decompressor)
	if err != nil {
		fmt.Printf("error  %v.\n", err)
	} else {
		fmt.Printf("ok  %v.\n", w)
	}
	decompressor.Close()

	fmt.Printf("enen %s", deCompressedData.String())
}

func test() {
	inData, _ := ioutil.ReadFile("kafka_demo.json")
	compressedData := new(bytes.Buffer)
	compress(inData, compressedData, 9)

	ioutil.WriteFile("compressed.dat", compressedData.Bytes(), os.ModeAppend)

	deCompressedData := new(bytes.Buffer)
	decompress(compressedData, deCompressedData)
	log.Print(deCompressedData)
}

func compress(src []byte, dest io.Writer, level int) {
	compressor, _ := flate.NewWriter(dest, level)
	compressor.Write(src)
	compressor.Close()
}
func decompress(src io.Reader, dest io.Writer) {
	decompressor := flate.NewReader(src)
	io.Copy(dest, decompressor)
	decompressor.Close()
}
