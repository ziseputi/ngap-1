// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DLForwardingPresentDlForwardingProposed Enumerated = 0
)

type DLForwarding struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
