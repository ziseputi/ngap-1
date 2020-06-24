// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ConfidentialityProtectionResultPresentPerformed    Enumerated = 0
	ConfidentialityProtectionResultPresentNotPerformed Enumerated = 1
)

type ConfidentialityProtectionResult struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
