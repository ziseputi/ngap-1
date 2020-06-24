package main

import (
	"fmt"
	"os"
	"reflect"

	"ngap/ngapType"
)

func main() {
	obj := ngapType.NGAPPDU{}
	objType := reflect.TypeOf(obj)
	var listObj []string
	var strWrite [10]string
	levelWrite := 1
	listObj, strWrite = GenerateEncode(objType.Name(), objType, listObj, strWrite, levelWrite)

	filePointer, _ := os.Create("test/test0.txt")
	filePointer.WriteString(strWrite[0])
	filePointer, _ = os.Create("test/test1.txt")
	filePointer.WriteString(strWrite[1])
	filePointer, _ = os.Create("test/test2.txt")
	filePointer.WriteString(strWrite[2])
	filePointer, _ = os.Create("test/test3.txt")
	filePointer.WriteString(strWrite[3])
	filePointer, _ = os.Create("test/test4.txt")
	filePointer.WriteString(strWrite[4])
	filePointer, _ = os.Create("test/test5.txt")
	filePointer.WriteString(strWrite[5])
	filePointer, _ = os.Create("test/test6.txt")
	filePointer.WriteString(strWrite[6])
	filePointer, _ = os.Create("test/test7.txt")
	filePointer.WriteString(strWrite[7])
	filePointer, _ = os.Create("test/test8.txt")
	filePointer.WriteString(strWrite[8])
	filePointer, _ = os.Create("test/test9.txt")
	filePointer.WriteString(strWrite[9])
}

func InListObject(listObj []string, objName string) (res bool) {
	for i := 0; i < len(listObj); i++ {
		if objName == listObj[i] {
			return true
		}
	}
	return false
}

func GenerateEncode(objName string, objType reflect.Type, listObj []string, strWrite [10]string, levelWrite int) (newListObj []string, newStrWrite [10]string) {
	if !InListObject(listObj, objName) {
		listObj = append(listObj, objName)
		if objName != objType.Name() {
			fmt.Println("1", objType)
			strWrite = AppendToStrWrite(objName+"\n", levelWrite, strWrite)
			return GenerateEncode(objType.Name(), objType, listObj, strWrite, levelWrite+1)
		} else if objType.NumField() == 0 {
			fmt.Println("2", objType)
			strWrite = AppendToStrWrite(objName+"\n", levelWrite, strWrite)
			return listObj, strWrite
		} else {
			strWrite = AppendToStrWrite(objType.Name()+"\n", levelWrite, strWrite)
			fmt.Println("3", objType)
			for i := 0; i < objType.NumField(); i++ {
				fmt.Println(i, "-", objType.Field(i).Type.Kind())
				if objType.Field(i).Type.Kind() == reflect.Ptr {
					switch objType.Field(i).Type.Elem().Name() {
					case "int", "BitString", "OctetString", "ObjectIdentifier", "Enumerated", "Integer", "PrintableString", "Criticality", "Presence", "PrivateIEID", "ProcedureCode", "ProtocolExtensionID", "ProtocolIEID", "TriggeringMessage":
					default:
						listObj, strWrite = GenerateEncode(objType.Field(i).Type.Elem().Name(), objType.Field(i).Type.Elem(), listObj, strWrite, levelWrite+1)
					}
				} else if objType.Field(i).Type.Kind() == reflect.Slice {
					// switch objType.Field(i).Type.Elem().Name() {
					// case "int", "BitString", "OctetString", "ObjectIdentifier", "Enumerated", "Integer", "PrintableString", "Criticality", "Presence", "PrivateIEID", "ProcedureCode", "ProtocolExtensionID", "ProtocolIEID", "TriggeringMessage":
					// 	strWrite = AppendToStrWrite(objName+"\n", levelWrite, strWrite)
					// default:
					// 	strWrite = AppendToStrWrite(objName+"\n", levelWrite, strWrite)
					// 	listObj, strWrite = GenerateEncode(objType.Field(i).Type.Elem().Name(), objType.Field(i).Type.Elem(), listObj, strWrite, levelWrite+1)
					// }
				} else {
					switch objType.Field(i).Type.Name() {
					case "int", "BitString", "OctetString", "ObjectIdentifier", "Enumerated", "Integer", "PrintableString", "Criticality", "Presence", "PrivateIEID", "ProcedureCode", "ProtocolExtensionID", "ProtocolIEID", "TriggeringMessage":
					default:
						listObj, strWrite = GenerateEncode(objType.Field(i).Type.Name(), objType.Field(i).Type, listObj, strWrite, levelWrite+1)
					}
				}
			}
		}
	}
	return listObj, strWrite
}

func AppendToStrWrite(str string, level int, strWrite [10]string) (res [10]string) {
	switch level {
	case 1, 2, 3:
		strWrite[0] += str
	case 4, 5, 6:
		strWrite[1] += str
	case 7:
		strWrite[2] += str
	case 8:
		strWrite[3] += str
	case 9:
		strWrite[4] += str
	case 10:
		strWrite[5] += str
	case 11:
		strWrite[6] += str
	case 12:
		strWrite[7] += str
	case 13:
		strWrite[8] += str
	default:
		strWrite[9] += str
	}
	return strWrite
}
