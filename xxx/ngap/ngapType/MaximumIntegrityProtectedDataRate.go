// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	MaximumIntegrityProtectedDataRatePresentBitrate64kbs  Enumerated = 0
	MaximumIntegrityProtectedDataRatePresentMaximumUERate Enumerated = 1
)

type MaximumIntegrityProtectedDataRate struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
