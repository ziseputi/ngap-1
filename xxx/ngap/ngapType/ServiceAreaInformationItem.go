// Created By HaoDHH-245789 VHT2020
package ngapType

type ServiceAreaInformationItem struct {
	PLMNIdentity   PLMNIdentity
	AllowedTACs    *AllowedTACs                                                `vht:"optional"`
	NotAllowedTACs *NotAllowedTACs                                             `vht:"optional"`
	IEExtensions   *ProtocolExtensionContainerServiceAreaInformationItemExtIEs `vht:"optional"`
}
