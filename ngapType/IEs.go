package ngapType

type AdditionalDLUPTNLInformationForHOList struct {
	List []AdditionalDLUPTNLInformationForHOItem `vht:"sizeMin:1,sizeMax:MaxnoofMultiConnectivityMinusOne"`
}

type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation        UPTransportLayerInformation
	AdditionalQosFlowSetupResponseList     QosFlowListWithDataForwarding
	AdditionalDLForwardingUPTNLInformation *UPTransportLayerInformation
	IEExtensions                           *ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs
}

type ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs struct {
	List []AdditionalDLUPTNLInformationForHOItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AdditionalDLUPTNLInformationForHOItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValueChoiceNothing int = 0
)

type AdditionalDLUPTNLInformationForHOItemExtIEsExtensionValue struct {
	Choice int
}

const (
	AdditionalQosFlowInformationMoreLikely Enumerated = 0
)

type AdditionalQosFlowInformation struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        PriorityLevelARP
	PreEmptionCapability    PreEmptionCapability
	PreEmptionVulnerability PreEmptionVulnerability
	IEExtensions            *ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs
}

type ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs struct {
	List []AllocationAndRetentionPriorityExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AllocationAndRetentionPriorityExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AllocationAndRetentionPriorityExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AllocationAndRetentionPriorityExtIEsExtensionValueChoiceNothing int = 0
)

type AllocationAndRetentionPriorityExtIEsExtensionValue struct {
	Choice int
}

type AllowedNSSAI struct {
	List []AllowedNSSAIItem `vht:"sizeMin:1,sizeMax:MaxnoofAllowedS-NSSAIs"`
}

type AllowedNSSAIItem struct {
	SNSSAI       SNSSAI
	IEExtensions *ProtocolExtensionContainerAllowedNSSAIItemExtIEs
}

type ProtocolExtensionContainerAllowedNSSAIItemExtIEs struct {
	List []AllowedNSSAIItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AllowedNSSAIItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                          `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AllowedNSSAIItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AllowedNSSAIItemExtIEsExtensionValueChoiceNothing int = 0
)

type AllowedNSSAIItemExtIEsExtensionValue struct {
	Choice int
}

type AllowedTACs struct {
	List []TAC `vht:"sizeMin:1,sizeMax:MaxnoofAllowedAreas"`
}

type AMFName struct {
	Value PrintableString `vht:"sizeExt,sizeMin:1,sizeMax:150"`
}

const (
	AMFPagingTargetChoiceGlobalRANNodeID  int = 0
	AMFPagingTargetChoiceTAI              int = 1
	AMFPagingTargetChoiceChoiceExtensions int = 2
)

type AMFPagingTarget struct {
	Choice           int
	GlobalRANNodeID  *GlobalRANNodeID
	TAI              *TAI
	ChoiceExtensions *ProtocolIESingleContainerAMFPagingTargetExtIEs
}

type ProtocolIESingleContainerAMFPagingTargetExtIEs struct {
	List []AMFPagingTargetExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type AMFPagingTargetExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	TypeValue    AMFPagingTargetExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	AMFPagingTargetExtIEsTypeValueChoiceNothing int = 0
)

type AMFPagingTargetExtIEsTypeValue struct {
	Choice int
}

type AMFPointer struct {
	Value BitString `vht:"sizeMin:6,sizeMax:6"`
}

type AMFRegionID struct {
	Value BitString `vht:"sizeMin:8,sizeMax:8"`
}

type AMFSetID struct {
	Value BitString `vht:"sizeMin:10,sizeMax:10"`
}

type AMFTNLAssociationSetupList struct {
	List []AMFTNLAssociationSetupItem `vht:"sizeMin:1,sizeMax:MaxnoofTNLAssociations"`
}

type AMFTNLAssociationSetupItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs
}

type ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs struct {
	List []AMFTNLAssociationSetupItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AMFTNLAssociationSetupItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AMFTNLAssociationSetupItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AMFTNLAssociationSetupItemExtIEsExtensionValueChoiceNothing int = 0
)

type AMFTNLAssociationSetupItemExtIEsExtensionValue struct {
	Choice int
}

type AMFTNLAssociationToAddList struct {
	List []AMFTNLAssociationToAddItem `vht:"sizeMin:1,sizeMax:MaxnoofTNLAssociations"`
}

type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	TNLAssociationUsage      *TNLAssociationUsage
	TNLAddressWeightFactor   TNLAddressWeightFactor
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs
}

type ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs struct {
	List []AMFTNLAssociationToAddItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AMFTNLAssociationToAddItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AMFTNLAssociationToAddItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AMFTNLAssociationToAddItemExtIEsExtensionValueChoiceNothing int = 0
)

type AMFTNLAssociationToAddItemExtIEsExtensionValue struct {
	Choice int
}

type AMFTNLAssociationToRemoveList struct {
	List []AMFTNLAssociationToRemoveItem `vht:"sizeMin:1,sizeMax:MaxnoofTNLAssociations"`
}

type AMFTNLAssociationToRemoveItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs
}

type ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs struct {
	List []AMFTNLAssociationToRemoveItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AMFTNLAssociationToRemoveItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AMFTNLAssociationToRemoveItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AMFTNLAssociationToRemoveItemExtIEsExtensionValueChoiceTNLAssociationTransportLayerAddressNGRAN int = 0
)

type AMFTNLAssociationToRemoveItemExtIEsExtensionValue struct {
	Choice                                   int
	TNLAssociationTransportLayerAddressNGRAN *CPTransportLayerInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDTNLAssociationTransportLayerAddressNGRAN"`
}

type AMFTNLAssociationToUpdateList struct {
	List []AMFTNLAssociationToUpdateItem `vht:"sizeMin:1,sizeMax:MaxnoofTNLAssociations"`
}

type AMFTNLAssociationToUpdateItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	TNLAssociationUsage      *TNLAssociationUsage
	TNLAddressWeightFactor   *TNLAddressWeightFactor
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs
}

type ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs struct {
	List []AMFTNLAssociationToUpdateItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AMFTNLAssociationToUpdateItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AMFTNLAssociationToUpdateItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AMFTNLAssociationToUpdateItemExtIEsExtensionValueChoiceNothing int = 0
)

type AMFTNLAssociationToUpdateItemExtIEsExtensionValue struct {
	Choice int
}

type AMFUENGAPID struct {
	Value Integer `vht:"valueMin:0,valueMax:1099511627775"`
}

type AreaOfInterest struct {
	AreaOfInterestTAIList     *AreaOfInterestTAIList
	AreaOfInterestCellList    *AreaOfInterestCellList
	AreaOfInterestRANNodeList *AreaOfInterestRANNodeList
	IEExtensions              *ProtocolExtensionContainerAreaOfInterestExtIEs
}

type ProtocolExtensionContainerAreaOfInterestExtIEs struct {
	List []AreaOfInterestExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AreaOfInterestExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AreaOfInterestExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AreaOfInterestExtIEsExtensionValueChoiceNothing int = 0
)

type AreaOfInterestExtIEsExtensionValue struct {
	Choice int
}

type AreaOfInterestCellList struct {
	List []AreaOfInterestCellItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinAoI"`
}

type AreaOfInterestCellItem struct {
	NGRANCGI     NGRANCGI
	IEExtensions *ProtocolExtensionContainerAreaOfInterestCellItemExtIEs
}

type ProtocolExtensionContainerAreaOfInterestCellItemExtIEs struct {
	List []AreaOfInterestCellItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AreaOfInterestCellItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AreaOfInterestCellItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AreaOfInterestCellItemExtIEsExtensionValueChoiceNothing int = 0
)

type AreaOfInterestCellItemExtIEsExtensionValue struct {
	Choice int
}

type AreaOfInterestList struct {
	List []AreaOfInterestItem `vht:"sizeMin:1,sizeMax:MaxnoofAoI"`
}

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest
	LocationReportingReferenceID LocationReportingReferenceID
	IEExtensions                 *ProtocolExtensionContainerAreaOfInterestItemExtIEs
}

type ProtocolExtensionContainerAreaOfInterestItemExtIEs struct {
	List []AreaOfInterestItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AreaOfInterestItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AreaOfInterestItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AreaOfInterestItemExtIEsExtensionValueChoiceNothing int = 0
)

type AreaOfInterestItemExtIEsExtensionValue struct {
	Choice int
}

type AreaOfInterestRANNodeList struct {
	List []AreaOfInterestRANNodeItem `vht:"sizeMin:1,sizeMax:MaxnoofRANNodeinAoI"`
}

type AreaOfInterestRANNodeItem struct {
	GlobalRANNodeID GlobalRANNodeID
	IEExtensions    *ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs
}

type ProtocolExtensionContainerAreaOfInterestRANNodeItemExtIEs struct {
	List []AreaOfInterestRANNodeItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AreaOfInterestRANNodeItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AreaOfInterestRANNodeItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AreaOfInterestRANNodeItemExtIEsExtensionValueChoiceNothing int = 0
)

type AreaOfInterestRANNodeItemExtIEsExtensionValue struct {
	Choice int
}

type AreaOfInterestTAIList struct {
	List []AreaOfInterestTAIItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIinAoI"`
}

type AreaOfInterestTAIItem struct {
	TAI          TAI
	IEExtensions *ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs
}

type ProtocolExtensionContainerAreaOfInterestTAIItemExtIEs struct {
	List []AreaOfInterestTAIItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AreaOfInterestTAIItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AreaOfInterestTAIItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AreaOfInterestTAIItemExtIEsExtensionValueChoiceNothing int = 0
)

type AreaOfInterestTAIItemExtIEsExtensionValue struct {
	Choice int
}

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells *AssistanceDataForRecommendedCells
	PagingAttemptInformation          *PagingAttemptInformation
	IEExtensions                      *ProtocolExtensionContainerAssistanceDataForPagingExtIEs
}

type ProtocolExtensionContainerAssistanceDataForPagingExtIEs struct {
	List []AssistanceDataForPagingExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AssistanceDataForPagingExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AssistanceDataForPagingExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AssistanceDataForPagingExtIEsExtensionValueChoiceNothing int = 0
)

type AssistanceDataForPagingExtIEsExtensionValue struct {
	Choice int
}

type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging RecommendedCellsForPaging
	IEExtensions              *ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs
}

type ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs struct {
	List []AssistanceDataForRecommendedCellsExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AssistanceDataForRecommendedCellsExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AssistanceDataForRecommendedCellsExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AssistanceDataForRecommendedCellsExtIEsExtensionValueChoiceNothing int = 0
)

type AssistanceDataForRecommendedCellsExtIEsExtensionValue struct {
	Choice int
}

type AssociatedQosFlowList struct {
	List []AssociatedQosFlowItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

const (
	QosFlowMappingIndicationUl Enumerated = 0
	QosFlowMappingIndicationDl Enumerated = 1
)

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        QosFlowIdentifier
	QosFlowMappingIndication *Enumerated `vht:"valueExt,valueMin:0,valueMax:2"`
	IEExtensions             *ProtocolExtensionContainerAssociatedQosFlowItemExtIEs
}

type ProtocolExtensionContainerAssociatedQosFlowItemExtIEs struct {
	List []AssociatedQosFlowItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type AssociatedQosFlowItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      AssociatedQosFlowItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	AssociatedQosFlowItemExtIEsExtensionValueChoiceNothing int = 0
)

type AssociatedQosFlowItemExtIEsExtensionValue struct {
	Choice int
}

type AveragingWindow struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:4095"`
}

type BitRate struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:4000000000000"`
}

const (
	BroadcastCancelledAreaListChoiceCellIDCancelledEUTRA          int = 0
	BroadcastCancelledAreaListChoiceTAICancelledEUTRA             int = 1
	BroadcastCancelledAreaListChoiceEmergencyAreaIDCancelledEUTRA int = 2
	BroadcastCancelledAreaListChoiceCellIDCancelledNR             int = 3
	BroadcastCancelledAreaListChoiceTAICancelledNR                int = 4
	BroadcastCancelledAreaListChoiceEmergencyAreaIDCancelledNR    int = 5
	BroadcastCancelledAreaListChoiceChoiceExtensions              int = 6
)

type BroadcastCancelledAreaList struct {
	Choice                        int
	CellIDCancelledEUTRA          *CellIDCancelledEUTRA
	TAICancelledEUTRA             *TAICancelledEUTRA
	EmergencyAreaIDCancelledEUTRA *EmergencyAreaIDCancelledEUTRA
	CellIDCancelledNR             *CellIDCancelledNR
	TAICancelledNR                *TAICancelledNR
	EmergencyAreaIDCancelledNR    *EmergencyAreaIDCancelledNR
	ChoiceExtensions              *ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs
}

type ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs struct {
	List []BroadcastCancelledAreaListExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type BroadcastCancelledAreaListExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    BroadcastCancelledAreaListExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	BroadcastCancelledAreaListExtIEsTypeValueChoiceNothing int = 0
)

type BroadcastCancelledAreaListExtIEsTypeValue struct {
	Choice int
}

const (
	BroadcastCompletedAreaListChoiceCellIDBroadcastEUTRA          int = 0
	BroadcastCompletedAreaListChoiceTAIBroadcastEUTRA             int = 1
	BroadcastCompletedAreaListChoiceEmergencyAreaIDBroadcastEUTRA int = 2
	BroadcastCompletedAreaListChoiceCellIDBroadcastNR             int = 3
	BroadcastCompletedAreaListChoiceTAIBroadcastNR                int = 4
	BroadcastCompletedAreaListChoiceEmergencyAreaIDBroadcastNR    int = 5
	BroadcastCompletedAreaListChoiceChoiceExtensions              int = 6
)

type BroadcastCompletedAreaList struct {
	Choice                        int
	CellIDBroadcastEUTRA          *CellIDBroadcastEUTRA
	TAIBroadcastEUTRA             *TAIBroadcastEUTRA
	EmergencyAreaIDBroadcastEUTRA *EmergencyAreaIDBroadcastEUTRA
	CellIDBroadcastNR             *CellIDBroadcastNR
	TAIBroadcastNR                *TAIBroadcastNR
	EmergencyAreaIDBroadcastNR    *EmergencyAreaIDBroadcastNR
	ChoiceExtensions              *ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs
}

type ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs struct {
	List []BroadcastCompletedAreaListExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type BroadcastCompletedAreaListExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    BroadcastCompletedAreaListExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	BroadcastCompletedAreaListExtIEsTypeValueChoiceNothing int = 0
)

type BroadcastCompletedAreaListExtIEsTypeValue struct {
	Choice int
}

type BroadcastPLMNList struct {
	List []BroadcastPLMNItem `vht:"sizeMin:1,sizeMax:MaxnoofBPLMNs"`
}

type BroadcastPLMNItem struct {
	PLMNIdentity        PLMNIdentity
	TAISliceSupportList SliceSupportList
	IEExtensions        *ProtocolExtensionContainerBroadcastPLMNItemExtIEs
}

type ProtocolExtensionContainerBroadcastPLMNItemExtIEs struct {
	List []BroadcastPLMNItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type BroadcastPLMNItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      BroadcastPLMNItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	BroadcastPLMNItemExtIEsExtensionValueChoiceNothing int = 0
)

type BroadcastPLMNItemExtIEsExtensionValue struct {
	Choice int
}

const (
	CancelAllWarningMessagesTrue Enumerated = 0
)

type CancelAllWarningMessages struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type CancelledCellsInEAIEUTRA struct {
	List []CancelledCellsInEAIEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinEAI"`
}

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs
}

type ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs struct {
	List []CancelledCellsInEAIEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CancelledCellsInEAIEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CancelledCellsInEAIEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CancelledCellsInEAIEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type CancelledCellsInEAIEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type CancelledCellsInEAINR struct {
	List []CancelledCellsInEAINRItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinEAI"`
}

type CancelledCellsInEAINRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs
}

type ProtocolExtensionContainerCancelledCellsInEAINRItemExtIEs struct {
	List []CancelledCellsInEAINRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CancelledCellsInEAINRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CancelledCellsInEAINRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CancelledCellsInEAINRItemExtIEsExtensionValueChoiceNothing int = 0
)

type CancelledCellsInEAINRItemExtIEsExtensionValue struct {
	Choice int
}

type CancelledCellsInTAIEUTRA struct {
	List []CancelledCellsInTAIEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinTAI"`
}

type CancelledCellsInTAIEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs
}

type ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs struct {
	List []CancelledCellsInTAIEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CancelledCellsInTAIEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CancelledCellsInTAIEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CancelledCellsInTAIEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type CancelledCellsInTAIEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type CancelledCellsInTAINR struct {
	List []CancelledCellsInTAINRItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinTAI"`
}

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs
}

type ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs struct {
	List []CancelledCellsInTAINRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CancelledCellsInTAINRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CancelledCellsInTAINRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CancelledCellsInTAINRItemExtIEsExtensionValueChoiceNothing int = 0
)

type CancelledCellsInTAINRItemExtIEsExtensionValue struct {
	Choice int
}

const (
	CauseChoiceRadioNetwork     int = 0
	CauseChoiceTransport        int = 1
	CauseChoiceNas              int = 2
	CauseChoiceProtocol         int = 3
	CauseChoiceMisc             int = 4
	CauseChoiceChoiceExtensions int = 5
)

type Cause struct {
	Choice           int
	RadioNetwork     *CauseRadioNetwork
	Transport        *CauseTransport
	Nas              *CauseNas
	Protocol         *CauseProtocol
	Misc             *CauseMisc
	ChoiceExtensions *ProtocolIESingleContainerCauseExtIEs
}

type ProtocolIESingleContainerCauseExtIEs struct {
	List []CauseExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type CauseExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality          `vht:"Reference:ProtocolIEID"`
	TypeValue    CauseExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	CauseExtIEsTypeValueChoiceNothing int = 0
)

type CauseExtIEsTypeValue struct {
	Choice int
}

const (
	CauseMiscControlProcessingOverload             Enumerated = 0
	CauseMiscNotEnoughUserPlaneProcessingResources Enumerated = 1
	CauseMiscHardwareFailure                       Enumerated = 2
	CauseMiscOmIntervention                        Enumerated = 3
	CauseMiscUnknownPLMN                           Enumerated = 4
	CauseMiscUnspecified                           Enumerated = 5
)

type CauseMisc struct {
	Value Enumerated `vht:"valueMin:0,valueMax:5"`
}

const (
	CauseNasNormalRelease         Enumerated = 0
	CauseNasAuthenticationFailure Enumerated = 1
	CauseNasDeregister            Enumerated = 2
	CauseNasUnspecified           Enumerated = 3
)

type CauseNas struct {
	Value Enumerated `vht:"valueMin:0,valueMax:3"`
}

const (
	CauseProtocolTransferSyntaxError                          Enumerated = 0
	CauseProtocolAbstractSyntaxErrorReject                    Enumerated = 1
	CauseProtocolAbstractSyntaxErrorIgnoreAndNotify           Enumerated = 2
	CauseProtocolMessageNotCompatibleWithReceiverState        Enumerated = 3
	CauseProtocolSemanticError                                Enumerated = 4
	CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage Enumerated = 5
	CauseProtocolUnspecified                                  Enumerated = 6
)

type CauseProtocol struct {
	Value Enumerated `vht:"valueMin:0,valueMax:6"`
}

const (
	CauseRadioNetworkUnspecified                                              Enumerated = 0
	CauseRadioNetworkTxnrelocoverallExpiry                                    Enumerated = 1
	CauseRadioNetworkSuccessfulHandover                                       Enumerated = 2
	CauseRadioNetworkReleaseDueToNgranGeneratedReason                         Enumerated = 3
	CauseRadioNetworkReleaseDueTo5gcGeneratedReason                           Enumerated = 4
	CauseRadioNetworkHandoverCancelled                                        Enumerated = 5
	CauseRadioNetworkPartialHandover                                          Enumerated = 6
	CauseRadioNetworkHoFailureInTarget5GCNgranNodeOrTargetSystem              Enumerated = 7
	CauseRadioNetworkHoTargetNotAllowed                                       Enumerated = 8
	CauseRadioNetworkTngrelocoverallExpiry                                    Enumerated = 9
	CauseRadioNetworkTngrelocprepExpiry                                       Enumerated = 10
	CauseRadioNetworkCellNotAvailable                                         Enumerated = 11
	CauseRadioNetworkUnknownTargetID                                          Enumerated = 12
	CauseRadioNetworkNoRadioResourcesAvailableInTargetCell                    Enumerated = 13
	CauseRadioNetworkUnknownLocalUENGAPID                                     Enumerated = 14
	CauseRadioNetworkInconsistentRemoteUENGAPID                               Enumerated = 15
	CauseRadioNetworkHandoverDesirableForRadioReason                          Enumerated = 16
	CauseRadioNetworkTimeCriticalHandover                                     Enumerated = 17
	CauseRadioNetworkResourceOptimisationHandover                             Enumerated = 18
	CauseRadioNetworkReduceLoadInServingCell                                  Enumerated = 19
	CauseRadioNetworkUserInactivity                                           Enumerated = 20
	CauseRadioNetworkRadioConnectionWithUeLost                                Enumerated = 21
	CauseRadioNetworkRadioResourcesNotAvailable                               Enumerated = 22
	CauseRadioNetworkInvalidQosCombination                                    Enumerated = 23
	CauseRadioNetworkFailureInRadioInterfaceProcedure                         Enumerated = 24
	CauseRadioNetworkInteractionWithOtherProcedure                            Enumerated = 25
	CauseRadioNetworkUnknownPDUSessionID                                      Enumerated = 26
	CauseRadioNetworkUnkownQosFlowID                                          Enumerated = 27
	CauseRadioNetworkMultiplePDUSessionIDInstances                            Enumerated = 28
	CauseRadioNetworkMultipleQosFlowIDInstances                               Enumerated = 29
	CauseRadioNetworkEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported Enumerated = 30
	CauseRadioNetworkNgIntraSystemHandoverTriggered                           Enumerated = 31
	CauseRadioNetworkNgInterSystemHandoverTriggered                           Enumerated = 32
	CauseRadioNetworkXnHandoverTriggered                                      Enumerated = 33
	CauseRadioNetworkNotSupported5QIValue                                     Enumerated = 34
	CauseRadioNetworkUeContextTransfer                                        Enumerated = 35
	CauseRadioNetworkImsVoiceEpsFallbackOrRatFallbackTriggered                Enumerated = 36
	CauseRadioNetworkUpIntegrityProtectionNotPossible                         Enumerated = 37
	CauseRadioNetworkUpConfidentialityProtectionNotPossible                   Enumerated = 38
	CauseRadioNetworkSliceNotSupported                                        Enumerated = 39
	CauseRadioNetworkUeInRrcInactiveStateNotReachable                         Enumerated = 40
	CauseRadioNetworkRedirection                                              Enumerated = 41
	CauseRadioNetworkResourcesNotAvailableForTheSlice                         Enumerated = 42
	CauseRadioNetworkUeMaxIntegrityProtectedDataRateReason                    Enumerated = 43
	CauseRadioNetworkReleaseDueToCnDetectedMobility                           Enumerated = 44
	CauseRadioNetworkN26InterfaceNotAvailable                                 Enumerated = 45
	CauseRadioNetworkReleaseDueToPreEmption                                   Enumerated = 46
	CauseRadioNetworkMultipleLocationReportingReferenceIDInstances            Enumerated = 47
)

