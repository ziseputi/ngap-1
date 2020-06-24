// Created By HaoDHH-245789 VHT2020
package ngapType

type PrivateMessageIEs struct {
	Id          PrivateIEID
	Criticality Criticality
	Value       PrivateMessageIEsValue `vht:"openType,referenceFieldName:Id"`
}

const (
	PrivateMessageIEsPresentNothing int = iota /* No components present */
)

type PrivateMessageIEsValue struct {
	Present int
}
