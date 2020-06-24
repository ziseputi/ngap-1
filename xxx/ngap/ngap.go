
package ngap

import (
	"vht5gc/lib/ngap/ngapEncode"
	"vht5gc/lib/ngap/ngapType"
)

func Encode (value ngapType.NGAPPDU) ([]byte, error) {
	binData, _, err := ngapEncode.NGAPPDU(value, nil, 8)
	return binData, err
}

func Decode (value []byte) (ngapType.NGAPPDU, error) {
	return ngapType.NGAPPDU{}, nil
}