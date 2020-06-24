// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	HandoverTypePresentIntra5gs    Enumerated = 0
	HandoverTypePresentFivegsToEps Enumerated = 1
	HandoverTypePresentEpsTo5gs    Enumerated = 2
)

type HandoverType struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:2"`
}
