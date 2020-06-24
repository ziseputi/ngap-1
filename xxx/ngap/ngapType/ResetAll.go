// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ResetAllPresentResetAll Enumerated = 0
)

type ResetAll struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
