// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PreEmptionCapabilityPresentShallNotTriggerPreEmption Enumerated = 0
	PreEmptionCapabilityPresentMayTriggerPreEmption      Enumerated = 1
)

type PreEmptionCapability struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
