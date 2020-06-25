package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// BuildPDU()
	BuildIEs()
}

func BuildPDU() {
	dataPDUDe, _ := ioutil.ReadFile("PDUDescriptions.asn")
	dataPDUDes := strings.Split(string(dataPDUDe), "\n")
	str := "// Created By HaoDHH-245789 VHT2020\npackage ngapBuild\n\nimport \"ngap/ngapType\"\n"
	for iDe := 300; iDe < len(dataPDUDes); iDe++ {
		dataPDUDeLine0 := strings.Fields(dataPDUDes[iDe])
		if len(dataPDUDeLine0) == 4 && dataPDUDeLine0[1] == "NGAP-ELEMENTARY-PROCEDURE" {
			dataPDUDeLine1 := strings.Fields(dataPDUDes[iDe+1])
			dataPDUDeLine2 := strings.Fields(dataPDUDes[iDe+2])
			dataPDUDeLine3 := strings.Fields(dataPDUDes[iDe+3])
			dataPDUDeLine4 := strings.Fields(dataPDUDes[iDe+4])
			dataPDUDeLine5 := strings.Fields(dataPDUDes[iDe+5])
			var pduPresent, pduPCode, pduCri, pduValuePresent string
			if dataPDUDeLine1[0] == "INITIATING" {
				pduPresent = "InitiatingMessage"
				pduValuePresent = GetStrOutOfDashB(dataPDUDeLine1[2])
				if dataPDUDeLine2[0] == "PROCEDURE" {
					pduPCode = GetStrOutOfDashB2(dataPDUDeLine2[2])
					pduCri = GetStrOutOfDashB(dataPDUDeLine3[1])
					str += BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent)
					continue
				} else if dataPDUDeLine2[0] == "SUCCESSFUL" {
					if dataPDUDeLine3[0] == "PROCEDURE" {
						pduPCode = GetStrOutOfDashB2(dataPDUDeLine3[2])
						pduCri = GetStrOutOfDashB(dataPDUDeLine4[1])
						str += BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent)

						pduPresent = "SuccessfulOutcome"
						pduValuePresent = GetStrOutOfDashB(dataPDUDeLine2[2])
						pduPCode = GetStrOutOfDashB2(dataPDUDeLine3[2])
						pduCri = GetStrOutOfDashB(dataPDUDeLine4[1])
						str += BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent)
						continue
					} else if dataPDUDeLine3[0] == "UNSUCCESSFUL" {
						pduPCode = GetStrOutOfDashB2(dataPDUDeLine4[2])
						pduCri = GetStrOutOfDashB(dataPDUDeLine5[1])
						str += BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent)

						pduPresent = "SuccessfulOutcome"
						pduValuePresent = GetStrOutOfDashB(dataPDUDeLine2[2])
						pduPCode = GetStrOutOfDashB2(dataPDUDeLine4[2])
						pduCri = GetStrOutOfDashB(dataPDUDeLine5[1])
						str += BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent)

						pduPresent = "UnsuccessfulOutcome"
						pduValuePresent = GetStrOutOfDashB(dataPDUDeLine3[2])
						pduPCode = GetStrOutOfDashB2(dataPDUDeLine4[2])
						pduCri = GetStrOutOfDashB(dataPDUDeLine5[1])
						str += BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent)
						continue
					} else {
						fmt.Println(iDe, "-", dataPDUDeLine0)
					}
				} else {
					fmt.Println(iDe, "-", dataPDUDeLine0)
				}
			} else {
				fmt.Println(iDe, "-", dataPDUDeLine0)
			}
		}
	}
	filePointer, _ := os.Create("../ngapBuild/PDU.go")
	filePointer.WriteString(str)
}

func BuildPDUDescription(pduPresent, pduPCode, pduCri, pduValuePresent string) string {
	str := "\nfunc " + pduValuePresent + " () (pdu ngapType.NGAPPDU) {\n" +
		"\tpdu.Present = ngapType.NGAPPDUPresent" + pduPresent + "\n" +
		"\tpdu." + pduPresent + " = new(ngapType." + pduPresent + ")\n" +
		"\tmsg := pdu." + pduPresent + "\n" +
		"\tmsg.ProcedureCode.Value = ngapType.ProcedureCode" + pduPCode + "\n" +
		"\tmsg.Criticality.Value = ngapType.Criticality" + pduCri + "\n" +
		"\tmsg.Value.Present = ngapType." + pduPresent + "ValuePresent" + pduValuePresent + "\n" +
		"\tmsg.Value." + pduValuePresent + " = new(ngapType." + pduValuePresent + ")\n" +
		"\tmsgProtocolIEs := &msg.Value." + pduValuePresent + ".ProtocolIEs\n" +
		"\tprotocolIEs := ngapType." + pduValuePresent + "IEs{}\n"
	str += BuildPDUContent(pduValuePresent)
	str += "\n\treturn\n}\n"
	return str
}

