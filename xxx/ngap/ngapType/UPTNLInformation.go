// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	UPTNLInformationPresentNothing int = iota /* No components present */
	UPTNLInformationPresentSingleTNLInformation
	UPTNLInformationPresentMultipleTNLInformation
	UPTNLInformationPresentChoiceExtensions
)

type UPTNLInformation struct {
	Present                int
	SingleTNLInformation   *SingleTNLInformation   `vht:"valueExt"`
	MultipleTNLInformation *MultipleTNLInformation `vht:"valueExt"`
	ChoiceExtensions       *ProtocolIESingleContainerUPTNLInformationExtIEs
}
