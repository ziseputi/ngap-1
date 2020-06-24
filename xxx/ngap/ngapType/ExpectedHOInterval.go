// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ExpectedHOIntervalPresentSec15    Enumerated = 0
	ExpectedHOIntervalPresentSec30    Enumerated = 1
	ExpectedHOIntervalPresentSec60    Enumerated = 2
	ExpectedHOIntervalPresentSec90    Enumerated = 3
	ExpectedHOIntervalPresentSec120   Enumerated = 4
	ExpectedHOIntervalPresentSec180   Enumerated = 5
	ExpectedHOIntervalPresentLongTime Enumerated = 6
)

type ExpectedHOInterval struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:6"`
}
