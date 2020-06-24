// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	IntegrityProtectionIndicationPresentRequired  Enumerated = 0
	IntegrityProtectionIndicationPresentPreferred Enumerated = 1
	IntegrityProtectionIndicationPresentNotNeeded Enumerated = 2
)

type IntegrityProtectionIndication struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:2"`
}
