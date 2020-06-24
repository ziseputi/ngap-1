// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	EmergencyServiceTargetCNPresentFiveGC Enumerated = 0
	EmergencyServiceTargetCNPresentEpc    Enumerated = 1
)

type EmergencyServiceTargetCN struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
