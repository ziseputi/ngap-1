// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CauseMiscPresentControlProcessingOverload             Enumerated = 0
	CauseMiscPresentNotEnoughUserPlaneProcessingResources Enumerated = 1
	CauseMiscPresentHardwareFailure                       Enumerated = 2
	CauseMiscPresentOmIntervention                        Enumerated = 3
	CauseMiscPresentUnknownPLMN                           Enumerated = 4
	CauseMiscPresentUnspecified                           Enumerated = 5
)

type CauseMisc struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:5"`
}
