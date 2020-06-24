package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// EncodePDUDescriptions()
	// EncodePDUContents()
	EncodeIEs()
}

//
// func EncodePDUDescriptions() () {
// 	data, _ := ioutil.ReadFile("PDUDescriptions.asn")
// 	datas := strings.Split(string(data), "\n")
// 	str := "package ngapEncode\n\n" +
// 		"import (\n" +
// 		"\t\"errors\"\n\n" +
// 		"\t\"ngap/ngapType\"\n" +
// 		")\n\n" +
// 		"func NGAPPDU (value ngapType.NGAPPDU, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tbinData, bitEnd = EncodePresent(value.Present, 3, binData, bitEnd)\n" +
// 		"\tswitch value.Present {\n" +
// 		"\tcase ngapType.NGAPPDUPresentInitiatingMessage:\n" +
// 		"\t\tif value.InitiatingMessage == nil {\n" +
// 		"\t\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: InitiatingMessage: NIL\")\n" +
// 		"\t\t}\n" +
// 		"\t\tbinData, bitEnd, err = InitiatingMessage(*value.InitiatingMessage, binData, bitEnd)\n" +
// 		"\t\tif err != nil {\n" +
// 		"\t\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: \" + err.Error())\n" +
// 		"\t\t}\n" +
// 		"\tcase ngapType.NGAPPDUPresentSuccessfulOutcome:\n" +
// 		"\t\tif value.SuccessfulOutcome == nil {\n" +
// 		"\t\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: SuccessfulOutcome: NIL\")\n" +
// 		"\t\t}\n" +
// 		"\t\tbinData, bitEnd, err = SuccessfulOutcome(*value.SuccessfulOutcome, binData, bitEnd)\n" +
// 		"\t\tif err != nil {\n" +
// 		"\t\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: \" + err.Error())\n" +
// 		"\t\t}\n" +
// 		"\tcase ngapType.NGAPPDUPresentUnsuccessfulOutcome:\n" +
// 		"\t\tif value.UnsuccessfulOutcome == nil {\n" +
// 		"\t\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: UnsuccessfulOutcome: NIL\")\n" +
// 		"\t\t}\n" +
// 		"\t\tbinData, bitEnd, err = UnsuccessfulOutcome(*value.UnsuccessfulOutcome, binData, bitEnd)\n" +
// 		"\t\tif err != nil {\n" +
// 		"\t\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: \" + err.Error())\n" +
// 		"\t\t}\n" +
// 		"\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"NGAPPDU: Present: INVALID\")\n" +
// 		"\t}\n\treturn binData, bitEnd, err\n" +
// 		"}\n"
// 	str += "\nfunc InitiatingMessage (value ngapType.InitiatingMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tswitch value.ProcedureCode.Value {\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) > 2 {
// 			if datax[2] == "::=" && datax[1] == "NGAP-ELEMENTARY-PROCEDURE" {
// 				dataxx := strings.Fields(datas[i+1])
// 				if dataxx[0] == "INITIATING" {
// 					dataxxx := strings.Fields(datas[i+2])
// 					if dataxxx[0] == "PROCEDURE" {
// 						str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 						str += "\t\tif value.Value.Present != ngapType.InitiatingMessageValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 						str += "\t\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: ProcedureCode: INVALID\")\n"
// 						str += "\t\t}\n"
// 					} else {
// 						dataxxx = strings.Fields(datas[i+3])
// 						if dataxxx[0] == "PROCEDURE" {
// 							str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 							str += "\t\tif value.Value.Present != ngapType.InitiatingMessageValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 							str += "\t\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: ProcedureCode: INVALID\")\n"
// 							str += "\t\t}\n"
// 						} else {
// 							dataxxx = strings.Fields(datas[i+4])
// 							if dataxxx[0] == "PROCEDURE" {
// 								str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 								str += "\t\tif value.Value.Present != ngapType.InitiatingMessageValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 								str += "\t\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: ProcedureCode: INVALID\")\n"
// 								str += "\t\t}\n"
// 							} else {
// 								fmt.Println(i, "-", datax)
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	str += "\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: ProcedureCode: INVALID\")\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd, err = ProcedureCode(value.ProcedureCode, binData, bitEnd)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\tvar binValue []byte\n" +
// 		"\tbinValue, bitEnd, err = InitiatingMessageValue(value.Value, binValue, 8)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"InitiatingMessage: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 		"}\n" +
// 		"\nfunc InitiatingMessageValue (value ngapType.InitiatingMessageValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tswitch value.Present {\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) > 2 {
// 			if datax[2] == "::=" && datax[1] == "NGAP-ELEMENTARY-PROCEDURE" {
// 				dataxx := strings.Fields(datas[i+1])
// 				if dataxx[0] == "INITIATING" {
// 					str += "\tcase ngapType.InitiatingMessageValuePresent" + GetStrOutOfDashE(dataxx[2]) + ":\n" +
// 						"\t\tif value." + GetStrOutOfDashE(dataxx[2]) + " == nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"InitiatingMessageValue: " + GetStrOutOfDashE(dataxx[2]) + ": NIL\")\n" +
// 						"\t\t}\n" +
// 						"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[2]) + "(*value." + GetStrOutOfDashE(dataxx[2]) + ", binData, bitEnd)\n" +
// 						"\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"InitiatingMessageValue: \" + err.Error())\n" +
// 						"\t\t}\n"
// 				}
// 			}
// 		}
// 	}
// 	str += "\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"InitiatingMessageValue: Present: INVALID\")\n" +
// 		"\t}\n" +
// 		"\treturn binData, bitEnd, err\n" +
// 		"}\n"
//
// 	str += "\nfunc SuccessfulOutcome (value ngapType.SuccessfulOutcome, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tswitch value.ProcedureCode.Value {\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) > 2 {
// 			if datax[2] == "::=" && datax[1] == "NGAP-ELEMENTARY-PROCEDURE" {
// 				dataxx := strings.Fields(datas[i+1])
// 				if dataxx[0] == "SUCCESSFUL" {
// 					dataxxx := strings.Fields(datas[i+2])
// 					if dataxxx[0] == "PROCEDURE" {
// 						str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 						str += "\t\tif value.Value.Present != ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 						str += "\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 						str += "\t\t}\n"
// 					} else {
// 						dataxxx = strings.Fields(datas[i+3])
// 						if dataxxx[0] == "PROCEDURE" {
// 							str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 							str += "\t\tif value.Value.Present != ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 							str += "\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 							str += "\t\t}\n"
// 						} else {
// 							dataxxx = strings.Fields(datas[i+4])
// 							if dataxxx[0] == "PROCEDURE" {
// 								str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 								str += "\t\tif value.Value.Present != ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 								str += "\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 								str += "\t\t}\n"
// 							} else {
// 								fmt.Println(i, "-", datax)
// 							}
// 						}
// 					}
// 				} else {
// 					dataxx := strings.Fields(datas[i+2])
// 					if dataxx[0] == "SUCCESSFUL" {
// 						dataxxx := strings.Fields(datas[i+3])
// 						if dataxxx[0] == "PROCEDURE" {
// 							str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 							str += "\t\tif value.Value.Present != ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 							str += "\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 							str += "\t\t}\n"
// 						} else {
// 							dataxxx = strings.Fields(datas[i+4])
// 							if dataxxx[0] == "PROCEDURE" {
// 								str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 								str += "\t\tif value.Value.Present != ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 								str += "\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 								str += "\t\t}\n"
// 							} else {
// 								dataxxx = strings.Fields(datas[i+5])
// 								if dataxxx[0] == "PROCEDURE" {
// 									str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 									str += "\t\tif value.Value.Present != ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 									str += "\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 									str += "\t\t}\n"
// 								} else {
// 									fmt.Println(i, "-", datax)
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	str += "\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: ProcedureCode: INVALID\")\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd, err = ProcedureCode(value.ProcedureCode, binData, bitEnd)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\tvar binValue []byte\n" +
// 		"\tbinValue, bitEnd, err = SuccessfulOutcomeValue(value.Value, binValue, 8)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcome: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 		"}\n" +
// 		"\nfunc SuccessfulOutcomeValue (value ngapType.SuccessfulOutcomeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tswitch value.Present {\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) > 2 {
// 			if datax[2] == "::=" && datax[1] == "NGAP-ELEMENTARY-PROCEDURE" {
// 				dataxx := strings.Fields(datas[i+1])
// 				if dataxx[0] == "SUCCESSFUL" {
// 					str += "\tcase ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + ":\n" +
// 						"\t\tif value." + GetStrOutOfDashE(dataxx[2]) + " == nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcomeValue: " + GetStrOutOfDashE(dataxx[2]) + ": NIL\")\n" +
// 						"\t\t}\n" +
// 						"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[2]) + "(*value." + GetStrOutOfDashE(dataxx[2]) + ", binData, bitEnd)\n" +
// 						"\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcomeValue: \" + err.Error())\n" +
// 						"\t\t}\n"
// 				} else {
// 					dataxx := strings.Fields(datas[i+2])
// 					if dataxx[0] == "SUCCESSFUL" {
// 						str += "\tcase ngapType.SuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + ":\n" +
// 							"\t\tif value." + GetStrOutOfDashE(dataxx[2]) + " == nil {\n" +
// 							"\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcomeValue: " + GetStrOutOfDashE(dataxx[2]) + ": NIL\")\n" +
// 							"\t\t}\n" +
// 							"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[2]) + "(*value." + GetStrOutOfDashE(dataxx[2]) + ", binData, bitEnd)\n" +
// 							"\t\tif err != nil {\n" +
// 							"\t\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcomeValue: \" + err.Error())\n" +
// 							"\t\t}\n"
// 					}
// 				}
// 			}
// 		}
// 	}
// 	str += "\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"SuccessfulOutcomeValue: Present: INVALID\")\n" +
// 		"\t}\n" +
// 		"\treturn binData, bitEnd, err\n" +
// 		"}\n"
//
// 	str += "\nfunc UnsuccessfulOutcome (value ngapType.UnsuccessfulOutcome, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tswitch value.ProcedureCode.Value {\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) > 2 {
// 			if datax[2] == "::=" && datax[1] == "NGAP-ELEMENTARY-PROCEDURE" {
// 				dataxx := strings.Fields(datas[i+1])
// 				if dataxx[0] == "UNSUCCESSFUL" {
// 					dataxxx := strings.Fields(datas[i+2])
// 					if dataxxx[0] == "PROCEDURE" {
// 						str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 						str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 						str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 						str += "\t\t}\n"
// 					} else {
// 						dataxxx = strings.Fields(datas[i+3])
// 						if dataxxx[0] == "PROCEDURE" {
// 							str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 							str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 							str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 							str += "\t\t}\n"
// 						} else {
// 							dataxxx = strings.Fields(datas[i+4])
// 							if dataxxx[0] == "PROCEDURE" {
// 								str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 								str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 								str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 								str += "\t\t}\n"
// 							} else {
// 								fmt.Println(i, "-", datax)
// 							}
// 						}
// 					}
// 				} else {
// 					dataxx := strings.Fields(datas[i+2])
// 					if dataxx[0] == "UNSUCCESSFUL" {
// 						dataxxx := strings.Fields(datas[i+3])
// 						if dataxxx[0] == "PROCEDURE" {
// 							str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 							str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 							str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 							str += "\t\t}\n"
// 						} else {
// 							dataxxx = strings.Fields(datas[i+4])
// 							if dataxxx[0] == "PROCEDURE" {
// 								str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 								str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 								str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 								str += "\t\t}\n"
// 							} else {
// 								dataxxx = strings.Fields(datas[i+5])
// 								if dataxxx[0] == "PROCEDURE" {
// 									str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 									str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 									str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 									str += "\t\t}\n"
// 								} else {
// 									fmt.Println(i, "-", datax)
// 								}
// 							}
// 						}
// 					} else {
// 						dataxx := strings.Fields(datas[i+3])
// 						if dataxx[0] == "UNSUCCESSFUL" {
// 							dataxxx := strings.Fields(datas[i+4])
// 							if dataxxx[0] == "PROCEDURE" {
// 								str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 								str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 								str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 								str += "\t\t}\n"
// 							} else {
// 								dataxxx = strings.Fields(datas[i+5])
// 								if dataxxx[0] == "PROCEDURE" {
// 									str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 									str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 									str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 									str += "\t\t}\n"
// 								} else {
// 									dataxxx = strings.Fields(datas[i+6])
// 									if dataxxx[0] == "PROCEDURE" {
// 										str += "\tcase ngapType.ProcedureCode" + GetStrOutOfDashE2(dataxxx[2]) + ":\n"
// 										str += "\t\tif value.Value.Present != ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + " {\n"
// 										str += "\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n"
// 										str += "\t\t}\n"
// 									} else {
// 										fmt.Println(i, "-", datax)
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	str += "\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: ProcedureCode: INVALID\")\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd, err = ProcedureCode(value.ProcedureCode, binData, bitEnd)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\tvar binValue []byte\n" +
// 		"\tbinValue, bitEnd, err = UnsuccessfulOutcomeValue(value.Value, binValue, 8)\n" +
// 		"\tif err != nil {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcome: \" + err.Error())\n" +
// 		"\t}\n" +
// 		"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 		"}\n" +
// 		"\nfunc UnsuccessfulOutcomeValue (value ngapType.UnsuccessfulOutcomeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tvar err error\n" +
// 		"\tswitch value.Present {\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) > 2 {
// 			if datax[2] == "::=" && datax[1] == "NGAP-ELEMENTARY-PROCEDURE" {
// 				dataxx := strings.Fields(datas[i+1])
// 				if dataxx[0] == "UNSUCCESSFUL" {
// 					str += "\tcase ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + ":\n" +
// 						"\t\tif value." + GetStrOutOfDashE(dataxx[2]) + " == nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: " + GetStrOutOfDashE(dataxx[2]) + ": NIL\")\n" +
// 						"\t\t}\n" +
// 						"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[2]) + "(*value." + GetStrOutOfDashE(dataxx[2]) + ", binData, bitEnd)\n" +
// 						"\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: \" + err.Error())\n" +
// 						"\t\t}\n"
// 				} else {
// 					dataxx := strings.Fields(datas[i+2])
// 					if dataxx[0] == "UNSUCCESSFUL" {
// 						str += "\tcase ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + ":\n" +
// 							"\t\tif value." + GetStrOutOfDashE(dataxx[2]) + " == nil {\n" +
// 							"\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: " + GetStrOutOfDashE(dataxx[2]) + ": NIL\")\n" +
// 							"\t\t}\n" +
// 							"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[2]) + "(*value." + GetStrOutOfDashE(dataxx[2]) + ", binData, bitEnd)\n" +
// 							"\t\tif err != nil {\n" +
// 							"\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: \" + err.Error())\n" +
// 							"\t\t}\n"
// 					} else {
// 						dataxx := strings.Fields(datas[i+3])
// 						if dataxx[0] == "UNSUCCESSFUL" {
// 							str += "\tcase ngapType.UnsuccessfulOutcomeValuePresent" + GetStrOutOfDashE(dataxx[2]) + ":\n" +
// 								"\t\tif value." + GetStrOutOfDashE(dataxx[2]) + " == nil {\n" +
// 								"\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: " + GetStrOutOfDashE(dataxx[2]) + ": NIL\")\n" +
// 								"\t\t}\n" +
// 								"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[2]) + "(*value." + GetStrOutOfDashE(dataxx[2]) + ", binData, bitEnd)\n" +
// 								"\t\tif err != nil {\n" +
// 								"\t\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: \" + err.Error())\n" +
// 								"\t\t}\n"
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	str += "\tdefault:\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"UnsuccessfulOutcomeValue: Present: INVALID\")\n" +
// 		"\t}\n" +
// 		"\treturn binData, bitEnd, err\n" +
// 		"}\n"
//
// 	filePointer, _ := os.Create("../ngapEncode/PDUDescriptions.go")
// 	filePointer.WriteString(str)
// }
//
// func EncodePDUContents() {
// 	data, _ := ioutil.ReadFile("PDUContents.asn")
// 	datas := strings.Split(string(data), "\n")
// 	str1 := "package ngapEncode\n\nimport (\n\t\"errors\"\n\n\t\"ngap/ngapType\"\n)\n"
// 	str2 := "package ngapEncode\n\nimport (\n\t\"errors\"\n\t\"strconv\"\n\n\t\"ngap/ngapType\"\n)\n"
// 	str3 := "package ngapEncode\n\nimport (\n\t\"errors\"\n\t\"strconv\"\n\n\t\"ngap/ngapType\"\n)\n"
// 	str4 := "package ngapEncode\n\nimport (\n\t\"errors\"\n\t\"strconv\"\n\n\t\"ngap/ngapType\"\n)\n"
// 	for i := 300; i < len(datas); i++ {
// 		datax := strings.Fields(datas[i])
// 		if len(datax) == 4 && datax[1] == "::=" && datax[2] == "SEQUENCE" {
// 			dataxs := strings.Fields(datas[i+1])
// 			if dataxs[0] == "protocolIEs" {
// 				str1 += "\nfunc " + datax[0] + "(value ngapType." + datax[0] + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tbinData, bitEnd, err = ProtocolIEContainer" + datax[0] + "IEs(value.ProtocolIEs, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + ": \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\treturn binData, bitEnd, err\n" +
// 					"}\n"
// 				str2 += "\nfunc ProtocolIEContainer" + datax[0] + "IEs(value ngapType.ProtocolIEContainer" + datax[0] + "IEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tnumIEs := len(value.List)\n" +
// 					"\tbinData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"ProtocolIEContainer" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\tfor i := 0; i < numIEs; i++ {\n" +
// 					"\t\tvar binIEs []byte\n" +
// 					"\t\tbinIEs, bitEnd, err = " + datax[0] + "IEs(value.List[i], binIEs, bitEnd)\n" +
// 					"\t\tif err != nil {\n" +
// 					"\t\t\treturn binData, bitEnd, errors.New(\"ProtocolIEContainer" + datax[0] + "IEs: [\" + strconv.Itoa(i) + \"]: \" + err.Error())\n" +
// 					"\t\t}\n" +
// 					"\t\tbinData = append(binData, binIEs...)\n" +
// 					"\t}\n" +
// 					"\treturn binData, bitEnd, err\n" +
// 					"}\n"
// 				str3 += "\nfunc " + datax[0] + "IEs(value ngapType." + datax[0] + "IEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tswitch value.ProtocolIEID.Value {\n"
// 				x := 0
// 				for {
// 					dataxx := strings.Fields(datas[i+6+x])
// 					if dataxx[0] == "..." {
// 						break
// 					}
// 					str3 += "\tcase ngapType.ProtocolIEID" + GetStrOutOfDashE2(dataxx[2]) + ":\n" +
// 						"\t\tif value.TypeValue.Present != ngapType." + datax[0] + "IEsTypeValuePresent" + GetStrOutOfDashE2(dataxx[2]) + " {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: ProtocolIEID: INVALID\")\n" +
// 						"\t\t}\n"
// 					x++
// 				}
// 				str3 += "\tdefault:\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: ProtocolIEID: INVALID\")\n" +
// 					"\t}\n" +
// 					"\tbinData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\tvar binValue []byte\n" +
// 					"\tbinValue, bitEnd, err = " + datax[0] + "IEsTypeValue(value.TypeValue, binValue, 8)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 					"}\n"
// 				str4 += "\nfunc " + datax[0] + "IEsTypeValue(value ngapType." + datax[0] + "IEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tswitch value.Present {\n"
// 				x = 0
// 				for {
// 					dataxx := strings.Fields(datas[i+6+x])
// 					if dataxx[0] == "..." {
// 						break
// 					}
// 					str4 += "\tcase ngapType." + datax[0] + "IEsTypeValuePresent" + GetStrOutOfDashE2(dataxx[2]) + ":\n" +
// 						"\t\tif value." + GetStrOutOfDashE2(dataxx[2]) + " == nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEsTypeValue: " + GetStrOutOfDashE2(dataxx[2]) + ": NIL\")\n" +
// 						"\t\t}\n"
// 					if len(dataxx) == 11 {
// 						str4 += "\t\tbinData, bitEnd, err = NGAPMessage(*value.NGAPMessage, binData, bitEnd)\n"
// 					} else {
// 						str4 += "\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[6]) + "(*value." + GetStrOutOfDashE2(dataxx[2]) + ", binData, bitEnd)\n"
// 					}
// 					str4 += "\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEsTypeValue: \" + err.Error())\n" +
// 						"\t\t}\n"
// 					x++
// 				}
// 				str4 += "\tdefault:\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEsTypeValue: Present: INVALID\")\n" +
// 					"\t}\n" +
// 					"\treturn binData, bitEnd, err\n" +
// 					"}\n"
// 			} else if dataxs[0] == "privateIEs" {
// 				str1 += "\nfunc " + datax[0] + "(value ngapType." + datax[0] + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tbinData, bitEnd, err = PrivateIEContainer" + datax[0] + "IEs(value.PrivateIEs, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + ": \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\treturn binData, bitEnd, err\n" +
// 					"}\n"
// 				str2 += "\nfunc PrivateIEContainer" + datax[0] + "IEs(value ngapType.PrivateIEContainer" + datax[0] + "IEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tnumIEs := len(value.List)\n" +
// 					"\tbinData, bitEnd, err = EncodeNumberIEsOfPrivateIEContainer(numIEs, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"PrivateIEContainer" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\tfor i := 0; i < numIEs; i++ {\n" +
// 					"\t\tvar binIEs []byte\n" +
// 					"\t\tbinIEs, bitEnd, err = " + datax[0] + "IEs(value.List[i], binIEs, bitEnd)\n" +
// 					"\t\tif err != nil {\n" +
// 					"\t\t\treturn binData, bitEnd, errors.New(\"PrivateIEContainer" + datax[0] + "IEs: [\" + strconv.Itoa(i) + \"]: \" + err.Error())\n" +
// 					"\t\t}\n" +
// 					"\t\tbinData = append(binData, binIEs...)\n" +
// 					"\t}\n" +
// 					"\treturn binData, bitEnd, err\n" +
// 					"}\n"
// 				str3 += "\nfunc " + datax[0] + "IEs(value ngapType." + datax[0] + "IEs, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\t// switch value.PrivateIEID.Value {\n"
// 				x := 0
// 				for {
// 					dataxx := strings.Fields(datas[i+6+x])
// 					if dataxx[0] == "..." {
// 						break
// 					}
// 					str3 += "\t// case ngapType.PrivateIEID" + GetStrOutOfDashE2(dataxx[2]) + ":\n" +
// 						"\t\t// if value.TypeValue.Present != ngapType." + datax[0] + "IEsTypeValuePresent" + GetStrOutOfDashE2(dataxx[2]) + " {\n" +
// 						"\t\t\t// return binData, bitEnd, errors.New(\"" + datax[0] + "IEs: PrivateIEID: INVALID\")\n" +
// 						"\t\t// }\n"
// 					x++
// 				}
// 				str3 += "\t// default:\n" +
// 					"\t\t// return binData, bitEnd, errors.New(\"" + datax[0] + "IEs: PrivateIEID: INVALID\")\n" +
// 					"\t// }\n" +
// 					"\tbinData, bitEnd, err = PrivateIEID(value.PrivateIEID, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\tvar binValue []byte\n" +
// 					"\tbinValue, bitEnd, err = " + datax[0] + "IEsTypeValue(value.TypeValue, binValue, 8)\n" +
// 					"\tif err != nil {\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEs: \" + err.Error())\n" +
// 					"\t}\n" +
// 					"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 					"}\n"
// 				str4 += "\nfunc " + datax[0] + "IEsTypeValue(value ngapType." + datax[0] + "IEsTypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 					"\tvar err error\n" +
// 					"\tswitch value.Present {\n"
// 				x = 0
// 				for {
// 					dataxx := strings.Fields(datas[i+6+x])
// 					if dataxx[0] == "..." {
// 						break
// 					}
// 					str4 += "\tcase ngapType." + datax[0] + "IEsTypeValuePresent" + GetStrOutOfDashE2(dataxx[2]) + ":\n" +
// 						"\t\tif value." + GetStrOutOfDashE2(dataxx[2]) + " == nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEsTypeValue: " + GetStrOutOfDashE2(dataxx[2]) + ": NIL\")\n" +
// 						"\t\t}\n"
// 					if len(dataxx) == 11 {
// 						str4 += "\t\tbinData, bitEnd, err = NGAPMessage(*value.NGAPMessage, binData, bitEnd)\n"
// 					} else {
// 						str4 += "\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[6]) + "(*value." + GetStrOutOfDashE2(dataxx[2]) + ", binData, bitEnd)\n"
// 					}
// 					str4 += "\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEsTypeValue: \" + err.Error())\n" +
// 						"\t\t}\n"
// 					x++
// 				}
// 				str4 += "\tdefault:\n" +
// 					"\t\treturn binData, bitEnd, errors.New(\"" + datax[0] + "IEsTypeValue: Present: INVALID\")\n" +
// 					"\t}\n" +
// 					"\treturn binData, bitEnd, err\n" +
// 					"}\n"
// 			}
// 		}
// 	}
// 	filePointer, _ := os.Create("../ngapEncode/PDUContents.go")
// 	filePointer.WriteString(str1)
// 	filePointer, _ = os.Create("../ngapEncode/PDUContentss.go")
// 	filePointer.WriteString(str2)
// 	filePointer, _ = os.Create("../ngapEncode/PDUContentsss.go")
// 	filePointer.WriteString(str3)
// 	filePointer, _ = os.Create("../ngapEncode/PDUContentssss.go")
// 	filePointer.WriteString(str4)
// }