type CauseRadioNetwork struct {
	Value Enumerated `vht:"valueMin:0,valueMax:47"`
}

const (
	CauseTransportTransportResourceUnavailable Enumerated = 0
	CauseTransportUnspecified                  Enumerated = 1
)

type CauseTransport struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type CellIDBroadcastEUTRA struct {
	List []CellIDBroadcastEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofCellIDforWarning"`
}

type CellIDBroadcastEUTRAItem struct {
	EUTRACGI     EUTRACGI
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs
}

type ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs struct {
	List []CellIDBroadcastEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CellIDBroadcastEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CellIDBroadcastEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CellIDBroadcastEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type CellIDBroadcastEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type CellIDBroadcastNR struct {
	List []CellIDBroadcastNRItem `vht:"sizeMin:1,sizeMax:MaxnoofCellIDforWarning"`
}

type CellIDBroadcastNRItem struct {
	NRCGI        NRCGI
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs
}

type ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs struct {
	List []CellIDBroadcastNRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CellIDBroadcastNRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CellIDBroadcastNRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CellIDBroadcastNRItemExtIEsExtensionValueChoiceNothing int = 0
)

type CellIDBroadcastNRItemExtIEsExtensionValue struct {
	Choice int
}

type CellIDCancelledEUTRA struct {
	List []CellIDCancelledEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofCellIDforWarning"`
}

type CellIDCancelledEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs
}

type ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs struct {
	List []CellIDCancelledEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CellIDCancelledEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CellIDCancelledEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CellIDCancelledEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type CellIDCancelledEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type CellIDCancelledNR struct {
	List []CellIDCancelledNRItem `vht:"sizeMin:1,sizeMax:MaxnoofCellIDforWarning"`
}

type CellIDCancelledNRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledNRItemExtIEs
}

type ProtocolExtensionContainerCellIDCancelledNRItemExtIEs struct {
	List []CellIDCancelledNRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CellIDCancelledNRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CellIDCancelledNRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CellIDCancelledNRItemExtIEsExtensionValueChoiceNothing int = 0
)

type CellIDCancelledNRItemExtIEsExtensionValue struct {
	Choice int
}

const (
	CellIDListForRestartChoiceEUTRACGIListforRestart int = 0
	CellIDListForRestartChoiceNRCGIListforRestart    int = 1
	CellIDListForRestartChoiceChoiceExtensions       int = 2
)

type CellIDListForRestart struct {
	Choice                 int
	EUTRACGIListforRestart *EUTRACGIList
	NRCGIListforRestart    *NRCGIList
	ChoiceExtensions       *ProtocolIESingleContainerCellIDListForRestartExtIEs
}

type ProtocolIESingleContainerCellIDListForRestartExtIEs struct {
	List []CellIDListForRestartExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type CellIDListForRestartExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	TypeValue    CellIDListForRestartExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	CellIDListForRestartExtIEsTypeValueChoiceNothing int = 0
)

type CellIDListForRestartExtIEsTypeValue struct {
	Choice int
}

const (
	CellSizeVerysmall Enumerated = 0
	CellSizeSmall     Enumerated = 1
	CellSizeMedium    Enumerated = 2
	CellSizeLarge     Enumerated = 3
)

type CellSize struct {
	Value Enumerated `vht:"valueMin:0,valueMax:3"`
}

type CellType struct {
	CellSize     CellSize
	IEExtensions *ProtocolExtensionContainerCellTypeExtIEs
}

type ProtocolExtensionContainerCellTypeExtIEs struct {
	List []CellTypeExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CellTypeExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CellTypeExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CellTypeExtIEsExtensionValueChoiceNothing int = 0
)

type CellTypeExtIEsExtensionValue struct {
	Choice int
}

type CNAssistedRANTuning struct {
	ExpectedUEBehaviour *ExpectedUEBehaviour
	IEExtensions        *ProtocolExtensionContainerCNAssistedRANTuningExtIEs
}

type ProtocolExtensionContainerCNAssistedRANTuningExtIEs struct {
	List []CNAssistedRANTuningExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CNAssistedRANTuningExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CNAssistedRANTuningExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CNAssistedRANTuningExtIEsExtensionValueChoiceNothing int = 0
)

type CNAssistedRANTuningExtIEsExtensionValue struct {
	Choice int
}

type CNTypeRestrictionsForEquivalent struct {
	List []CNTypeRestrictionsForEquivalentItem `vht:"sizeMin:1,sizeMax:MaxnoofEPLMNs"`
}

const (
	CnTypeEpcForbidden    Enumerated = 0
	CnTypeFiveGCForbidden Enumerated = 1
)

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity PLMNIdentity
	CnType       Enumerated `vht:"valueExt,valueMin:0,valueMax:2"`
	IEExtensions *ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs
}

type ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs struct {
	List []CNTypeRestrictionsForEquivalentItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CNTypeRestrictionsForEquivalentItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CNTypeRestrictionsForEquivalentItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CNTypeRestrictionsForEquivalentItemExtIEsExtensionValueChoiceNothing int = 0
)

type CNTypeRestrictionsForEquivalentItemExtIEsExtensionValue struct {
	Choice int
}

const (
	CNTypeRestrictionsForServingEpcForbidden Enumerated = 0
)

type CNTypeRestrictionsForServing struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type CommonNetworkInstance struct {
	Value OctetString
}

type CompletedCellsInEAIEUTRA struct {
	List []CompletedCellsInEAIEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinEAI"`
}

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI     EUTRACGI
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs
}

type ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs struct {
	List []CompletedCellsInEAIEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CompletedCellsInEAIEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CompletedCellsInEAIEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CompletedCellsInEAIEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type CompletedCellsInEAIEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type CompletedCellsInEAINR struct {
	List []CompletedCellsInEAINRItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinEAI"`
}

type CompletedCellsInEAINRItem struct {
	NRCGI        NRCGI
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs
}

type ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs struct {
	List []CompletedCellsInEAINRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CompletedCellsInEAINRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CompletedCellsInEAINRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CompletedCellsInEAINRItemExtIEsExtensionValueChoiceNothing int = 0
)

type CompletedCellsInEAINRItemExtIEsExtensionValue struct {
	Choice int
}

type CompletedCellsInTAIEUTRA struct {
	List []CompletedCellsInTAIEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinTAI"`
}

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI     EUTRACGI
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs
}

type ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs struct {
	List []CompletedCellsInTAIEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CompletedCellsInTAIEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CompletedCellsInTAIEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CompletedCellsInTAIEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type CompletedCellsInTAIEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type CompletedCellsInTAINR struct {
	List []CompletedCellsInTAINRItem `vht:"sizeMin:1,sizeMax:MaxnoofCellinTAI"`
}

type CompletedCellsInTAINRItem struct {
	NRCGI        NRCGI
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs
}

type ProtocolExtensionContainerCompletedCellsInTAINRItemExtIEs struct {
	List []CompletedCellsInTAINRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CompletedCellsInTAINRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CompletedCellsInTAINRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CompletedCellsInTAINRItemExtIEsExtensionValueChoiceNothing int = 0
)

type CompletedCellsInTAINRItemExtIEsExtensionValue struct {
	Choice int
}

const (
	ConcurrentWarningMessageIndTrue Enumerated = 0
)

type ConcurrentWarningMessageInd struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	ConfidentialityProtectionIndicationRequired  Enumerated = 0
	ConfidentialityProtectionIndicationPreferred Enumerated = 1
	ConfidentialityProtectionIndicationNotNeeded Enumerated = 2
)

type ConfidentialityProtectionIndication struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	ConfidentialityProtectionResultPerformed    Enumerated = 0
	ConfidentialityProtectionResultNotPerformed Enumerated = 1
)

type ConfidentialityProtectionResult struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            UEIdentityIndexValue
	UESpecificDRX                   *PagingDRX
	PeriodicRegistrationUpdateTimer PeriodicRegistrationUpdateTimer
	MICOModeIndication              *MICOModeIndication
	TAIListForInactive              TAIListForInactive
	ExpectedUEBehaviour             *ExpectedUEBehaviour
	IEExtensions                    *ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs
}

type ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs struct {
	List []CoreNetworkAssistanceInformationForInactiveExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CoreNetworkAssistanceInformationForInactiveExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValueChoiceNothing int = 0
)

type CoreNetworkAssistanceInformationForInactiveExtIEsExtensionValue struct {
	Choice int
}

type COUNTValueForPDCPSN12 struct {
	PDCPSN12     Integer `vht:"valueMin:0,valueMax:4095"`
	HFNPDCPSN12  Integer `vht:"valueMin:0,valueMax:1048575"`
	IEExtensions *ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs
}

type ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs struct {
	List []COUNTValueForPDCPSN12ExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type COUNTValueForPDCPSN12ExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      COUNTValueForPDCPSN12ExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	COUNTValueForPDCPSN12ExtIEsExtensionValueChoiceNothing int = 0
)

type COUNTValueForPDCPSN12ExtIEsExtensionValue struct {
	Choice int
}

type COUNTValueForPDCPSN18 struct {
	PDCPSN18     Integer `vht:"valueMin:0,valueMax:262143"`
	HFNPDCPSN18  Integer `vht:"valueMin:0,valueMax:16383"`
	IEExtensions *ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs
}

type ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs struct {
	List []COUNTValueForPDCPSN18ExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type COUNTValueForPDCPSN18ExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      COUNTValueForPDCPSN18ExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	COUNTValueForPDCPSN18ExtIEsExtensionValueChoiceNothing int = 0
)

type COUNTValueForPDCPSN18ExtIEsExtensionValue struct {
	Choice int
}

const (
	CPTransportLayerInformationChoiceEndpointIPAddress int = 0
	CPTransportLayerInformationChoiceChoiceExtensions  int = 1
)

type CPTransportLayerInformation struct {
	Choice            int
	EndpointIPAddress *TransportLayerAddress
	ChoiceExtensions  *ProtocolIESingleContainerCPTransportLayerInformationExtIEs
}

type ProtocolIESingleContainerCPTransportLayerInformationExtIEs struct {
	List []CPTransportLayerInformationExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type CPTransportLayerInformationExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	TypeValue    CPTransportLayerInformationExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	CPTransportLayerInformationExtIEsTypeValueChoiceEndpointIPAddressAndPort int = 0
)

type CPTransportLayerInformationExtIEsTypeValue struct {
	Choice                   int
	EndpointIPAddressAndPort *EndpointIPAddressAndPort `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDEndpointIPAddressAndPort"`
}

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode
	TriggeringMessage         *TriggeringMessage
	ProcedureCriticality      *Criticality
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList
	IEExtensions              *ProtocolExtensionContainerCriticalityDiagnosticsExtIEs
}

type ProtocolExtensionContainerCriticalityDiagnosticsExtIEs struct {
	List []CriticalityDiagnosticsExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CriticalityDiagnosticsExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CriticalityDiagnosticsExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CriticalityDiagnosticsExtIEsExtensionValueChoiceNothing int = 0
)

type CriticalityDiagnosticsExtIEsExtensionValue struct {
	Choice int
}

type CriticalityDiagnosticsIEList struct {
	List []CriticalityDiagnosticsIEItem `vht:"sizeMin:1,sizeMax:MaxnoofErrors"`
}

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality
	IEID          ProtocolIEID
	TypeOfError   TypeOfError
	IEExtensions  *ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs
}

type ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs struct {
	List []CriticalityDiagnosticsIEItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type CriticalityDiagnosticsIEItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      CriticalityDiagnosticsIEItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	CriticalityDiagnosticsIEItemExtIEsExtensionValueChoiceNothing int = 0
)

type CriticalityDiagnosticsIEItemExtIEsExtensionValue struct {
	Choice int
}

type DataCodingScheme struct {
	Value BitString `vht:"sizeMin:8,sizeMax:8"`
}

const (
	DataForwardingAcceptedDataForwardingAccepted Enumerated = 0
)

type DataForwardingAccepted struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	DataForwardingNotPossibleDataForwardingNotPossible Enumerated = 0
)

type DataForwardingNotPossible struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type DataForwardingResponseDRBList struct {
	List []DataForwardingResponseDRBItem `vht:"sizeMin:1,sizeMax:MaxnoofDRBs"`
}

type DataForwardingResponseDRBItem struct {
	DRBID                        DRBID
	DLForwardingUPTNLInformation *UPTransportLayerInformation
	ULForwardingUPTNLInformation *UPTransportLayerInformation
	IEExtensions                 *ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs
}

type ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs struct {
	List []DataForwardingResponseDRBItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DataForwardingResponseDRBItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DataForwardingResponseDRBItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DataForwardingResponseDRBItemExtIEsExtensionValueChoiceNothing int = 0
)

type DataForwardingResponseDRBItemExtIEsExtensionValue struct {
	Choice int
}

type DataForwardingResponseERABList struct {
	List []DataForwardingResponseERABListItem `vht:"sizeMin:1,sizeMax:MaxnoofE-RABs"`
}

type DataForwardingResponseERABListItem struct {
	ERABID       ERABID
	DLForwarding *DLForwarding
	IEExtensions *ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs
}

type ProtocolExtensionContainerDataForwardingResponseERABListItemExtIEs struct {
	List []DataForwardingResponseERABListItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DataForwardingResponseERABListItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DataForwardingResponseERABListItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DataForwardingResponseERABListItemExtIEsExtensionValueChoiceNothing int = 0
)

type DataForwardingResponseERABListItemExtIEsExtensionValue struct {
	Choice int
}

const (
	DelayCriticalDelayCritical    Enumerated = 0
	DelayCriticalNonDelayCritical Enumerated = 1
)

type DelayCritical struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

const (
	DLForwardingDlForwardingProposed Enumerated = 0
)

type DLForwarding struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	DLNGUTNLInformationReusedTrue Enumerated = 0
)

type DLNGUTNLInformationReused struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	DirectForwardingPathAvailabilityDirectPathAvailable Enumerated = 0
)

type DirectForwardingPathAvailability struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type DRBID struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:32"`
}

type DRBsSubjectToStatusTransferList struct {
	List []DRBsSubjectToStatusTransferItem `vht:"sizeMin:1,sizeMax:MaxnoofDRBs"`
}

type DRBsSubjectToStatusTransferItem struct {
	DRBID        DRBID
	DRBStatusUL  DRBStatusUL
	DRBStatusDL  DRBStatusDL
	IEExtensions *ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs
}

type ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs struct {
	List []DRBsSubjectToStatusTransferItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DRBsSubjectToStatusTransferItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DRBsSubjectToStatusTransferItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DRBsSubjectToStatusTransferItemExtIEsExtensionValueChoiceOldAssociatedQosFlowListULendmarkerexpected int = 0
)

type DRBsSubjectToStatusTransferItemExtIEsExtensionValue struct {
	Choice                                      int
	OldAssociatedQosFlowListULendmarkerexpected *AssociatedQosFlowList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDOldAssociatedQosFlowListULendmarkerexpected"`
}

const (
	DRBStatusDLChoiceDRBStatusDL12    int = 0
	DRBStatusDLChoiceDRBStatusDL18    int = 1
	DRBStatusDLChoiceChoiceExtensions int = 2
)

type DRBStatusDL struct {
	Choice           int
	DRBStatusDL12    *DRBStatusDL12
	DRBStatusDL18    *DRBStatusDL18
	ChoiceExtensions *ProtocolIESingleContainerDRBStatusDLExtIEs
}

type ProtocolIESingleContainerDRBStatusDLExtIEs struct {
	List []DRBStatusDLExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type DRBStatusDLExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    DRBStatusDLExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DRBStatusDLExtIEsTypeValueChoiceNothing int = 0
)

type DRBStatusDLExtIEsTypeValue struct {
	Choice int
}

type DRBStatusDL12 struct {
	DLCOUNTValue COUNTValueForPDCPSN12
	IEExtensions *ProtocolExtensionContainerDRBStatusDL12ExtIEs
}

type ProtocolExtensionContainerDRBStatusDL12ExtIEs struct {
	List []DRBStatusDL12ExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DRBStatusDL12ExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DRBStatusDL12ExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DRBStatusDL12ExtIEsExtensionValueChoiceNothing int = 0
)

type DRBStatusDL12ExtIEsExtensionValue struct {
	Choice int
}

type DRBStatusDL18 struct {
	DLCOUNTValue COUNTValueForPDCPSN18
	IEExtensions *ProtocolExtensionContainerDRBStatusDL18ExtIEs
}

type ProtocolExtensionContainerDRBStatusDL18ExtIEs struct {
	List []DRBStatusDL18ExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DRBStatusDL18ExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DRBStatusDL18ExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DRBStatusDL18ExtIEsExtensionValueChoiceNothing int = 0
)

type DRBStatusDL18ExtIEsExtensionValue struct {
	Choice int
}

const (
	DRBStatusULChoiceDRBStatusUL12    int = 0
	DRBStatusULChoiceDRBStatusUL18    int = 1
	DRBStatusULChoiceChoiceExtensions int = 2
)

type DRBStatusUL struct {
	Choice           int
	DRBStatusUL12    *DRBStatusUL12
	DRBStatusUL18    *DRBStatusUL18
	ChoiceExtensions *ProtocolIESingleContainerDRBStatusULExtIEs
}

type ProtocolIESingleContainerDRBStatusULExtIEs struct {
	List []DRBStatusULExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type DRBStatusULExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                `vht:"Reference:ProtocolIEID"`
	TypeValue    DRBStatusULExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	DRBStatusULExtIEsTypeValueChoiceNothing int = 0
)

type DRBStatusULExtIEsTypeValue struct {
	Choice int
}

type DRBStatusUL12 struct {
	ULCOUNTValue              COUNTValueForPDCPSN12
	ReceiveStatusOfULPDCPSDUs *BitString `vht:"valueMin:1,valueMax:2048"`
	IEExtensions              *ProtocolExtensionContainerDRBStatusUL12ExtIEs
}

type ProtocolExtensionContainerDRBStatusUL12ExtIEs struct {
	List []DRBStatusUL12ExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DRBStatusUL12ExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DRBStatusUL12ExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DRBStatusUL12ExtIEsExtensionValueChoiceNothing int = 0
)

type DRBStatusUL12ExtIEsExtensionValue struct {
	Choice int
}

type DRBStatusUL18 struct {
	ULCOUNTValue              COUNTValueForPDCPSN18
	ReceiveStatusOfULPDCPSDUs *BitString `vht:"valueMin:1,valueMax:131072"`
	IEExtensions              *ProtocolExtensionContainerDRBStatusUL18ExtIEs
}

type ProtocolExtensionContainerDRBStatusUL18ExtIEs struct {
	List []DRBStatusUL18ExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DRBStatusUL18ExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DRBStatusUL18ExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DRBStatusUL18ExtIEsExtensionValueChoiceNothing int = 0
)

type DRBStatusUL18ExtIEsExtensionValue struct {
	Choice int
}

type DRBsToQosFlowsMappingList struct {
	List []DRBsToQosFlowsMappingItem `vht:"sizeMin:1,sizeMax:MaxnoofDRBs"`
}

type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID
	AssociatedQosFlowList AssociatedQosFlowList
	IEExtensions          *ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs
}

type ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs struct {
	List []DRBsToQosFlowsMappingItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type DRBsToQosFlowsMappingItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      DRBsToQosFlowsMappingItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	DRBsToQosFlowsMappingItemExtIEsExtensionValueChoiceNothing int = 0
)

