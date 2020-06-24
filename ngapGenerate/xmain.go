package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var strExt string

func main() {
	// NGAPCommonDataTypes()
	// NGAPConstants()
	// NGAPContainers()
	NGAPIEs()
	// NGAPPDUContents()
	// NGAPPDUDescriptions()
}

func NGAPCommonDataTypes() {
	data, _ := ioutil.ReadFile("CommonDataTypes.asn")
	datas := strings.Split(string(data), "\n")
	str := "package ngapType\n"
	for i := 0; i < len(datas); i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 2 {
			if dataLine[1] == "::=" && dataLine[2] == "ENUMERATED" {
				strx, out := ENUMERATED(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "CHOICE" {
				strx, out := CHOICE(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "INTEGER" {
				strx, out := INTEGER(datas, i)
				str += strx
				i = out
			}
		}
	}
	filePointer, _ := os.Create("../CommonDataTypes.go")
	filePointer.WriteString(str)
}

func NGAPConstants() {
	// TODO
}

func NGAPContainers() {
	// TODO
}

func NGAPIEs() {
	strExt = ""
	data, _ := ioutil.ReadFile("IEs.asn")
	datas := strings.Split(string(data), "\n")
	str := "package ngapType\n"
	kxx := 0
	for i := 0; i < len(datas); i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 2 {
			if dataLine[1] == "::=" && dataLine[2] == "ENUMERATED" {
				strx, out := ENUMERATED(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "INTEGER" {
				strx, out := INTEGER(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "OCTET" && dataLine[3] == "STRING" {
				strx, out := OCTETSTRING(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "BIT" && dataLine[3] == "STRING" {
				strx, out := BITSTRING(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "PrintableString" {
				strx, out := PrintableString(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "CHOICE" {
				strx, out := CHOICE(datas, i)
				str += strx
				i = out
			}
			if dataLine[2] == "::=" && dataLine[1] == "NGAP-PROTOCOL-IES" {
				strx, out := NGAPPROTOCOLIES(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "SEQUENCE" {
				strx, out := SEQUENCE(datas, i)
				str += strx
				i = out
			}
			if dataLine[2] == "::=" && dataLine[1] == "NGAP-PROTOCOL-EXTENSION" {
				strx, out := NGAPPROTOCOLEXTENSION(datas, i)
				str += strx
				i = out
			}
		}
	}
	fmt.Println(kxx)
	str += strExt
	filePointer, _ := os.Create("../ngapType/IEsEnumerated.go")
	filePointer.WriteString(str)
}

func PDUDescriptions() {
	data, _ := ioutil.ReadFile("ngapASN1/PDUDescriptions.asn1")
	datas := strings.Split(string(data), "\n")
	str := "package ngapType\n" +
		"\nconst (\n\tNGAPPDUChoiceInitiatingMessage int = 0\n\tNGAPPDUChoiceSuccessfulOutcome int = 1\n" +
		"\tNGAPPDUChoiceUnsuccessfulOutcome int = 2\n\t/* Extensions may appear below */\n)\n" +
		"\ntype NGAPPDU struct {\n\tChoice int\n\tInitiatingMessage *InitiatingMessage\n" +
		"\tSuccessfulOutcome *SuccessfulOutcome\n\tUnsuccessfulOutcome *UnsuccessfulOutcome\n}\n" +
		"\ntype InitiatingMessage struct {\n\tProcedureCode ProcedureCode\n\tCriticality Criticality `vht:" +
		"\"Reference:ProcedureCode\"`\n\tValue InitiatingMessageValue `vht:\"Reference:ProcedureCode\"`\n}\n" +
		"\ntype SuccessfulOutcome struct {\n\tProcedureCode ProcedureCode\n\tCriticality Criticality `vht:" +
		"\"Reference:ProcedureCode\"`\n\tValue SuccessfulOutcomeValue `vht:\"Reference:ProcedureCode\"`\n}\n" +
		"\ntype UnsuccessfulOutcome struct {\n\tProcedureCode ProcedureCode\n\tCriticality Criticality `vht:" +
		"\"Reference:ProcedureCode\"`\n\tValue UnsuccessfulOutcomeValue `vht:\"Reference:ProcedureCode\"`\n}\n"
	// INITIATING MESSAGE
	str += "\nconst (\n"
	j := 0
	for i := 300; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "NGAP-ELEMENTARY-PROCEDURE" && dataLine[2] == "::=" {
				dataLine1 := strings.Fields(datas[i+1])
				if dataLine1[0] == "INITIATING" {
					str += "\tInitiatingMessageValueChoice" + dataLine1[2] + " int = " + strconv.Itoa(j) + "\n"
					j++
				}
			}
		}
	}
	str += ")\n"
	str += "\ntype InitiatingMessageValue struct {\n\tChoice int\n"
	for i := 300; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "NGAP-ELEMENTARY-PROCEDURE" && dataLine[2] == "::=" {
				dataLine1 := strings.Fields(datas[i+1])
				var ProcedureCode string
				var Criticality string
				if dataLine1[0] == "INITIATING" {
					dataLine2 := strings.Fields(datas[i+2])
					if dataLine2[0] == "SUCCESSFUL" {
						dataLine3 := strings.Fields(datas[i+3])
						if dataLine3[0] == "UNSUCCESSFUL" {
							ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+4])[2])
							Criticality = strings.Title(strings.Fields(datas[i+5])[1])
						} else {
							ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+3])[2])
							Criticality = strings.Title(strings.Fields(datas[i+4])[1])
						}
					} else if dataLine2[0] == "UNSUCCESSFUL" {
						ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+3])[2])
						Criticality = strings.Title(strings.Fields(datas[i+4])[1])
					} else {
						ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+2])[2])
						Criticality = strings.Title(strings.Fields(datas[i+3])[1])
					}
					str += "\t" + dataLine1[2] + " *" + dataLine1[2] + " `vht:\"Presence:PresenceOptional," +
						"Criticality:Criticality" + Criticality + ",ProcedureCode:ProcedureCode" + ProcedureCode + "\"`\n"
					j++
				}
			}
		}
	}
	str += "}\n"
	// SUCCESSFUL OUTCOME
	str += "\nconst (\n"
	j = 0
	for i := 300; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "NGAP-ELEMENTARY-PROCEDURE" && dataLine[2] == "::=" {
				dataLine1 := strings.Fields(datas[i+1])
				if dataLine1[0] == "SUCCESSFUL" {
					str += "\tSuccessfulOutcomeValueChoice" + dataLine1[2] + " int = " + strconv.Itoa(j) + "\n"
					j++
				} else {
					dataLine2 := strings.Fields(datas[i+2])
					if dataLine2[0] == "SUCCESSFUL" {
						str += "\tSuccessfulOutcomeValueChoice" + dataLine2[2] + " int = " + strconv.Itoa(j) + "\n"
						j++
					}
				}
			}
		}
	}
	str += ")\n"
	str += "\ntype SuccessfulOutcomeValue struct {\n\tChoice int\n"
	for i := 300; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "NGAP-ELEMENTARY-PROCEDURE" && dataLine[2] == "::=" {
				dataLine1 := strings.Fields(datas[i+1])
				var ProcedureCode string
				var Criticality string
				if dataLine1[0] == "SUCCESSFUL" {
					dataLine2 := strings.Fields(datas[i+2])
					if dataLine2[0] == "UNSUCCESSFUL" {
						ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+3])[2])
						Criticality = strings.Title(strings.Fields(datas[i+4])[1])
					} else {
						ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+2])[2])
						Criticality = strings.Title(strings.Fields(datas[i+3])[1])
					}
					str += "\t" + dataLine1[2] + " *" + dataLine1[2] + " `vht:\"Presence:PresenceOptional," +
						"Criticality:Criticality" + Criticality + ",ProcedureCode:ProcedureCode" + ProcedureCode + "\"`\n"
					j++
				} else {
					dataLine2 := strings.Fields(datas[i+2])
					if dataLine2[0] == "SUCCESSFUL" {
						dataLine3 := strings.Fields(datas[i+3])
						if dataLine3[0] == "UNSUCCESSFUL" {
							ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+4])[2])
							Criticality = strings.Title(strings.Fields(datas[i+5])[1])
						} else {
							ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+3])[2])
							Criticality = strings.Title(strings.Fields(datas[i+4])[1])
						}
						str += "\t" + dataLine2[2] + " *" + dataLine2[2] + " `vht:\"Presence:PresenceOptional," +
							"Criticality:Criticality" + Criticality + ",ProcedureCode:ProcedureCode" + ProcedureCode + "\"`\n"
						j++
					}
				}
			}
		}
	}
	str += "}\n"
	// UNSUCCESSFUL OUTCOME
	str += "\nconst (\n"
	j = 0
	for i := 300; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "NGAP-ELEMENTARY-PROCEDURE" && dataLine[2] == "::=" {
				dataLine1 := strings.Fields(datas[i+1])
				if dataLine1[0] == "UNSUCCESSFUL" {
					str += "\tUnsuccessfulOutcomeValueChoice" + dataLine1[2] + " int = " + strconv.Itoa(j) + "\n"
					j++
				} else {
					dataLine2 := strings.Fields(datas[i+2])
					if dataLine2[0] == "UNSUCCESSFUL" {
						str += "\tUnsuccessfulOutcomeValueChoice" + dataLine2[2] + " int = " + strconv.Itoa(j) + "\n"
						j++
					} else {
						dataLine3 := strings.Fields(datas[i+3])
						if dataLine3[0] == "UNSUCCESSFUL" {
							str += "\tUnsuccessfulOutcomeValueChoice" + dataLine3[2] + " int = " + strconv.Itoa(j) + "\n"
							j++
						}
					}
				}
			}
		}
	}
	str += ")\n"
	str += "\ntype UnsuccessfulOutcomeValue struct {\n\tChoice int\n"
	for i := 300; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "NGAP-ELEMENTARY-PROCEDURE" && dataLine[2] == "::=" {
				dataLine1 := strings.Fields(datas[i+1])
				var ProcedureCode string
				var Criticality string
				if dataLine1[0] == "UNSUCCESSFUL" {
					ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+2])[2])
					Criticality = strings.Title(strings.Fields(datas[i+3])[1])
					str += "\t" + dataLine1[2] + " *" + dataLine1[2] + " `vht:\"Presence:PresenceOptional," +
						"Criticality:Criticality" + Criticality + ",ProcedureCode:ProcedureCode" + ProcedureCode + "\"`\n"
					j++
				} else {
					dataLine2 := strings.Fields(datas[i+2])
					if dataLine2[0] == "UNSUCCESSFUL" {
						ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+3])[2])
						Criticality = strings.Title(strings.Fields(datas[i+4])[1])
						str += "\t" + dataLine2[2] + " *" + dataLine2[2] + " `vht:\"Presence:PresenceOptional," +
							"Criticality:Criticality" + Criticality + ",ProcedureCode:ProcedureCode" + ProcedureCode + "\"`\n"
						j++
					} else {
						dataLine3 := strings.Fields(datas[i+3])
						if dataLine3[0] == "UNSUCCESSFUL" {
							ProcedureCode = GetStrOutOfDash2(strings.Fields(datas[i+4])[2])
							Criticality = strings.Title(strings.Fields(datas[i+5])[1])
							str += "\t" + dataLine3[2] + " *" + dataLine3[2] + " `vht:\"Presence:PresenceOptional," +
								"Criticality:Criticality" + Criticality + ",ProcedureCode:ProcedureCode" + ProcedureCode + "\"`\n"
							j++
						}
					}
				}
			}
		}
	}
	str += "}\n"
	filePointer, _ := os.Create("../../ngapType/PDUDescriptions.go")
	filePointer.WriteString(str)
}