func EncodeIEs() {
	data, _ := ioutil.ReadFile("IEs.asn")
	datas := strings.Split(string(data), "\n")
	str := "package ngapEncode\n\nimport (\n\t\"errors\"\n\t\"strconv\"\n\n\t\"ngap/ngapType\"\n)\n"
	for i := 120; i < len(datas); i++ {
		datax := strings.Fields(datas[i])
		if len(datax) > 2 {
			// if datax[1] == "::=" && datax[2] == "ENUMERATED" {
			// 	strx, o := EncodeENUMERATED(datas, i)
			// 	str += strx
			// 	i = o
			// }
			// if datax[1] == "::=" && datax[2] == "INTEGER" {
			// 	strx, o := EncodeINTEGER(datas, i)
			// 	str += strx
			// 	i = o
			// }
			// if datax[1] == "::=" && datax[2] == "PrintableString" {
			// 	strx, o := EncodePrintableString(datas, i)
			// 	str += strx
			// 	i = o
			// }
			// if datax[1] == "::=" && datax[2] == "OCTET" {
			// 	strx, o := EncodeOCTET(datas, i)
			// 	str += strx
			// 	i = o
			// }
			// if datax[1] == "::=" && datax[2] == "BIT" {
			// 	strx, o := EncodeBIT(datas, i)
			// 	str += strx
			// 	i = o
			// }
			// if datax[1] == "::=" && datax[2] == "OBJECT" {
			// 	strx, o := EncodeOBJECT(datas, i)
			// 	str += strx
			// 	i = o
			// }
			if datax[1] == "::=" && datax[2] == "CHOICE" {
				strx, o := EncodeCHOICE(datas, i)
				str += strx
				i = o
			}
			// if datax[1] == "::=" && datax[2] == "SEQUENCE" {
			// 	strx, o := EncodeSEQUENCE(datas, i)
			// 	str += strx
			// 	i = o
			// }
		}
	}
	// filePointer, _ := os.Create("../ngapEncode/IEsPresent.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsPresentExtension.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsPresentExtensions.go")
	filePointer, _ := os.Create("../ngapEncode/IEsPresentExtensionss.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsSequence.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsSequenceExtension.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsSequenceExtensions.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsSequenceExtensionss.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsSequenceOf.go")
	// filePointer, _ := os.Create("../ngapEncode/IEsCommon.go")
	filePointer.WriteString(str)
}

