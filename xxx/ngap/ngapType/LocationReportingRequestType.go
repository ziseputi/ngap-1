// Created By HaoDHH-245789 VHT2020
package ngapType

type LocationReportingRequestType struct {
	EventType                                 EventType
	ReportArea                                ReportArea
	AreaOfInterestList                        *AreaOfInterestList                                           `vht:"optional"`
	LocationReportingReferenceIDToBeCancelled *LocationReportingReferenceID                                 `vht:"optional"`
	IEExtensions                              *ProtocolExtensionContainerLocationReportingRequestTypeExtIEs `vht:"optional"`
}