func NGAPPDUContents() {
	data, _ := ioutil.ReadFile("PDUContents.asn")
	datas := strings.Split(string(data), "\n")
	str := "package main\n"
	for i := 0; i < len(datas); i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 2 {
			if dataLine[2] == "::=" && dataLine[1] == "NGAP-PROTOCOL-IES" {
				strx, out := NGAPPROTOCOLIES2(datas, i)
				str += strx
				i = out
			}
			if dataLine[1] == "::=" && dataLine[2] == "SEQUENCE" {
				strx, out := SEQUENCE(datas, i)
				str += strx
				i = out
			}
			if dataLine[2] == "::=" && dataLine[1] == "NGAP-PRIVATE-IES" {
				strx, out := NGAPPRIVATEIES(datas, i)
				str += strx
				i = out
			}
		}
	}
	filePointer, _ := os.Create("PDUContents.go")
	filePointer.WriteString(str)
}

func PDUContents() {
	data, _ := ioutil.ReadFile("ngapASN1/PDUContents.asn1")
	datas := strings.Split(string(data), "\n")
	str := "package ngapType\n"
	for i := 335; i < len(datas)-3; i++ {
		dataLine := strings.Fields(datas[i])
		if len(dataLine) > 3 {
			if dataLine[1] == "::=" && dataLine[2] == "SEQUENCE" {
				str += "\ntype " + dataLine[0] + " struct {\n"
				dataLine1 := strings.Fields(datas[i+1])
				dataLine5 := strings.Fields(datas[i+5])
				str += "\t" + strings.Title(dataLine1[0]) + " " + GetStrOutOfDash(dataLine1[1]) + GetStrOutOfDash(dataLine5[0]) + "\n"
				str += "}\n"
				str += "\ntype " + GetStrOutOfDash(dataLine1[1]) + GetStrOutOfDash(dataLine5[0]) + " struct {\n"
				str += "\tList []" + GetStrOutOfDash(dataLine5[0]) + " `vht:\"sizeMin:0,sizeMax:65535\"`\n"
				str += "}\n"
				str += "\ntype " + GetStrOutOfDash(dataLine5[0]) + " struct {\n"
				str += "\tProtocolIEID ProtocolIEID\n"
				str += "\tCriticality Criticality `vht:\"Reference:ProtocolIEID\"`\n"
				str += "\tValue " + GetStrOutOfDash(dataLine5[0]) + "Value `vht:\"Reference:ProtocolIEID\"`\n"
				str += "}\n"
				str += "\nconst (\n"
				j, k := 0, 0
				for {
					dataLineX := strings.Fields(datas[i+6+j])
					if len(dataLineX) < 2 {
						break
					} else if dataLineX[0] != "{" {
						j++
						continue
					}
					str += "\t" + GetStrOutOfDash(dataLine5[0]) + "ValueChoice" + GetStrOutOfDash2(strings.Title(dataLineX[2])) + "" +
						" int = " + strconv.Itoa(k) + "\n"
					j++
					k++
				}
				str += ")\n"
				str += "\ntype " + GetStrOutOfDash(dataLine5[0]) + "Value struct {\n"
				str += "\tChoice int\n"
				j = 0
				for {
					dataLineX := strings.Fields(datas[i+6+j])
					if len(dataLineX) < 2 {
						break
					} else if dataLineX[0] != "{" {
						j++
						continue
					}
					str += "\t" + GetStrOutOfDash2(strings.Title(dataLineX[2])) + " *" + GetStrOutOfDash(strings.Title(dataLineX[6])) + "" +
						" `vht:\"Presence:Presence" + strings.Title(dataLineX[8]) + ",Criticality:Criticality" + strings.Title(dataLineX[4]) + "" +
						",ProtocolIEID:ProtocolIEID" + GetStrOutOfDash(strings.Title(dataLineX[6])) + "\"`" + " \n"
					j++
				}
				str += "}\n"
			}
		}
	}
	filePointer, _ := os.Create("../../ngapType/PDUContents.go")
	filePointer.WriteString(str)
}