type DRBsToQosFlowsMappingItemExtIEsExtensionValue struct {
	Choice int
}

type Dynamic5QIDescriptor struct {
	PriorityLevelQos       PriorityLevelQos
	PacketDelayBudget      PacketDelayBudget
	PacketErrorRate        PacketErrorRate
	FiveQI                 *FiveQI
	DelayCritical          *DelayCritical
	AveragingWindow        *AveragingWindow
	MaximumDataBurstVolume *MaximumDataBurstVolume
	IEExtensions           *ProtocolExtensionContainerDynamic5QIDescriptorExtIEs
}

type ProtocolExtensionContainerDynamic5QIDescriptorExtIEs struct {
	List []Dynamic5QIDescriptorExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type Dynamic5QIDescriptorExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      Dynamic5QIDescriptorExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	Dynamic5QIDescriptorExtIEsExtensionValueChoiceNothing int = 0
)

type Dynamic5QIDescriptorExtIEsExtensionValue struct {
	Choice int
}

type EmergencyAreaID struct {
	Value OctetString `vht:"sizeMin:3,sizeMax:3"`
}

type EmergencyAreaIDBroadcastEUTRA struct {
	List []EmergencyAreaIDBroadcastEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofEmergencyAreaID"`
}

type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID
	CompletedCellsInEAIEUTRA CompletedCellsInEAIEUTRA
	IEExtensions             *ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs
}

type ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs struct {
	List []EmergencyAreaIDBroadcastEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EmergencyAreaIDBroadcastEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type EmergencyAreaIDBroadcastEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type EmergencyAreaIDBroadcastNR struct {
	List []EmergencyAreaIDBroadcastNRItem `vht:"sizeMin:1,sizeMax:MaxnoofEmergencyAreaID"`
}

type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       EmergencyAreaID
	CompletedCellsInEAINR CompletedCellsInEAINR
	IEExtensions          *ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs
}

type ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs struct {
	List []EmergencyAreaIDBroadcastNRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EmergencyAreaIDBroadcastNRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EmergencyAreaIDBroadcastNRItemExtIEsExtensionValueChoiceNothing int = 0
)

type EmergencyAreaIDBroadcastNRItemExtIEsExtensionValue struct {
	Choice int
}

type EmergencyAreaIDCancelledEUTRA struct {
	List []EmergencyAreaIDCancelledEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofEmergencyAreaID"`
}

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID
	CancelledCellsInEAIEUTRA CancelledCellsInEAIEUTRA
	IEExtensions             *ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs
}

type ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs struct {
	List []EmergencyAreaIDCancelledEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EmergencyAreaIDCancelledEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type EmergencyAreaIDCancelledEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type EmergencyAreaIDCancelledNR struct {
	List []EmergencyAreaIDCancelledNRItem `vht:"sizeMin:1,sizeMax:MaxnoofEmergencyAreaID"`
}

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       EmergencyAreaID
	CancelledCellsInEAINR CancelledCellsInEAINR
	IEExtensions          *ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs
}

type ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs struct {
	List []EmergencyAreaIDCancelledNRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EmergencyAreaIDCancelledNRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EmergencyAreaIDCancelledNRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EmergencyAreaIDCancelledNRItemExtIEsExtensionValueChoiceNothing int = 0
)

type EmergencyAreaIDCancelledNRItemExtIEsExtensionValue struct {
	Choice int
}

type EmergencyAreaIDList struct {
	List []EmergencyAreaID `vht:"sizeMin:1,sizeMax:MaxnoofEmergencyAreaID"`
}

type EmergencyAreaIDListForRestart struct {
	List []EmergencyAreaID `vht:"sizeMin:1,sizeMax:MaxnoofEAIforRestart"`
}

type EmergencyFallbackIndicator struct {
	EmergencyFallbackRequestIndicator EmergencyFallbackRequestIndicator
	EmergencyServiceTargetCN          *EmergencyServiceTargetCN
	IEExtensions                      *ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs
}

type ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs struct {
	List []EmergencyFallbackIndicatorExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EmergencyFallbackIndicatorExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EmergencyFallbackIndicatorExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EmergencyFallbackIndicatorExtIEsExtensionValueChoiceNothing int = 0
)

type EmergencyFallbackIndicatorExtIEsExtensionValue struct {
	Choice int
}

const (
	EmergencyFallbackRequestIndicatorEmergencyFallbackRequested Enumerated = 0
)

type EmergencyFallbackRequestIndicator struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	EmergencyServiceTargetCNFiveGC Enumerated = 0
	EmergencyServiceTargetCNEpc    Enumerated = 1
)

type EmergencyServiceTargetCN struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type ENDCSONConfigurationTransfer struct {
	Value OctetString
}

type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress
	PortNumber        PortNumber
	IEExtensions      *ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs
}

type ProtocolExtensionContainerEndpointIPAddressAndPortExtIEs struct {
	List []EndpointIPAddressAndPortExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EndpointIPAddressAndPortExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EndpointIPAddressAndPortExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EndpointIPAddressAndPortExtIEsExtensionValueChoiceNothing int = 0
)

type EndpointIPAddressAndPortExtIEsExtensionValue struct {
	Choice int
}

type EquivalentPLMNs struct {
	List []PLMNIdentity `vht:"sizeMin:1,sizeMax:MaxnoofEPLMNs"`
}

type EPSTAC struct {
	Value OctetString `vht:"sizeMin:2,sizeMax:2"`
}

type EPSTAI struct {
	PLMNIdentity PLMNIdentity
	EPSTAC       EPSTAC
	IEExtensions *ProtocolExtensionContainerEPSTAIExtIEs
}

type ProtocolExtensionContainerEPSTAIExtIEs struct {
	List []EPSTAIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EPSTAIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EPSTAIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EPSTAIExtIEsExtensionValueChoiceNothing int = 0
)

type EPSTAIExtIEsExtensionValue struct {
	Choice int
}

type ERABID struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:15"`
}

type ERABInformationList struct {
	List []ERABInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofE-RABs"`
}

type ERABInformationItem struct {
	ERABID       ERABID
	DLForwarding *DLForwarding
	IEExtensions *ProtocolExtensionContainerERABInformationItemExtIEs
}

type ProtocolExtensionContainerERABInformationItemExtIEs struct {
	List []ERABInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ERABInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ERABInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ERABInformationItemExtIEsExtensionValueChoiceNothing int = 0
)

type ERABInformationItemExtIEsExtensionValue struct {
	Choice int
}

type EUTRACellIdentity struct {
	Value BitString `vht:"sizeMin:28,sizeMax:28"`
}

type EUTRACGI struct {
	PLMNIdentity      PLMNIdentity
	EUTRACellIdentity EUTRACellIdentity
	IEExtensions      *ProtocolExtensionContainerEUTRACGIExtIEs
}

type ProtocolExtensionContainerEUTRACGIExtIEs struct {
	List []EUTRACGIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type EUTRACGIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      EUTRACGIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	EUTRACGIExtIEsExtensionValueChoiceNothing int = 0
)

type EUTRACGIExtIEsExtensionValue struct {
	Choice int
}

type EUTRACGIList struct {
	List []EUTRACGI `vht:"sizeMin:1,sizeMax:MaxnoofCellsinngeNB"`
}

type EUTRACGIListForWarning struct {
	List []EUTRACGI `vht:"sizeMin:1,sizeMax:MaxnoofCellIDforWarning"`
}

type EUTRAencryptionAlgorithms struct {
	Value BitString `vht:"sizeExx,sizeMin:16,sizeMax:16"`
}

type EUTRAintegrityProtectionAlgorithms struct {
	Value BitString `vht:"sizeExx,sizeMin:16,sizeMax:16"`
}

const (
	EventTypeDirect                          Enumerated = 0
	EventTypeChangeOfServeCell               Enumerated = 1
	EventTypeUePresenceInAreaOfInterest      Enumerated = 2
	EventTypeStopChangeOfServeCell           Enumerated = 3
	EventTypeStopUePresenceInAreaOfInterest  Enumerated = 4
	EventTypeCancelLocationReportingForTheUe Enumerated = 5
)

type EventType struct {
	Value Enumerated `vht:"valueMin:0,valueMax:5"`
}

type ExpectedActivityPeriod struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:30|40|50|60|80|100|120|150|180|181"`
}

const (
	ExpectedHOIntervalSec15    Enumerated = 0
	ExpectedHOIntervalSec30    Enumerated = 1
	ExpectedHOIntervalSec60    Enumerated = 2
	ExpectedHOIntervalSec90    Enumerated = 3
	ExpectedHOIntervalSec120   Enumerated = 4
	ExpectedHOIntervalSec180   Enumerated = 5
	ExpectedHOIntervalLongTime Enumerated = 6
)

type ExpectedHOInterval struct {
	Value Enumerated `vht:"valueMin:0,valueMax:6"`
}

type ExpectedIdlePeriod struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:30|40|50|60|80|100|120|150|180|181"`
}

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 *ExpectedActivityPeriod
	ExpectedIdlePeriod                     *ExpectedIdlePeriod
	SourceOfUEActivityBehaviourInformation *SourceOfUEActivityBehaviourInformation
	IEExtensions                           *ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs
}

type ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs struct {
	List []ExpectedUEActivityBehaviourExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ExpectedUEActivityBehaviourExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ExpectedUEActivityBehaviourExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ExpectedUEActivityBehaviourExtIEsExtensionValueChoiceNothing int = 0
)

type ExpectedUEActivityBehaviourExtIEsExtensionValue struct {
	Choice int
}

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour *ExpectedUEActivityBehaviour
	ExpectedHOInterval          *ExpectedHOInterval
	ExpectedUEMobility          *ExpectedUEMobility
	ExpectedUEMovingTrajectory  *ExpectedUEMovingTrajectory
	IEExtensions                *ProtocolExtensionContainerExpectedUEBehaviourExtIEs
}

type ProtocolExtensionContainerExpectedUEBehaviourExtIEs struct {
	List []ExpectedUEBehaviourExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ExpectedUEBehaviourExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ExpectedUEBehaviourExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ExpectedUEBehaviourExtIEsExtensionValueChoiceNothing int = 0
)

type ExpectedUEBehaviourExtIEsExtensionValue struct {
	Choice int
}

const (
	ExpectedUEMobilityStationary Enumerated = 0
	ExpectedUEMobilityMobile     Enumerated = 1
)

type ExpectedUEMobility struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type ExpectedUEMovingTrajectory struct {
	List []ExpectedUEMovingTrajectoryItem `vht:"sizeMin:1,sizeMax:MaxnoofCellsUEMovingTrajectory"`
}

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         NGRANCGI
	TimeStayedInCell *Integer `vht:"valueMin:0,valueMax:4095"`
	IEExtensions     *ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs
}

type ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs struct {
	List []ExpectedUEMovingTrajectoryItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ExpectedUEMovingTrajectoryItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ExpectedUEMovingTrajectoryItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ExpectedUEMovingTrajectoryItemExtIEsExtensionValueChoiceNothing int = 0
)

type ExpectedUEMovingTrajectoryItemExtIEsExtensionValue struct {
	Choice int
}

type ExtendedRATRestrictionInformation struct {
	PrimaryRATRestriction   BitString `vht:"valueExt,valueMin:8,valueMax:8"`
	SecondaryRATRestriction BitString `vht:"valueExt,valueMin:8,valueMax:8"`
	IEExtensions            *ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs
}

type ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs struct {
	List []ExtendedRATRestrictionInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ExtendedRATRestrictionInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ExtendedRATRestrictionInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ExtendedRATRestrictionInformationExtIEsExtensionValueChoiceNothing int = 0
)

type ExtendedRATRestrictionInformationExtIEsExtensionValue struct {
	Choice int
}

type ExtendedRNCID struct {
	Value Integer `vht:"valueMin:4096,valueMax:65535"`
}

type FiveGSTMSI struct {
	AMFSetID     AMFSetID
	AMFPointer   AMFPointer
	FiveGTMSI    FiveGTMSI
	IEExtensions *ProtocolExtensionContainerFiveGSTMSIExtIEs
}

type ProtocolExtensionContainerFiveGSTMSIExtIEs struct {
	List []FiveGSTMSIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type FiveGSTMSIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      FiveGSTMSIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	FiveGSTMSIExtIEsExtensionValueChoiceNothing int = 0
)

type FiveGSTMSIExtIEsExtensionValue struct {
	Choice int
}

type FiveGTMSI struct {
	Value OctetString `vht:"sizeMin:4,sizeMax:4"`
}

type FiveQI struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:255"`
}

type ForbiddenAreaInformation struct {
	List []ForbiddenAreaInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofEPLMNsPlusOne"`
}

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  PLMNIdentity
	ForbiddenTACs ForbiddenTACs
	IEExtensions  *ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs
}

type ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs struct {
	List []ForbiddenAreaInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ForbiddenAreaInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ForbiddenAreaInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ForbiddenAreaInformationItemExtIEsExtensionValueChoiceNothing int = 0
)

type ForbiddenAreaInformationItemExtIEsExtensionValue struct {
	Choice int
}

type ForbiddenTACs struct {
	List []TAC `vht:"sizeMin:1,sizeMax:MaxnoofForbTACs"`
}

type GBRQosInformation struct {
	MaximumFlowBitRateDL    BitRate
	MaximumFlowBitRateUL    BitRate
	GuaranteedFlowBitRateDL BitRate
	GuaranteedFlowBitRateUL BitRate
	NotificationControl     *NotificationControl
	MaximumPacketLossRateDL *PacketLossRate
	MaximumPacketLossRateUL *PacketLossRate
	IEExtensions            *ProtocolExtensionContainerGBRQosInformationExtIEs
}

type ProtocolExtensionContainerGBRQosInformationExtIEs struct {
	List []GBRQosInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type GBRQosInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      GBRQosInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	GBRQosInformationExtIEsExtensionValueChoiceNothing int = 0
)

type GBRQosInformationExtIEsExtensionValue struct {
	Choice int
}

type GlobalGNBID struct {
	PLMNIdentity PLMNIdentity
	GNBID        GNBID
	IEExtensions *ProtocolExtensionContainerGlobalGNBIDExtIEs
}

type ProtocolExtensionContainerGlobalGNBIDExtIEs struct {
	List []GlobalGNBIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type GlobalGNBIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      GlobalGNBIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	GlobalGNBIDExtIEsExtensionValueChoiceNothing int = 0
)

type GlobalGNBIDExtIEsExtensionValue struct {
	Choice int
}

type GlobalN3IWFID struct {
	PLMNIdentity PLMNIdentity
	N3IWFID      N3IWFID
	IEExtensions *ProtocolExtensionContainerGlobalN3IWFIDExtIEs
}

type ProtocolExtensionContainerGlobalN3IWFIDExtIEs struct {
	List []GlobalN3IWFIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type GlobalN3IWFIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      GlobalN3IWFIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	GlobalN3IWFIDExtIEsExtensionValueChoiceNothing int = 0
)

type GlobalN3IWFIDExtIEsExtensionValue struct {
	Choice int
}

type GlobalNgENBID struct {
	PLMNIdentity PLMNIdentity
	NgENBID      NgENBID
	IEExtensions *ProtocolExtensionContainerGlobalNgENBIDExtIEs
}

type ProtocolExtensionContainerGlobalNgENBIDExtIEs struct {
	List []GlobalNgENBIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type GlobalNgENBIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      GlobalNgENBIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	GlobalNgENBIDExtIEsExtensionValueChoiceNothing int = 0
)

type GlobalNgENBIDExtIEsExtensionValue struct {
	Choice int
}

const (
	GlobalRANNodeIDChoiceGlobalGNBID      int = 0
	GlobalRANNodeIDChoiceGlobalNgENBID    int = 1
	GlobalRANNodeIDChoiceGlobalN3IWFID    int = 2
	GlobalRANNodeIDChoiceChoiceExtensions int = 3
)

type GlobalRANNodeID struct {
	Choice           int
	GlobalGNBID      *GlobalGNBID
	GlobalNgENBID    *GlobalNgENBID
	GlobalN3IWFID    *GlobalN3IWFID
	ChoiceExtensions *ProtocolIESingleContainerGlobalRANNodeIDExtIEs
}

type ProtocolIESingleContainerGlobalRANNodeIDExtIEs struct {
	List []GlobalRANNodeIDExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type GlobalRANNodeIDExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	TypeValue    GlobalRANNodeIDExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	GlobalRANNodeIDExtIEsTypeValueChoiceNothing int = 0
)

type GlobalRANNodeIDExtIEsTypeValue struct {
	Choice int
}

const (
	GNBIDChoiceGNBID            int = 0
	GNBIDChoiceChoiceExtensions int = 1
)

type GNBID struct {
	Choice           int
	GNBID            *BitString `vht:"valueMin:22,valueMax:32"`
	ChoiceExtensions *ProtocolIESingleContainerGNBIDExtIEs
}

type ProtocolIESingleContainerGNBIDExtIEs struct {
	List []GNBIDExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type GNBIDExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality          `vht:"Reference:ProtocolIEID"`
	TypeValue    GNBIDExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	GNBIDExtIEsTypeValueChoiceNothing int = 0
)

type GNBIDExtIEsTypeValue struct {
	Choice int
}

type GTPTEID struct {
	Value OctetString `vht:"sizeMin:4,sizeMax:4"`
}

type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress
	GTPTEID               GTPTEID
	IEExtensions          *ProtocolExtensionContainerGTPTunnelExtIEs
}

type ProtocolExtensionContainerGTPTunnelExtIEs struct {
	List []GTPTunnelExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type GTPTunnelExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      GTPTunnelExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	GTPTunnelExtIEsExtensionValueChoiceNothing int = 0
)

type GTPTunnelExtIEsExtensionValue struct {
	Choice int
}

type GUAMI struct {
	PLMNIdentity PLMNIdentity
	AMFRegionID  AMFRegionID
	AMFSetID     AMFSetID
	AMFPointer   AMFPointer
	IEExtensions *ProtocolExtensionContainerGUAMIExtIEs
}

type ProtocolExtensionContainerGUAMIExtIEs struct {
	List []GUAMIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type GUAMIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      GUAMIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	GUAMIExtIEsExtensionValueChoiceNothing int = 0
)

type GUAMIExtIEsExtensionValue struct {
	Choice int
}

const (
	GUAMITypeNative Enumerated = 0
	GUAMITypeMapped Enumerated = 1
)

type GUAMIType struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  *UPTransportLayerInformation
	QosFlowToBeForwardedList      *QosFlowToBeForwardedList
	DataForwardingResponseDRBList *DataForwardingResponseDRBList
	IEExtensions                  *ProtocolExtensionContainerHandoverCommandTransferExtIEs
}

type ProtocolExtensionContainerHandoverCommandTransferExtIEs struct {
	List []HandoverCommandTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type HandoverCommandTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      HandoverCommandTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	HandoverCommandTransferExtIEsExtensionValueChoiceAdditionalDLForwardingUPTNLInformation int = 0
	HandoverCommandTransferExtIEsExtensionValueChoiceULForwardingUPTNLInformation           int = 1
	HandoverCommandTransferExtIEsExtensionValueChoiceAdditionalULForwardingUPTNLInformation int = 2
	HandoverCommandTransferExtIEsExtensionValueChoiceDataForwardingResponseERABList         int = 3
)

type HandoverCommandTransferExtIEsExtensionValue struct {
	Choice                                 int
	AdditionalDLForwardingUPTNLInformation *QosFlowPerTNLInformationList    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDAdditionalDLForwardingUPTNLInformation"`
	ULForwardingUPTNLInformation           *UPTransportLayerInformation     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDULForwardingUPTNLInformation"`
	AdditionalULForwardingUPTNLInformation *UPTransportLayerInformationList `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDAdditionalULForwardingUPTNLInformation"`
	DataForwardingResponseERABList         *DataForwardingResponseERABList  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDDataForwardingResponseERABList"`
}

const (
	HandoverFlagHandoverPreparation Enumerated = 0
)

type HandoverFlag struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type HandoverPreparationUnsuccessfulTransfer struct {
	Cause        Cause
	IEExtensions *ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs
}

type ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs struct {
	List []HandoverPreparationUnsuccessfulTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type HandoverPreparationUnsuccessfulTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      HandoverPreparationUnsuccessfulTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	HandoverPreparationUnsuccessfulTransferExtIEsExtensionValueChoiceNothing int = 0
)

type HandoverPreparationUnsuccessfulTransferExtIEsExtensionValue struct {
	Choice int
}

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         UPTransportLayerInformation
	DLForwardingUPTNLInformation  *UPTransportLayerInformation
	SecurityResult                *SecurityResult
	QosFlowSetupResponseList      QosFlowListWithDataForwarding
	QosFlowFailedToSetupList      *QosFlowListWithCause
	DataForwardingResponseDRBList *DataForwardingResponseDRBList
	IEExtensions                  *ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs
}

type ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs struct {
	List []HandoverRequestAcknowledgeTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type HandoverRequestAcknowledgeTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      HandoverRequestAcknowledgeTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	HandoverRequestAcknowledgeTransferExtIEsExtensionValueChoiceAdditionalDLUPTNLInformationForHOList  int = 0
	HandoverRequestAcknowledgeTransferExtIEsExtensionValueChoiceULForwardingUPTNLInformation           int = 1
	HandoverRequestAcknowledgeTransferExtIEsExtensionValueChoiceAdditionalULForwardingUPTNLInformation int = 2
	HandoverRequestAcknowledgeTransferExtIEsExtensionValueChoiceDataForwardingResponseERABList         int = 3
)

