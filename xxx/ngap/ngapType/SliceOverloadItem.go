// Created By HaoDHH-245789 VHT2020
package ngapType

type SliceOverloadItem struct {
	SNSSAI       SNSSAI                                             `vht:"valueExt"`
	IEExtensions *ProtocolExtensionContainerSliceOverloadItemExtIEs `vht:"optional"`
}