func ENUMERATED(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	ext := ""
	if len(datax) == 3 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue Enumerated\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 4 && datax[3] == "{" {
		str = "\nconst (\n"
		x := 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "..." {
				ext = "valueExt,"
				break
			} else if dataxx[0] == "}" {
				break
			}
			str += "\t" + GetStrOutOfDash(datax[0]) + GetStrOutOfDash(strings.Split(dataxx[0], ",")[0]) + " Enumerated = " + strconv.Itoa(x) + "\n"
			x++
		}
		str += ")\n"
		str += "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue Enumerated `vht:\""+ext+"valueMin:0,valueMax:" + strconv.Itoa(x-1) + "\"`\n"
		str += "}\n"
		return str, in + x
	} else {
		str = "\nconst (\n"
		x := 0
		for {
			if datax[x+4] == "..." {
				ext = "valueExt,"
				break
			} else if datax[x+4] == "}" {
				break
			}
			str += "\t" + GetStrOutOfDash(datax[0]) + GetStrOutOfDash(strings.Split(datax[x+4], ",")[0]) + " Enumerated = " + strconv.Itoa(x) + "\n"
			x++
		}
		str += ")\n"
		str += "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue Enumerated `vht:\""+ext+"valueMin:0,valueMax:" + strconv.Itoa(x-1) + "\"`\n"
		str += "}\n"
		return str, in
	}
}

