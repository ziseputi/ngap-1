// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	SONInformationRequestPresentXnTNLConfigurationInfo Enumerated = 0
)

type SONInformationRequest struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
