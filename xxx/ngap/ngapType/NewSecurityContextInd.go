// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	NewSecurityContextIndPresentTrue Enumerated = 0
)

type NewSecurityContextInd struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