func INTEGER(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	if len(datax) == 3 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue Integer\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 4 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[1], strings.Split(strings.Split(datax[3], "..")[1], ")")[0]
		str += "\tValue Integer `vht:\"valueMin:" + min + ",valueMax:" + max + "\"`\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 5 && datax[4] == "...)" {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[1], strings.Split(strings.Split(datax[3], "..")[1], ",")[0]
		str += "\tValue Integer `vht:\"valueExt,valueMin:" + min + ",valueMax:" + max + "\"`\n"
		str += "}\n"
		return str, in
	} else {
		fmt.Println(datax)
		return str, in
	}
}

func OCTETSTRING(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	if len(datax) == 4 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue OctetString\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 5 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		dataxx := strings.Split(datax[4], "..")
		if len(dataxx) > 1 {
			min, max := strings.Split(strings.Split(datax[4], "..")[0], "(")[2], strings.Split(strings.Split(datax[4], "..")[1], ")")[0]
			str += "\tValue OctetString `vht:\"sizeExt,sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		} else {
			min, max := strings.Split(strings.Split(datax[4], ")")[0], "(")[2], strings.Split(strings.Split(datax[4], ")")[0], "(")[2]
			str += "\tValue OctetString `vht:\"sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		}
		str += "}\n"
		return str, in
	} else {
		fmt.Println(datax)
		return str, in
	}
}

func BITSTRING(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	if len(datax) == 4 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue BitString\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 5 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		dataxx := strings.Split(datax[4], "..")
		if len(dataxx) > 1 {
			min, max := strings.Split(strings.Split(datax[4], "..")[0], "(")[2], strings.Split(strings.Split(datax[4], "..")[1], ")")[0]
			str += "\tValue BitString `vht:\"sizeExt,sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		} else {
			min, max := strings.Split(strings.Split(datax[4], ")")[0], "(")[2], strings.Split(strings.Split(datax[4], ")")[0], "(")[2]
			str += "\tValue BitString `vht:\"sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		}
		str += "}\n"
		return str, in
	} else if len(datax) == 6 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		dataxx := strings.Split(datax[4], "..")
		if len(dataxx) > 1 {
			min, max := strings.Split(strings.Split(datax[4], "..")[0], "(")[2], strings.Split(strings.Split(datax[4], "..")[1], ",")[0]
			str += "\tValue BitString `vht:\"sizeExx,sizeExt,sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		} else {
			min, max := strings.Split(strings.Split(datax[4], ",")[0], "(")[2], strings.Split(strings.Split(datax[4], ",")[0], "(")[2]
			str += "\tValue BitString `vht:\"sizeExx,sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		}
		str += "}\n"
		return str, in
	} else {
		fmt.Println(datax)
		return str, in
	}
}

func PrintableString(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	if len(datax) == 3 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tValue PrintableString\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 4 {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[2], strings.Split(strings.Split(datax[3], "..")[1], ")")[0]
		str += "\tValue PrintableString `vht:\"sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		str += "}\n"
		return str, in
	} else if len(datax) == 5 && datax[4] == "...))" {
		str = "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[2], strings.Split(strings.Split(datax[3], "..")[1], ",")[0]
		str += "\tValue PrintableString `vht:\"sizeExt,sizeMin:" + min + ",sizeMax:" + max + "\"`\n"
		str += "}\n"
		return str, in
	} else {
		fmt.Println(datax)
		return str, in
	}
}

