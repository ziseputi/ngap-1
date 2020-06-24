// Created By HaoDHH-245789 VHT2020
package ngapType

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells *AssistanceDataForRecommendedCells                       `vht:"valueExt,optional"`
	PagingAttemptInformation          *PagingAttemptInformation                                `vht:"valueExt,optional"`
	IEExtensions                      *ProtocolExtensionContainerAssistanceDataForPagingExtIEs `vht:"optional"`
}
