// Created By HaoDHH-245789 VHT2020
package ngapType

type TAICancelledEUTRAItem struct {
	TAI                      TAI `vht:"valueExt"`
	CancelledCellsInTAIEUTRA CancelledCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs `vht:"optional"`
}