type HandoverRequestAcknowledgeTransferExtIEsExtensionValue struct {
	Choice                                 int
	AdditionalDLUPTNLInformationForHOList  *AdditionalDLUPTNLInformationForHOList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDAdditionalDLUPTNLInformationForHOList"`
	ULForwardingUPTNLInformation           *UPTransportLayerInformation           `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDULForwardingUPTNLInformation"`
	AdditionalULForwardingUPTNLInformation *UPTransportLayerInformationList       `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDAdditionalULForwardingUPTNLInformation"`
	DataForwardingResponseERABList         *DataForwardingResponseERABList        `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDDataForwardingResponseERABList"`
}

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability
	IEExtensions                     *ProtocolExtensionContainerHandoverRequiredTransferExtIEs
}

type ProtocolExtensionContainerHandoverRequiredTransferExtIEs struct {
	List []HandoverRequiredTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type HandoverRequiredTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      HandoverRequiredTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	HandoverRequiredTransferExtIEsExtensionValueChoiceNothing int = 0
)

type HandoverRequiredTransferExtIEsExtensionValue struct {
	Choice int
}

type HandoverResourceAllocationUnsuccessfulTransfer struct {
	Cause                  Cause
	CriticalityDiagnostics *CriticalityDiagnostics
	IEExtensions           *ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs
}

type ProtocolExtensionContainerHandoverResourceAllocationUnsuccessfulTransferExtIEs struct {
	List []HandoverResourceAllocationUnsuccessfulTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type HandoverResourceAllocationUnsuccessfulTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValueChoiceNothing int = 0
)

type HandoverResourceAllocationUnsuccessfulTransferExtIEsExtensionValue struct {
	Choice int
}

const (
	HandoverTypeIntra5gs      Enumerated = 0
	HandoverTypeFivegsToEps   Enumerated = 1
	HandoverTypeEpsTo5gs      Enumerated = 2
	HandoverTypeFivegsToUtran Enumerated = 3
)

type HandoverType struct {
	Value Enumerated `vht:"valueMin:0,valueMax:3"`
}

const (
	IMSVoiceSupportIndicatorSupported    Enumerated = 0
	IMSVoiceSupportIndicatorNotSupported Enumerated = 1
)

type IMSVoiceSupportIndicator struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type IndexToRFSP struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:256"`
}

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  RecommendedCellsForPaging
	RecommendRANNodesForPaging RecommendedRANNodesForPaging
	IEExtensions               *ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs
}

type ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs struct {
	List []InfoOnRecommendedCellsAndRANNodesForPagingExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type InfoOnRecommendedCellsAndRANNodesForPagingExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValueChoiceNothing int = 0
)

type InfoOnRecommendedCellsAndRANNodesForPagingExtIEsExtensionValue struct {
	Choice int
}

const (
	IntegrityProtectionIndicationRequired  Enumerated = 0
	IntegrityProtectionIndicationPreferred Enumerated = 1
	IntegrityProtectionIndicationNotNeeded Enumerated = 2
)

type IntegrityProtectionIndication struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	IntegrityProtectionResultPerformed    Enumerated = 0
	IntegrityProtectionResultNotPerformed Enumerated = 1
)

type IntegrityProtectionResult struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type IntendedNumberOfPagingAttempts struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:16"`
}

type InterfacesToTrace struct {
	Value BitString `vht:"sizeMin:8,sizeMax:8"`
}

type LAC struct {
	Value OctetString `vht:"sizeMin:2,sizeMax:2"`
}

type LAI struct {
	PLMNidentity PLMNIdentity
	LAC          LAC
	IEExtensions *ProtocolExtensionContainerLAIExtIEs
}

type ProtocolExtensionContainerLAIExtIEs struct {
	List []LAIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type LAIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      LAIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	LAIExtIEsExtensionValueChoiceNothing int = 0
)

type LAIExtIEsExtensionValue struct {
	Choice int
}

const (
	LastVisitedCellInformationChoiceNGRANCell        int = 0
	LastVisitedCellInformationChoiceEUTRANCell       int = 1
	LastVisitedCellInformationChoiceUTRANCell        int = 2
	LastVisitedCellInformationChoiceGERANCell        int = 3
	LastVisitedCellInformationChoiceChoiceExtensions int = 4
)

type LastVisitedCellInformation struct {
	Choice           int
	NGRANCell        *LastVisitedNGRANCellInformation
	EUTRANCell       *LastVisitedEUTRANCellInformation
	UTRANCell        *LastVisitedUTRANCellInformation
	GERANCell        *LastVisitedGERANCellInformation
	ChoiceExtensions *ProtocolIESingleContainerLastVisitedCellInformationExtIEs
}

type ProtocolIESingleContainerLastVisitedCellInformationExtIEs struct {
	List []LastVisitedCellInformationExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type LastVisitedCellInformationExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                               `vht:"Reference:ProtocolIEID"`
	TypeValue    LastVisitedCellInformationExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	LastVisitedCellInformationExtIEsTypeValueChoiceNothing int = 0
)

type LastVisitedCellInformationExtIEsTypeValue struct {
	Choice int
}

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation
	IEExtensions               *ProtocolExtensionContainerLastVisitedCellItemExtIEs
}

type ProtocolExtensionContainerLastVisitedCellItemExtIEs struct {
	List []LastVisitedCellItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type LastVisitedCellItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      LastVisitedCellItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	LastVisitedCellItemExtIEsExtensionValueChoiceNothing int = 0
)

type LastVisitedCellItemExtIEsExtensionValue struct {
	Choice int
}

type LastVisitedEUTRANCellInformation struct {
	Value OctetString
}

type LastVisitedGERANCellInformation struct {
	Value OctetString
}

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          NGRANCGI
	CellType                              CellType
	TimeUEStayedInCell                    TimeUEStayedInCell
	TimeUEStayedInCellEnhancedGranularity *TimeUEStayedInCellEnhancedGranularity
	HOCauseValue                          *Cause
	IEExtensions                          *ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs
}

type ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs struct {
	List []LastVisitedNGRANCellInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type LastVisitedNGRANCellInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      LastVisitedNGRANCellInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	LastVisitedNGRANCellInformationExtIEsExtensionValueChoiceNothing int = 0
)

type LastVisitedNGRANCellInformationExtIEsExtensionValue struct {
	Choice int
}

type LastVisitedUTRANCellInformation struct {
	Value OctetString
}

const (
	LocationReportingAdditionalInfoIncludePSCell Enumerated = 0
)

type LocationReportingAdditionalInfo struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type LocationReportingReferenceID struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:64"`
}

type LocationReportingRequestType struct {
	EventType                                 EventType
	ReportArea                                ReportArea
	AreaOfInterestList                        *AreaOfInterestList
	LocationReportingReferenceIDToBeCancelled *LocationReportingReferenceID
	IEExtensions                              *ProtocolExtensionContainerLocationReportingRequestTypeExtIEs
}

type ProtocolExtensionContainerLocationReportingRequestTypeExtIEs struct {
	List []LocationReportingRequestTypeExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type LocationReportingRequestTypeExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      LocationReportingRequestTypeExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	LocationReportingRequestTypeExtIEsExtensionValueChoiceLocationReportingAdditionalInfo int = 0
)

type LocationReportingRequestTypeExtIEsExtensionValue struct {
	Choice                          int
	LocationReportingAdditionalInfo *LocationReportingAdditionalInfo `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDLocationReportingAdditionalInfo"`
}

type MaskedIMEISV struct {
	Value BitString `vht:"sizeMin:64,sizeMax:64"`
}

type MaximumDataBurstVolume struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:2000000"`
}

type MessageIdentifier struct {
	Value BitString `vht:"sizeMin:16,sizeMax:16"`
}

const (
	MaximumIntegrityProtectedDataRateBitrate64kbs  Enumerated = 0
	MaximumIntegrityProtectedDataRateMaximumUERate Enumerated = 1
)

type MaximumIntegrityProtectedDataRate struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

const (
	MICOModeIndicationTrue Enumerated = 0
)

type MICOModeIndication struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type MobilityRestrictionList struct {
	ServingPLMN              PLMNIdentity
	EquivalentPLMNs          *EquivalentPLMNs
	RATRestrictions          *RATRestrictions
	ForbiddenAreaInformation *ForbiddenAreaInformation
	ServiceAreaInformation   *ServiceAreaInformation
	IEExtensions             *ProtocolExtensionContainerMobilityRestrictionListExtIEs
}

type ProtocolExtensionContainerMobilityRestrictionListExtIEs struct {
	List []MobilityRestrictionListExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type MobilityRestrictionListExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      MobilityRestrictionListExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	MobilityRestrictionListExtIEsExtensionValueChoiceLastEUTRANPLMNIdentity          int = 0
	MobilityRestrictionListExtIEsExtensionValueChoiceCNTypeRestrictionsForServing    int = 1
	MobilityRestrictionListExtIEsExtensionValueChoiceCNTypeRestrictionsForEquivalent int = 2
)

type MobilityRestrictionListExtIEsExtensionValue struct {
	Choice                          int
	LastEUTRANPLMNIdentity          *PLMNIdentity                    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDLastEUTRANPLMNIdentity"`
	CNTypeRestrictionsForServing    *CNTypeRestrictionsForServing    `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDCNTypeRestrictionsForServing"`
	CNTypeRestrictionsForEquivalent *CNTypeRestrictionsForEquivalent `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDCNTypeRestrictionsForEquivalent"`
}

const (
	N3IWFIDChoiceN3IWFID          int = 0
	N3IWFIDChoiceChoiceExtensions int = 1
)

type N3IWFID struct {
	Choice           int
	N3IWFID          *BitString `vht:"valueMin:16,valueMax:16"`
	ChoiceExtensions *ProtocolIESingleContainerN3IWFIDExtIEs
}

type ProtocolIESingleContainerN3IWFIDExtIEs struct {
	List []N3IWFIDExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type N3IWFIDExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	TypeValue    N3IWFIDExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	N3IWFIDExtIEsTypeValueChoiceNothing int = 0
)

type N3IWFIDExtIEsTypeValue struct {
	Choice int
}

type NASPDU struct {
	Value OctetString
}

type NASSecurityParametersFromNGRAN struct {
	Value OctetString
}

type NetworkInstance struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:256"`
}

const (
	NewSecurityContextIndTrue Enumerated = 0
)

type NewSecurityContextInd struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type NextHopChainingCount struct {
	Value Integer `vht:"valueMin:0,valueMax:7"`
}

const (
	NextPagingAreaScopeSame    Enumerated = 0
	NextPagingAreaScopeChanged Enumerated = 1
)

type NextPagingAreaScope struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

const (
	NgENBIDChoiceMacroNgENBID      int = 0
	NgENBIDChoiceShortMacroNgENBID int = 1
	NgENBIDChoiceLongMacroNgENBID  int = 2
	NgENBIDChoiceChoiceExtensions  int = 3
)

type NgENBID struct {
	Choice            int
	MacroNgENBID      *BitString `vht:"valueMin:20,valueMax:20"`
	ShortMacroNgENBID *BitString `vht:"valueMin:18,valueMax:18"`
	LongMacroNgENBID  *BitString `vht:"valueMin:21,valueMax:21"`
	ChoiceExtensions  *ProtocolIESingleContainerNgENBIDExtIEs
}

type ProtocolIESingleContainerNgENBIDExtIEs struct {
	List []NgENBIDExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type NgENBIDExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality            `vht:"Reference:ProtocolIEID"`
	TypeValue    NgENBIDExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NgENBIDExtIEsTypeValueChoiceNothing int = 0
)

type NgENBIDExtIEsTypeValue struct {
	Choice int
}

const (
	NGRANCGIChoiceNRCGI            int = 0
	NGRANCGIChoiceEUTRACGI         int = 1
	NGRANCGIChoiceChoiceExtensions int = 2
)

type NGRANCGI struct {
	Choice           int
	NRCGI            *NRCGI
	EUTRACGI         *EUTRACGI
	ChoiceExtensions *ProtocolIESingleContainerNGRANCGIExtIEs
}

type ProtocolIESingleContainerNGRANCGIExtIEs struct {
	List []NGRANCGIExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type NGRANCGIExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	TypeValue    NGRANCGIExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	NGRANCGIExtIEsTypeValueChoiceNothing int = 0
)

type NGRANCGIExtIEsTypeValue struct {
	Choice int
}

type NGRANTNLAssociationToRemoveList struct {
	List []NGRANTNLAssociationToRemoveItem `vht:"sizeMin:1,sizeMax:MaxnoofTNLAssociations"`
}

type NGRANTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress    CPTransportLayerInformation
	TNLAssociationTransportLayerAddressAMF *CPTransportLayerInformation
	IEExtensions                           *ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs
}

type ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs struct {
	List []NGRANTNLAssociationToRemoveItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type NGRANTNLAssociationToRemoveItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      NGRANTNLAssociationToRemoveItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	NGRANTNLAssociationToRemoveItemExtIEsExtensionValueChoiceNothing int = 0
)

type NGRANTNLAssociationToRemoveItemExtIEsExtensionValue struct {
	Choice int
}

type NGRANTraceID struct {
	Value OctetString `vht:"sizeMin:8,sizeMax:8"`
}

type NonDynamic5QIDescriptor struct {
	FiveQI                 FiveQI
	PriorityLevelQos       *PriorityLevelQos
	AveragingWindow        *AveragingWindow
	MaximumDataBurstVolume *MaximumDataBurstVolume
	IEExtensions           *ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs
}

type ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs struct {
	List []NonDynamic5QIDescriptorExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type NonDynamic5QIDescriptorExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      NonDynamic5QIDescriptorExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	NonDynamic5QIDescriptorExtIEsExtensionValueChoiceNothing int = 0
)

type NonDynamic5QIDescriptorExtIEsExtensionValue struct {
	Choice int
}

type NotAllowedTACs struct {
	List []TAC `vht:"sizeMin:1,sizeMax:MaxnoofAllowedAreas"`
}

const (
	NotificationCauseFulfilled    Enumerated = 0
	NotificationCauseNotFulfilled Enumerated = 1
)

type NotificationCause struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

const (
	NotificationControlNotificationRequested Enumerated = 0
)

type NotificationControl struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type NRCellIdentity struct {
	Value BitString `vht:"sizeMin:36,sizeMax:36"`
}

type NRCGI struct {
	PLMNIdentity   PLMNIdentity
	NRCellIdentity NRCellIdentity
	IEExtensions   *ProtocolExtensionContainerNRCGIExtIEs
}

type ProtocolExtensionContainerNRCGIExtIEs struct {
	List []NRCGIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type NRCGIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      NRCGIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	NRCGIExtIEsExtensionValueChoiceNothing int = 0
)

type NRCGIExtIEsExtensionValue struct {
	Choice int
}

type NRCGIList struct {
	List []NRCGI `vht:"sizeMin:1,sizeMax:MaxnoofCellsingNB"`
}

type NRCGIListForWarning struct {
	List []NRCGI `vht:"sizeMin:1,sizeMax:MaxnoofCellIDforWarning"`
}

type NRencryptionAlgorithms struct {
	Value BitString `vht:"sizeExx,sizeMin:16,sizeMax:16"`
}

type NRintegrityProtectionAlgorithms struct {
	Value BitString `vht:"sizeExx,sizeMin:16,sizeMax:16"`
}

type NRPPaPDU struct {
	Value OctetString
}

type NumberOfBroadcasts struct {
	Value Integer `vht:"valueMin:0,valueMax:65535"`
}

type NumberOfBroadcastsRequested struct {
	Value Integer `vht:"valueMin:0,valueMax:65535"`
}

const (
	OverloadActionRejectNonEmergencyMoDt                                    Enumerated = 0
	OverloadActionRejectRrcCrSignalling                                     Enumerated = 1
	OverloadActionPermitEmergencySessionsAndMobileTerminatedServicesOnly    Enumerated = 2
	OverloadActionPermitHighPrioritySessionsAndMobileTerminatedServicesOnly Enumerated = 3
)

type OverloadAction struct {
	Value Enumerated `vht:"valueMin:0,valueMax:3"`
}

const (
	OverloadResponseChoiceOverloadAction   int = 0
	OverloadResponseChoiceChoiceExtensions int = 1
)

type OverloadResponse struct {
	Choice           int
	OverloadAction   *OverloadAction
	ChoiceExtensions *ProtocolIESingleContainerOverloadResponseExtIEs
}

type ProtocolIESingleContainerOverloadResponseExtIEs struct {
	List []OverloadResponseExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type OverloadResponseExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                     `vht:"Reference:ProtocolIEID"`
	TypeValue    OverloadResponseExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	OverloadResponseExtIEsTypeValueChoiceNothing int = 0
)

type OverloadResponseExtIEsTypeValue struct {
	Choice int
}

type OverloadStartNSSAIList struct {
	List []OverloadStartNSSAIItem `vht:"sizeMin:1,sizeMax:MaxnoofSliceItems"`
}

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   SliceOverloadList
	SliceOverloadResponse               *OverloadResponse
	SliceTrafficLoadReductionIndication *TrafficLoadReductionIndication
	IEExtensions                        *ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs
}

type ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs struct {
	List []OverloadStartNSSAIItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type OverloadStartNSSAIItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      OverloadStartNSSAIItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	OverloadStartNSSAIItemExtIEsExtensionValueChoiceNothing int = 0
)

type OverloadStartNSSAIItemExtIEsExtensionValue struct {
	Choice int
}

type PacketDelayBudget struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:1023"`
}

type PacketErrorRate struct {
	PERScalar    Integer `vht:"valueExt,valueMin:0,valueMax:9"`
	PERExponent  Integer `vht:"valueExt,valueMin:0,valueMax:9"`
	IEExtensions *ProtocolExtensionContainerPacketErrorRateExtIEs
}

type ProtocolExtensionContainerPacketErrorRateExtIEs struct {
	List []PacketErrorRateExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PacketErrorRateExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PacketErrorRateExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PacketErrorRateExtIEsExtensionValueChoiceNothing int = 0
)

type PacketErrorRateExtIEsExtensionValue struct {
	Choice int
}

type PacketLossRate struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:1000"`
}

type PagingAttemptInformation struct {
	PagingAttemptCount             PagingAttemptCount
	IntendedNumberOfPagingAttempts IntendedNumberOfPagingAttempts
	NextPagingAreaScope            *NextPagingAreaScope
	IEExtensions                   *ProtocolExtensionContainerPagingAttemptInformationExtIEs
}

type ProtocolExtensionContainerPagingAttemptInformationExtIEs struct {
	List []PagingAttemptInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PagingAttemptInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PagingAttemptInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PagingAttemptInformationExtIEsExtensionValueChoiceNothing int = 0
)

type PagingAttemptInformationExtIEsExtensionValue struct {
	Choice int
}

type PagingAttemptCount struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:16"`
}

const (
	PagingDRXV32  Enumerated = 0
	PagingDRXV64  Enumerated = 1
	PagingDRXV128 Enumerated = 2
	PagingDRXV256 Enumerated = 3
)

type PagingDRX struct {
	Value Enumerated `vht:"valueMin:0,valueMax:3"`
}

const (
	PagingOriginNon3gpp Enumerated = 0
)

type PagingOrigin struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	PagingPriorityPriolevel1 Enumerated = 0
	PagingPriorityPriolevel2 Enumerated = 1
	PagingPriorityPriolevel3 Enumerated = 2
	PagingPriorityPriolevel4 Enumerated = 3
	PagingPriorityPriolevel5 Enumerated = 4
	PagingPriorityPriolevel6 Enumerated = 5
	PagingPriorityPriolevel7 Enumerated = 6
	PagingPriorityPriolevel8 Enumerated = 7
)

type PagingPriority struct {
	Value Enumerated `vht:"valueMin:0,valueMax:7"`
}

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation
	SecurityIndication    *SecurityIndication
	IEExtensions          *ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs
}

type ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs struct {
	List []PathSwitchRequestAcknowledgeTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PathSwitchRequestAcknowledgeTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PathSwitchRequestAcknowledgeTransferExtIEsExtensionValueChoiceAdditionalNGUUPTNLInformation int = 0
)

type PathSwitchRequestAcknowledgeTransferExtIEsExtensionValue struct {
	Choice                        int
	AdditionalNGUUPTNLInformation *UPTransportLayerInformationPairList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDAdditionalNGUUPTNLInformation"`
}

type PathSwitchRequestSetupFailedTransfer struct {
	Cause        Cause
	IEExtensions *ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs
}

type ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs struct {
	List []PathSwitchRequestSetupFailedTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PathSwitchRequestSetupFailedTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PathSwitchRequestSetupFailedTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PathSwitchRequestSetupFailedTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PathSwitchRequestSetupFailedTransferExtIEsExtensionValue struct {
	Choice int
}

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        UPTransportLayerInformation
	DLNGUTNLInformationReused    *DLNGUTNLInformationReused
	UserPlaneSecurityInformation *UserPlaneSecurityInformation
	QosFlowAcceptedList          QosFlowAcceptedList
	IEExtensions                 *ProtocolExtensionContainerPathSwitchRequestTransferExtIEs
}

type ProtocolExtensionContainerPathSwitchRequestTransferExtIEs struct {
	List []PathSwitchRequestTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PathSwitchRequestTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PathSwitchRequestTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PathSwitchRequestTransferExtIEsExtensionValueChoiceAdditionalDLQosFlowPerTNLInformation int = 0
)