//
// func EncodeENUMERATED(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	ext := ""
// 	if len(datax) == 3 {
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + "(value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tvar err error" +
// 			"\tbinData, bitEnd, err = EncodeEnumeratedNo(value.Value, binData, bitEnd)\n" +
// 			"\tif err != nil {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \"+ err.Error())\n" +
// 			"\t}\n" +
// 			"\treturn binData, bitEnd, err\n" +
// 			"}\n"
// 		return str, in
// 	} else if len(datax) == 4 && datax[3] == "{" {
// 		x, xs := 0, 0
// 		for {
// 			dataxx := strings.Fields(datas[in+1+x])
// 			if dataxx[0] == "..." {
// 				ext = "Ext"
// 				break
// 			} else if dataxx[0] == "}" {
// 				break
// 			}
// 			x++
// 		}
// 		if x < 1 {
// 			xs = 1
// 		} else if x < 3 {
// 			xs = 2
// 		} else if x < 7 {
// 			xs = 3
// 		} else if x < 15 {
// 			xs = 4
// 		} else if x < 31 {
// 			xs = 5
// 		} else if x < 63 {
// 			xs = 6
// 		} else {
// 			x = 7
// 		}
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tif value.Value < 0 || value.Value > " + strconv.Itoa(x) + " {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": Enumerated: valueMin:0,valueMax:" + strconv.Itoa(x) + "\")\n" +
// 			"\t}\n" +
// 			"\tbinData, bitEnd = EncodeEnumerated" + ext + "(value.Value, " + strconv.Itoa(xs) + ", binData, bitEnd)\n" +
// 			"\treturn binData, bitEnd, nil\n" +
// 			"}\n"
// 		return str, in + x
// 	} else {
// 		x, xs := 0, 0
// 		for {
// 			if datax[x+4] == "..." {
// 				ext = "Ext"
// 				break
// 			} else if datax[x+4] == "}" {
// 				break
// 			}
// 			x++
// 		}
// 		if x < 1 {
// 			xs = 1
// 		} else if x < 3 {
// 			xs = 2
// 		} else if x < 7 {
// 			xs = 3
// 		} else if x < 15 {
// 			xs = 4
// 		} else if x < 31 {
// 			xs = 5
// 		} else if x < 63 {
// 			xs = 6
// 		} else {
// 			x = 7
// 		}
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tif value.Value < 0 || value.Value > " + strconv.Itoa(x) + " {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": Enumerated: valueMin:0,valueMax:" + strconv.Itoa(x) + "\")\n" +
// 			"\t}\n" +
// 			"\tbinData, bitEnd = EncodeEnumerated" + ext + "(value.Value, " + strconv.Itoa(xs) + ", binData, bitEnd)\n" +
// 			"\treturn binData, bitEnd, nil\n" +
// 			"}\n"
// 		return str, in
// 	}
// }
//
// func EncodeINTEGER(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	if len(datax) == 3 {
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + "(value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tvar err error" +
// 			"\tbinData, bitEnd, err = EncodeIntegerNo(value.Value, binData, bitEnd)\n" +
// 			"\tif err != nil {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \"+ err.Error())\n" +
// 			"\t}\n" +
// 			"\treturn binData, bitEnd, err\n" +
// 			"}\n"
// 		return str, in
// 	} else if len(datax) == 4 {
// 		min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[1], strings.Split(strings.Split(datax[3], "..")[1], ")")[0]
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tif value.Value < " + min + " || value.Value > " + max + " {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": Integer: valueMin:" + min + ",valueMax:" + max + "\")\n" +
// 			"\t}\n" +
// 			"\tbinData, bitEnd = EncodeInteger(value.Value, 0, binData, bitEnd)\n" +
// 			"\treturn binData, bitEnd, nil\n" +
// 			"}\n"
// 		return str, in
// 	} else if len(datax) == 5 && datax[4] == "...)" {
// 		min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[1], strings.Split(strings.Split(datax[3], "..")[1], ",")[0]
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tif value.Value < " + min + " || value.Value > " + max + " {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": Integer: valueMin:" + min + ",valueMax:" + max + "\")\n" +
// 			"\t}\n" +
// 			"\tbinData, bitEnd = EncodeIntegerExt(value.Value, 0, binData, bitEnd)\n" +
// 			"\treturn binData, bitEnd, nil\n" +
// 			"}\n"
// 		return str, in
// 	}
// 	return str, in
// }
//
// func EncodePrintableString(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[2], strings.Split(strings.Split(datax[3], "..")[1], ",")[0]
// 	str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tif len(value.Value) < " + min + " || len(value.Value) > " + max + " {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": PrintableString: sizeMin:" + min + ",sizeMax:" + max + "\")\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd = EncodePrintableString(value.Value," + min + ", " + max + ", binData, bitEnd)\n" +
// 		"\treturn binData, bitEnd, nil\n" +
// 		"}\n"
// 	return str, in
// }
//
// func EncodeOCTET(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	if len(datax) == 4 {
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + "(value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tvar err error\n" +
// 			"\tbinData, bitEnd, err = EncodeOctetStringNo(value.Value, binData, bitEnd)\n" +
// 			"\tif err != nil {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \"+ err.Error())\n" +
// 			"\t}\n" +
// 			"\treturn binData, bitEnd, err\n" +
// 			"}\n"
// 		return str, in
// 	} else if len(datax) == 5 {
// 		dataxx := strings.Split(datax[4], "..")
// 		if len(dataxx) < 2 {
// 			mimax := strings.Split(strings.Split(datax[4], ")")[0], "(")[2]
// 			str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 				"\tif len(value.Value) != " + mimax + " {\n" +
// 				"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": OctetString: sizeMin:" + mimax + ",sizeMax:" + mimax + "\")\n" +
// 				"\t}\n" +
// 				"\tbinData, bitEnd = EncodeOctetStringFix(value.Value, binData, bitEnd)\n" +
// 				"\treturn binData, bitEnd, nil\n" +
// 				"}\n"
// 			return str, in
// 		} else {
// 			min, max := strings.Split(dataxx[0], "(")[2], strings.Split(dataxx[1], ")")[0]
// 			str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 				"\tif len(value.Value) < " + min + " || len(value.Value) > " + max + " {\n" +
// 				"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": OctetString: sizeMin:" + min + ",sizeMax:" + max + "\")\n" +
// 				"\t}\n" +
// 				"\tbinData, bitEnd = EncodeOctetStringExt(value.Value," + min + ", " + max + ", binData, bitEnd)\n" +
// 				"\treturn binData, bitEnd, nil\n" +
// 				"}\n"
// 			return str, in
// 		}
//
// 	}
// 	return str, in
// }
//
// func EncodeBIT(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	if len(datax) == 4 {
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + "(value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tvar err error\n" +
// 			"\tbinData, bitEnd, err = EncodeBitStringNo(value.Value, binData, bitEnd)\n" +
// 			"\tif err != nil {\n" +
// 			"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \"+ err.Error())\n" +
// 			"\t}\n" +
// 			"\treturn binData, bitEnd, err\n" +
// 			"}\n"
// 		return str, in
// 	} else if len(datax) == 5 {
// 		dataxx := strings.Split(datax[4], "..")
// 		if len(dataxx) < 2 {
// 			mimax := strings.Split(strings.Split(datax[4], ")")[0], "(")[2]
// 			mimmax, _ := strconv.Atoi(mimax)
// 			mimmax = (mimmax-1)/8 + 1
// 			str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 				"\tif value.Value.BitLength != " + mimax + " || len(value.Value.Bytes) != " + strconv.Itoa(mimmax) + " {\n" +
// 				"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": BitString: sizeMin:" + mimax + ",sizeMax:" + mimax + "\")\n" +
// 				"\t}\n" +
// 				"\tbinData, bitEnd = EncodeBitStringFix(value.Value, binData, bitEnd)\n" +
// 				"\treturn binData, bitEnd, nil\n" +
// 				"}\n"
// 			return str, in
// 		} else {
// 			min, max := strings.Split(dataxx[0], "(")[2], strings.Split(dataxx[1], ")")[0]
// 			str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 				"\tif value.Value.BitLength < " + min + " || value.Value.BitLength > " + max + " || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {\n" +
// 				"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": BitString: sizeMin:" + min + ",sizeMax:" + max + "\")\n" +
// 				"\t}\n" +
// 				"\tbinData, bitEnd = EncodeBitStringExt(value.Value, binData, bitEnd)\n" +
// 				"\treturn binData, bitEnd, nil\n" +
// 				"}\n"
// 			return str, in
// 		}
// 	} else if len(datax) == 6 {
// 		dataxx := strings.Split(datax[4], "..")
// 		if len(dataxx) < 2 {
// 			mimax := strings.Split(strings.Split(datax[4], ",")[0], "(")[2]
// 			mimmax, _ := strconv.Atoi(mimax)
// 			mimmax = (mimmax-1)/8 + 1
// 			str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 				"\tif value.Value.BitLength != " + mimax + " || len(value.Value.Bytes) != " + strconv.Itoa(mimmax) + " {\n" +
// 				"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": BitString: sizeMin:" + mimax + ",sizeMax:" + mimax + "\")\n" +
// 				"\t}\n" +
// 				"\tbinData, bitEnd = EncodeBitStringFixExx(value.Value, binData, bitEnd)\n" +
// 				"\treturn binData, bitEnd, nil\n" +
// 				"}\n"
// 			return str, in
// 		} else {
// 			min, max := strings.Split(dataxx[0], "(")[2], strings.Split(dataxx[1], ",")[0]
// 			str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 				"\tif value.Value.BitLength < " + min + " || value.Value.BitLength > " + max + " || len(value.Value.Bytes) != int((value.Value.BitLength-1)/8+1) {\n" +
// 				"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": BitString: sizeMin:" + min + ",sizeMax:" + max + "\")\n" +
// 				"\t}\n" +
// 				"\tbinData, bitEnd = EncodeBitStringExtExx(value.Value, binData, bitEnd)\n" +
// 				"\treturn binData, bitEnd, nil\n" +
// 				"}\n"
// 			return str, in
// 		}
// 	}
// 	return str, in
// }
//
// func EncodeOBJECT(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[2], strings.Split(strings.Split(datax[3], "..")[1], ",")[0]
// 	str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		"\tif len(value.Value) < " + min + " || len(value.Value) > " + max + " {\n" +
// 		"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": PrintableString: sizeMin:" + min + ",sizeMax:" + max + "\")\n" +
// 		"\t}\n" +
// 		"\tbinData, bitEnd = EncodeObjectIdentifier(value.Value," + min + ", " + max + ", binData, bitEnd)\n" +
// 		"\treturn binData, bitEnd, nil\n" +
// 		"}\n"
// 	return str, in
// }

func EncodeCHOICE(datas []string, in int) (str string, out int) {
	datax := strings.Fields(datas[in])
	xx, xs := 0, 0
	for {
		dataxx := strings.Fields(datas[in+1+xx])
		if len(dataxx) == 1 {
			break
		}
		xx++
	}
	if xx < 3 {
		xs = 1
	} else if xx < 5 {
		xs = 2
	} else if xx < 9 {
		xs = 3
	} else if xx < 17 {
		xs = 4
	} else if xx < 33 {
		xs = 5
	} else if xx < 65 {
		xs = 6
	} else if xx < 129 {
		xs = 7
	}
	str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
		"\tvar err error\n" +
		"\tbinData, bitEnd = EncodePresent(value.Present, "+strconv.Itoa(xs)+", binData, bitEnd)\n" +
		"\tswitch value.Present {\n"
	choiceExt := ""
	for i := 0; i < xx; i++ {
		dataxx := strings.Fields(datas[in+1+i])
		if dataxx[0] == "choice-Extensions" {
			str += "\tcase ngapType." + GetStrOutOfDashE(datax[0]) + "Present" + GetStrOutOfDashE(dataxx[0]) + ":\n" +
				"\t\tif value." + GetStrOutOfDashE(dataxx[0]) + " == nil {\n" +
				"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": NIL\")\n" +
				"\t\t}\n"
			chExt := GetStrOutOfDashE(dataxx[1]) + GetStrOutOfDashE(strings.Split(strings.Split(dataxx[3],"{")[1], "}")[0])
			str += "\t\tbinData, bitEnd, err = " + chExt + "(*value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
				"\t\tif err != nil {\n" +
				"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
				"\t\t}\n"
			choiceExt = GetStrOutOfDashE(strings.Split(strings.Split(dataxx[3],"{")[1], "}")[0])
		} else if dataxx[1] == "BIT" {
			if len(dataxx) == 4 {
				dataxs := strings.Split(dataxx[3], "..")
				if len(dataxs) < 2 {
					str += "\tcase ngapType." + GetStrOutOfDashE(datax[0]) + "Present" + GetStrOutOfDashE(dataxx[0]) + ":\n" +
						"\t\tif value." + GetStrOutOfDashE(dataxx[0]) + " == nil {\n" +
						"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": NIL\")\n" +
						"\t\t}\n"
					mimax := strings.Split(strings.Split(dataxx[3], "(")[2], ")")[0]
					mimmax, _ := strconv.Atoi(mimax)
					mimmax = (mimmax-1)/8+1
					str+= "\t\tif value."+GetStrOutOfDashE(dataxx[0])+".BitLength != "+mimax+" || len(value."+GetStrOutOfDashE(dataxx[0])+".Bytes) != "+strconv.Itoa(mimmax)+" {\n" +
						"\t\t\treturn binData, bitEnd, errors.New(\""+GetStrOutOfDashE(datax[0])+": "+GetStrOutOfDashE(dataxx[0])+": BitString: sizeMin:"+mimax+",sizeMax:"+mimax+"\")\n" +
						"\t\t}\n" +
						"\t\tbinData, bitEnd = EncodeBitStringFix(*value."+GetStrOutOfDashE(dataxx[0])+", binData, bitEnd)\n"
				} else {
					str += "\tcase ngapType." + GetStrOutOfDashE(datax[0]) + "Present" + GetStrOutOfDashE(dataxx[0]) + ":\n" +
						"\t\tif value." + GetStrOutOfDashE(dataxx[0]) + " == nil {\n" +
						"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": NIL\")\n" +
						"\t\t}\n"
					min, max := strings.Split(dataxs[0], "(")[2], strings.Split(dataxs[1], ")")[0]
					str+= "\t\tif value."+ GetStrOutOfDashE(dataxx[0]) +".BitLength < "+min+" || value."+ GetStrOutOfDashE(dataxx[0]) +".BitLength > "+max+" || len(value."+ GetStrOutOfDashE(dataxx[0]) +".Bytes) != int((value."+ GetStrOutOfDashE(dataxx[0]) +".BitLength-1)/8+1) {\n" +
						"\t\t\treturn binData, bitEnd, errors.New(\""+GetStrOutOfDashE(datax[0])+": "+GetStrOutOfDashE(dataxx[0])+": BitString: sizeMin:"+min+",sizeMax:"+max+"\")\n" +
						"\t\t}\n" +
						"\t\tbinData, bitEnd = EncodeBitStringExt(*value."+GetStrOutOfDashE(dataxx[0])+", binData, bitEnd)\n"
				}
			} else {
				fmt.Println(dataxx)
			}
		} else {
			str += "\tcase ngapType." + GetStrOutOfDashE(datax[0]) + "Present" + GetStrOutOfDashE(dataxx[0]) + ":\n" +
				"\t\tif value." + GetStrOutOfDashE(dataxx[0]) + " == nil {\n" +
				"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": NIL\")\n" +
				"\t\t}\n" +
				"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(strings.Split(dataxx[1],",")[0]) + "(*value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
				"\t\tif err != nil {\n" +
				"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
				"\t\t}\n"
		}
	}
	str += "\tdefault:\n" +
		"\t\treturn binData, bitEnd, errors.New(\""+GetStrOutOfDashE(datax[0])+": Present: INVALID\")\n" +
		"\t}\n" +
		"\treturn binData, bitEnd, err\n" +
		"}\n"

	str = ""
	if choiceExt != "" {
		for i:=in; i<len(datas); i++ {
			datax = strings.Fields(datas[i])
			if len(datax) > 3 {
				if GetStrOutOfDashE(datax[0]) == choiceExt {
					str = "\nfunc ProtocolIESingleContainer"+choiceExt+"(value ngapType.ProtocolIESingleContainer"+choiceExt+", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
						"\tvar err error\n" +
						"\tnumIEs := len(value.List)\n" +
						"\tbinData, bitEnd, err = EncodeNumberIEsOfProtocolIESingleContainer(numIEs, binData, bitEnd)\n" +
						"\tif err != nil {\n" +
						"\t\treturn binData, bitEnd, errors.New(\"ProtocolIESingleContainer"+choiceExt+": \" + err.Error())\n" +
						"\t}\n" +
						"\tfor i := 0; i < numIEs; i++ {\n" +
						"\t\tvar binIEs []byte\n" +
						"\t\tbinIEs, bitEnd, err = "+choiceExt+"(value.List[i], binIEs, bitEnd)\n" +
						"\t\tif err != nil {\n" +
						"\t\t\treturn binData, bitEnd, errors.New(\"ProtocolIESingleContainer"+choiceExt+": [\" + strconv.Itoa(i) + \"]: \" + err.Error())\n" +
						"\t\t}\n" +
						"\t\tbinData = append(binData, binIEs...)\n" +
						"\t}\n" +
						"\treturn binData, bitEnd, err\n" +
						"}\n"

					xs := 0
					for {
						dataxs := strings.Fields(datas[i+1+xs])
						if len(dataxs) <= 1 {
							break
						}
						xs++
					}

					str = "\nfunc "+choiceExt+"(value ngapType."+choiceExt+", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
						"\tvar err error\n" +
						"\tswitch value.ProtocolIEID.Value {\n"
					for is:=0; is<xs; is++ {
						dataxs := strings.Fields(datas[i+1+is])
						str += "\tcase ngapType.ProtocolIEID"+ GetStrOutOfDashE2(dataxs[2])+":\n" +
							"\t\tif value.TypeValue.Present != ngapType."+choiceExt+"TypeValuePresent"+ GetStrOutOfDashE2(dataxs[2])+" {\n" +
							"\t\t\treturn binData, bitEnd, errors.New(\""+choiceExt+": ProtocolIEID: INVALID\")\n" +
							"\t\t}\n"
					}
					str += "\tdefault:\n" +
						"\t\treturn binData, bitEnd, errors.New(\""+choiceExt+": ProtocolIEID: INVALID\")\n" +
						"\t}\n" +
						"\tbinData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)\n" +
						"\tif err != nil {\n" +
						"\t\treturn binData, bitEnd, errors.New(\""+choiceExt+": \" + err.Error())\n" +
						"\t}\n" +
						"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
						"\tif err != nil {\n" +
						"\t\treturn binData, bitEnd, errors.New(\""+choiceExt+": \" + err.Error())\n" +
						"\t}\n" +
						"\tvar binValue []byte\n" +
						"\tbinValue, bitEnd, err = "+choiceExt+"TypeValue(value.TypeValue, binValue, 8)\n" +
						"\tif err != nil {\n" +
						"\t\treturn binData, bitEnd, errors.New(\""+choiceExt+": \" + err.Error())\n" +
						"\t}\n" +
						"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
						"}\n"

					str = "\nfunc "+choiceExt+"TypeValue(value ngapType."+choiceExt+"TypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
						"\tvar err error\n" +
						"\tswitch value.Present {\n"
					for is:=0; is<xs; is++ {
						dataxs := strings.Fields(datas[i+1+is])
						str += "\tcase ngapType."+choiceExt+"TypeValuePresent"+ GetStrOutOfDashE2(dataxs[2])+":\n" +
							"\t\tif value."+ GetStrOutOfDashE2(dataxs[2])+" == nil {\n" +
							"\t\t\treturn binData, bitEnd, errors.New(\""+choiceExt+"TypeValue: "+ GetStrOutOfDashE2(dataxs[2])+": NIL\")\n" +
							"\t\t}\n" +
							"\t\tbinData, bitEnd, err = "+ GetStrOutOfDashE(dataxs[6])+"(*value."+ GetStrOutOfDashE2(dataxs[2])+", binData, bitEnd)\n" +
							"\t\tif err != nil {\n" +
							"\t\t\treturn binData, bitEnd, errors.New(\""+choiceExt+"TypeValue: \" + err.Error())\n" +
							"\t\t}\n"
					}
					str += "\tdefault:\n" +
						"\t\treturn binData, bitEnd, errors.New(\""+choiceExt+"TypeValue: Present: INVALID\")\n" +
						"\t}\n" +
						"\treturn binData, bitEnd, err\n" +
						"}\n"
				}
			}
		}
	}
	return str, in
}

//
// func EncodeSEQUENCE(datas []string, in int) (str string, out int) {
// 	datax := strings.Fields(datas[in])
// 	if len(datax) == 6 && datax[4] == "OF" {
// 		// min, max := strings.Split(strings.Split(datax[3], "..")[0], "(")[2], strings.Split(strings.Split(datax[3], "..")[1], ")")[0]
// 		// str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 		// 	"\tvar err error\n" +
// 		// 	"\tnumItem := len(value.List)\n" +
// 		// 	"\tif numItem < " + min + " || numItem > ngapType." + GetStrOutOfDashE(max) + " {\n" +
// 		// 	"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": List: valueExt,sizeLB:" + min + ",sizeUB:" + max + "\")\n" +
// 		// 	"\t}\n" +
// 		// 	"\tbinData, bitEnd = EncodeNumberItemSequenceOf(numItem, " + min + ", ngapType."+GetStrOutOfDashE(max)+", binData, bitEnd)\n" +
// 		// 	"\tfor i:=0; i<numItem; i++ {\n" +
// 		// 	"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(datax[5]) + "(value.List[i], binData, bitEnd)\n" +
// 		// 	"\t\tif err != nil {\n" +
// 		// 	"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": [\"+strconv.Itoa(i)+\"]: \" + err.Error())\n" +
// 		// 	"\t\t}\n" +
// 		// 	"\t}\n" +
// 		// 	"\treturn binData, bitEnd, err\n" +
// 		// 	"}\n"
// 	} else if len(datax) == 4 {
// 		str += "\nfunc " + GetStrOutOfDashE(datax[0]) + " (value ngapType." + GetStrOutOfDashE(datax[0]) + ", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 			"\tvar err error\n" +
// 			"\tvar option uint64 = 0\n"
// 		xn := 0
// 		for {
// 			dataxx := strings.Fields(datas[in+1+xn])
// 			if len(dataxx) == 1 {
// 				break
// 			}
// 			xn++
// 		}
// 		xk := 0
// 		for xi := xn - 1; xi >= 0; xi-- {
// 			dataxx := strings.Fields(datas[in+1+xi])
// 			if dataxx[len(dataxx)-1] == "OPTIONAL," || dataxx[len(dataxx)-2] == "OPTIONAL" {
// 				str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + " != nil {\n" +
// 					"\t\toption += 1<<" + strconv.Itoa(xk) + "\n" +
// 					"\t}\n"
// 				xk++
// 			}
// 		}
// 		str += "\tbinData, bitEnd = EncodeOption(option, " + strconv.Itoa(xk+1) + ", binData, bitEnd)\n"
// 		ext, exx := "", ""
// 		for xi := 0; xi < xn; xi++ {
// 			dataxx := strings.Fields(datas[in+1+xi])
// 			if dataxx[len(dataxx)-1] == "OPTIONAL," || dataxx[len(dataxx)-1] == "OPTIONAL" || dataxx[len(dataxx)-2] == "OPTIONAL" {
// 				if dataxx[0] == "iE-Extensions" {
// 					ext = GetStrOutOfDashE(strings.Split(strings.Split(dataxx[3], "{")[1], "}")[0])
// 					str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + " != nil {\n" +
// 						"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[1]) + ext + "(*value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
// 						"\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
// 						"\t\t}\n" +
// 						"\t}\n"
// 				} else if dataxx[1] == "BIT" {
// 					min, max := strings.Split(strings.Split(dataxx[3], "..")[0], "(")[2], strings.Split(strings.Split(dataxx[3], "..")[1], ")")[0]
// 					str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + ".BitLength < " + min + " || value." + GetStrOutOfDashE(dataxx[0]) + ".BitLength > " + max + " || len(value." + GetStrOutOfDashE(dataxx[0]) + ".Bytes) != int((value." + GetStrOutOfDashE(dataxx[0]) + ".BitLength-1)/8+1) {\n" +
// 						"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": BitString: sizeMin:" + min + ",sizeMax:" + max + "\")\n" +
// 						"\t}\n" +
// 						"\tbinData, bitEnd = EncodeBitStringExt(*value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n"
// 				} else if dataxx[1] == "OCTET" {
// 					fmt.Println(dataxx)
// 				} else if dataxx[1] == "INTEGER" || dataxx[1] == "INTEGER," {
// 					if len(dataxx) == 4 {
// 						min, max := strings.Split(strings.Split(dataxx[2], "..")[0], "(")[1], strings.Split(strings.Split(dataxx[2], "..")[1], ")")[0]
// 						str += "\tif *value." + GetStrOutOfDashE(dataxx[0]) + " < " + min + " || *value." + GetStrOutOfDashE(dataxx[0]) + " > " + max + " {\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": Integer: valueMin:" + min + ",valueMax:" + max + "\")\n" +
// 							"\t}\n" +
// 							"\tbinData, bitEnd = EncodeInteger(*value." + GetStrOutOfDashE(dataxx[0]) + ", 0, binData, bitEnd)\n"
// 					} else {
// 						fmt.Println(dataxx)
// 					}
// 				} else if dataxx[1] == "ENUMERATED" || dataxx[1] == "ENUMERATED," {
// 					max := strconv.Itoa(len(dataxx) - 6)
// 					str += "\tif *value." + GetStrOutOfDashE(dataxx[0]) + " < 0 || *value." + GetStrOutOfDashE(dataxx[0]) + " > " + max + " {\n" +
// 						"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": Enumerated: valueMin:0,valueMax:" + max + "\")\n" +
// 						"\t}\n" +
// 						"\tbinData, bitEnd = EncodeEnumeratedExt(*value." + GetStrOutOfDashE(dataxx[0]) + ", 0, binData, bitEnd)\n"
// 				} else {
// 					str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + " != nil {\n" +
// 						"\t\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[1]) + "(*value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
// 						"\t\tif err != nil {\n" +
// 						"\t\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
// 						"\t\t}\n" +
// 						"\t}\n"
// 				}
// 			} else {
// 				if dataxx[0] == "protocolIEs" {
// 					exx = GetStrOutOfDashE(strings.Split(strings.Split(dataxx[3], "{")[1], "}")[0])
// 					str += "\tbinData, bitEnd, err = " + GetStrOutOfDashE(dataxx[1]) + exx + "(value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
// 						"\tif err != nil {\n" +
// 						"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
// 						"\t}\n"
// 				} else if dataxx[1] == "BIT" {
// 					str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + ".BitLength != 8 || len(value." + GetStrOutOfDashE(dataxx[0]) + ".Bytes) != 1 {\n" +
// 						"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": BitString: sizeMin:8,sizeMax:8\")\n" +
// 						"\t}\n" +
// 						"\tbinData, bitEnd = EncodeBitStringFixExx(value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n"
// 				} else if dataxx[1] == "OCTET" {
// 					if dataxx[3] == "(CONTAINING" {
// 						str += "\tbinData, bitEnd, err = EncodeOctetStringContaining(value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
// 							"\tif err != nil {\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
// 							"\t}\n"
// 					} else {
// 						str += "\tif len(value." + GetStrOutOfDashE(dataxx[0]) + ") != 4 {\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": OctetString: sizeMin:4,sizeMax:4\")\n" +
// 							"\t}\n" +
// 							"\tbinData, bitEnd = EncodeOctetStringFix(value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n"
// 					}
// 				} else if dataxx[1] == "INTEGER" || dataxx[1] == "INTEGER," {
// 					if len(dataxx) == 3 {
// 						min, max := strings.Split(strings.Split(dataxx[2], "..")[0], "(")[1], strings.Split(strings.Split(dataxx[2], "..")[1], ")")[0]
// 						str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + " < " + min + " || value." + GetStrOutOfDashE(dataxx[0]) + " > " + max + " {\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": Integer: valueMin:" + min + ",valueMax:" + max + "\")\n" +
// 							"\t}\n" +
// 							"\tbinData, bitEnd = EncodeInteger(value." + GetStrOutOfDashE(dataxx[0]) + ", 0, binData, bitEnd)\n"
// 					} else if len(dataxx) == 4 {
// 						min, max := strings.Split(strings.Split(dataxx[2], "..")[0], "(")[1], strings.Split(strings.Split(dataxx[2], "..")[1], ",")[0]
// 						str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + " < " + min + " || value." + GetStrOutOfDashE(dataxx[0]) + " > " + max + " {\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": Integer: valueMin:" + min + ",valueMax:" + max + "\")\n" +
// 							"\t}\n" +
// 							"\tbinData, bitEnd = EncodeIntegerExt(value." + GetStrOutOfDashE(dataxx[0]) + ", 0, binData, bitEnd)\n"
// 					} else {
// 						fmt.Println(dataxx)
// 					}
// 				} else if dataxx[1] == "ENUMERATED" || dataxx[1] == "ENUMERATED," {
// 					max := strconv.Itoa(len(dataxx) - 5)
// 					str += "\tif value." + GetStrOutOfDashE(dataxx[0]) + " < 0 || value." + GetStrOutOfDashE(dataxx[0]) + " > " + max + " {\n" +
// 						"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": " + GetStrOutOfDashE(dataxx[0]) + ": Enumerated: valueMin:0,valueMax:" + max + "\")\n" +
// 						"\t}\n" +
// 						"\tbinData, bitEnd = EncodeEnumeratedExt(value." + GetStrOutOfDashE(dataxx[0]) + ", 0, binData, bitEnd)\n"
// 				} else {
// 					str += "\tbinData, bitEnd, err = " + GetStrOutOfDashE(strings.Split(dataxx[1], ",")[0]) + "(value." + GetStrOutOfDashE(dataxx[0]) + ", binData, bitEnd)\n" +
// 						"\tif err != nil {\n" +
// 						"\t\treturn binData, bitEnd, errors.New(\"" + GetStrOutOfDashE(datax[0]) + ": \" + err.Error())\n" +
// 						"\t}\n"
// 				}
// 			}
// 		}
// 		str += "\treturn binData, bitEnd, err\n}\n"
// 		if ext != "" {
// 			for is:=in; is<len(datas); is++ {
// 				datax = strings.Fields(datas[is])
// 				if len(datax) > 3 {
// 					if GetStrOutOfDashE(datax[0]) == ext {
// 						// str = "\nfunc ProtocolExtensionContainer"+ext+"(value ngapType.ProtocolExtensionContainer"+ext+", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 						// 	"\tvar err error\n" +
// 						// 	"\tnumIEs := len(value.List)\n" +
// 						// 	"\tbinData, bitEnd, err = EncodeNumberIEsOfProtocolExtensionContainer(numIEs, binData, bitEnd)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\"ProtocolExtensionContainer"+ext+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\tfor i := 0; i < numIEs; i++ {\n" +
// 						// 	"\t\tvar binIEs []byte\n" +
// 						// 	"\t\tbinIEs, bitEnd, err = "+ext+"(value.List[i], binIEs, bitEnd)\n" +
// 						// 	"\t\tif err != nil {\n" +
// 						// 	"\t\t\treturn binData, bitEnd, errors.New(\"ProtocolExtensionContainer"+ext+": [\" + strconv.Itoa(i) + \"]: \" + err.Error())\n" +
// 						// 	"\t\t}\n" +
// 						// 	"\t\tbinData = append(binData, binIEs...)\n" +
// 						// 	"\t}\n" +
// 						// 	"\treturn binData, bitEnd, err\n" +
// 						// 	"}\n"
// 						// xs, xsk := 0, 0
// 						// str = "\nfunc "+ext+"(value ngapType."+ext+", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 						// 	"\tvar err error\n"
// 						// for {
// 						// 	dataxs := strings.Fields(datas[is+1+xs])
// 						// 	if len(dataxs) == 1 {
// 						// 		break
// 						// 	}
// 						// 	if xsk == 0 {
// 						// 		str += "\tswitch value.ProtocolExtensionID.Value {\n"
// 						// 		xsk = 1
// 						// 	}
// 						// 	str += "\tcase ngapType.ProtocolIEID"+GetStrOutOfDashE2(dataxs[2])+":\n" +
// 						// 		"\t\tif value.ExtensionValue.Present != ngapType."+ext+"ExtensionValuePresent"+GetStrOutOfDashE2(dataxs[2])+" {\n" +
// 						// 		"\t\t\treturn binData, bitEnd, errors.New(\""+ext+": ProtocolExtensionID: INVALID\")\n" +
// 						// 		"\t\t}\n"
// 						// 	xs++
// 						// }
// 						// if xsk == 1 {
// 						// 	str += "\tdefault:\n" +
// 						// 		"\t\treturn binData, bitEnd, errors.New(\""+ext+": ProtocolExtensionID: INVALID\")\n" +
// 						// 		"\t}\n"
// 						// }
// 						//
// 						// str += "\tbinData, bitEnd, err = ProtocolExtensionID(value.ProtocolExtensionID, binData, bitEnd)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+ext+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+ext+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\tvar binValue []byte\n" +
// 						// 	"\tbinValue, bitEnd, err = "+ext+"ExtensionValue(value.ExtensionValue, binValue, 8)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+ext+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 						// 	"}\n"
//
// 						xs := 0
// 						str = "\nfunc "+ext+"ExtensionValue(value ngapType."+ext+"ExtensionValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 							"\tvar err error\n"
// 						str += "\tswitch value.Present {\n"
// 						for {
// 							dataxs := strings.Fields(datas[is+1+xs])
// 							if len(dataxs) == 1 {
// 								break
// 							}
// 							str += "\tcase ngapType."+ext+"ExtensionValuePresent"+GetStrOutOfDashE2(dataxs[2])+":\n" +
// 								"\t\tif value."+GetStrOutOfDashE2(dataxs[2])+" == nil {\n" +
// 								"\t\t\treturn binData, bitEnd, errors.New(\""+ext+"ExtensionValue: "+GetStrOutOfDashE2(dataxs[2])+": NIL\")\n" +
// 								"\t\t}\n" +
// 								"\t\tbinData, bitEnd, err = "+GetStrOutOfDashE(dataxs[6])+"(*value."+GetStrOutOfDashE2(dataxs[2])+", binData, bitEnd)\n" +
// 								"\t\tif err != nil {\n" +
// 								"\t\t\treturn binData, bitEnd, errors.New(\""+ext+"ExtensionValue: \" + err.Error())\n" +
// 								"\t\t}\n"
// 							xs++
// 						}
// 						str += "\tdefault:\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\""+ext+"ExtensionValue: Present: INVALID\")\n" +
// 							"\t}\n"
// 						str += "\treturn binData, bitEnd, err\n" +
// 							"}\n"
//
// 					}
// 				}
// 			}
//
// 		}
// 		if exx != "" {
// 			for is:=in; is<len(datas); is++ {
// 				datax = strings.Fields(datas[is])
// 				if len(datax) > 3 {
// 					if GetStrOutOfDashE(datax[0]) == exx {
// 						// str = "\nfunc ProtocolIEContainer"+exx+"(value ngapType.ProtocolIEContainer"+exx+", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 						// 	"\tvar err error\n" +
// 						// 	"\tnumIEs := len(value.List)\n" +
// 						// 	"\tbinData, bitEnd, err = EncodeNumberIEsOfProtocolIEContainer(numIEs, binData, bitEnd)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\"ProtocolIEContainer"+exx+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\tfor i := 0; i < numIEs; i++ {\n" +
// 						// 	"\t\tvar binIEs []byte\n" +
// 						// 	"\t\tbinIEs, bitEnd, err = "+exx+"(value.List[i], binIEs, bitEnd)\n" +
// 						// 	"\t\tif err != nil {\n" +
// 						// 	"\t\t\treturn binData, bitEnd, errors.New(\"ProtocolIEContainer"+exx+": [\" + strconv.Itoa(i) + \"]: \" + err.Error())\n" +
// 						// 	"\t\t}\n" +
// 						// 	"\t\tbinData = append(binData, binIEs...)\n" +
// 						// 	"\t}\n" +
// 						// 	"\treturn binData, bitEnd, err\n" +
// 						// 	"}\n"
// 						// xs := 0
// 						// str = "\nfunc "+exx+"(value ngapType."+exx+", binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 						// 	"\tvar err error\n" +
// 						// 	"\tswitch value.ProtocolIEID.Value {\n"
// 						// for {
// 						// 	dataxs := strings.Fields(datas[is+1+xs])
// 						// 	if len(dataxs) == 1 {
// 						// 		break
// 						// 	}
// 						// 	str += "\tcase ngapType.ProtocolIEID"+GetStrOutOfDashE2(dataxs[2])+":\n" +
// 						// 		"\t\tif value.TypeValue.Present != ngapType."+exx+"TypeValuePresent"+GetStrOutOfDashE2(dataxs[2])+" {\n" +
// 						// 		"\t\t\treturn binData, bitEnd, errors.New(\""+exx+": ProtocolIEID: INVALID\")\n" +
// 						// 		"\t\t}\n"
// 						// 	xs++
// 						// }
// 						// str += "\tdefault:\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+exx+": ProtocolIEID: INVALID\")\n" +
// 						// 	"\t}\n" +
// 						// 	"\tbinData, bitEnd, err = ProtocolIEID(value.ProtocolIEID, binData, bitEnd)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+exx+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\tbinData, bitEnd, err = Criticality(value.Criticality, binData, bitEnd)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+exx+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\tvar binValue []byte\n" +
// 						// 	"\tbinValue, bitEnd, err = "+exx+"TypeValue(value.TypeValue, binValue, 8)\n" +
// 						// 	"\tif err != nil {\n" +
// 						// 	"\t\treturn binData, bitEnd, errors.New(\""+exx+": \" + err.Error())\n" +
// 						// 	"\t}\n" +
// 						// 	"\treturn append(binData, EncodeLengthValue(binValue)...), bitEnd, err\n" +
// 						// 	"}\n"
// 						xs := 0
// 						str = "\nfunc "+exx+"TypeValue(value ngapType."+exx+"TypeValue, binData []byte, bitEnd uint64) ([]byte, uint64, error) {\n" +
// 							"\tvar err error\n"
// 						str += "\tswitch value.Present {\n"
// 						for {
// 							dataxs := strings.Fields(datas[is+1+xs])
// 							if len(dataxs) == 1 {
// 								break
// 							}
// 							str += "\tcase ngapType."+exx+"TypeValuePresent"+GetStrOutOfDashE2(dataxs[2])+":\n" +
// 								"\t\tif value."+GetStrOutOfDashE2(dataxs[2])+" == nil {\n" +
// 								"\t\t\treturn binData, bitEnd, errors.New(\""+exx+"TypeValue: "+GetStrOutOfDashE2(dataxs[2])+": NIL\")\n" +
// 								"\t\t}\n" +
// 								"\t\tbinData, bitEnd, err = "+GetStrOutOfDashE(dataxs[6])+"(*value."+GetStrOutOfDashE2(dataxs[2])+", binData, bitEnd)\n" +
// 								"\t\tif err != nil {\n" +
// 								"\t\t\treturn binData, bitEnd, errors.New(\""+exx+"TypeValue: \" + err.Error())\n" +
// 								"\t\t}\n"
// 							xs++
// 						}
// 						str += "\tdefault:\n" +
// 							"\t\treturn binData, bitEnd, errors.New(\""+exx+"TypeValue: Present: INVALID\")\n" +
// 							"\t}\n"
// 						str += "\treturn binData, bitEnd, err\n" +
// 							"}\n"
//
// 					}
// 				}
// 			}
//
// 		}
// 	} else {
// 		fmt.Println(in, "-", datax)
// 	}
// 	return str, in
// }

func GetStrOutOfDashE(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 0; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}

func GetStrOutOfDashE2(str string) (newStr string) {
	strs := strings.Split(strings.Title(str), "-")
	for i := 1; i < len(strs); i++ {
		newStr += strs[i]
	}
	return
}