func CHOICE(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	str = "\nconst (\n"
	x := 0
	for {
		dataxx := strings.Fields(datas[in+1+x])
		if dataxx[0] == "..." || dataxx[0] == "}" {
			break
		}
		str += "\t" + GetStrOutOfDash(datax[0]) + "Choice" + GetStrOutOfDash(dataxx[0]) + " int = " + strconv.Itoa(x) + "\n"
		x++
	}
	str += ")\n"
	str += "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
	str += "\tChoice int\n"
	x = 0
	for {
		dataxx := strings.Fields(datas[in+1+x])
		if dataxx[0] == "..." || dataxx[0] == "}" {
			break
		} else if dataxx[1] == "INTEGER" || dataxx[1] == "INTEGER," {
			if len(dataxx) == 2 {
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Integer\n"
			} else if len(dataxx) == 3 {
				min, max := strings.Split(strings.Split(dataxx[2], "..")[0], "(")[1], strings.Split(strings.Split(dataxx[2], "..")[1], ")")[0]
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Integer `vht:\"valueMin:" + min + ",valueMax:" + max + "\"`\n"
			} else {
				fmt.Println(dataxx)
			}
		} else if dataxx[1] == "ENUMERATED" || dataxx[1] == "ENUMERATED," {
			if len(dataxx) == 2 {
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Enumerated\n"
			} else {
				fmt.Println(dataxx)
			}
		} else if dataxx[1] == "OBJECT" && (dataxx[2] == "IDENTIFIER" || dataxx[2] == "IDENTIFIER,") {
			if len(dataxx) == 3 {
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *ObjectIdentifier\n"
			} else {
				fmt.Println(dataxx)
			}
		} else if dataxx[1] == "OCTET" && (dataxx[2] == "STRING" || dataxx[2] == "STRING,") {
			if len(dataxx) == 3 {
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *OctetString\n"
			} else {
				fmt.Println(dataxx)
			}
		} else if dataxx[1] == "BIT" && (dataxx[2] == "STRING" || dataxx[2] == "STRING,") {
			if len(dataxx) == 3 {
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *BitString\n"
			} else if len(dataxx) == 4 {
				min, max := "", ""
				if len(strings.Split(dataxx[3], "..")) > 1 {
					min, max = strings.Split(strings.Split(dataxx[3], "..")[0], "(")[2], strings.Split(strings.Split(dataxx[3], "..")[1], ")")[0]
				} else {
					min, max = strings.Split(strings.Split(dataxx[3], ")")[0], "(")[2], strings.Split(strings.Split(dataxx[3], ")")[0], "(")[2]
				}
				str += "\t" + GetStrOutOfDash(dataxx[0]) + " *BitString `vht:\"valueMin:" + min + ",valueMax:" + max + "\"`\n"
			} else {
				fmt.Println(dataxx)
			}
		} else if len(dataxx) == 2 {
			str += "\t" + GetStrOutOfDash(dataxx[0]) + " *" + GetStrOutOfDash(strings.Split(dataxx[1],",")[0]) + "\n"
		} else if len(dataxx) == 5 {
			str += "\t" + GetStrOutOfDash(dataxx[0]) + " *" + GetStrOutOfDash(dataxx[1]) + GetStrOutOfDash(strings.Split(strings.Split(dataxx[3],"}")[0],"{")[1]) + "\n"
		} else {
			fmt.Println(dataxx)
		}
		x++
	}
	str += "}\n"
	return str, in + x
}

