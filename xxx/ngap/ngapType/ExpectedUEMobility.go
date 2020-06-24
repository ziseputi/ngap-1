// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ExpectedUEMobilityPresentStationary Enumerated = 0
	ExpectedUEMobilityPresentMobile     Enumerated = 1
)

type ExpectedUEMobility struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
