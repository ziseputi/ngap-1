// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	SONInformationPresentNothing int = iota /* No components present */
	SONInformationPresentSONInformationRequest
	SONInformationPresentSONInformationReply
	SONInformationPresentChoiceExtensions
)

type SONInformation struct {
	Present               int
	SONInformationRequest *SONInformationRequest
	SONInformationReply   *SONInformationReply `vht:"valueExt"`
	ChoiceExtensions      *ProtocolIESingleContainerSONInformationExtIEs
}