func SEQUENCE(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	x := 0
	if len(datax) == 6 && datax[4] == "OF" {
		str += "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tList []" + GetStrOutOfDash(datax[5]) + "`vht:\"" +
			"sizeMin:" + strings.Title(GetMinMaxSEQUENCEOF(datax[3])[0]) + "," +
			"sizeMax:" + strings.Title(GetMinMaxSEQUENCEOF(datax[3])[1]) + "\"`\n"
		str += "}\n"
	} else if len(datax) == 4 {
		str += "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" {
				break
			} else if dataxx[0] == "..." || dataxx[0] == "--" {
				x++
				continue
			} else if len(dataxx) == 2 {
				if dataxx[1] == "INTEGER," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " Integer\n"
				} else if dataxx[1] == "ENUMERATED," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " Enumerated\n"
				} else {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " " + GetStrOutOfDash(strings.Split(dataxx[1], ",")[0]) + "\n"
				}
				x++
			} else if len(dataxx) == 3 {
				if dataxx[1] == "INTEGER" && dataxx[2] == "OPTIONAL," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Integer\n"
				} else if dataxx[1] == "ENUMERATED" && dataxx[2] == "OPTIONAL," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Enumerated\n"
				} else if dataxx[2] == "OPTIONAL," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " *" + GetStrOutOfDash(dataxx[1]) + "\n"
				} else if dataxx[1] == "BIT" && dataxx[2] == "STRING," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " BitString\n"
				} else if dataxx[1] == "OCTET" && dataxx[2] == "STRING," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " OctetString\n"
				} else if dataxx[1] == "OBJECT" && dataxx[2] == "IDENTIFIER," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " ObjectIdentifier\n"
				} else if dataxx[1] == "INTEGER" {
					min, max := strings.Split(strings.Split(dataxx[2], "..")[0],"(")[1], strings.Split(strings.Split(dataxx[2], "..")[1],")")[0]
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " Integer `vht:\"valueMin:"+min+",valueMax:"+max+"\"`\n"
				} else {
					fmt.Println(in+1+x, "-", dataxx)
				}
				x++
			} else if len(dataxx) == 4 {
				if dataxx[1] == "INTEGER" {
					if dataxx[3] == "OPTIONAL," {
						min, max := strings.Split(strings.Split(dataxx[2], "..")[0],"(")[1], strings.Split(strings.Split(dataxx[2], "..")[1],")")[0]
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Integer `vht:\"valueMin:"+min+",valueMax:"+max+"\"`\n"
					} else if dataxx[3] == "...)," {
						min, max := strings.Split(strings.Split(dataxx[2], "..")[0],"(")[1], strings.Split(strings.Split(dataxx[2], "..")[1],",")[0]
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " Integer `vht:\"valueExt,valueMin:"+min+",valueMax:"+max+"\"`\n"
					} else {
						fmt.Println(in+1+x, "-", dataxx)
					}
				} else if dataxx[1] == "OCTET" && dataxx[2] == "STRING" {
					mimax := strings.Split(strings.Split(dataxx[3], "(")[2],")")[0]
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " OctetString `vht:\"sizeMin:"+mimax+",sizeMax:"+mimax+"\"`\n"
				} else {
					fmt.Println(in+1+x, "-", dataxx)
				}
				x++
			} else if len(dataxx) == 5 {
				if dataxx[1] == "BIT" && dataxx[2] == "STRING" {
					if dataxx[4] == "OPTIONAL," {
						min, max := strings.Split(strings.Split(dataxx[3], "..")[0],"(")[2], strings.Split(strings.Split(dataxx[3], "..")[1],")")[0]
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " *BitString `vht:\"valueMin:"+min+",valueMax:"+max+"\"`\n"
					} else if dataxx[4] == "...))," {
						mimax := strings.Split(strings.Split(dataxx[3], "(")[2],",")[0]
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " BitString `vht:\"valueExt,valueMin:"+mimax+",valueMax:"+mimax+"\"`\n"
					} else {
						fmt.Println(in+1+x, "-", dataxx)
					}
				} else if dataxx[1] == "OCTET" && dataxx[2] == "STRING" {
					if dataxx[4] == "OPTIONAL," {
						min, max := strings.Split(strings.Split(dataxx[3], "..")[0],"(")[2], strings.Split(strings.Split(dataxx[3], "..")[1],")")[0]
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " *OctetString `vht:\"valueMin:"+min+",valueMax:"+max+"\"`\n"
					} else if dataxx[4] == "...))," {
						mimax := strings.Split(strings.Split(dataxx[3], "(")[2],",")[0]
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " OctetString `vht:\"valueExt,valueMin:"+mimax+",valueMax:"+mimax+"\"`\n"
					} else if dataxx[3] == "(CONTAINING" {
						str += "\t" + GetStrOutOfDash(dataxx[0]) + " OctetString `vht:\"Reference:"+GetStrOutOfDash(strings.Split(dataxx[4],")")[0])+"\"`\n"
					} else {
						fmt.Println(in+1+x, "-", dataxx)
					}
				} else if dataxx[0] == "protocolIEs" && dataxx[1] == "ProtocolIE-Container" {
					str += "\tProtocolIEs ProtocolIEContainer"+GetStrOutOfDash(strings.Split(strings.Split(dataxx[3],"{")[1], "}")[0])+"\n"
				} else if dataxx[0] == "privateIEs" && dataxx[1] == "PrivateIE-Container" {
					str += "\tPrivateIEs PrivateIEContainer"+GetStrOutOfDash(strings.Split(strings.Split(dataxx[3],"{")[1], "}")[0])+"\n"
				} else {
					fmt.Println(in+1+x, "-", dataxx)
				}
				x++
			} else if len(dataxx) == 6 {
				if dataxx[0] == "iE-Extensions" {
					str += "\tIEExtensions *" + dataxx[1] + GetStrOutOfDash(strings.Split(strings.Split(dataxx[3], "{")[1], "}")[0]) + "\n"
				} else {
					fmt.Println(in+1+x, "-", dataxx)
				}
				x++
			} else if dataxx[1] == "ENUMERATED" {
				if dataxx[len(dataxx)-1] == "OPTIONAL," {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " *Enumerated `vht:\"valueExt,valueMin:0,valueMax:"+strconv.Itoa(len(dataxx)-6)+"\"`\n"
					strExt += "\nconst (\n"
					for xl := 3; xl<len(dataxx)-3; xl++ {
						strExt += "\t"+GetStrOutOfDash(dataxx[0]) + GetStrOutOfDash(strings.Split(dataxx[xl],",")[0]) + " Enumerated = " + strconv.Itoa(xl-3) + "\n"
					}
					strExt += ")\n"
				} else {
					str += "\t" + GetStrOutOfDash(dataxx[0]) + " Enumerated `vht:\"valueExt,valueMin:0,valueMax:"+strconv.Itoa(len(dataxx)-5)+"\"`\n"
					strExt += "\nconst (\n"
					for xl := 3; xl<len(dataxx)-2; xl++ {
						strExt += "\t"+GetStrOutOfDash(dataxx[0]) + GetStrOutOfDash(strings.Split(dataxx[xl],",")[0]) + " Enumerated = " + strconv.Itoa(xl-3) + "\n"
					}
					strExt += ")\n"
				}
				x++
			} else {
				fmt.Println(in+1+x, "----", dataxx)
				x++
			}
		}
		str += "}\n"
	} else {
		fmt.Println(datax)
	}
	return str, in+x
}

