// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	MICOModeIndicationPresentTrue Enumerated = 0
)

type MICOModeIndication struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