func BuildPDUContent(pduValuePresent string) (str string) {
	dataPDUCo, _ := ioutil.ReadFile("PDUContents.asn")
	dataPDUCos := strings.Split(string(dataPDUCo), "\n")
	for iCo := 300; iCo < len(dataPDUCos); iCo++ {
		dataPDUCoLine0 := strings.Fields(dataPDUCos[iCo])
		if len(dataPDUCoLine0) > 3 && GetStrOutOfDashB(dataPDUCoLine0[0]) == (pduValuePresent+"IEs") {
			x := 0
			for {
				dataPDUCoLinex := strings.Fields(dataPDUCos[iCo+1+x])
				if len(dataPDUCoLinex) == 1 {
					break
				}
				// 	{ ID id-AMF-UE-NGAP-ID CRITICALITY ignore TYPE AMF-UE-NGAP-ID PRESENCE mandatory }|
				str += "\n\tprotocolIEs = ngapType." + pduValuePresent + "IEs{}\n" +
					"\tprotocolIEs.ProtocolIEID.Value = ngapType.ProtocolIEID" + GetStrOutOfDashB2(dataPDUCoLinex[2]) + "\n" +
					"\tprotocolIEs.Criticality.Value = ngapType.Criticality" + GetStrOutOfDashB(dataPDUCoLinex[4]) + "\n" +
					"\tprotocolIEs.TypeValue.Present = ngapType." + pduValuePresent + "IEsTypeValuePresent" + GetStrOutOfDashB2(dataPDUCoLinex[2]) + "\n" +
					"\tprotocolIEs.TypeValue." + GetStrOutOfDashB2(dataPDUCoLinex[2]) + " = new(ngapType." + GetStrOutOfDashB(dataPDUCoLinex[6]) + ")\n" +
					"\t*protocolIEs.TypeValue." + GetStrOutOfDashB2(dataPDUCoLinex[2]) + " = " + GetStrOutOfDashB(dataPDUCoLinex[6]) + "()\n" +
					"\tmsgProtocolIEs.List = append(msgProtocolIEs.List, protocolIEs)\n"
				x++
			}
			break
		}
	}

	return str
}

