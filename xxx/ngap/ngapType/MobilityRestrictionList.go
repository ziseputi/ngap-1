// Created By HaoDHH-245789 VHT2020
package ngapType

type MobilityRestrictionList struct {
	ServingPLMN              PLMNIdentity
	EquivalentPLMNs          *EquivalentPLMNs                                         `vht:"optional"`
	RATRestrictions          *RATRestrictions                                         `vht:"optional"`
	ForbiddenAreaInformation *ForbiddenAreaInformation                                `vht:"optional"`
	ServiceAreaInformation   *ServiceAreaInformation                                  `vht:"optional"`
	IEExtensions             *ProtocolExtensionContainerMobilityRestrictionListExtIEs `vht:"optional"`
}
