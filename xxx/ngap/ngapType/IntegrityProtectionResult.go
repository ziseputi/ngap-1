// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	IntegrityProtectionResultPresentPerformed    Enumerated = 0
	IntegrityProtectionResultPresentNotPerformed Enumerated = 1
)

type IntegrityProtectionResult struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
