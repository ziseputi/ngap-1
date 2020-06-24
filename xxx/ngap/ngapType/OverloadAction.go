// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	OverloadActionPresentRejectNonEmergencyMoDt                                    Enumerated = 0
	OverloadActionPresentRejectRrcCrSignalling                                     Enumerated = 1
	OverloadActionPresentPermitEmergencySessionsAndMobileTerminatedServicesOnly    Enumerated = 2
	OverloadActionPresentPermitHighPrioritySessionsAndMobileTerminatedServicesOnly Enumerated = 3
)

type OverloadAction struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:3"`
}
