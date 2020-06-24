// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CauseRadioNetworkPresentUnspecified                                              Enumerated = 0
	CauseRadioNetworkPresentTxnrelocoverallExpiry                                    Enumerated = 1
	CauseRadioNetworkPresentSuccessfulHandover                                       Enumerated = 2
	CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason                         Enumerated = 3
	CauseRadioNetworkPresentReleaseDueTo5gcGeneratedReason                           Enumerated = 4
	CauseRadioNetworkPresentHandoverCancelled                                        Enumerated = 5
	CauseRadioNetworkPresentPartialHandover                                          Enumerated = 6
	CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem              Enumerated = 7
	CauseRadioNetworkPresentHoTargetNotAllowed                                       Enumerated = 8
	CauseRadioNetworkPresentTngrelocoverallExpiry                                    Enumerated = 9
	CauseRadioNetworkPresentTngrelocprepExpiry                                       Enumerated = 10
	CauseRadioNetworkPresentCellNotAvailable                                         Enumerated = 11
	CauseRadioNetworkPresentUnknownTargetID                                          Enumerated = 12
	CauseRadioNetworkPresentNoRadioResourcesAvailableInTargetCell                    Enumerated = 13
	CauseRadioNetworkPresentUnknownLocalUENGAPID                                     Enumerated = 14
	CauseRadioNetworkPresentInconsistentRemoteUENGAPID                               Enumerated = 15
	CauseRadioNetworkPresentHandoverDesirableForRadioReason                          Enumerated = 16
	CauseRadioNetworkPresentTimeCriticalHandover                                     Enumerated = 17
	CauseRadioNetworkPresentResourceOptimisationHandover                             Enumerated = 18
	CauseRadioNetworkPresentReduceLoadInServingCell                                  Enumerated = 19
	CauseRadioNetworkPresentUserInactivity                                           Enumerated = 20
	CauseRadioNetworkPresentRadioConnectionWithUeLost                                Enumerated = 21
	CauseRadioNetworkPresentRadioResourcesNotAvailable                               Enumerated = 22
	CauseRadioNetworkPresentInvalidQosCombination                                    Enumerated = 23
	CauseRadioNetworkPresentFailureInRadioInterfaceProcedure                         Enumerated = 24
	CauseRadioNetworkPresentInteractionWithOtherProcedure                            Enumerated = 25
	CauseRadioNetworkPresentUnknownPDUSessionID                                      Enumerated = 26
	CauseRadioNetworkPresentUnkownQosFlowID                                          Enumerated = 27
	CauseRadioNetworkPresentMultiplePDUSessionIDInstances                            Enumerated = 28
	CauseRadioNetworkPresentMultipleQosFlowIDInstances                               Enumerated = 29
	CauseRadioNetworkPresentEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported Enumerated = 30
	CauseRadioNetworkPresentNgIntraSystemHandoverTriggered                           Enumerated = 31
	CauseRadioNetworkPresentNgInterSystemHandoverTriggered                           Enumerated = 32
	CauseRadioNetworkPresentXnHandoverTriggered                                      Enumerated = 33
	CauseRadioNetworkPresentNotSupported5QIValue                                     Enumerated = 34
	CauseRadioNetworkPresentUeContextTransfer                                        Enumerated = 35
	CauseRadioNetworkPresentImsVoiceEpsFallbackOrRatFallbackTriggered                Enumerated = 36
	CauseRadioNetworkPresentUpIntegrityProtectionNotPossible                         Enumerated = 37
	CauseRadioNetworkPresentUpConfidentialityProtectionNotPossible                   Enumerated = 38
	CauseRadioNetworkPresentSliceNotSupported                                        Enumerated = 39
	CauseRadioNetworkPresentUeInRrcInactiveStateNotReachable                         Enumerated = 40
	CauseRadioNetworkPresentRedirection                                              Enumerated = 41
	CauseRadioNetworkPresentResourcesNotAvailableForTheSlice                         Enumerated = 42
	CauseRadioNetworkPresentUeMaxIntegrityProtectedDataRateReason                    Enumerated = 43
	CauseRadioNetworkPresentReleaseDueToCnDetectedMobility                           Enumerated = 44
)

type CauseRadioNetwork struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:44"`
}
