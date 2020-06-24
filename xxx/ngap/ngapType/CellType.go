// Created By HaoDHH-245789 VHT2020
package ngapType

type CellType struct {
	CellSize     CellSize
	IEExtensions *ProtocolExtensionContainerCellTypeExtIEs `vht:"optional"`
}