func BuildIEs() {
	data, _ := ioutil.ReadFile("IEs.asn")
	datas := strings.Split(string(data), "\n")
	str := "// Created By HaoDHH-245789 VHT2020\npackage ngapBuild\n\nimport \"ngap/ngapType\"\n"
	for i := 0; i < len(datas); i++ {
		data0 := strings.Fields(datas[i])
		if len(data0) > 2 {
			if data0[1] == "::=" && data0[2] == "ENUMERATED" {
				if len(data0) == 4 {
					str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
					x := 1
					for {
						datax := strings.Fields(datas[i+x])
						if datax[0] == "..." || datax[0] == "}" {
							break
						}
						str += "\tvalue.Value = ngapType." + GetStrOutOfDashB(data0[0]) + GetStrOutOfDashB(strings.Split(datax[0], ",")[0]) + "\n"
						x++
					}
					str += "\treturn\n}\n"
				} else {
					str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
					x := 4
					for {
						if data0[x] == "..." || data0[x] == "}" {
							break
						}
						str += "\tvalue.Value = ngapType." + GetStrOutOfDashB(data0[0]) + GetStrOutOfDashB(strings.Split(data0[x], ",")[0]) + "\n"
						x++
					}
					str += "\treturn\n}\n"
				}
			} else
			if data0[1] == "::=" && data0[2] == "INTEGER" {
				min, max := "", ""
				if len(data0) == 4 {
					min, max = strings.Split(strings.Split(data0[3], "..")[0], "(")[1], strings.Split(strings.Split(data0[3], "..")[1], ")")[0]
				} else if len(data0) == 5 {
					min, max = strings.Split(strings.Split(data0[3], "..")[0], "(")[1], strings.Split(strings.Split(data0[3], "..")[1], ",")[0]
				}
				str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
				str += "\tvalue.Value = " + min + "\n"
				str += "\tvalue.Value = " + max + "\n"
				str += "\treturn\n}\n"
			} else
			if data0[1] == "::=" && data0[2] == "BIT" {
				if len(data0) == 5 {
					data0x := strings.Split(data0[4], "..")
					if len(data0x) < 2 {
						minmax := strings.Split(strings.Split(data0[4], "(")[2], ")")[0]
						str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
						str += "\tvalue.Value.BitLength = " + minmax + "\n"
						mimax, _ := strconv.Atoi(minmax)
						bytesss := ""
						if mimax < 9 {
							bytesss = "171"
						} else if mimax < 17 {
							bytesss = "171,255"
						} else if mimax < 25 {
							bytesss = "171,101,255"
						} else if mimax < 33 {
							bytesss = "171,101,102,255"
						} else if mimax < 41 {
							bytesss = "171,101,102,103,255"
						} else if mimax < 49 {
							bytesss = "171,101,102,103,104,255"
						} else if mimax < 57 {
							bytesss = "171,101,102,103,104,105,255"
						} else if mimax < 65 {
							bytesss = "171,101,102,103,104,105,106,255"
						} else if mimax < 73 {
							bytesss = "171,101,102,103,104,105,106,107,255"
						} else if mimax < 81 {
							bytesss = "171,101,102,103,104,105,106,107,108,255"
						} else if mimax == 160 {
							bytesss = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255"
						} else if mimax == 256 {
							bytesss = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255"
						} else {
							bytesss = "x"
						}
						str += "\tvalue.Value.Bytes = []byte{" + bytesss + "}\n"
						str += "\treturn\n}\n"
					} else {
						min, max := strings.Split(data0x[0], "(")[2], strings.Split(data0x[1], ")")[0]
						str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
						minx, _ := strconv.Atoi(min)
						bytesmin := ""
						if minx < 9 {
							bytesmin = "171"
						} else if minx < 17 {
							bytesmin = "171,255"
						} else if minx < 25 {
							bytesmin = "171,101,255"
						} else if minx < 33 {
							bytesmin = "171,101,102,255"
						} else if minx < 41 {
							bytesmin = "171,101,102,103,255"
						} else if minx < 49 {
							bytesmin = "171,101,102,103,104,255"
						} else if minx < 57 {
							bytesmin = "171,101,102,103,104,105,255"
						} else if minx < 65 {
							bytesmin = "171,101,102,103,104,105,106,255"
						} else if minx < 73 {
							bytesmin = "171,101,102,103,104,105,106,107,255"
						} else if minx < 81 {
							bytesmin = "171,101,102,103,104,105,106,107,108,255"
						} else if minx == 160 {
							bytesmin = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255"
						} else if minx == 256 {
							bytesmin = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255"
						} else {
							bytesmin = "x"
						}
						maxx, _ := strconv.Atoi(max)
						bytesmax := ""
						if maxx < 9 {
							bytesmax = "171"
						} else if maxx < 17 {
							bytesmax = "171,255"
						} else if maxx < 25 {
							bytesmax = "171,101,255"
						} else if maxx < 33 {
							bytesmax = "171,101,102,255"
						} else if maxx < 41 {
							bytesmax = "171,101,102,103,255"
						} else if maxx < 49 {
							bytesmax = "171,101,102,103,104,255"
						} else if maxx < 57 {
							bytesmax = "171,101,102,103,104,105,255"
						} else if maxx < 65 {
							bytesmax = "171,101,102,103,104,105,106,255"
						} else if maxx < 73 {
							bytesmax = "171,101,102,103,104,105,106,107,255"
						} else if maxx < 81 {
							bytesmax = "171,101,102,103,104,105,106,107,108,255"
						} else if maxx == 160 {
							bytesmax = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255"
						} else if maxx == 256 {
							bytesmax = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255"
						} else {
							bytesmax = "x"
						}
						str += "\tvalue.Value.BitLength = " + min + "\n"
						str += "\tvalue.Value.Bytes = []byte{" + bytesmin + "}\n"
						str += "\tvalue.Value.BitLength = " + max + "\n"
						str += "\tvalue.Value.Bytes = []byte{" + bytesmax + "}\n"
						str += "\treturn\n}\n"
					}
				} else if len(data0) == 6 {
					data0x := strings.Split(data0[4], "..")
					if len(data0x) < 2 {
						minmax := strings.Split(strings.Split(data0[4], "(")[2], ",")[0]
						str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
						str += "\tvalue.Value.BitLength = " + minmax + "\n"
						mimax, _ := strconv.Atoi(minmax)
						bytesss := ""
						if mimax < 9 {
							bytesss = "171"
						} else if mimax < 17 {
							bytesss = "171,255"
						} else if mimax < 25 {
							bytesss = "171,101,255"
						} else if mimax < 33 {
							bytesss = "171,101,102,255"
						} else if mimax < 41 {
							bytesss = "171,101,102,103,255"
						} else if mimax < 49 {
							bytesss = "171,101,102,103,104,255"
						} else if mimax < 57 {
							bytesss = "171,101,102,103,104,105,255"
						} else if mimax < 65 {
							bytesss = "171,101,102,103,104,105,106,255"
						} else if mimax < 73 {
							bytesss = "171,101,102,103,104,105,106,107,255"
						} else if mimax < 81 {
							bytesss = "171,101,102,103,104,105,106,107,108,255"
						} else if mimax == 160 {
							bytesss = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255"
						} else if mimax == 256 {
							bytesss = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255"
						} else {
							bytesss = "x"
						}
						str += "\tvalue.Value.Bytes = []byte{" + bytesss + "}\n"
						str += "\treturn\n}\n"
					} else {
						min, max := strings.Split(data0x[0], "(")[2], strings.Split(data0x[1], ",")[0]
						str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
						minx, _ := strconv.Atoi(min)
						bytesmin := ""
						if minx < 9 {
							bytesmin = "171"
						} else if minx < 17 {
							bytesmin = "171,255"
						} else if minx < 25 {
							bytesmin = "171,101,255"
						} else if minx < 33 {
							bytesmin = "171,101,102,255"
						} else if minx < 41 {
							bytesmin = "171,101,102,103,255"
						} else if minx < 49 {
							bytesmin = "171,101,102,103,104,255"
						} else if minx < 57 {
							bytesmin = "171,101,102,103,104,105,255"
						} else if minx < 65 {
							bytesmin = "171,101,102,103,104,105,106,255"
						} else if minx < 73 {
							bytesmin = "171,101,102,103,104,105,106,107,255"
						} else if minx < 81 {
							bytesmin = "171,101,102,103,104,105,106,107,108,255"
						} else if minx == 160 {
							bytesmin = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255"
						} else if minx == 256 {
							bytesmin = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255"
						} else {
							bytesmin = "x"
						}
						maxx, _ := strconv.Atoi(max)
						bytesmax := ""
						if maxx < 9 {
							bytesmax = "171"
						} else if maxx < 17 {
							bytesmax = "171,255"
						} else if maxx < 25 {
							bytesmax = "171,101,255"
						} else if maxx < 33 {
							bytesmax = "171,101,102,255"
						} else if maxx < 41 {
							bytesmax = "171,101,102,103,255"
						} else if maxx < 49 {
							bytesmax = "171,101,102,103,104,255"
						} else if maxx < 57 {
							bytesmax = "171,101,102,103,104,105,255"
						} else if maxx < 65 {
							bytesmax = "171,101,102,103,104,105,106,255"
						} else if maxx < 73 {
							bytesmax = "171,101,102,103,104,105,106,107,255"
						} else if maxx < 81 {
							bytesmax = "171,101,102,103,104,105,106,107,108,255"
						} else if maxx == 160 {
							bytesmax = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255"
						} else if maxx == 256 {
							bytesmax = "171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,101,102,103,104,105,106,107,108,255,171,255"
						} else {
							bytesmax = "x"
						}
						str += "\tvalue.Value.BitLength = " + min + "\n"
						str += "\tvalue.Value.Bytes = []byte{" + bytesmin + "}\n"
						str += "\tvalue.Value.BitLength = " + max + "\n"
						str += "\tvalue.Value.Bytes = []byte{" + bytesmax + "}\n"
						str += "\treturn\n}\n"
					}
				}
			} else
			if data0[1] == "::=" && data0[2] == "OCTET" {
				if len(data0) == 4 {
					str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
					str += "\tvalue.Value = []byte{171, 101, 102, 103, 104, 105, 106, 107, 108, 255}\n"
					str += "\tvalue.Value = []byte{171, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 255}\n"
					str += "\treturn\n}\n"
				} else if len(data0) == 5 {
					data0x := strings.Split(data0[4], "..")
					if len(data0x) < 2 {
						mimax := strings.Split(strings.Split(data0[4], "(")[2], ")")[0]
						mimaxx, _ := strconv.Atoi(mimax)
						bytessss := "171"
						for im := 1; im < mimaxx; im++ {
							if im%64 == 0 {
								bytessss += ",\n\t\t" + strconv.Itoa(im%256)
							} else {
								bytessss += "," + strconv.Itoa(im%256)
							}
						}
						str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
						str += "\tvalue.Value = []byte{" + bytessss + "}\n"
						str += "\treturn\n}\n"
					} else {
						min, max := strings.Split(data0x[0], "(")[2], strings.Split(data0x[1], ")")[0]
						minx, _ := strconv.Atoi(min)
						bytesmin := "171"
						for im := 1; im < minx; im++ {
							bytesmin += ",1" + strconv.Itoa(im)
						}
						maxx, _ := strconv.Atoi(max)
						bytesmax := "219"
						for im := 1; im < maxx; im++ {
							if im%64 == 0 {
								bytesmax += ",\n\t\t" + strconv.Itoa(im%256)
							} else {
								bytesmax += "," + strconv.Itoa(im%256)
							}
						}
						str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
						str += "\tvalue.Value = []byte{" + bytesmin + "}\n"
						str += "\tvalue.Value = []byte{" + bytesmax + "}\n"
						str += "\treturn\n}\n"
					}
				} else {
					fmt.Println(data0)
				}
			} else
			if data0[1] == "::=" && data0[2] == "OBJECT" {
				fmt.Println(data0)
			} else
			if data0[1] == "::=" && data0[2] == "PrintableString" {
				str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
				str += "\tvalue.Value = \"-haodhh-vht5gc-\"\n"
				str += "\tvalue.Value = \"-haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc--haodhh-vht5gc-\"\n"
				str += "\treturn\n}\n"
			} else
			if data0[1] == "::=" && data0[2] == "CHOICE" {
				str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
				xs := 1
				for {
					datax := strings.Fields(datas[i+xs])
					if len(datax) <= 1 || datax[0] == "choice-Extensions" {
						break
					}
					if datax[1] == "BIT" {
						dataxx := strings.Split(datax[3], "..")
						if len(dataxx) < 2 {
							mimax := strings.Split(strings.Split(datax[3], "(")[2], ")")[0]
							mima, _ := strconv.Atoi(mimax)
							str += "\tvalue.Present = ngapType." + GetStrOutOfDashB(data0[0]) + "Present" + GetStrOutOfDashB(datax[0]) + "\n" +
								"\tvalue." + GetStrOutOfDashB(datax[0]) + " = new(ngapType.BitString)\n" +
								"\tvalue."+ GetStrOutOfDashB(datax[0]) +".BitLength = "+mimax+"\n"
								if mima < 9 {
									str += "\tvalue."+ GetStrOutOfDashB(datax[0]) +".Bytes = []byte{171}\n"
								} else if mima < 17 {
									str += "\tvalue."+ GetStrOutOfDashB(datax[0]) +".Bytes = []byte{171,219}\n"
								} else if mima < 25 {
									str += "\tvalue."+ GetStrOutOfDashB(datax[0]) +".Bytes = []byte{171,219,255}\n"
								}

						} else {
							min, max := strings.Split(dataxx[0], "(")[2], strings.Split(dataxx[1], ")")[0]
							str += "\tvalue.Present = ngapType." + GetStrOutOfDashB(data0[0]) + "Present" + GetStrOutOfDashB(datax[0]) + "\n" +
								"\tvalue." + GetStrOutOfDashB(datax[0]) + " = new(ngapType.BitString)\n" +
								"\tvalue."+ GetStrOutOfDashB(datax[0]) +".BitLength = "+min+"\n" +
								"\tvalue."+ GetStrOutOfDashB(datax[0]) +".Bytes = []byte{171,219,255}\n" +
								"\tvalue."+ GetStrOutOfDashB(datax[0]) +".BitLength = "+max+"\n" +
								"\tvalue."+ GetStrOutOfDashB(datax[0]) +".Bytes = []byte{171,219,217,255}\n"
						}
					} else {
						str += "\tvalue.Present = ngapType." + GetStrOutOfDashB(data0[0]) + "Present" + GetStrOutOfDashB(datax[0]) + "\n" +
							"\tvalue." + GetStrOutOfDashB(datax[0]) + " = new(ngapType." + GetStrOutOfDashB(strings.Split(datax[1], ",")[0]) + ")\n" +
							"\t*value." + GetStrOutOfDashB(datax[0]) + " = " + GetStrOutOfDashB(strings.Split(datax[1], ",")[0]) + "()\n"
					}
					xs++
				}
				str += "\treturn\n}\n"
			} else
			if data0[1] == "::=" && data0[2] == "SEQUENCE" {
				if len(data0) == 4 {
					str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n"
					xs := 1
					for {
						dataxs := strings.Fields(datas[i+xs])
						if len(dataxs) <= 1 {
							break
						}
						if dataxs[len(dataxs)-1] == "OPTIONAL," {
							if dataxs[0] == "iE-Extensions" {
							} else if dataxs[0] == "protocolIEs" {
							} else if dataxs[1] == "INTEGER" {

							} else if dataxs[1] == "ENUMERATED" {

							} else if dataxs[1] == "BIT" {

							} else if dataxs[1] == "OCTET" {

							} else {
								str += "\tvalue." + GetStrOutOfDashB(dataxs[0]) + " = new(ngapType." + GetStrOutOfDashB(dataxs[1]) + ")\n" +
									"\t*value." + GetStrOutOfDashB(dataxs[0]) + " = " + GetStrOutOfDashB(dataxs[1]) + "()\n"
							}
						} else {
							if dataxs[0] == "iE-Extensions" {
							} else if dataxs[0] == "protocolIEs" {
							} else if dataxs[1] == "INTEGER" {

							} else if dataxs[1] == "ENUMERATED" {

							} else if dataxs[1] == "BIT" {

							} else if dataxs[1] == "OCTET" {

							} else {
								str += "\tvalue." + GetStrOutOfDashB(dataxs[0]) + " = " + GetStrOutOfDashB(strings.Split(dataxs[1], ",")[0]) + "()\n"
							}
						}
						xs++
					}
					str += "\treturn\n}\n"
				} else
				if len(data0) == 6 {
					min, max := strings.Split(strings.Split(data0[3], "..")[0], "(")[2], strings.Split(strings.Split(data0[3], "..")[1], ")")[0]
					str += "\nfunc " + GetStrOutOfDashB(data0[0]) + "() (value ngapType." + GetStrOutOfDashB(data0[0]) + ") {\n" +
						"\titem := " + GetStrOutOfDashB(data0[5]) + "()\n"
					minx, _ := strconv.Atoi(min)
					appendmin := ""
					maxx := GetMaxConstants(max)
					appendmax := ""
					for im := 0; im < minx; im++ {
						appendmin += ",item"
					}
					for im := 0; im < maxx; im++ {
						if im%32 == 0 {
							appendmax += ",\n\t\titem"
						} else {
							appendmax += ",item"
						}
					}
					str += "\tvalue.List = append(value.List" + appendmin + ")\n" +
						"\tvalue.List = append(value.List" + appendmax + ")\n" +
						"\treturn\n}\n"
				} else {
					fmt.Println(data0)
				}
			} else
			{
				// fmt.Println(data0)
			}
		}
	}
	filePointer, _ := os.Create("../ngapBuild/IEs.go")
	filePointer.WriteString(str)
}

func GetStrOutOfDashB(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 0; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}

// for name begin with "id-"
func GetStrOutOfDashB2(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 1; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}

func GetMaxConstants(str string) int {
	data, _ := ioutil.ReadFile("Constants.asn")
	datas := strings.Split(string(data), "\n")
	for i := 0; i < len(datas); i++ {
		datax := strings.Fields(datas[i])
		if len(datax) == 4 {
			if datax[0] == str {
				val, _ := strconv.Atoi(datax[3])
				if val > 128 {
					return 128
				}
				return val
			}
		}
	}
	return 0
}