func NGAPPROTOCOLEXTENSION(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	x := 0
	if len(datax) == 4 {
		str += "\ntype ProtocolExtensionContainer" + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tList []" + GetStrOutOfDash(datax[0]) + " `vht:\"sizeMin:0,sizeMax:MaxProtocolExtensions\"`\n"
		str += "}\n"
		str += "\ntype " + GetStrOutOfDash(datax[0]) + " struct {\n"
		str += "\tProtocolExtensionID ProtocolExtensionID\n"
		str += "\tCriticality Criticality `vht:\"Reference:ProtocolExtensionID\"`\n"
		str += "\tExtensionValue " + GetStrOutOfDash(datax[0]) + "ExtensionValue `vht:\"Reference:ProtocolExtensionID\"`\n"
		str += "}\n"
		str += "\nconst (\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" {
				if x == 1 {
					str += "\t" + GetStrOutOfDash(datax[0]) + "ExtensionValueChoiceNothing int = 0\n"
				}
				break
			} else if len(dataxx) == 1 {
				x++
				continue
			} else if len(dataxx) == 10 {
				str += "\t" + GetStrOutOfDash(datax[0]) + "ExtensionValueChoice" + GetStrOutOfDash2(dataxx[2]) + " int = " + strconv.Itoa(x) + "\n"
				x++
			} else if len(dataxx) == 13 {
				str += "\t" + GetStrOutOfDash(datax[0]) + "ExtensionValueChoice" + GetStrOutOfDash2(dataxx[2]) + " int = " + strconv.Itoa(x) + "\n"
				x++
			} else {
				fmt.Println(in+1+x, "-", dataxx)
				x++
			}
		}
		str += ")\n"
		str += "\ntype " + GetStrOutOfDash(datax[0]) + "ExtensionValue struct {\n"
		str += "\tChoice int\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" {
				break
			} else if len(dataxx) == 1 {
				x++
				continue
			} else if len(dataxx) == 10 {
				str += "\t" + GetStrOutOfDash2(dataxx[2]) + " *" + GetStrOutOfDash(dataxx[6]) + " `vht:\"Presence:Presence" + strings.Title(dataxx[8]) + ",Criticality:Criticality" + strings.Title(dataxx[4]) + ",ProtocolExtensionID:ProtocolExtensionID" + GetStrOutOfDash2(dataxx[2]) + "\"`\n"
				x++
			}  else if len(dataxx) == 13 {
				if dataxx[6] == "OCTET" && dataxx[7] == "STRING" {
					str += "\t" + GetStrOutOfDash2(dataxx[2]) + " *OctetString `vht:\"Presence:Presence" + strings.Title(dataxx[11]) + ",Criticality:Criticality" + strings.Title(dataxx[4]) + ",ProtocolExtensionID:ProtocolExtensionID" + GetStrOutOfDash2(dataxx[2]) + ",Reference:" + GetStrOutOfDash(strings.Split(dataxx[9],")")[0]) +"\"`\n"
				} else {
					fmt.Println(in+1+x, "-", dataxx)
				}
				x++
			} else {
				fmt.Println(in+1+x, "-", dataxx)
				x++
			}
		}
		str += "}\n"
	} else {
		fmt.Println(in, "-", datax)
	}
	return str, in+x
}

func NGAPPROTOCOLIES(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	var x int
	if len(datax) == 4 {
		str += "\ntype ProtocolIESingleContainer" + GetStrOutOfDash(strings.Title(datax[0])) + " struct {\n"
		str += "\tList []" + GetStrOutOfDash(strings.Title(datax[0])) + " `vht:\"valueMin:0,valueMax:MaxProtocolIEs\"`\n"
		str += "}\n"

		str += "\ntype " + GetStrOutOfDash(strings.Title(datax[0])) + " struct {\n"
		str += "\tProtocolIEID ProtocolIEID\n"
		str += "\tCriticality Criticality `vht:\"Reference:ProtocolIEID\"`\n"
		str += "\tTypeValue " + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValue `vht:\"Reference:ProtocolIEID\"`\n"
		str += "}\n"
		str += "\nconst (\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" {
				if x == 1 {
					str += "\t" + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValueChoiceNothing int = 0\n"
				}
				break
			} else if dataxx[0] == "..." {
				x++
				continue
			} else if len(dataxx) == 10 {
				str += "\t" + GetStrOutOfDash(datax[0]) + "TypeValueChoice" + GetStrOutOfDash2(dataxx[2]) + " int = " + strconv.Itoa(x) + "\n"
				x++
			} else {
				fmt.Println(in+1+x, "-", dataxx)
				x++
			}
		}
		str += ")\n"
		str += "\ntype " + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValue struct {\n"
		str += "\tChoice int\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" || dataxx[0] == "..." {
				break
			} else if len(dataxx) == 10 {
				Presence := dataxx[8]
				Criticality := dataxx[4]
				str += "\t" + GetStrOutOfDash2(dataxx[2]) + " *" + GetStrOutOfDash(dataxx[6]) + " `vht:\"Presence:Presence" + strings.Title(Presence) + ",Criticality:Criticality" + strings.Title(Criticality) + ",ProtocolIEID:ProtocolIEID" + GetStrOutOfDash2(dataxx[2]) + "\"`\n"
				x++
			} else {
				fmt.Println(in+1+x, "-", dataxx)
				x++
			}
		}
		str += "}\n"
	} else {
		fmt.Println(datax)
	}
	return str, in+x
}

