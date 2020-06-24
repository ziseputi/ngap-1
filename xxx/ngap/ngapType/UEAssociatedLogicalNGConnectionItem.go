// Created By HaoDHH-245789 VHT2020
package ngapType

type UEAssociatedLogicalNGConnectionItem struct {
	AMFUENGAPID  *AMFUENGAPID                                                         `vht:"optional"`
	RANUENGAPID  *RANUENGAPID                                                         `vht:"optional"`
	IEExtensions *ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs `vht:"optional"`
}
