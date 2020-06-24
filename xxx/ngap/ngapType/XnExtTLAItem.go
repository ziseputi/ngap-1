// Created By HaoDHH-245789 VHT2020
package ngapType

type XnExtTLAItem struct {
	IPsecTLA     *TransportLayerAddress                        `vht:"optional"`
	GTPTLAs      *XnGTPTLAs                                    `vht:"optional"`
	IEExtensions *ProtocolExtensionContainerXnExtTLAItemExtIEs `vht:"optional"`
}
