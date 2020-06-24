// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CancelAllWarningMessagesPresentTrue Enumerated = 0
)

type CancelAllWarningMessages struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