func NGAPPROTOCOLIES2(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	var x int
	if len(datax) == 4 {
		str += "\ntype ProtocolIEContainer" + GetStrOutOfDash(strings.Title(datax[0])) + " struct {\n"
		str += "\tList []" + GetStrOutOfDash(strings.Title(datax[0])) + " `vht:\"valueMin:0,valueMax:MaxPrivateIEs\"`\n"
		str += "}\n"

		str += "\ntype " + GetStrOutOfDash(strings.Title(datax[0])) + " struct {\n"
		str += "\tProtocolIEID ProtocolIEID\n"
		str += "\tCriticality Criticality `vht:\"Reference:ProtocolIEID\"`\n"
		str += "\tTypeValue " + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValue `vht:\"Reference:ProtocolIEID\"`\n"
		str += "}\n"
		str += "\nconst (\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" {
				if x == 1 {
					str += "\t" + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValueChoiceNothing int = 0\n"
				}
				break
			} else if dataxx[0] == "..." {
				x++
				continue
			} else if len(dataxx) == 10 {
				str += "\t" + GetStrOutOfDash(datax[0]) + "TypeValueChoice" + GetStrOutOfDash2(dataxx[2]) + " int = " + strconv.Itoa(x) + "\n"
				x++
			} else if len(dataxx) == 11 {
				if dataxx[6] == "OCTET" && dataxx[7] == "STRING" {
					str += "\t" + GetStrOutOfDash(datax[0]) + "TypeValueChoice" + GetStrOutOfDash2(dataxx[2]) + " int = " + strconv.Itoa(x) + "\n"
				} else {
					fmt.Println(in+1+x, "-1", dataxx)
				}
				x++
			} else {
				fmt.Println(in+1+x, "-1", dataxx)
				x++
			}
		}
		str += ")\n"
		str += "\ntype " + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValue struct {\n"
		str += "\tChoice int\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" || dataxx[0] == "..." {
				break
			} else if len(dataxx) == 10 {
				Presence := dataxx[8]
				Criticality := dataxx[4]
				str += "\t" + GetStrOutOfDash2(dataxx[2]) + " *" + GetStrOutOfDash(dataxx[6]) + " `vht:\"Presence:Presence" + strings.Title(Presence) + ",Criticality:Criticality" + strings.Title(Criticality) + ",ProtocolIEID:ProtocolIEID" + GetStrOutOfDash2(dataxx[2]) + "\"`\n"
				x++
			} else if len(dataxx) == 11 {
				if dataxx[6] == "OCTET" && dataxx[7] == "STRING" {
					Presence := dataxx[9]
					Criticality := dataxx[4]
					str += "\t" + GetStrOutOfDash2(dataxx[2]) + " *OctetString `vht:\"Presence:Presence" + strings.Title(Presence) + ",Criticality:Criticality" + strings.Title(Criticality) + ",ProtocolIEID:ProtocolIEID" + GetStrOutOfDash2(dataxx[2]) + "\"`\n"
				} else {
					fmt.Println(in+1+x, "-1", dataxx)
				}
				x++
			} else {
				fmt.Println(in+1+x, "-2", dataxx)
				x++
			}
		}
		str += "}\n"
	} else {
		fmt.Println(datax)
	}
	return str, in+x
}

func NGAPPRIVATEIES(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	var x int
	if len(datax) == 4 {
		str += "\ntype PrivateIEContainer" + GetStrOutOfDash(strings.Title(datax[0])) + " struct {\n"
		str += "\tList []" + GetStrOutOfDash(strings.Title(datax[0])) + " `vht:\"valueMin:0,valueMax:MaxProtocolIEs\"`\n"
		str += "}\n"

		str += "\ntype " + GetStrOutOfDash(strings.Title(datax[0])) + " struct {\n"
		str += "\tPrivateIEID PrivateIEID\n"
		str += "\tCriticality Criticality `vht:\"Reference:ProtocolIEID\"`\n"
		str += "\tTypeValue " + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValue `vht:\"Reference:ProtocolIEID\"`\n"
		str += "}\n"
		str += "\nconst (\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" {
				if x == 1 {
					str += "\t" + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValueChoiceNothing int = 0\n"
				}
				break
			} else if dataxx[0] == "..." {
				x++
				continue
			} else if len(dataxx) == 10 {
				str += "\t" + GetStrOutOfDash(datax[0]) + "TypeValueChoice" + GetStrOutOfDash2(dataxx[2]) + " int = " + strconv.Itoa(x) + "\n"
				x++
			} else {
				fmt.Println(in+1+x, "-1", dataxx)
				x++
			}
		}
		str += ")\n"
		str += "\ntype " + GetStrOutOfDash(strings.Title(datax[0])) + "TypeValue struct {\n"
		str += "\tChoice int\n"
		x = 0
		for {
			dataxx := strings.Fields(datas[in+1+x])
			if dataxx[0] == "}" || dataxx[0] == "..." {
				break
			} else if len(dataxx) == 10 {
				Presence := dataxx[8]
				Criticality := dataxx[4]
				str += "\t" + GetStrOutOfDash2(dataxx[2]) + " *" + GetStrOutOfDash(dataxx[6]) + " `vht:\"Presence:Presence" + strings.Title(Presence) + ",Criticality:Criticality" + strings.Title(Criticality) + ",ProtocolIEID:ProtocolIEID" + GetStrOutOfDash2(dataxx[2]) + "\"`\n"
				x++
			} else {
				fmt.Println(in+1+x, "-2", dataxx)
				x++
			}
		}
		str += "}\n"
	} else {
		fmt.Println(datax)
	}
	return str, in+x
}

func NGAPELEMENTARYPROCEDURE(datas []string, in int) (str string, out int) {
	return str, in
}

func GetMinMaxSEQUENCEOF(str string) (size []string) {
	strs := strings.Split(str, "(")
	strs = strings.Split(strs[2], ")")
	size = strings.Split(strs[0], "..")
	return
}

func GetStrOutOfDash(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 0; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}

// for name begin with "id-"
func GetStrOutOfDash2(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 1; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}
