package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	CheckPDU()
}

func CheckPDU() {
	data, _ := ioutil.ReadFile("PDUDescriptions.asn")
	datas := strings.Split(string(data), "\n")
	str := "package ngapEncode\n\nimport (\n\t\"errors\"\n\n\t\"ngap/ngapType\"\n)\n"
	str += "\nfunc CheckMandatory(pdu ngapType.NGAPPDU) error {\n" +
		"\tswitch pdu.Present {\n" +
		"\tcase ngapType.NGAPPDUPresentInitiatingMessage:\n" +
		"\t\tswitch pdu.InitiatingMessage.Value.Present {\n"
	for i := 300; i < len(datas); i++ {
		data0 := strings.Fields(datas[i])
		if len(data0) == 4 && data0[1] == "NGAP-ELEMENTARY-PROCEDURE" && data0[2] == "::=" {
			data1 := strings.Fields(datas[i+1])
			str += "\t\tcase ngapType.InitiatingMessagePresent"+GetStrOutOfDashC(data1[2])+":\n" +
				"\t\t\treturn nil\n"
		}
	}
	str += "\t\tdefault:\n\t\t\treturn errors.New(\"NGAPPDU: InitiatingMessage: Present: INVALID\")\n\t\t}\n"
	str += "\tcase ngapType.NGAPPDUPresentSuccessfulOutcome:\n\t\tswitch pdu.SuccessfulOutcome.Value.Present {\n"
	for i := 300; i < len(datas); i++ {
		data0 := strings.Fields(datas[i])
		if len(data0) == 4 && data0[1] == "NGAP-ELEMENTARY-PROCEDURE" && data0[2] == "::=" {
			data2 := strings.Fields(datas[i+2])
			if data2[0] == "SUCCESSFUL" {
				str += "\t\tcase ngapType.SuccessfulOutcomePresent"+GetStrOutOfDashC(data2[2])+":\n" +
					"\t\t\treturn nil\n"
			}
		}
	}
	str += "\t\tdefault:\n\t\t\treturn errors.New(\"NGAPPDU: SuccessfulOutcome: Present: INVALID\")\n\t\t}\n"
	str += "\tcase ngapType.NGAPPDUPresentUnsuccessfulOutcome:\n\t\tswitch pdu.UnsuccessfulOutcome.Value.Present {\n"
	for i := 300; i < len(datas); i++ {
		data0 := strings.Fields(datas[i])
		if len(data0) == 4 && data0[1] == "NGAP-ELEMENTARY-PROCEDURE" && data0[2] == "::=" {
			data3 := strings.Fields(datas[i+3])
			if data3[0] == "UNSUCCESSFUL" {
				str += "\t\tcase ngapType.UnsuccessfulOutcomePresent"+GetStrOutOfDashC(data3[2])+":\n" +
					"\t\t\treturn nil\n"
			}
		}
	}
	str += "\t\tdefault:\n\t\t\treturn errors.New(\"NGAPPDU: UnsuccessfulOutcome: Present: INVALID\")\n\t\t}\n"
	str += "\tdefault:\n\t\treturn errors.New(\"NGAPPDU: Present: INVALID\")\n\t}\n}\n"
	filePointer, _ := os.Create("../ngapEncode/CheckMandatory.go")
	filePointer.WriteString(str)
}


func GetStrOutOfDashC(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 0; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}

// for name begin with "id-"
func GetStrOutOfDashC2(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 1; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}
