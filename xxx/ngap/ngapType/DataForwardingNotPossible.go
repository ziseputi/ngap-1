// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DataForwardingNotPossiblePresentDataForwardingNotPossible Enumerated = 0
)

type DataForwardingNotPossible struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
