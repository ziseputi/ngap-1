// Created By HaoDHH-245789 VHT2020
package ngapEncode

import "ngap/ngapType"

func EncodeUint64(value uint64, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	value = (value << (64 - lenValue)) >> (64 - lenValue)
	switch lenValue {
	case 1, 2, 3, 4, 5, 6, 7, 8:
		value = value << (8 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> bitEnd)
		}
		if (lenValue + bitEnd) > 8 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 8
		}
		return binData, lenValue + bitEnd
	case 9, 10, 11, 12, 13, 14, 15, 16:
		value = value << (16 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 8))
		}
		binData = append(binData, byte(value>>bitEnd))
		if (lenValue + bitEnd) > 16 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 16
		}
		return binData, lenValue + bitEnd - 8
	case 17, 18, 19, 20, 21, 22, 23, 24:
		value = value << (24 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 16))
		}
		binData = append(binData, byte(value>>(bitEnd+8)), byte(value>>bitEnd))
		if (lenValue + bitEnd) > 24 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 24
		}
		return binData, lenValue + bitEnd - 16
	case 25, 26, 27, 28, 29, 30, 31, 32:
		value = value << (32 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 24))
		}
		binData = append(binData, byte(value>>(bitEnd+16)), byte(value>>(bitEnd+8)), byte(value>>bitEnd))
		if (lenValue + bitEnd) > 32 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 32
		}
		return binData, lenValue + bitEnd - 24
	case 33, 34, 35, 36, 37, 38, 39, 40:
		value = value << (40 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 32))
		}
		binData = append(binData, byte(value>>(bitEnd+24)), byte(value>>(bitEnd+16)), byte(value>>(bitEnd+8)), byte(value>>bitEnd))
		if (lenValue + bitEnd) > 40 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 40
		}
		return binData, lenValue + bitEnd - 32
	case 41, 42, 43, 44, 45, 46, 47, 48:
		value = value << (48 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 40))
		}
		binData = append(binData, byte(value>>(bitEnd+32)), byte(value>>(bitEnd+24)), byte(value>>(bitEnd+16)), byte(value>>(bitEnd+8)), byte(value>>bitEnd))
		if (lenValue + bitEnd) > 48 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 48
		}
		return binData, lenValue + bitEnd - 40
	case 49, 50, 51, 52, 53, 54, 55, 56:
		value = value << (56 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 48))
		}
		binData = append(binData, byte(value>>(bitEnd+40)), byte(value>>(bitEnd+32)), byte(value>>(bitEnd+24)), byte(value>>(bitEnd+16)), byte(value>>(bitEnd+8)), byte(value>>bitEnd))
		if (lenValue + bitEnd) > 56 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 56
		}
		return binData, lenValue + bitEnd - 48
	case 57, 58, 59, 60, 61, 62, 63, 64:
		value = value << (64 - lenValue)
		if binData != nil {
			binData[len(binData)-1] += byte(value >> (bitEnd + 56))
		}
		binData = append(binData, byte(value>>(bitEnd+48)), byte(value>>(bitEnd+40)), byte(value>>(bitEnd+32)), byte(value>>(bitEnd+24)), byte(value>>(bitEnd+16)), byte(value>>(bitEnd+8)), byte(value>>bitEnd))
		if (lenValue + bitEnd) > 64 {
			return append(binData, byte(value<<(8-bitEnd))), lenValue + bitEnd - 64
		}
		return binData, lenValue + bitEnd - 56
	default:
		return binData, bitEnd
	}
}

