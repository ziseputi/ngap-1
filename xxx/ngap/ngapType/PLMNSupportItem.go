// Created By HaoDHH-245789 VHT2020
package ngapType

type PLMNSupportItem struct {
	PLMNIdentity     PLMNIdentity
	SliceSupportList SliceSupportList
	IEExtensions     *ProtocolExtensionContainerPLMNSupportItemExtIEs `vht:"optional"`
}
