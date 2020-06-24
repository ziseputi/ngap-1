// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ConcurrentWarningMessageIndPresentTrue Enumerated = 0
)

type ConcurrentWarningMessageInd struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
