// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	UserLocationInformationPresentNothing int = iota /* No components present */
	UserLocationInformationPresentUserLocationInformationEUTRA
	UserLocationInformationPresentUserLocationInformationNR
	UserLocationInformationPresentUserLocationInformationN3IWF
	UserLocationInformationPresentChoiceExtensions
)

type UserLocationInformation struct {
	Present                      int
	UserLocationInformationEUTRA *UserLocationInformationEUTRA `vht:"valueExt"`
	UserLocationInformationNR    *UserLocationInformationNR    `vht:"valueExt"`
	UserLocationInformationN3IWF *UserLocationInformationN3IWF `vht:"valueExt"`
	ChoiceExtensions             *ProtocolIESingleContainerUserLocationInformationExtIEs
}
