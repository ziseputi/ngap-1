// Created By HaoDHH-245789 VHT2020
package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"vht5gc/lib/ngap"
	"vht5gc/lib/ngap/ngapBuild"
)

func main () {
	pdu := ngapBuild.PathSwitchRequest()
	bin, err := ngap.Encode(pdu)
	fmt.Printf("%v\n%v\n", bin, err)
	hexStr := hex.EncodeToString(bin)
	fmt.Println(hexStr)

	filePointer, _ := os.Create("output/test.bin")
	filePointer.WriteString(hexStr)

	//fmt.Println(hex.DecodeString("0003e0ac106cb5000186a00007"))
}