func EncodeString(value string, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeStringExt(value string, lenValue int, lenBit uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	binData, bitEnd = EncodeUint64(uint64(lenValue), lenBit, binData, bitEnd)
	return append(binData, []byte(value)...), 8
}

func EncodeBitString(value ngapType.BitString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	switch value.BitLength {
	case 1, 2, 3, 4, 5, 6, 7, 8:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0])>>(8-value.BitLength), value.BitLength, binData, bitEnd)
	case 9, 10, 11, 12, 13, 14, 15, 16:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1])>>(16-value.BitLength), value.BitLength-8, binData, bitEnd)
	case 17, 18, 19, 20, 21, 22, 23, 24:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2])>>(24-value.BitLength), value.BitLength-16, binData, bitEnd)
	case 25, 26, 27, 28, 29, 30, 31, 32:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3])>>(32-value.BitLength), value.BitLength-24, binData, bitEnd)
	case 33, 34, 35, 36, 37, 38, 39, 40:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4])>>(42-value.BitLength), value.BitLength-32, binData, bitEnd)
	case 41, 42, 43, 44, 45, 46, 47, 48:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5])>>(48-value.BitLength), value.BitLength-40, binData, bitEnd)
	case 49, 50, 51, 52, 53, 54, 55, 56:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6])>>(56-value.BitLength), value.BitLength-48, binData, bitEnd)
	case 57, 58, 59, 60, 61, 62, 63, 64:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7])>>(64-value.BitLength), value.BitLength-56, binData, bitEnd)
	case 65, 66, 67, 68, 69, 70, 71, 72:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8])>>(72-value.BitLength), value.BitLength-64, binData, bitEnd)
	case 73, 74, 75, 76, 77, 78, 79, 80:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9])>>(80-value.BitLength), value.BitLength-72, binData, bitEnd)
	case 81, 82, 83, 84, 85, 86, 87, 88:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10])>>(88-value.BitLength), value.BitLength-80, binData, bitEnd)
	case 89, 90, 91, 92, 93, 94, 95, 96:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11])>>(96-value.BitLength), value.BitLength-88, binData, bitEnd)
	case 97, 98, 99, 100, 101, 102, 103, 104:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12])>>(104-value.BitLength), value.BitLength-96, binData, bitEnd)
	case 105, 106, 107, 108, 109, 110, 111, 112:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13])>>(112-value.BitLength), value.BitLength-104, binData, bitEnd)
	case 113, 114, 115, 116, 117, 118, 119, 120:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[14])>>(120-value.BitLength), value.BitLength-112, binData, bitEnd)
	case 121, 122, 123, 124, 125, 126, 127, 128:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[14]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[15])>>(128-value.BitLength), value.BitLength-120, binData, bitEnd)
	case 129, 130, 131, 132, 133, 134, 135, 136:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[14]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[15]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[16])>>(136-value.BitLength), value.BitLength-128, binData, bitEnd)
	case 137, 138, 139, 140, 141, 142, 143, 144:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[14]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[15]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[16]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[17])>>(144-value.BitLength), value.BitLength-136, binData, bitEnd)
	case 145, 146, 147, 148, 149, 150, 151, 152:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[14]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[15]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[16]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[17]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[18])>>(152-value.BitLength), value.BitLength-144, binData, bitEnd)
	case 153, 154, 155, 156, 157, 158, 159, 160:
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[7]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[8]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[9]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[10]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[11]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[12]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[13]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[14]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[15]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[16]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[17]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[18]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value.Bytes[19])>>(160-value.BitLength), value.BitLength-152, binData, bitEnd)
	}
	return binData, bitEnd
}

func EncodeBitStringExt(value ngapType.BitString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	binData, bitEnd = EncodeUint64(value.BitLength-1, 8, binData, bitEnd)
	return EncodeBitString(value, binData, 8)
}

func EncodeEnumerated(value ngapType.Enumerated, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeEnumeratedExt(value ngapType.Enumerated, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodeBytes(value []byte, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return append(binData, value...), 8
}

func EncodeOctetString(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return append(binData, value...), 8
}

func EncodeOctetStringEx(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	switch len(value) {
	case 1:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
	case 2:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
	case 3:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[2]), 8, binData, bitEnd)
	case 4:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[3]), 8, binData, bitEnd)
	case 5:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[4]), 8, binData, bitEnd)
	case 6:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[5]), 8, binData, bitEnd)
	case 7:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[6]), 8, binData, bitEnd)
	case 8:
		binData, bitEnd = EncodeUint64(uint64(value[0]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[1]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[2]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[3]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[4]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[5]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[6]), 8, binData, bitEnd)
		binData, bitEnd = EncodeUint64(uint64(value[7]), 8, binData, bitEnd)
	}
	return binData, bitEnd
}

func EncodeOctetStringExt(value ngapType.OctetString, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return append(binData, value...), 8
}

func EncodeObjectIdentifier(value ngapType.ObjectIdentifier, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return binData, bitEnd
}

func EncodePresent(value int, lenValue uint64, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return EncodeUint64(uint64(value)-1, lenValue, binData, bitEnd)
}

func EncodeLengthValue(binData []byte) []byte {
	if len(binData) < 128 {
		return append([]byte{byte(len(binData))}, binData...)
	}
	return append([]byte{byte(len(binData)/256 + 128), byte(len(binData))}, binData...)
}

func EncodeNumberProtocolIE(value int, binData []byte, bitEnd uint64) ([]byte, uint64) {
	return EncodeUint64(uint64(value), 24, binData, 8)
}

func ConvertInt64ToBytes(value int64) (bin []byte) {
	for {
		bin = append([]byte{byte(value)}, bin...)
		if value < 256 {
			return
		}
		value /= 256
	}
}
