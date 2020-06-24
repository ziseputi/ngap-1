// Created By HaoDHH-245789 VHT2020
package ngapType

type ERABInformationItem struct {
	ERABID       ERABID
	DLForwarding *DLForwarding                                        `vht:"optional"`
	IEExtensions *ProtocolExtensionContainerERABInformationItemExtIEs `vht:"optional"`
}
