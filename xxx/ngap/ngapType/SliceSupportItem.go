// Created By HaoDHH-245789 VHT2020
package ngapType

type SliceSupportItem struct {
	SNSSAI       SNSSAI                                            `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerSliceSupportItemExtIEs `vht:"optional"`
}
