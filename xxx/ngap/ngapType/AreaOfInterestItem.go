// Created By HaoDHH-245789 VHT2020
package ngapType

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest `vht:"valueExt"`
	LocationReportingReferenceID LocationReportingReferenceID
	IEExtensions                 *ProtocolExtensionContainerAreaOfInterestItemExtIEs `vht:"optional"`
}
