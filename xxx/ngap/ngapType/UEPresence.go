// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	UEPresencePresentIn      Enumerated = 0
	UEPresencePresentOut     Enumerated = 1
	UEPresencePresentUnknown Enumerated = 2
)

type UEPresence struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:2"`
}
