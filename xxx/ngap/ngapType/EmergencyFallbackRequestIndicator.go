// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	EmergencyFallbackRequestIndicatorPresentEmergencyFallbackRequested Enumerated = 0
)

type EmergencyFallbackRequestIndicator struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
