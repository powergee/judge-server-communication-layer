package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"compress/zlib"
	"bytes"
	"encoding/json"
)

func main() {
	data, _ := ioutil.ReadFile("encoded")
	fmt.Println(data)
	fmt.Println("Total size:", len(data))

	size := binary.BigEndian.Uint32(data[:4])
	fmt.Println(size)

	compressed := data[4:]
	reader := bytes.NewReader(compressed)
	z, _ := zlib.NewReader(reader)
	defer z.Close()
	decompressed, _ := ioutil.ReadAll(z)

	var parsed map[string]interface{}
	json.Unmarshal(decompressed, &parsed)
	
	for key, value := range parsed {
		fmt.Println(key, ":", value)
	}
}