type PathSwitchRequestTransferExtIEsExtensionValue struct {
	Choice                               int
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDAdditionalDLQosFlowPerTNLInformation"`
}

type PathSwitchRequestUnsuccessfulTransfer struct {
	Cause        Cause
	IEExtensions *ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs
}

type ProtocolExtensionContainerPathSwitchRequestUnsuccessfulTransferExtIEs struct {
	List []PathSwitchRequestUnsuccessfulTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PathSwitchRequestUnsuccessfulTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PathSwitchRequestUnsuccessfulTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL BitRate
	PDUSessionAggregateMaximumBitRateUL BitRate
	IEExtensions                        *ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs
}

type ProtocolExtensionContainerPDUSessionAggregateMaximumBitRateExtIEs struct {
	List []PDUSessionAggregateMaximumBitRateExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionAggregateMaximumBitRateExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionAggregateMaximumBitRateExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionAggregateMaximumBitRateExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionAggregateMaximumBitRateExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionID struct {
	Value Integer `vht:"valueMin:0,valueMax:255"`
}

type PDUSessionResourceAdmittedList struct {
	List []PDUSessionResourceAdmittedItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceAdmittedItem struct {
	PDUSessionID                       PDUSessionID
	HandoverRequestAcknowledgeTransfer OctetString `vht:"Reference:HandoverRequestAcknowledgeTransfer"`
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs struct {
	List []PDUSessionResourceAdmittedItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceAdmittedItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceAdmittedItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceAdmittedItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceAdmittedItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToModifyListModCfm struct {
	List []PDUSessionResourceFailedToModifyItemModCfm `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToModifyItemModCfm struct {
	PDUSessionID                                           PDUSessionID
	PDUSessionResourceModifyIndicationUnsuccessfulTransfer OctetString `vht:"Reference:PDUSessionResourceModifyIndicationUnsuccessfulTransfer"`
	IEExtensions                                           *ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs struct {
	List []PDUSessionResourceFailedToModifyItemModCfmExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToModifyItemModCfmExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToModifyItemModCfmExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToModifyListModRes struct {
	List []PDUSessionResourceFailedToModifyItemModRes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToModifyItemModRes struct {
	PDUSessionID                                 PDUSessionID
	PDUSessionResourceModifyUnsuccessfulTransfer OctetString `vht:"Reference:PDUSessionResourceModifyUnsuccessfulTransfer"`
	IEExtensions                                 *ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModResExtIEs struct {
	List []PDUSessionResourceFailedToModifyItemModResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToModifyItemModResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToModifyItemModResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToSetupListCxtFail struct {
	List []PDUSessionResourceFailedToSetupItemCxtFail `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToSetupItemCxtFail struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer OctetString `vht:"Reference:PDUSessionResourceSetupUnsuccessfulTransfer"`
	IEExtensions                                *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtFailExtIEs struct {
	List []PDUSessionResourceFailedToSetupItemCxtFailExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToSetupItemCxtFailExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToSetupItemCxtFailExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToSetupListCxtRes struct {
	List []PDUSessionResourceFailedToSetupItemCxtRes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToSetupItemCxtRes struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer OctetString `vht:"Reference:PDUSessionResourceSetupUnsuccessfulTransfer"`
	IEExtensions                                *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs struct {
	List []PDUSessionResourceFailedToSetupItemCxtResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToSetupItemCxtResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToSetupItemCxtResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToSetupListHOAck struct {
	List []PDUSessionResourceFailedToSetupItemHOAck `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToSetupItemHOAck struct {
	PDUSessionID                                   PDUSessionID
	HandoverResourceAllocationUnsuccessfulTransfer OctetString `vht:"Reference:HandoverResourceAllocationUnsuccessfulTransfer"`
	IEExtensions                                   *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemHOAckExtIEs struct {
	List []PDUSessionResourceFailedToSetupItemHOAckExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToSetupItemHOAckExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToSetupItemHOAckExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToSetupListPSReq struct {
	List []PDUSessionResourceFailedToSetupItemPSReq `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToSetupItemPSReq struct {
	PDUSessionID                         PDUSessionID
	PathSwitchRequestSetupFailedTransfer OctetString `vht:"Reference:PathSwitchRequestSetupFailedTransfer"`
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemPSReqExtIEs struct {
	List []PDUSessionResourceFailedToSetupItemPSReqExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToSetupItemPSReqExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToSetupItemPSReqExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceFailedToSetupListSURes struct {
	List []PDUSessionResourceFailedToSetupItemSURes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceFailedToSetupItemSURes struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer OctetString `vht:"Reference:PDUSessionResourceSetupUnsuccessfulTransfer"`
	IEExtensions                                *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemSUResExtIEs struct {
	List []PDUSessionResourceFailedToSetupItemSUResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceFailedToSetupItemSUResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceFailedToSetupItemSUResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceHandoverList struct {
	List []PDUSessionResourceHandoverItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceHandoverItem struct {
	PDUSessionID            PDUSessionID
	HandoverCommandTransfer OctetString `vht:"Reference:HandoverCommandTransfer"`
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs struct {
	List []PDUSessionResourceHandoverItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceHandoverItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceHandoverItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceHandoverItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceHandoverItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceInformationList struct {
	List []PDUSessionResourceInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceInformationItem struct {
	PDUSessionID              PDUSessionID
	QosFlowInformationList    QosFlowInformationList
	DRBsToQosFlowsMappingList *DRBsToQosFlowsMappingList
	IEExtensions              *ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceInformationItemExtIEs struct {
	List []PDUSessionResourceInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceInformationItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceInformationItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceListCxtRelCpl struct {
	List []PDUSessionResourceItemCxtRelCpl `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceItemCxtRelCpl struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs struct {
	List []PDUSessionResourceItemCxtRelCplExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceItemCxtRelCplExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceItemCxtRelCplExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceItemCxtRelCplExtIEsExtensionValueChoicePDUSessionResourceReleaseResponseTransfer int = 0
)

type PDUSessionResourceItemCxtRelCplExtIEsExtensionValue struct {
	Choice                                    int
	PDUSessionResourceReleaseResponseTransfer *OctetString `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDPDUSessionResourceReleaseResponseTransfer,Reference:PDUSessionResourceReleaseResponseTransfer"`
}

type PDUSessionResourceListCxtRelReq struct {
	List []PDUSessionResourceItemCxtRelReq `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs struct {
	List []PDUSessionResourceItemCxtRelReqExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceItemCxtRelReqExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceItemCxtRelReqExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceItemCxtRelReqExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceItemCxtRelReqExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceListHORqd struct {
	List []PDUSessionResourceItemHORqd `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceItemHORqd struct {
	PDUSessionID             PDUSessionID
	HandoverRequiredTransfer OctetString `vht:"Reference:HandoverRequiredTransfer"`
	IEExtensions             *ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs struct {
	List []PDUSessionResourceItemHORqdExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceItemHORqdExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceItemHORqdExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceItemHORqdExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceItemHORqdExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList      QosFlowModifyConfirmList
	ULNGUUPTNLInformation         UPTransportLayerInformation
	AdditionalNGUUPTNLInformation *UPTransportLayerInformationPairList
	QosFlowFailedToModifyList     *QosFlowListWithCause
	IEExtensions                  *ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs struct {
	List []PDUSessionResourceModifyConfirmTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyConfirmTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyConfirmTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyConfirmTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceModifyConfirmTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceModifyIndicationUnsuccessfulTransfer struct {
	Cause        Cause
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs struct {
	List []PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceModifyRequestTransfer struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs
}

type ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs struct {
	List []PDUSessionResourceModifyRequestTransferIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type PDUSessionResourceModifyRequestTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                         `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceModifyRequestTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoicePDUSessionAggregateMaximumBitRate int = 0
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoiceULNGUUPTNLModifyList              int = 1
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoiceNetworkInstance                   int = 2
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoiceQosFlowAddOrModifyRequestList     int = 3
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoiceQosFlowToReleaseList              int = 4
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoiceAdditionalULNGUUPTNLInformation   int = 5
	PDUSessionResourceModifyRequestTransferIEsTypeValueChoiceCommonNetworkInstance             int = 6
)

type PDUSessionResourceModifyRequestTransferIEsTypeValue struct {
	Choice                            int
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionAggregateMaximumBitRate"`
	ULNGUUPTNLModifyList              *ULNGUUPTNLModifyList              `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDULNGUUPTNLModifyList"`
	NetworkInstance                   *NetworkInstance                   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNetworkInstance"`
	QosFlowAddOrModifyRequestList     *QosFlowAddOrModifyRequestList     `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDQosFlowAddOrModifyRequestList"`
	QosFlowToReleaseList              *QosFlowListWithCause              `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDQosFlowToReleaseList"`
	AdditionalULNGUUPTNLInformation   *UPTransportLayerInformationList   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAdditionalULNGUUPTNLInformation"`
	CommonNetworkInstance             *CommonNetworkInstance             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCommonNetworkInstance"`
}

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation                *UPTransportLayerInformation
	ULNGUUPTNLInformation                *UPTransportLayerInformation
	QosFlowAddOrModifyResponseList       *QosFlowAddOrModifyResponseList
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList
	QosFlowFailedToAddOrModifyList       *QosFlowListWithCause
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs struct {
	List []PDUSessionResourceModifyResponseTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyResponseTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyResponseTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyResponseTransferExtIEsExtensionValueChoiceAdditionalNGUUPTNLInformation int = 0
)

type PDUSessionResourceModifyResponseTransferExtIEsExtensionValue struct {
	Choice                        int
	AdditionalNGUUPTNLInformation *UPTransportLayerInformationPairList `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDAdditionalNGUUPTNLInformation"`
}

type PDUSessionResourceModifyIndicationTransfer struct {
	DLQosFlowPerTNLInformation           QosFlowPerTNLInformation
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs struct {
	List []PDUSessionResourceModifyIndicationTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyIndicationTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyIndicationTransferExtIEsExtensionValueChoiceSecondaryRATUsageInformation int = 0
	PDUSessionResourceModifyIndicationTransferExtIEsExtensionValueChoiceSecurityResult               int = 1
)

type PDUSessionResourceModifyIndicationTransferExtIEsExtensionValue struct {
	Choice                       int
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSecondaryRATUsageInformation"`
	SecurityResult               *SecurityResult               `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSecurityResult"`
}

type PDUSessionResourceModifyListModCfm struct {
	List []PDUSessionResourceModifyItemModCfm `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceModifyItemModCfm struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceModifyConfirmTransfer OctetString `vht:"Reference:PDUSessionResourceModifyConfirmTransfer"`
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs struct {
	List []PDUSessionResourceModifyItemModCfmExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyItemModCfmExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyItemModCfmExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyItemModCfmExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceModifyItemModCfmExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceModifyListModInd struct {
	List []PDUSessionResourceModifyItemModInd `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceModifyItemModInd struct {
	PDUSessionID                               PDUSessionID
	PDUSessionResourceModifyIndicationTransfer OctetString `vht:"Reference:PDUSessionResourceModifyIndicationTransfer"`
	IEExtensions                               *ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs struct {
	List []PDUSessionResourceModifyItemModIndExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyItemModIndExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyItemModIndExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyItemModIndExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceModifyItemModIndExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceModifyListModReq struct {
	List []PDUSessionResourceModifyItemModReq `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            PDUSessionID
	NASPDU                                  *NASPDU
	PDUSessionResourceModifyRequestTransfer OctetString `vht:"Reference:PDUSessionResourceModifyRequestTransfer"`
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs struct {
	List []PDUSessionResourceModifyItemModReqExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyItemModReqExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyItemModReqExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyItemModReqExtIEsExtensionValueChoiceSNSSAI int = 0
)

type PDUSessionResourceModifyItemModReqExtIEsExtensionValue struct {
	Choice int
	SNSSAI *SNSSAI `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDSNSSAI"`
}

type PDUSessionResourceModifyListModRes struct {
	List []PDUSessionResourceModifyItemModRes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceModifyResponseTransfer OctetString `vht:"Reference:PDUSessionResourceModifyResponseTransfer"`
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs struct {
	List []PDUSessionResourceModifyItemModResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyItemModResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyItemModResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyItemModResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceModifyItemModResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceModifyUnsuccessfulTransfer struct {
	Cause                  Cause
	CriticalityDiagnostics *CriticalityDiagnostics
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceModifyUnsuccessfulTransferExtIEs struct {
	List []PDUSessionResourceModifyUnsuccessfulTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceModifyUnsuccessfulTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceModifyUnsuccessfulTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceNotifyList struct {
	List []PDUSessionResourceNotifyItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceNotifyItem struct {
	PDUSessionID                     PDUSessionID
	PDUSessionResourceNotifyTransfer OctetString `vht:"Reference:PDUSessionResourceNotifyTransfer"`
	IEExtensions                     *ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceNotifyItemExtIEs struct {
	List []PDUSessionResourceNotifyItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceNotifyItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceNotifyItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceNotifyItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceNotifyItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceNotifyReleasedTransfer struct {
	Cause        Cause
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs struct {
	List []PDUSessionResourceNotifyReleasedTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceNotifyReleasedTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValueChoiceSecondaryRATUsageInformation int = 0
)

type PDUSessionResourceNotifyReleasedTransferExtIEsExtensionValue struct {
	Choice                       int
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSecondaryRATUsageInformation"`
}

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList
	QosFlowReleasedList *QosFlowListWithCause
	IEExtensions        *ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs struct {
	List []PDUSessionResourceNotifyTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceNotifyTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                          `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceNotifyTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceNotifyTransferExtIEsExtensionValueChoiceSecondaryRATUsageInformation int = 0
)

type PDUSessionResourceNotifyTransferExtIEsExtensionValue struct {
	Choice                       int
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSecondaryRATUsageInformation"`
}

type PDUSessionResourceReleaseCommandTransfer struct {
	Cause        Cause
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs struct {
	List []PDUSessionResourceReleaseCommandTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceReleaseCommandTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceReleaseCommandTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceReleaseCommandTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceReleaseCommandTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceReleasedListNot struct {
	List []PDUSessionResourceReleasedItemNot `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceReleasedItemNot struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceNotifyReleasedTransfer OctetString `vht:"Reference:PDUSessionResourceNotifyReleasedTransfer"`
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceReleasedItemNotExtIEs struct {
	List []PDUSessionResourceReleasedItemNotExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceReleasedItemNotExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceReleasedItemNotExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceReleasedItemNotExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceReleasedItemNotExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceReleasedListPSAck struct {
	List []PDUSessionResourceReleasedItemPSAck `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceReleasedItemPSAck struct {
	PDUSessionID                          PDUSessionID
	PathSwitchRequestUnsuccessfulTransfer OctetString `vht:"Reference:PathSwitchRequestUnsuccessfulTransfer"`
	IEExtensions                          *ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceReleasedItemPSAckExtIEs struct {
	List []PDUSessionResourceReleasedItemPSAckExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceReleasedItemPSAckExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceReleasedItemPSAckExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceReleasedItemPSAckExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceReleasedListPSFail struct {
	List []PDUSessionResourceReleasedItemPSFail `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceReleasedItemPSFail struct {
	PDUSessionID                          PDUSessionID
	PathSwitchRequestUnsuccessfulTransfer OctetString `vht:"Reference:PathSwitchRequestUnsuccessfulTransfer"`
	IEExtensions                          *ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceReleasedItemPSFailExtIEs struct {
	List []PDUSessionResourceReleasedItemPSFailExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceReleasedItemPSFailExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceReleasedItemPSFailExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceReleasedItemPSFailExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceReleasedListRelRes struct {
	List []PDUSessionResourceReleasedItemRelRes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceReleasedItemRelRes struct {
	PDUSessionID                              PDUSessionID
	PDUSessionResourceReleaseResponseTransfer OctetString `vht:"Reference:PDUSessionResourceReleaseResponseTransfer"`
	IEExtensions                              *ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs struct {
	List []PDUSessionResourceReleasedItemRelResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceReleasedItemRelResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceReleasedItemRelResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceReleasedItemRelResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceReleasedItemRelResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceReleaseResponseTransfer struct {
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs struct {
	List []PDUSessionResourceReleaseResponseTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceReleaseResponseTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceReleaseResponseTransferExtIEsExtensionValueChoiceSecondaryRATUsageInformation int = 0
)

type PDUSessionResourceReleaseResponseTransferExtIEsExtensionValue struct {
	Choice                       int
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSecondaryRATUsageInformation"`
}

type PDUSessionResourceSecondaryRATUsageList struct {
	List []PDUSessionResourceSecondaryRATUsageItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSecondaryRATUsageItem struct {
	PDUSessionID                        PDUSessionID
	SecondaryRATDataUsageReportTransfer OctetString `vht:"Reference:SecondaryRATDataUsageReportTransfer"`
	IEExtensions                        *ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSecondaryRATUsageItemExtIEs struct {
	List []PDUSessionResourceSecondaryRATUsageItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSecondaryRATUsageItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSecondaryRATUsageItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupListCxtReq struct {
	List []PDUSessionResourceSetupItemCxtReq `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSetupItemCxtReq struct {
	PDUSessionID                           PDUSessionID
	NASPDU                                 *NASPDU
	SNSSAI                                 SNSSAI
	PDUSessionResourceSetupRequestTransfer OctetString `vht:"Reference:PDUSessionResourceSetupRequestTransfer"`
	IEExtensions                           *ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupItemCxtReqExtIEs struct {
	List []PDUSessionResourceSetupItemCxtReqExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupItemCxtReqExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupItemCxtReqExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupItemCxtReqExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupListCxtRes struct {
	List []PDUSessionResourceSetupItemCxtRes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSetupItemCxtRes struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceSetupResponseTransfer OctetString `vht:"Reference:PDUSessionResourceSetupResponseTransfer"`
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs struct {
	List []PDUSessionResourceSetupItemCxtResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupItemCxtResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupItemCxtResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupItemCxtResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupItemCxtResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupListHOReq struct {
	List []PDUSessionResourceSetupItemHOReq `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            PDUSessionID
	SNSSAI                  SNSSAI
	HandoverRequestTransfer OctetString `vht:"Reference:PDUSessionResourceSetupRequestTransfer"`
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs struct {
	List []PDUSessionResourceSetupItemHOReqExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupItemHOReqExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                          `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupItemHOReqExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupItemHOReqExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupItemHOReqExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupListSUReq struct {
	List []PDUSessionResourceSetupItemSUReq `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           PDUSessionID
	PDUSessionNASPDU                       *NASPDU
	SNSSAI                                 SNSSAI
	PDUSessionResourceSetupRequestTransfer OctetString `vht:"Reference:PDUSessionResourceSetupRequestTransfer"`
	IEExtensions                           *ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs struct {
	List []PDUSessionResourceSetupItemSUReqExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupItemSUReqExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                          `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupItemSUReqExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupItemSUReqExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupItemSUReqExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupListSURes struct {
	List []PDUSessionResourceSetupItemSURes `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSetupItemSURes struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceSetupResponseTransfer OctetString `vht:"Reference:PDUSessionResourceSetupResponseTransfer"`
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs struct {
	List []PDUSessionResourceSetupItemSUResExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupItemSUResExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                          `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupItemSUResExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupItemSUResExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupItemSUResExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupRequestTransfer struct {
	ProtocolIEs ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs
}

type ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs struct {
	List []PDUSessionResourceSetupRequestTransferIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type PDUSessionResourceSetupRequestTransferIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                        `vht:"Reference:ProtocolIEID"`
	TypeValue    PDUSessionResourceSetupRequestTransferIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoicePDUSessionAggregateMaximumBitRate int = 0
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceULNGUUPTNLInformation             int = 1
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceAdditionalULNGUUPTNLInformation   int = 2
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceDataForwardingNotPossible         int = 3
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoicePDUSessionType                    int = 4
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceSecurityIndication                int = 5
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceNetworkInstance                   int = 6
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceQosFlowSetupRequestList           int = 7
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceCommonNetworkInstance             int = 8
	PDUSessionResourceSetupRequestTransferIEsTypeValueChoiceDirectForwardingPathAvailability  int = 9
)

type PDUSessionResourceSetupRequestTransferIEsTypeValue struct {
	Choice                            int
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionAggregateMaximumBitRate"`
	ULNGUUPTNLInformation             *UPTransportLayerInformation       `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDULNGUUPTNLInformation"`
	AdditionalULNGUUPTNLInformation   *UPTransportLayerInformationList   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDAdditionalULNGUUPTNLInformation"`
	DataForwardingNotPossible         *DataForwardingNotPossible         `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDDataForwardingNotPossible"`
	PDUSessionType                    *PDUSessionType                    `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDPDUSessionType"`
	SecurityIndication                *SecurityIndication                `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDSecurityIndication"`
	NetworkInstance                   *NetworkInstance                   `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDNetworkInstance"`
	QosFlowSetupRequestList           *QosFlowSetupRequestList           `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDQosFlowSetupRequestList"`
	CommonNetworkInstance             *CommonNetworkInstance             `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDCommonNetworkInstance"`
	DirectForwardingPathAvailability  *DirectForwardingPathAvailability  `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolIEID:ProtocolIEIDDirectForwardingPathAvailability"`
}

type PDUSessionResourceSetupResponseTransfer struct {
	DLQosFlowPerTNLInformation           QosFlowPerTNLInformation
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList
	SecurityResult                       *SecurityResult
	QosFlowFailedToSetupList             *QosFlowListWithCause
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs struct {
	List []PDUSessionResourceSetupResponseTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupResponseTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupResponseTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupResponseTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupResponseTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  Cause
	CriticalityDiagnostics *CriticalityDiagnostics
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs struct {
	List []PDUSessionResourceSetupUnsuccessfulTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSetupUnsuccessfulTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSetupUnsuccessfulTransferExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceSwitchedList struct {
	List []PDUSessionResourceSwitchedItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceSwitchedItem struct {
	PDUSessionID                         PDUSessionID
	PathSwitchRequestAcknowledgeTransfer OctetString `vht:"Reference:PathSwitchRequestAcknowledgeTransfer"`
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceSwitchedItemExtIEs struct {
	List []PDUSessionResourceSwitchedItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceSwitchedItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceSwitchedItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceSwitchedItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceSwitchedItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceToBeSwitchedDLList struct {
	List []PDUSessionResourceToBeSwitchedDLItem `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceToBeSwitchedDLItem struct {
	PDUSessionID              PDUSessionID
	PathSwitchRequestTransfer OctetString `vht:"Reference:PathSwitchRequestTransfer"`
	IEExtensions              *ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceToBeSwitchedDLItemExtIEs struct {
	List []PDUSessionResourceToBeSwitchedDLItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceToBeSwitchedDLItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceToBeSwitchedDLItemExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceToReleaseListHOCmd struct {
	List []PDUSessionResourceToReleaseItemHOCmd `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceToReleaseItemHOCmd struct {
	PDUSessionID                            PDUSessionID
	HandoverPreparationUnsuccessfulTransfer OctetString `vht:"Reference:HandoverPreparationUnsuccessfulTransfer"`
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceToReleaseItemHOCmdExtIEs struct {
	List []PDUSessionResourceToReleaseItemHOCmdExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceToReleaseItemHOCmdExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceToReleaseItemHOCmdExtIEsExtensionValue struct {
	Choice int
}

type PDUSessionResourceToReleaseListRelCmd struct {
	List []PDUSessionResourceToReleaseItemRelCmd `vht:"sizeMin:1,sizeMax:MaxnoofPDUSessions"`
}

type PDUSessionResourceToReleaseItemRelCmd struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceReleaseCommandTransfer OctetString `vht:"Reference:PDUSessionResourceReleaseCommandTransfer"`
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs
}

type ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs struct {
	List []PDUSessionResourceToReleaseItemRelCmdExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionResourceToReleaseItemRelCmdExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionResourceToReleaseItemRelCmdExtIEsExtensionValue struct {
	Choice int
}

const (
	PDUSessionTypeIpv4         Enumerated = 0
	PDUSessionTypeIpv6         Enumerated = 1
	PDUSessionTypeIpv4v6       Enumerated = 2
	PDUSessionTypeEthernet     Enumerated = 3
	PDUSessionTypeUnstructured Enumerated = 4
)

type PDUSessionType struct {
	Value Enumerated `vht:"valueMin:0,valueMax:4"`
}

const (
	RATTypeNr              Enumerated = 0
	RATTypeEutra           Enumerated = 1
	RATTypeNrUnlicensed    Enumerated = 2
	RATTypeEUtraUnlicensed Enumerated = 3
)

type PDUSessionUsageReport struct {
	RATType                   Enumerated `vht:"valueExt,valueMin:0,valueMax:4"`
	PDUSessionTimedReportList VolumeTimedReportList
	IEExtensions              *ProtocolExtensionContainerPDUSessionUsageReportExtIEs
}

type ProtocolExtensionContainerPDUSessionUsageReportExtIEs struct {
	List []PDUSessionUsageReportExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PDUSessionUsageReportExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PDUSessionUsageReportExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PDUSessionUsageReportExtIEsExtensionValueChoiceNothing int = 0
)

type PDUSessionUsageReportExtIEsExtensionValue struct {
	Choice int
}

type PeriodicRegistrationUpdateTimer struct {
	Value BitString `vht:"sizeMin:8,sizeMax:8"`
}

type PLMNIdentity struct {
	Value OctetString `vht:"sizeMin:3,sizeMax:3"`
}

type PLMNSupportList struct {
	List []PLMNSupportItem `vht:"sizeMin:1,sizeMax:MaxnoofPLMNs"`
}

type PLMNSupportItem struct {
	PLMNIdentity     PLMNIdentity
	SliceSupportList SliceSupportList
	IEExtensions     *ProtocolExtensionContainerPLMNSupportItemExtIEs
}

type ProtocolExtensionContainerPLMNSupportItemExtIEs struct {
	List []PLMNSupportItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type PLMNSupportItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      PLMNSupportItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	PLMNSupportItemExtIEsExtensionValueChoiceNothing int = 0
)

type PLMNSupportItemExtIEsExtensionValue struct {
	Choice int
}

type PortNumber struct {
	Value OctetString `vht:"sizeMin:2,sizeMax:2"`
}

const (
	PreEmptionCapabilityShallNotTriggerPreEmption Enumerated = 0
	PreEmptionCapabilityMayTriggerPreEmption      Enumerated = 1
)

type PreEmptionCapability struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

const (
	PreEmptionVulnerabilityNotPreEmptable Enumerated = 0
	PreEmptionVulnerabilityPreEmptable    Enumerated = 1
)

type PreEmptionVulnerability struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type PriorityLevelARP struct {
	Value Integer `vht:"valueMin:1,valueMax:15"`
}

type PriorityLevelQos struct {
	Value Integer `vht:"valueExt,valueMin:1,valueMax:127"`
}

const (
	PWSFailedCellIDListChoiceEUTRACGIPWSFailedList int = 0
	PWSFailedCellIDListChoiceNRCGIPWSFailedList    int = 1
	PWSFailedCellIDListChoiceChoiceExtensions      int = 2
)

type PWSFailedCellIDList struct {
	Choice                int
	EUTRACGIPWSFailedList *EUTRACGIList
	NRCGIPWSFailedList    *NRCGIList
	ChoiceExtensions      *ProtocolIESingleContainerPWSFailedCellIDListExtIEs
}

type ProtocolIESingleContainerPWSFailedCellIDListExtIEs struct {
	List []PWSFailedCellIDListExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type PWSFailedCellIDListExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                        `vht:"Reference:ProtocolIEID"`
	TypeValue    PWSFailedCellIDListExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	PWSFailedCellIDListExtIEsTypeValueChoiceNothing int = 0
)

type PWSFailedCellIDListExtIEsTypeValue struct {
	Choice int
}

const (
	QosCharacteristicsChoiceNonDynamic5QI    int = 0
	QosCharacteristicsChoiceDynamic5QI       int = 1
	QosCharacteristicsChoiceChoiceExtensions int = 2
)

type QosCharacteristics struct {
	Choice           int
	NonDynamic5QI    *NonDynamic5QIDescriptor
	Dynamic5QI       *Dynamic5QIDescriptor
	ChoiceExtensions *ProtocolIESingleContainerQosCharacteristicsExtIEs
}

type ProtocolIESingleContainerQosCharacteristicsExtIEs struct {
	List []QosCharacteristicsExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type QosCharacteristicsExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                       `vht:"Reference:ProtocolIEID"`
	TypeValue    QosCharacteristicsExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	QosCharacteristicsExtIEsTypeValueChoiceNothing int = 0
)

type QosCharacteristicsExtIEsTypeValue struct {
	Choice int
}

type QosFlowAcceptedList struct {
	List []QosFlowAcceptedItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowAcceptedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAcceptedItemExtIEs
}

type ProtocolExtensionContainerQosFlowAcceptedItemExtIEs struct {
	List []QosFlowAcceptedItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowAcceptedItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowAcceptedItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowAcceptedItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowAcceptedItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowAddOrModifyRequestList struct {
	List []QosFlowAddOrModifyRequestItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters *QosFlowLevelQosParameters
	ERABID                    *ERABID
	IEExtensions              *ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs
}

type ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs struct {
	List []QosFlowAddOrModifyRequestItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowAddOrModifyRequestItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowAddOrModifyRequestItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowAddOrModifyRequestItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowAddOrModifyRequestItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowAddOrModifyResponseList struct {
	List []QosFlowAddOrModifyResponseItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowAddOrModifyResponseItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs
}

type ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs struct {
	List []QosFlowAddOrModifyResponseItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowAddOrModifyResponseItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowAddOrModifyResponseItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowAddOrModifyResponseItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowAddOrModifyResponseItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowIdentifier struct {
	Value Integer `vht:"valueExt,valueMin:0,valueMax:63"`
}

type QosFlowInformationList struct {
	List []QosFlowInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowInformationItem struct {
	QosFlowIdentifier QosFlowIdentifier
	DLForwarding      *DLForwarding
	IEExtensions      *ProtocolExtensionContainerQosFlowInformationItemExtIEs
}

type ProtocolExtensionContainerQosFlowInformationItemExtIEs struct {
	List []QosFlowInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowInformationItemExtIEsExtensionValueChoiceULForwarding int = 0
)

type QosFlowInformationItemExtIEsExtensionValue struct {
	Choice       int
	ULForwarding *ULForwarding `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDULForwarding"`
}

type QosFlowLevelQosParameters struct {
	QosCharacteristics             QosCharacteristics
	AllocationAndRetentionPriority AllocationAndRetentionPriority
	GBRQosInformation              *GBRQosInformation
	ReflectiveQosAttribute         *ReflectiveQosAttribute
	AdditionalQosFlowInformation   *AdditionalQosFlowInformation
	IEExtensions                   *ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs
}

type ProtocolExtensionContainerQosFlowLevelQosParametersExtIEs struct {
	List []QosFlowLevelQosParametersExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowLevelQosParametersExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowLevelQosParametersExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowLevelQosParametersExtIEsExtensionValueChoiceQosMonitoringRequest int = 0
)

type QosFlowLevelQosParametersExtIEsExtensionValue struct {
	Choice               int
	QosMonitoringRequest *QosMonitoringRequest `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDQosMonitoringRequest"`
}

const (
	QosMonitoringRequestUl   Enumerated = 0
	QosMonitoringRequestDl   Enumerated = 1
	QosMonitoringRequestBoth Enumerated = 2
)

type QosMonitoringRequest struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

type QosFlowListWithCause struct {
	List []QosFlowWithCauseItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowWithCauseItem struct {
	QosFlowIdentifier QosFlowIdentifier
	Cause             Cause
	IEExtensions      *ProtocolExtensionContainerQosFlowWithCauseItemExtIEs
}

type ProtocolExtensionContainerQosFlowWithCauseItemExtIEs struct {
	List []QosFlowWithCauseItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowWithCauseItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowWithCauseItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowWithCauseItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowWithCauseItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowModifyConfirmList struct {
	List []QosFlowModifyConfirmItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs
}

type ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs struct {
	List []QosFlowModifyConfirmItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowModifyConfirmItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowModifyConfirmItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowModifyConfirmItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowModifyConfirmItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowNotifyList struct {
	List []QosFlowNotifyItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowNotifyItem struct {
	QosFlowIdentifier QosFlowIdentifier
	NotificationCause NotificationCause
	IEExtensions      *ProtocolExtensionContainerQosFlowNotifyItemExtIEs
}

type ProtocolExtensionContainerQosFlowNotifyItemExtIEs struct {
	List []QosFlowNotifyItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowNotifyItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowNotifyItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowNotifyItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowNotifyItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation
	AssociatedQosFlowList       AssociatedQosFlowList
	IEExtensions                *ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs
}

type ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs struct {
	List []QosFlowPerTNLInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowPerTNLInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowPerTNLInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowPerTNLInformationExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowPerTNLInformationExtIEsExtensionValue struct {
	Choice int
}

type QosFlowPerTNLInformationList struct {
	List []QosFlowPerTNLInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofMultiConnectivityMinusOne"`
}

type QosFlowPerTNLInformationItem struct {
	QosFlowPerTNLInformation QosFlowPerTNLInformation
	IEExtensions             *ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs
}

type ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs struct {
	List []QosFlowPerTNLInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowPerTNLInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowPerTNLInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowPerTNLInformationItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowPerTNLInformationItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowSetupRequestList struct {
	List []QosFlowSetupRequestItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier
	QosFlowLevelQosParameters QosFlowLevelQosParameters
	ERABID                    *ERABID
	IEExtensions              *ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs
}

type ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs struct {
	List []QosFlowSetupRequestItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowSetupRequestItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowSetupRequestItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowSetupRequestItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowSetupRequestItemExtIEsExtensionValue struct {
	Choice int
}

type QosFlowListWithDataForwarding struct {
	List []QosFlowItemWithDataForwarding `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowItemWithDataForwarding struct {
	QosFlowIdentifier      QosFlowIdentifier
	DataForwardingAccepted *DataForwardingAccepted
	IEExtensions           *ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs
}

type ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs struct {
	List []QosFlowItemWithDataForwardingExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowItemWithDataForwardingExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                       `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowItemWithDataForwardingExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowItemWithDataForwardingExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowItemWithDataForwardingExtIEsExtensionValue struct {
	Choice int
}

type QosFlowToBeForwardedList struct {
	List []QosFlowToBeForwardedItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs
}

type ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs struct {
	List []QosFlowToBeForwardedItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QosFlowToBeForwardedItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QosFlowToBeForwardedItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QosFlowToBeForwardedItemExtIEsExtensionValueChoiceNothing int = 0
)

type QosFlowToBeForwardedItemExtIEsExtensionValue struct {
	Choice int
}

type QoSFlowsUsageReportList struct {
	List []QoSFlowsUsageReportItem `vht:"sizeMin:1,sizeMax:MaxnoofQosFlows"`
}

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       QosFlowIdentifier
	RATType                 Enumerated `vht:"valueExt,valueMin:0,valueMax:4"`
	QoSFlowsTimedReportList VolumeTimedReportList
	IEExtensions            *ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs
}

type ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs struct {
	List []QoSFlowsUsageReportItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type QoSFlowsUsageReportItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                 `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      QoSFlowsUsageReportItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	QoSFlowsUsageReportItemExtIEsExtensionValueChoiceNothing int = 0
)

type QoSFlowsUsageReportItemExtIEsExtensionValue struct {
	Choice int
}

type RANNodeName struct {
	Value PrintableString `vht:"sizeExt,sizeMin:1,sizeMax:150"`
}

type RANPagingPriority struct {
	Value Integer `vht:"valueMin:1,valueMax:256"`
}

type RANStatusTransferTransparentContainer struct {
	DRBsSubjectToStatusTransferList DRBsSubjectToStatusTransferList
	IEExtensions                    *ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs
}

type ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs struct {
	List []RANStatusTransferTransparentContainerExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RANStatusTransferTransparentContainerExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RANStatusTransferTransparentContainerExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RANStatusTransferTransparentContainerExtIEsExtensionValueChoiceNothing int = 0
)

type RANStatusTransferTransparentContainerExtIEsExtensionValue struct {
	Choice int
}

type RANUENGAPID struct {
	Value Integer `vht:"valueMin:0,valueMax:4294967295"`
}

const (
	RATInformationUnlicensed Enumerated = 0
)

type RATInformation struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type RATRestrictions struct {
	List []RATRestrictionsItem `vht:"sizeMin:1,sizeMax:MaxnoofEPLMNsPlusOne"`
}

type RATRestrictionsItem struct {
	PLMNIdentity              PLMNIdentity
	RATRestrictionInformation RATRestrictionInformation
	IEExtensions              *ProtocolExtensionContainerRATRestrictionsItemExtIEs
}

type ProtocolExtensionContainerRATRestrictionsItemExtIEs struct {
	List []RATRestrictionsItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RATRestrictionsItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RATRestrictionsItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RATRestrictionsItemExtIEsExtensionValueChoiceExtendedRATRestrictionInformation int = 0
)

type RATRestrictionsItemExtIEsExtensionValue struct {
	Choice                            int
	ExtendedRATRestrictionInformation *ExtendedRATRestrictionInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDExtendedRATRestrictionInformation"`
}

type RATRestrictionInformation struct {
	Value BitString `vht:"sizeExx,sizeMin:8,sizeMax:8"`
}

type RecommendedCellsForPaging struct {
	RecommendedCellList RecommendedCellList
	IEExtensions        *ProtocolExtensionContainerRecommendedCellsForPagingExtIEs
}

type ProtocolExtensionContainerRecommendedCellsForPagingExtIEs struct {
	List []RecommendedCellsForPagingExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RecommendedCellsForPagingExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RecommendedCellsForPagingExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RecommendedCellsForPagingExtIEsExtensionValueChoiceNothing int = 0
)

type RecommendedCellsForPagingExtIEsExtensionValue struct {
	Choice int
}

type RecommendedCellList struct {
	List []RecommendedCellItem `vht:"sizeMin:1,sizeMax:MaxnoofRecommendedCells"`
}

type RecommendedCellItem struct {
	NGRANCGI         NGRANCGI
	TimeStayedInCell *Integer `vht:"valueMin:0,valueMax:4095"`
	IEExtensions     *ProtocolExtensionContainerRecommendedCellItemExtIEs
}

type ProtocolExtensionContainerRecommendedCellItemExtIEs struct {
	List []RecommendedCellItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RecommendedCellItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RecommendedCellItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RecommendedCellItemExtIEsExtensionValueChoiceNothing int = 0
)

type RecommendedCellItemExtIEsExtensionValue struct {
	Choice int
}

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList RecommendedRANNodeList
	IEExtensions           *ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs
}

type ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs struct {
	List []RecommendedRANNodesForPagingExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RecommendedRANNodesForPagingExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RecommendedRANNodesForPagingExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RecommendedRANNodesForPagingExtIEsExtensionValueChoiceNothing int = 0
)

type RecommendedRANNodesForPagingExtIEsExtensionValue struct {
	Choice int
}

type RecommendedRANNodeList struct {
	List []RecommendedRANNodeItem `vht:"sizeMin:1,sizeMax:MaxnoofRecommendedRANNodes"`
}

type RecommendedRANNodeItem struct {
	AMFPagingTarget AMFPagingTarget
	IEExtensions    *ProtocolExtensionContainerRecommendedRANNodeItemExtIEs
}

type ProtocolExtensionContainerRecommendedRANNodeItemExtIEs struct {
	List []RecommendedRANNodeItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RecommendedRANNodeItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RecommendedRANNodeItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RecommendedRANNodeItemExtIEsExtensionValueChoiceNothing int = 0
)

type RecommendedRANNodeItemExtIEsExtensionValue struct {
	Choice int
}

const (
	RedirectionVoiceFallbackPossible    Enumerated = 0
	RedirectionVoiceFallbackNotPossible Enumerated = 1
)

type RedirectionVoiceFallback struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

const (
	ReflectiveQosAttributeSubjectTo Enumerated = 0
)

type ReflectiveQosAttribute struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type RelativeAMFCapacity struct {
	Value Integer `vht:"valueMin:0,valueMax:255"`
}

const (
	ReportAreaCell Enumerated = 0
)

type ReportArea struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type RepetitionPeriod struct {
	Value Integer `vht:"valueMin:0,valueMax:131071"`
}

const (
	ResetAllResetAll Enumerated = 0
)

type ResetAll struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	ResetTypeChoiceNGInterface       int = 0
	ResetTypeChoicePartOfNGInterface int = 1
	ResetTypeChoiceChoiceExtensions  int = 2
)

type ResetType struct {
	Choice            int
	NGInterface       *ResetAll
	PartOfNGInterface *UEAssociatedLogicalNGConnectionList
	ChoiceExtensions  *ProtocolIESingleContainerResetTypeExtIEs
}

type ProtocolIESingleContainerResetTypeExtIEs struct {
	List []ResetTypeExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type ResetTypeExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	TypeValue    ResetTypeExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	ResetTypeExtIEsTypeValueChoiceNothing int = 0
)

type ResetTypeExtIEsTypeValue struct {
	Choice int
}

type RNCID struct {
	Value Integer `vht:"valueMin:0,valueMax:4095"`
}

type RoutingID struct {
	Value OctetString
}

type RRCContainer struct {
	Value OctetString
}

const (
	RRCEstablishmentCauseEmergency          Enumerated = 0
	RRCEstablishmentCauseHighPriorityAccess Enumerated = 1
	RRCEstablishmentCauseMtAccess           Enumerated = 2
	RRCEstablishmentCauseMoSignalling       Enumerated = 3
	RRCEstablishmentCauseMoData             Enumerated = 4
	RRCEstablishmentCauseMoVoiceCall        Enumerated = 5
	RRCEstablishmentCauseMoVideoCall        Enumerated = 6
	RRCEstablishmentCauseMoSMS              Enumerated = 7
	RRCEstablishmentCauseMpsPriorityAccess  Enumerated = 8
	RRCEstablishmentCauseMcsPriorityAccess  Enumerated = 9
	RRCEstablishmentCauseNotAvailable       Enumerated = 10
)

type RRCEstablishmentCause struct {
	Value Enumerated `vht:"valueMin:0,valueMax:10"`
}

const (
	RRCInactiveTransitionReportRequestSubsequentStateTransitionReport Enumerated = 0
	RRCInactiveTransitionReportRequestSingleRrcConnectedStateReport   Enumerated = 1
	RRCInactiveTransitionReportRequestCancelReport                    Enumerated = 2
)

type RRCInactiveTransitionReportRequest struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

const (
	RRCStateInactive  Enumerated = 0
	RRCStateConnected Enumerated = 1
)

type RRCState struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type RIMInformationTransfer struct {
	TargetRANNodeID TargetRANNodeID
	SourceRANNodeID SourceRANNodeID
	RIMInformation  RIMInformation
	IEExtensions    *ProtocolExtensionContainerRIMInformationTransferExtIEs
}

type ProtocolExtensionContainerRIMInformationTransferExtIEs struct {
	List []RIMInformationTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type RIMInformationTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      RIMInformationTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	RIMInformationTransferExtIEsExtensionValueChoiceNothing int = 0
)

type RIMInformationTransferExtIEsExtensionValue struct {
	Choice int
}

const (
	RIMRSDetectionRsDetected    Enumerated = 0
	RIMRSDetectionRsDisappeared Enumerated = 1
)

type RIMInformation struct {
	TargetgNBSetID GNBSetID
	RIMRSDetection Enumerated `vht:"valueExt,valueMin:0,valueMax:2"`
}

type GNBSetID struct {
	Value BitString `vht:"sizeMin:22,sizeMax:22"`
}

type SCTPTLAs struct {
	List []TransportLayerAddress `vht:"sizeMin:1,sizeMax:MaxnoofXnTLAs"`
}

type SD struct {
	Value OctetString `vht:"sizeMin:3,sizeMax:3"`
}

type SecondaryRATUsageInformation struct {
	PDUSessionUsageReport   *PDUSessionUsageReport
	QosFlowsUsageReportList *QoSFlowsUsageReportList
	IEExtensions            *ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs
}

type ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs struct {
	List []SecondaryRATUsageInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SecondaryRATUsageInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SecondaryRATUsageInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SecondaryRATUsageInformationExtIEsExtensionValueChoiceNothing int = 0
)

type SecondaryRATUsageInformationExtIEsExtensionValue struct {
	Choice int
}

type SecondaryRATDataUsageReportTransfer struct {
	SecondaryRATUsageInformation *SecondaryRATUsageInformation
	IEExtensions                 *ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs
}

type ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs struct {
	List []SecondaryRATDataUsageReportTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SecondaryRATDataUsageReportTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SecondaryRATDataUsageReportTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SecondaryRATDataUsageReportTransferExtIEsExtensionValueChoiceNothing int = 0
)

type SecondaryRATDataUsageReportTransferExtIEsExtensionValue struct {
	Choice int
}

type SecurityContext struct {
	NextHopChainingCount NextHopChainingCount
	NextHopNH            SecurityKey
	IEExtensions         *ProtocolExtensionContainerSecurityContextExtIEs
}

type ProtocolExtensionContainerSecurityContextExtIEs struct {
	List []SecurityContextExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SecurityContextExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SecurityContextExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SecurityContextExtIEsExtensionValueChoiceNothing int = 0
)

type SecurityContextExtIEsExtensionValue struct {
	Choice int
}

type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication
	MaximumIntegrityProtectedDataRateUL *MaximumIntegrityProtectedDataRate
	IEExtensions                        *ProtocolExtensionContainerSecurityIndicationExtIEs
}

type ProtocolExtensionContainerSecurityIndicationExtIEs struct {
	List []SecurityIndicationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SecurityIndicationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SecurityIndicationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SecurityIndicationExtIEsExtensionValueChoiceMaximumIntegrityProtectedDataRateDL int = 0
)

type SecurityIndicationExtIEsExtensionValue struct {
	Choice                              int
	MaximumIntegrityProtectedDataRateDL *MaximumIntegrityProtectedDataRate `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDMaximumIntegrityProtectedDataRateDL"`
}

type SecurityKey struct {
	Value BitString `vht:"sizeMin:256,sizeMax:256"`
}

type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult
	ConfidentialityProtectionResult ConfidentialityProtectionResult
	IEExtensions                    *ProtocolExtensionContainerSecurityResultExtIEs
}

type ProtocolExtensionContainerSecurityResultExtIEs struct {
	List []SecurityResultExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SecurityResultExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SecurityResultExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SecurityResultExtIEsExtensionValueChoiceNothing int = 0
)

type SecurityResultExtIEsExtensionValue struct {
	Choice int
}

type SerialNumber struct {
	Value BitString `vht:"sizeMin:16,sizeMax:16"`
}

type ServedGUAMIList struct {
	List []ServedGUAMIItem `vht:"sizeMin:1,sizeMax:MaxnoofServedGUAMIs"`
}

type ServedGUAMIItem struct {
	GUAMI         GUAMI
	BackupAMFName *AMFName
	IEExtensions  *ProtocolExtensionContainerServedGUAMIItemExtIEs
}

type ProtocolExtensionContainerServedGUAMIItemExtIEs struct {
	List []ServedGUAMIItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ServedGUAMIItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ServedGUAMIItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ServedGUAMIItemExtIEsExtensionValueChoiceGUAMIType int = 0
)

type ServedGUAMIItemExtIEsExtensionValue struct {
	Choice    int
	GUAMIType *GUAMIType `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDGUAMIType"`
}

type ServiceAreaInformation struct {
	List []ServiceAreaInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofEPLMNsPlusOne"`
}

type ServiceAreaInformationItem struct {
	PLMNIdentity   PLMNIdentity
	AllowedTACs    *AllowedTACs
	NotAllowedTACs *NotAllowedTACs
	IEExtensions   *ProtocolExtensionContainerServiceAreaInformationItemExtIEs
}

type ProtocolExtensionContainerServiceAreaInformationItemExtIEs struct {
	List []ServiceAreaInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ServiceAreaInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ServiceAreaInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ServiceAreaInformationItemExtIEsExtensionValueChoiceNothing int = 0
)

type ServiceAreaInformationItemExtIEsExtensionValue struct {
	Choice int
}

type SgNBUEX2APID struct {
	Value Integer `vht:"valueMin:0,valueMax:4294967295"`
}

type SliceOverloadList struct {
	List []SliceOverloadItem `vht:"sizeMin:1,sizeMax:MaxnoofSliceItems"`
}

type SliceOverloadItem struct {
	SNSSAI       SNSSAI
	IEExtensions *ProtocolExtensionContainerSliceOverloadItemExtIEs
}

type ProtocolExtensionContainerSliceOverloadItemExtIEs struct {
	List []SliceOverloadItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SliceOverloadItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                           `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SliceOverloadItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SliceOverloadItemExtIEsExtensionValueChoiceNothing int = 0
)

type SliceOverloadItemExtIEsExtensionValue struct {
	Choice int
}

type SliceSupportList struct {
	List []SliceSupportItem `vht:"sizeMin:1,sizeMax:MaxnoofSliceItems"`
}

type SliceSupportItem struct {
	SNSSAI       SNSSAI
	IEExtensions *ProtocolExtensionContainerSliceSupportItemExtIEs
}

type ProtocolExtensionContainerSliceSupportItemExtIEs struct {
	List []SliceSupportItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SliceSupportItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                          `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SliceSupportItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SliceSupportItemExtIEsExtensionValueChoiceNothing int = 0
)

type SliceSupportItemExtIEsExtensionValue struct {
	Choice int
}

type SNSSAI struct {
	SST          SST
	SD           *SD
	IEExtensions *ProtocolExtensionContainerSNSSAIExtIEs
}

type ProtocolExtensionContainerSNSSAIExtIEs struct {
	List []SNSSAIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SNSSAIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SNSSAIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SNSSAIExtIEsExtensionValueChoiceNothing int = 0
)

type SNSSAIExtIEsExtensionValue struct {
	Choice int
}

type SONConfigurationTransfer struct {
	TargetRANNodeID        TargetRANNodeID
	SourceRANNodeID        SourceRANNodeID
	SONInformation         SONInformation
	XnTNLConfigurationInfo *XnTNLConfigurationInfo
	IEExtensions           *ProtocolExtensionContainerSONConfigurationTransferExtIEs
}

type ProtocolExtensionContainerSONConfigurationTransferExtIEs struct {
	List []SONConfigurationTransferExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SONConfigurationTransferExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                  `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SONConfigurationTransferExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SONConfigurationTransferExtIEsExtensionValueChoiceNothing int = 0
)

type SONConfigurationTransferExtIEsExtensionValue struct {
	Choice int
}

const (
	SONInformationChoiceSONInformationRequest int = 0
	SONInformationChoiceSONInformationReply   int = 1
	SONInformationChoiceChoiceExtensions      int = 2
)

type SONInformation struct {
	Choice                int
	SONInformationRequest *SONInformationRequest
	SONInformationReply   *SONInformationReply
	ChoiceExtensions      *ProtocolIESingleContainerSONInformationExtIEs
}

type ProtocolIESingleContainerSONInformationExtIEs struct {
	List []SONInformationExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type SONInformationExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                   `vht:"Reference:ProtocolIEID"`
	TypeValue    SONInformationExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	SONInformationExtIEsTypeValueChoiceNothing int = 0
)

type SONInformationExtIEsTypeValue struct {
	Choice int
}

type SONInformationReply struct {
	XnTNLConfigurationInfo *XnTNLConfigurationInfo
	IEExtensions           *ProtocolExtensionContainerSONInformationReplyExtIEs
}

type ProtocolExtensionContainerSONInformationReplyExtIEs struct {
	List []SONInformationReplyExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SONInformationReplyExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SONInformationReplyExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SONInformationReplyExtIEsExtensionValueChoiceNothing int = 0
)

type SONInformationReplyExtIEsExtensionValue struct {
	Choice int
}

const (
	SONInformationRequestXnTNLConfigurationInfo Enumerated = 0
)

type SONInformationRequest struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      RRCContainer
	PDUSessionResourceInformationList *PDUSessionResourceInformationList
	ERABInformationList               *ERABInformationList
	TargetCellID                      NGRANCGI
	IndexToRFSP                       *IndexToRFSP
	UEHistoryInformation              UEHistoryInformation
	IEExtensions                      *ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs
}

type ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs struct {
	List []SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValueChoiceSgNBUEX2APID int = 0
)

type SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEsExtensionValue struct {
	Choice       int
	SgNBUEX2APID *SgNBUEX2APID `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSgNBUEX2APID"`
}

const (
	SourceOfUEActivityBehaviourInformationSubscriptionInformation Enumerated = 0
	SourceOfUEActivityBehaviourInformationStatistics              Enumerated = 1
)

type SourceOfUEActivityBehaviourInformation struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type SourceRANNodeID struct {
	GlobalRANNodeID GlobalRANNodeID
	SelectedTAI     TAI
	IEExtensions    *ProtocolExtensionContainerSourceRANNodeIDExtIEs
}

type ProtocolExtensionContainerSourceRANNodeIDExtIEs struct {
	List []SourceRANNodeIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SourceRANNodeIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SourceRANNodeIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SourceRANNodeIDExtIEsExtensionValueChoiceNothing int = 0
)

type SourceRANNodeIDExtIEsExtensionValue struct {
	Choice int
}

type SourceToTargetTransparentContainer struct {
	Value OctetString
}

type SourceToTargetAMFInformationReroute struct {
	ConfiguredNSSAI     *ConfiguredNSSAI
	RejectedNSSAIinPLMN *RejectedNSSAIinPLMN
	RejectedNSSAIinTA   *RejectedNSSAIinTA
	IEExtensions        *ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs
}

type ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs struct {
	List []SourceToTargetAMFInformationRerouteExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SourceToTargetAMFInformationRerouteExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SourceToTargetAMFInformationRerouteExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SourceToTargetAMFInformationRerouteExtIEsExtensionValueChoiceNothing int = 0
)

type SourceToTargetAMFInformationRerouteExtIEsExtensionValue struct {
	Choice int
}

const (
	SRVCCOperationPossiblePossible Enumerated = 0
)

type SRVCCOperationPossible struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type ConfiguredNSSAI struct {
	Value OctetString `vht:"sizeMin:128,sizeMax:128"`
}

type RejectedNSSAIinPLMN struct {
	Value OctetString `vht:"sizeMin:32,sizeMax:32"`
}

type RejectedNSSAIinTA struct {
	Value OctetString `vht:"sizeMin:32,sizeMax:32"`
}

type SST struct {
	Value OctetString `vht:"sizeMin:1,sizeMax:1"`
}

type SupportedTAList struct {
	List []SupportedTAItem `vht:"sizeMin:1,sizeMax:MaxnoofTACs"`
}

type SupportedTAItem struct {
	TAC               TAC
	BroadcastPLMNList BroadcastPLMNList
	IEExtensions      *ProtocolExtensionContainerSupportedTAItemExtIEs
}

type ProtocolExtensionContainerSupportedTAItemExtIEs struct {
	List []SupportedTAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type SupportedTAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      SupportedTAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	SupportedTAItemExtIEsExtensionValueChoiceRATInformation int = 0
)

type SupportedTAItemExtIEsExtensionValue struct {
	Choice         int
	RATInformation *RATInformation `vht:"Presence:PresenceOptional,Criticality:CriticalityReject,ProtocolExtensionID:ProtocolExtensionIDRATInformation"`
}

type TAC struct {
	Value OctetString `vht:"sizeMin:3,sizeMax:3"`
}

type TAI struct {
	PLMNIdentity PLMNIdentity
	TAC          TAC
	IEExtensions *ProtocolExtensionContainerTAIExtIEs
}

type ProtocolExtensionContainerTAIExtIEs struct {
	List []TAIExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAIExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAIExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAIExtIEsExtensionValueChoiceNothing int = 0
)

type TAIExtIEsExtensionValue struct {
	Choice int
}

type TAIBroadcastEUTRA struct {
	List []TAIBroadcastEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIforWarning"`
}

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI
	CompletedCellsInTAIEUTRA CompletedCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs
}

type ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs struct {
	List []TAIBroadcastEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAIBroadcastEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAIBroadcastEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAIBroadcastEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type TAIBroadcastEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type TAIBroadcastNR struct {
	List []TAIBroadcastNRItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIforWarning"`
}

type TAIBroadcastNRItem struct {
	TAI                   TAI
	CompletedCellsInTAINR CompletedCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAIBroadcastNRItemExtIEs
}

type ProtocolExtensionContainerTAIBroadcastNRItemExtIEs struct {
	List []TAIBroadcastNRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAIBroadcastNRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAIBroadcastNRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAIBroadcastNRItemExtIEsExtensionValueChoiceNothing int = 0
)

type TAIBroadcastNRItemExtIEsExtensionValue struct {
	Choice int
}

type TAICancelledEUTRA struct {
	List []TAICancelledEUTRAItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIforWarning"`
}

type TAICancelledEUTRAItem struct {
	TAI                      TAI
	CancelledCellsInTAIEUTRA CancelledCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs
}

type ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs struct {
	List []TAICancelledEUTRAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAICancelledEUTRAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAICancelledEUTRAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAICancelledEUTRAItemExtIEsExtensionValueChoiceNothing int = 0
)

type TAICancelledEUTRAItemExtIEsExtensionValue struct {
	Choice int
}

type TAICancelledNR struct {
	List []TAICancelledNRItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIforWarning"`
}

type TAICancelledNRItem struct {
	TAI                   TAI
	CancelledCellsInTAINR CancelledCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAICancelledNRItemExtIEs
}

type ProtocolExtensionContainerTAICancelledNRItemExtIEs struct {
	List []TAICancelledNRItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAICancelledNRItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAICancelledNRItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAICancelledNRItemExtIEsExtensionValueChoiceNothing int = 0
)

type TAICancelledNRItemExtIEsExtensionValue struct {
	Choice int
}

type TAIListForInactive struct {
	List []TAIListForInactiveItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIforInactive"`
}

type TAIListForInactiveItem struct {
	TAI          TAI
	IEExtensions *ProtocolExtensionContainerTAIListForInactiveItemExtIEs
}

type ProtocolExtensionContainerTAIListForInactiveItemExtIEs struct {
	List []TAIListForInactiveItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAIListForInactiveItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAIListForInactiveItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAIListForInactiveItemExtIEsExtensionValueChoiceNothing int = 0
)

type TAIListForInactiveItemExtIEsExtensionValue struct {
	Choice int
}

type TAIListForPaging struct {
	List []TAIListForPagingItem `vht:"sizeMin:1,sizeMax:MaxnoofTAIforPaging"`
}

type TAIListForPagingItem struct {
	TAI          TAI
	IEExtensions *ProtocolExtensionContainerTAIListForPagingItemExtIEs
}

type ProtocolExtensionContainerTAIListForPagingItemExtIEs struct {
	List []TAIListForPagingItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TAIListForPagingItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TAIListForPagingItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TAIListForPagingItemExtIEsExtensionValueChoiceNothing int = 0
)

type TAIListForPagingItemExtIEsExtensionValue struct {
	Choice int
}

type TAIListForRestart struct {
	List []TAI `vht:"sizeMin:1,sizeMax:MaxnoofTAIforRestart"`
}

type TAIListForWarning struct {
	List []TAI `vht:"sizeMin:1,sizeMax:MaxnoofTAIforWarning"`
}

type TargeteNBID struct {
	GlobalENBID    GlobalNgENBID
	SelectedEPSTAI EPSTAI
	IEExtensions   *ProtocolExtensionContainerTargeteNBIDExtIEs
}

type ProtocolExtensionContainerTargeteNBIDExtIEs struct {
	List []TargeteNBIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TargeteNBIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TargeteNBIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TargeteNBIDExtIEsExtensionValueChoiceNothing int = 0
)

type TargeteNBIDExtIEsExtensionValue struct {
	Choice int
}

const (
	TargetIDChoiceTargetRANNodeID  int = 0
	TargetIDChoiceTargeteNBID      int = 1
	TargetIDChoiceChoiceExtensions int = 2
)

type TargetID struct {
	Choice           int
	TargetRANNodeID  *TargetRANNodeID
	TargeteNBID      *TargeteNBID
	ChoiceExtensions *ProtocolIESingleContainerTargetIDExtIEs
}

type ProtocolIESingleContainerTargetIDExtIEs struct {
	List []TargetIDExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type TargetIDExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality             `vht:"Reference:ProtocolIEID"`
	TypeValue    TargetIDExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	TargetIDExtIEsTypeValueChoiceTargetRNCID int = 0
)

type TargetIDExtIEsTypeValue struct {
	Choice      int
	TargetRNCID *TargetRNCID `vht:"Presence:PresenceMandatory,Criticality:CriticalityReject,ProtocolIEID:ProtocolIEIDTargetRNCID"`
}

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer RRCContainer
	IEExtensions *ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs
}

type ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs struct {
	List []TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValueChoiceNothing int = 0
)

type TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEsExtensionValue struct {
	Choice int
}

type TargetRANNodeID struct {
	GlobalRANNodeID GlobalRANNodeID
	SelectedTAI     TAI
	IEExtensions    *ProtocolExtensionContainerTargetRANNodeIDExtIEs
}

type ProtocolExtensionContainerTargetRANNodeIDExtIEs struct {
	List []TargetRANNodeIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TargetRANNodeIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TargetRANNodeIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TargetRANNodeIDExtIEsExtensionValueChoiceNothing int = 0
)

type TargetRANNodeIDExtIEsExtensionValue struct {
	Choice int
}

type TargetRNCID struct {
	LAI           LAI
	RNCID         RNCID
	ExtendedRNCID *ExtendedRNCID
	IEExtensions  *ProtocolExtensionContainerTargetRNCIDExtIEs
}

type ProtocolExtensionContainerTargetRNCIDExtIEs struct {
	List []TargetRNCIDExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TargetRNCIDExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                     `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TargetRNCIDExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TargetRNCIDExtIEsExtensionValueChoiceNothing int = 0
)

type TargetRNCIDExtIEsExtensionValue struct {
	Choice int
}

type TargetToSourceTransparentContainer struct {
	Value OctetString
}

const (
	TimerApproachForGUAMIRemovalApplyTimer Enumerated = 0
)

type TimerApproachForGUAMIRemoval struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type TimeStamp struct {
	Value OctetString `vht:"sizeMin:4,sizeMax:4"`
}

const (
	TimeToWaitV1s  Enumerated = 0
	TimeToWaitV2s  Enumerated = 1
	TimeToWaitV5s  Enumerated = 2
	TimeToWaitV10s Enumerated = 3
	TimeToWaitV20s Enumerated = 4
	TimeToWaitV60s Enumerated = 5
)

type TimeToWait struct {
	Value Enumerated `vht:"valueMin:0,valueMax:5"`
}

type TimeUEStayedInCell struct {
	Value Integer `vht:"valueMin:0,valueMax:4095"`
}

type TimeUEStayedInCellEnhancedGranularity struct {
	Value Integer `vht:"valueMin:0,valueMax:40950"`
}

type TNLAddressWeightFactor struct {
	Value Integer `vht:"valueMin:0,valueMax:255"`
}

type TNLAssociationList struct {
	List []TNLAssociationItem `vht:"sizeMin:1,sizeMax:MaxnoofTNLAssociations"`
}

type TNLAssociationItem struct {
	TNLAssociationAddress CPTransportLayerInformation
	Cause                 Cause
	IEExtensions          *ProtocolExtensionContainerTNLAssociationItemExtIEs
}

type ProtocolExtensionContainerTNLAssociationItemExtIEs struct {
	List []TNLAssociationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TNLAssociationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                            `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TNLAssociationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TNLAssociationItemExtIEsExtensionValueChoiceNothing int = 0
)

type TNLAssociationItemExtIEsExtensionValue struct {
	Choice int
}

const (
	TNLAssociationUsageUe    Enumerated = 0
	TNLAssociationUsageNonUe Enumerated = 1
	TNLAssociationUsageBoth  Enumerated = 2
)

type TNLAssociationUsage struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

type TraceActivation struct {
	NGRANTraceID                   NGRANTraceID
	InterfacesToTrace              InterfacesToTrace
	TraceDepth                     TraceDepth
	TraceCollectionEntityIPAddress TransportLayerAddress
	IEExtensions                   *ProtocolExtensionContainerTraceActivationExtIEs
}

type ProtocolExtensionContainerTraceActivationExtIEs struct {
	List []TraceActivationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type TraceActivationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      TraceActivationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	TraceActivationExtIEsExtensionValueChoiceNothing int = 0
)

type TraceActivationExtIEsExtensionValue struct {
	Choice int
}

const (
	TraceDepthMinimum                               Enumerated = 0
	TraceDepthMedium                                Enumerated = 1
	TraceDepthMaximum                               Enumerated = 2
	TraceDepthMinimumWithoutVendorSpecificExtension Enumerated = 3
	TraceDepthMediumWithoutVendorSpecificExtension  Enumerated = 4
	TraceDepthMaximumWithoutVendorSpecificExtension Enumerated = 5
)

type TraceDepth struct {
	Value Enumerated `vht:"valueMin:0,valueMax:5"`
}

type TrafficLoadReductionIndication struct {
	Value Integer `vht:"valueMin:1,valueMax:99"`
}

type TransportLayerAddress struct {
	Value BitString `vht:"sizeExx,sizeExt,sizeMin:1,sizeMax:160"`
}

const (
	TypeOfErrorNotUnderstood Enumerated = 0
	TypeOfErrorMissing       Enumerated = 1
)

type TypeOfError struct {
	Value Enumerated `vht:"valueMin:0,valueMax:1"`
}

type UEAggregateMaximumBitRate struct {
	UEAggregateMaximumBitRateDL BitRate
	UEAggregateMaximumBitRateUL BitRate
	IEExtensions                *ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs
}

type ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs struct {
	List []UEAggregateMaximumBitRateExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UEAggregateMaximumBitRateExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UEAggregateMaximumBitRateExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UEAggregateMaximumBitRateExtIEsExtensionValueChoiceNothing int = 0
)

type UEAggregateMaximumBitRateExtIEsExtensionValue struct {
	Choice int
}

type UEAssociatedLogicalNGConnectionList struct {
	List []UEAssociatedLogicalNGConnectionItem `vht:"sizeMin:1,sizeMax:MaxnoofNGConnectionsToReset"`
}

type UEAssociatedLogicalNGConnectionItem struct {
	AMFUENGAPID  *AMFUENGAPID
	RANUENGAPID  *RANUENGAPID
	IEExtensions *ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs
}

type ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs struct {
	List []UEAssociatedLogicalNGConnectionItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UEAssociatedLogicalNGConnectionItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UEAssociatedLogicalNGConnectionItemExtIEsExtensionValueChoiceNothing int = 0
)

type UEAssociatedLogicalNGConnectionItemExtIEsExtensionValue struct {
	Choice int
}

const (
	UEContextRequestRequested Enumerated = 0
)

type UEContextRequest struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type UEHistoryInformation struct {
	List []LastVisitedCellItem `vht:"sizeMin:1,sizeMax:MaxnoofCellsinUEHistoryInfo"`
}

const (
	UEIdentityIndexValueChoiceIndexLength10    int = 0
	UEIdentityIndexValueChoiceChoiceExtensions int = 1
)

type UEIdentityIndexValue struct {
	Choice           int
	IndexLength10    *BitString `vht:"valueMin:10,valueMax:10"`
	ChoiceExtensions *ProtocolIESingleContainerUEIdentityIndexValueExtIEs
}

type ProtocolIESingleContainerUEIdentityIndexValueExtIEs struct {
	List []UEIdentityIndexValueExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type UEIdentityIndexValueExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                         `vht:"Reference:ProtocolIEID"`
	TypeValue    UEIdentityIndexValueExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEIdentityIndexValueExtIEsTypeValueChoiceNothing int = 0
)

type UEIdentityIndexValueExtIEsTypeValue struct {
	Choice int
}

const (
	UENGAPIDsChoiceUENGAPIDPair     int = 0
	UENGAPIDsChoiceAMFUENGAPID      int = 1
	UENGAPIDsChoiceChoiceExtensions int = 2
)

type UENGAPIDs struct {
	Choice           int
	UENGAPIDPair     *UENGAPIDPair
	AMFUENGAPID      *AMFUENGAPID
	ChoiceExtensions *ProtocolIESingleContainerUENGAPIDsExtIEs
}

type ProtocolIESingleContainerUENGAPIDsExtIEs struct {
	List []UENGAPIDsExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type UENGAPIDsExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality              `vht:"Reference:ProtocolIEID"`
	TypeValue    UENGAPIDsExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UENGAPIDsExtIEsTypeValueChoiceNothing int = 0
)

type UENGAPIDsExtIEsTypeValue struct {
	Choice int
}

type UENGAPIDPair struct {
	AMFUENGAPID  AMFUENGAPID
	RANUENGAPID  RANUENGAPID
	IEExtensions *ProtocolExtensionContainerUENGAPIDPairExtIEs
}

type ProtocolExtensionContainerUENGAPIDPairExtIEs struct {
	List []UENGAPIDPairExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UENGAPIDPairExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UENGAPIDPairExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UENGAPIDPairExtIEsExtensionValueChoiceNothing int = 0
)

type UENGAPIDPairExtIEsExtensionValue struct {
	Choice int
}

const (
	UEPagingIdentityChoiceFiveGSTMSI       int = 0
	UEPagingIdentityChoiceChoiceExtensions int = 1
)

type UEPagingIdentity struct {
	Choice           int
	FiveGSTMSI       *FiveGSTMSI
	ChoiceExtensions *ProtocolIESingleContainerUEPagingIdentityExtIEs
}

type ProtocolIESingleContainerUEPagingIdentityExtIEs struct {
	List []UEPagingIdentityExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type UEPagingIdentityExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                     `vht:"Reference:ProtocolIEID"`
	TypeValue    UEPagingIdentityExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UEPagingIdentityExtIEsTypeValueChoiceNothing int = 0
)

type UEPagingIdentityExtIEsTypeValue struct {
	Choice int
}

const (
	UEPresenceIn      Enumerated = 0
	UEPresenceOut     Enumerated = 1
	UEPresenceUnknown Enumerated = 2
)

type UEPresence struct {
	Value Enumerated `vht:"valueMin:0,valueMax:2"`
}

type UEPresenceInAreaOfInterestList struct {
	List []UEPresenceInAreaOfInterestItem `vht:"sizeMin:1,sizeMax:MaxnoofAoI"`
}

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID LocationReportingReferenceID
	UEPresence                   UEPresence
	IEExtensions                 *ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs
}

type ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs struct {
	List []UEPresenceInAreaOfInterestItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UEPresenceInAreaOfInterestItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                        `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UEPresenceInAreaOfInterestItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UEPresenceInAreaOfInterestItemExtIEsExtensionValueChoiceNothing int = 0
)

type UEPresenceInAreaOfInterestItemExtIEsExtensionValue struct {
	Choice int
}

type UERadioCapability struct {
	Value OctetString
}

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA
	IEExtensions                      *ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs
}

type ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs struct {
	List []UERadioCapabilityForPagingExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UERadioCapabilityForPagingExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                    `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UERadioCapabilityForPagingExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UERadioCapabilityForPagingExtIEsExtensionValueChoiceNothing int = 0
)

type UERadioCapabilityForPagingExtIEsExtensionValue struct {
	Choice int
}

type UERadioCapabilityForPagingOfNR struct {
	Value OctetString
}

type UERadioCapabilityForPagingOfEUTRA struct {
	Value OctetString
}

const (
	UERetentionInformationUesRetained Enumerated = 0
)

type UERetentionInformation struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             NRencryptionAlgorithms
	NRintegrityProtectionAlgorithms    NRintegrityProtectionAlgorithms
	EUTRAencryptionAlgorithms          EUTRAencryptionAlgorithms
	EUTRAintegrityProtectionAlgorithms EUTRAintegrityProtectionAlgorithms
	IEExtensions                       *ProtocolExtensionContainerUESecurityCapabilitiesExtIEs
}

type ProtocolExtensionContainerUESecurityCapabilitiesExtIEs struct {
	List []UESecurityCapabilitiesExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UESecurityCapabilitiesExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UESecurityCapabilitiesExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UESecurityCapabilitiesExtIEsExtensionValueChoiceNothing int = 0
)

type UESecurityCapabilitiesExtIEsExtensionValue struct {
	Choice int
}

type ULNGUUPTNLModifyList struct {
	List []ULNGUUPTNLModifyItem `vht:"sizeMin:1,sizeMax:MaxnoofMultiConnectivity"`
}

type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation
	DLNGUUPTNLInformation UPTransportLayerInformation
	IEExtensions          *ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs
}

type ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs struct {
	List []ULNGUUPTNLModifyItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type ULNGUUPTNLModifyItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      ULNGUUPTNLModifyItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	ULNGUUPTNLModifyItemExtIEsExtensionValueChoiceNothing int = 0
)

type ULNGUUPTNLModifyItemExtIEsExtensionValue struct {
	Choice int
}

type UnavailableGUAMIList struct {
	List []UnavailableGUAMIItem `vht:"sizeMin:1,sizeMax:MaxnoofServedGUAMIs"`
}

type UnavailableGUAMIItem struct {
	GUAMI                        GUAMI
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval
	BackupAMFName                *AMFName
	IEExtensions                 *ProtocolExtensionContainerUnavailableGUAMIItemExtIEs
}

type ProtocolExtensionContainerUnavailableGUAMIItemExtIEs struct {
	List []UnavailableGUAMIItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UnavailableGUAMIItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                              `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UnavailableGUAMIItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UnavailableGUAMIItemExtIEsExtensionValueChoiceNothing int = 0
)

type UnavailableGUAMIItemExtIEsExtensionValue struct {
	Choice int
}

const (
	ULForwardingUlForwardingProposed Enumerated = 0
)

type ULForwarding struct {
	Value Enumerated `vht:"valueMin:0,valueMax:0"`
}

const (
	UPTransportLayerInformationChoiceGTPTunnel        int = 0
	UPTransportLayerInformationChoiceChoiceExtensions int = 1
)

type UPTransportLayerInformation struct {
	Choice           int
	GTPTunnel        *GTPTunnel
	ChoiceExtensions *ProtocolIESingleContainerUPTransportLayerInformationExtIEs
}

type ProtocolIESingleContainerUPTransportLayerInformationExtIEs struct {
	List []UPTransportLayerInformationExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type UPTransportLayerInformationExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                                `vht:"Reference:ProtocolIEID"`
	TypeValue    UPTransportLayerInformationExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UPTransportLayerInformationExtIEsTypeValueChoiceNothing int = 0
)

type UPTransportLayerInformationExtIEsTypeValue struct {
	Choice int
}

type UPTransportLayerInformationList struct {
	List []UPTransportLayerInformationItem `vht:"sizeMin:1,sizeMax:MaxnoofMultiConnectivityMinusOne"`
}

type UPTransportLayerInformationItem struct {
	NGUUPTNLInformation UPTransportLayerInformation
	IEExtensions        *ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs
}

type ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs struct {
	List []UPTransportLayerInformationItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UPTransportLayerInformationItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                         `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UPTransportLayerInformationItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UPTransportLayerInformationItemExtIEsExtensionValueChoiceNothing int = 0
)

type UPTransportLayerInformationItemExtIEsExtensionValue struct {
	Choice int
}

type UPTransportLayerInformationPairList struct {
	List []UPTransportLayerInformationPairItem `vht:"sizeMin:1,sizeMax:MaxnoofMultiConnectivityMinusOne"`
}

type UPTransportLayerInformationPairItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation
	DLNGUUPTNLInformation UPTransportLayerInformation
	IEExtensions          *ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs
}

type ProtocolExtensionContainerUPTransportLayerInformationPairItemExtIEs struct {
	List []UPTransportLayerInformationPairItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UPTransportLayerInformationPairItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                             `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UPTransportLayerInformationPairItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UPTransportLayerInformationPairItemExtIEsExtensionValueChoiceNothing int = 0
)

type UPTransportLayerInformationPairItemExtIEsExtensionValue struct {
	Choice int
}

const (
	UserLocationInformationChoiceUserLocationInformationEUTRA int = 0
	UserLocationInformationChoiceUserLocationInformationNR    int = 1
	UserLocationInformationChoiceUserLocationInformationN3IWF int = 2
	UserLocationInformationChoiceChoiceExtensions             int = 3
)

type UserLocationInformation struct {
	Choice                       int
	UserLocationInformationEUTRA *UserLocationInformationEUTRA
	UserLocationInformationNR    *UserLocationInformationNR
	UserLocationInformationN3IWF *UserLocationInformationN3IWF
	ChoiceExtensions             *ProtocolIESingleContainerUserLocationInformationExtIEs
}

type ProtocolIESingleContainerUserLocationInformationExtIEs struct {
	List []UserLocationInformationExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type UserLocationInformationExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                            `vht:"Reference:ProtocolIEID"`
	TypeValue    UserLocationInformationExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	UserLocationInformationExtIEsTypeValueChoiceNothing int = 0
)

type UserLocationInformationExtIEsTypeValue struct {
	Choice int
}

type UserLocationInformationEUTRA struct {
	EUTRACGI     EUTRACGI
	TAI          TAI
	TimeStamp    *TimeStamp
	IEExtensions *ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs
}

type ProtocolExtensionContainerUserLocationInformationEUTRAExtIEs struct {
	List []UserLocationInformationEUTRAExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UserLocationInformationEUTRAExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UserLocationInformationEUTRAExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UserLocationInformationEUTRAExtIEsExtensionValueChoicePSCellInformation int = 0
)

type UserLocationInformationEUTRAExtIEsExtensionValue struct {
	Choice            int
	PSCellInformation *NGRANCGI `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDPSCellInformation"`
}

type UserLocationInformationN3IWF struct {
	IPAddress    TransportLayerAddress
	PortNumber   PortNumber
	IEExtensions *ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs
}

type ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs struct {
	List []UserLocationInformationN3IWFExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UserLocationInformationN3IWFExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UserLocationInformationN3IWFExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UserLocationInformationN3IWFExtIEsExtensionValueChoiceNothing int = 0
)

type UserLocationInformationN3IWFExtIEsExtensionValue struct {
	Choice int
}

type UserLocationInformationNR struct {
	NRCGI        NRCGI
	TAI          TAI
	TimeStamp    *TimeStamp
	IEExtensions *ProtocolExtensionContainerUserLocationInformationNRExtIEs
}

type ProtocolExtensionContainerUserLocationInformationNRExtIEs struct {
	List []UserLocationInformationNRExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UserLocationInformationNRExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                   `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UserLocationInformationNRExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UserLocationInformationNRExtIEsExtensionValueChoicePSCellInformation int = 0
)

type UserLocationInformationNRExtIEsExtensionValue struct {
	Choice            int
	PSCellInformation *NGRANCGI `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDPSCellInformation"`
}

type UserPlaneSecurityInformation struct {
	SecurityResult     SecurityResult
	SecurityIndication SecurityIndication
	IEExtensions       *ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs
}

type ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs struct {
	List []UserPlaneSecurityInformationExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type UserPlaneSecurityInformationExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      UserPlaneSecurityInformationExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	UserPlaneSecurityInformationExtIEsExtensionValueChoiceNothing int = 0
)

type UserPlaneSecurityInformationExtIEsExtensionValue struct {
	Choice int
}

type VolumeTimedReportList struct {
	List []VolumeTimedReportItem `vht:"sizeMin:1,sizeMax:MaxnoofTimePeriods"`
}

type VolumeTimedReportItem struct {
	StartTimeStamp OctetString `vht:"sizeMin:4,sizeMax:4"`
	EndTimeStamp   OctetString `vht:"sizeMin:4,sizeMax:4"`
	UsageCountUL   Integer     `vht:"valueMin:0,valueMax:18446744073709551615"`
	UsageCountDL   Integer     `vht:"valueMin:0,valueMax:18446744073709551615"`
	IEExtensions   *ProtocolExtensionContainerVolumeTimedReportItemExtIEs
}

type ProtocolExtensionContainerVolumeTimedReportItemExtIEs struct {
	List []VolumeTimedReportItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type VolumeTimedReportItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                               `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      VolumeTimedReportItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	VolumeTimedReportItemExtIEsExtensionValueChoiceNothing int = 0
)

type VolumeTimedReportItemExtIEsExtensionValue struct {
	Choice int
}

type WarningAreaCoordinates struct {
	Value OctetString `vht:"sizeExt,sizeMin:1,sizeMax:1024"`
}

const (
	WarningAreaListChoiceEUTRACGIListForWarning int = 0
	WarningAreaListChoiceNRCGIListForWarning    int = 1
	WarningAreaListChoiceTAIListForWarning      int = 2
	WarningAreaListChoiceEmergencyAreaIDList    int = 3
	WarningAreaListChoiceChoiceExtensions       int = 4
)

type WarningAreaList struct {
	Choice                 int
	EUTRACGIListForWarning *EUTRACGIListForWarning
	NRCGIListForWarning    *NRCGIListForWarning
	TAIListForWarning      *TAIListForWarning
	EmergencyAreaIDList    *EmergencyAreaIDList
	ChoiceExtensions       *ProtocolIESingleContainerWarningAreaListExtIEs
}

type ProtocolIESingleContainerWarningAreaListExtIEs struct {
	List []WarningAreaListExtIEs `vht:"valueMin:0,valueMax:MaxProtocolIEs"`
}

type WarningAreaListExtIEs struct {
	ProtocolIEID ProtocolIEID
	Criticality  Criticality                    `vht:"Reference:ProtocolIEID"`
	TypeValue    WarningAreaListExtIEsTypeValue `vht:"Reference:ProtocolIEID"`
}

const (
	WarningAreaListExtIEsTypeValueChoiceNothing int = 0
)

type WarningAreaListExtIEsTypeValue struct {
	Choice int
}

type WarningMessageContents struct {
	Value OctetString `vht:"sizeExt,sizeMin:1,sizeMax:9600"`
}

type WarningSecurityInfo struct {
	Value OctetString `vht:"sizeMin:50,sizeMax:50"`
}

type WarningType struct {
	Value OctetString `vht:"sizeMin:2,sizeMax:2"`
}

type XnExtTLAs struct {
	List []XnExtTLAItem `vht:"sizeMin:1,sizeMax:MaxnoofXnExtTLAs"`
}

type XnExtTLAItem struct {
	IPsecTLA     *TransportLayerAddress
	GTPTLAs      *XnGTPTLAs
	IEExtensions *ProtocolExtensionContainerXnExtTLAItemExtIEs
}

type ProtocolExtensionContainerXnExtTLAItemExtIEs struct {
	List []XnExtTLAItemExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type XnExtTLAItemExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                      `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      XnExtTLAItemExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	XnExtTLAItemExtIEsExtensionValueChoiceSCTPTLAs int = 0
)

type XnExtTLAItemExtIEsExtensionValue struct {
	Choice   int
	SCTPTLAs *SCTPTLAs `vht:"Presence:PresenceOptional,Criticality:CriticalityIgnore,ProtocolExtensionID:ProtocolExtensionIDSCTPTLAs"`
}

type XnGTPTLAs struct {
	List []TransportLayerAddress `vht:"sizeMin:1,sizeMax:MaxnoofXnGTP-TLAs"`
}

type XnTLAs struct {
	List []TransportLayerAddress `vht:"sizeMin:1,sizeMax:MaxnoofXnTLAs"`
}

type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses         XnTLAs
	XnExtendedTransportLayerAddresses *XnExtTLAs
	IEExtensions                      *ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs
}

type ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs struct {
	List []XnTNLConfigurationInfoExtIEs `vht:"sizeMin:0,sizeMax:MaxProtocolExtensions"`
}

type XnTNLConfigurationInfoExtIEs struct {
	ProtocolExtensionID ProtocolExtensionID
	Criticality         Criticality                                `vht:"Reference:ProtocolExtensionID"`
	ExtensionValue      XnTNLConfigurationInfoExtIEsExtensionValue `vht:"Reference:ProtocolExtensionID"`
}

const (
	XnTNLConfigurationInfoExtIEsExtensionValueChoiceNothing int = 0
)

type XnTNLConfigurationInfoExtIEsExtensionValue struct {
	Choice int
}
