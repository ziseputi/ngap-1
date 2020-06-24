// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ReflectiveQosAttributePresentSubjectTo Enumerated = 0
)

type ReflectiveQosAttribute struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
