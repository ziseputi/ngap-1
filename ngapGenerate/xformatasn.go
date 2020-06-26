package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("PDUContents.asn")
	datas := strings.Split(string(data), "\n")
	for i:=0; i<len(datas); i++ {
		datai := strings.Fields(datas[i])
		if len(datai) > 1 && datai[0] == "protocolIEs" && len(datai) != 7 {
			fmt.Println(i, "-", datai)
		}
	}
}
