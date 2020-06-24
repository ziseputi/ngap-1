// Created By HaoDHH-245789 VHT2020
package ngapType

type CoreNetworkAssistanceInformation struct {
	UEIdentityIndexValue            UEIdentityIndexValue `vht:"valueLB:0,valueUB:1"`
	UESpecificDRX                   *PagingDRX           `vht:"optional"`
	PeriodicRegistrationUpdateTimer PeriodicRegistrationUpdateTimer
	MICOModeIndication              *MICOModeIndication `vht:"optional"`
	TAIListForInactive              TAIListForInactive
	ExpectedUEBehaviour             *ExpectedUEBehaviour                                              `vht:"valueExt,optional"`
	IEExtensions                    *ProtocolExtensionContainerCoreNetworkAssistanceInformationExtIEs `vht:"optional"`
}
