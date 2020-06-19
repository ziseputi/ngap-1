// Created By HaoDHH-245789 VHT2020
package ngapEncode

import "ngap/ngapType"

// BIT STRING

func EncodeBitString(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeBitStringExt(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeBitStringExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeBitStringExtExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// OCTET STRING

func EncodeOctetString(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeOctetStringExt(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeOctetStringExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeOctetStringExtExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// OBJECT IDENTIFIER

func EncodeObjectIdentifier(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeObjectIdentifierExt(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeObjectIdentifierExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeObjectIdentifierExtExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// ENUMERATED

func EncodeEnumerated(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeEnumeratedExt(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeEnumeratedExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeEnumeratedExtExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// INTEGER

func EncodeInteger(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeIntegerExt(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeIntegerExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeIntegerExtExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// PrintableString

func EncodePrintableString(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodePrintableStringExt(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodePrintableStringExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodePrintableStringExtExx(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// Choice

func EncodeChoice(value int, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Criticality

func Criticality(value ngapType.Criticality, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Presence

func Presence(value ngapType.Presence, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// PrivateIEID

func PrivateIEID(value ngapType.PrivateIEID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// ProcedureCode

func ProcedureCode(value ngapType.ProcedureCode, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// ProtocolExtensionID

func ProtocolExtensionID(value ngapType.ProtocolExtensionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// ProtocolIEID

func ProtocolIEID(value ngapType.ProtocolExtensionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// TriggeringMessage

func TriggeringMessage(value ngapType.ProtocolExtensionID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Lists

func EncodeMaxnoofLists(value ngapType.Integer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Extension constants

func EncodeMaxExtensionConstants(value ngapType.Integer, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}
