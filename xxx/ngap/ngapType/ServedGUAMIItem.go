// Created By HaoDHH-245789 VHT2020
package ngapType

type ServedGUAMIItem struct {
	GUAMI         GUAMI                                            `vht:"valueExt"`
	BackupAMFName *AMFName                                         `vht:"sizeExt,sizeLB:1,sizeUB:150,optional"`
	IEExtensions  *ProtocolExtensionContainerServedGUAMIItemExtIEs `vht:"optional"`
}
