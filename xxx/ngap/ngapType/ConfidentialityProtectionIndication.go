// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ConfidentialityProtectionIndicationPresentRequired  Enumerated = 0
	ConfidentialityProtectionIndicationPresentPreferred Enumerated = 1
	ConfidentialityProtectionIndicationPresentNotNeeded Enumerated = 2
)

type ConfidentialityProtectionIndication struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:2"`
}
