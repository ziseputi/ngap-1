// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	QosCharacteristicsPresentNothing int = iota /* No components present */
	QosCharacteristicsPresentNonDynamic5QI
	QosCharacteristicsPresentDynamic5QI
	QosCharacteristicsPresentChoiceExtensions
)

type QosCharacteristics struct {
	Present          int
	NonDynamic5QI    *NonDynamic5QIDescriptor `vht:"valueExt"`
	Dynamic5QI       *Dynamic5QIDescriptor    `vht:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerQosCharacteristicsExtIEs
}
