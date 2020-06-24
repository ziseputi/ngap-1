// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	RRCStatePresentInactive  Enumerated = 0
	RRCStatePresentConnected Enumerated = 1
)

type RRCState struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
