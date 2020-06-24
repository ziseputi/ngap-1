// Created By HaoDHH-245789 VHT2020
package ngapType

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      RRCContainer
	PDUSessionResourceInformationList *PDUSessionResourceInformationList `vht:"optional"`
	ERABInformationList               *ERABInformationList               `vht:"optional"`
	TargetCellID                      NGRANCGI                           `vht:"valueLB:0,valueUB:2"`
	IndexToRFSP                       *IndexToRFSP                       `vht:"optional"`
	UEHistoryInformation              UEHistoryInformation
	IEExtensions                      *ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs `vht:"optional"`
}
