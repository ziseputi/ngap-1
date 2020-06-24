// Created By HaoDHH-245789 VHT2020
package ngapEncode

import "ngap/ngapType"

// BIT STRING

// (SIZE(mimax))
func EncodeBitStringFix(value ngapType.BitString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// (SIZE(min..max))
func EncodeBitStringExt(value ngapType.BitString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// (SIZE(mimax, ...))
func EncodeBitStringFixExx(value ngapType.BitString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// (SIZE(min..max, ...))
func EncodeBitStringExtExx(value ngapType.BitString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// OCTET STRING

// NO TAG
func EncodeOctetStringNo(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// (SIZE(mimax))
func EncodeOctetStringFix(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// (SIZE(min..max))
func EncodeOctetStringExt(value ngapType.OctetString, lenMin uint64, lenMax uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// CONTAINING
func EncodeOctetStringContaining(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

func EncodeOctetStringExtExx(value ngapType.OctetString, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// OBJECT IDENTIFIER

func EncodeObjectIdentifier(value ngapType.ObjectIdentifier, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeObjectIdentifierExt(value ngapType.ObjectIdentifier, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeObjectIdentifierExx(value ngapType.ObjectIdentifier, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeObjectIdentifierExtExx(value ngapType.ObjectIdentifier, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// ENUMERATED

// { value1, value2, value3 }
func EncodeEnumerated(value ngapType.Enumerated, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// { value1, value2, ... }
func EncodeEnumeratedExt(value ngapType.Enumerated, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// INTEGER

// (min..max)
func EncodeInteger(value ngapType.Integer, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// (min..max, ...)
func EncodeIntegerExt(value ngapType.Integer, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// PrintableString

// (SIZE(min..max, ...))
func EncodePrintableString(value ngapType.PrintableString, lenMin uint64, lenMax uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// Choice

func EncodePresent(value int, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// Option

func EncodeOption(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

// Length Value

func EncodeLengthValue(binData []byte) []byte {
	return binData
}

// Length Extension Value

func EncodeLengthExtensionValue(binData []byte) []byte {
	return binData
}

// Length Type Value

func EncodeLengthTypeValue(binData []byte) []byte {
	return binData
}

// Number IEs of ProtocolIEContainer

func EncodeNumberIEsOfProtocolIEContainer(value int, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Number IEs of ProtocolIESingleContainer

func EncodeNumberIEsOfProtocolIESingleContainer(value int, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Number IEs of ProtocolExtensionContainer

func EncodeNumberIEsOfProtocolExtensionContainer(value int, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Number IEs of PrivateIEContainer

func EncodeNumberIEsOfPrivateIEContainer(value int, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// Number Item SEQUENCE OF

func EncodeNumberItemSequenceOf(value int, lenMin int, lenMax int, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
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

func ProtocolIEID(value ngapType.ProtocolIEID, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}

// TriggeringMessage

func TriggeringMessage(value ngapType.TriggeringMessage, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
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

// NGAPMessage

func NGAPMessage(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64, error) {
	return binData, bitEnd, nil
}
