// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	AdditionalQosFlowInformationPresentMoreLikely Enumerated = 0
)

type AdditionalQosFlowInformation struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
