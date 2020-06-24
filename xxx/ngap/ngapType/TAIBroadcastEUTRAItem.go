// Created By HaoDHH-245789 VHT2020
package ngapType

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI `vht:"valueExt"`
	CompletedCellsInTAIEUTRA CompletedCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs `vht:"optional"`
}
