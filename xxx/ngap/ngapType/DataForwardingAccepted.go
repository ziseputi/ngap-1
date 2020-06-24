// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DataForwardingAcceptedPresentDataForwardingAccepted Enumerated = 0
)

type DataForwardingAccepted struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
