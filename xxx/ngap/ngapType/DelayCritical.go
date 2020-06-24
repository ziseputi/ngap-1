// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DelayCriticalPresentDelayCritical    Enumerated = 0
	DelayCriticalPresentNonDelayCritical Enumerated = 1
)

type DelayCritical struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
