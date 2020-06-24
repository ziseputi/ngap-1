// Created By HaoDHH-245789 VHT2020
package ngapType

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          NGRANCGI `vht:"valueLB:0,valueUB:2"`
	CellType                              CellType `vht:"valueExt"`
	TimeUEStayedInCell                    TimeUEStayedInCell
	TimeUEStayedInCellEnhancedGranularity *TimeUEStayedInCellEnhancedGranularity                           `vht:"optional"`
	HOCauseValue                          *Cause                                                           `vht:"valueLB:0,valueUB:5,optional"`
	IEExtensions                          *ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs `vht:"optional"`
}
