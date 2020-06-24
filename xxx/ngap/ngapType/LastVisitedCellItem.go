// Created By HaoDHH-245789 VHT2020
package ngapType

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation                           `vht:"valueLB:0,valueUB:4"`
	IEExtensions               *ProtocolExtensionContainerLastVisitedCellItemExtIEs `vht:"optional"`
}
