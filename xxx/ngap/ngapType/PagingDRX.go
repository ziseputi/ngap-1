// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PagingDRXPresentV32  Enumerated = 0
	PagingDRXPresentV64  Enumerated = 1
	PagingDRXPresentV128 Enumerated = 2
	PagingDRXPresentV256 Enumerated = 3
)

type PagingDRX struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:3"`
}
