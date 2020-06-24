// Created By HaoDHH-245789 VHT2020
package ngapType

type AreaOfInterest struct {
	AreaOfInterestTAIList     *AreaOfInterestTAIList                          `vht:"optional"`
	AreaOfInterestCellList    *AreaOfInterestCellList                         `vht:"optional"`
	AreaOfInterestRANNodeList *AreaOfInterestRANNodeList                      `vht:"optional"`
	IEExtensions              *ProtocolExtensionContainerAreaOfInterestExtIEs `vht:"optional"`
}
