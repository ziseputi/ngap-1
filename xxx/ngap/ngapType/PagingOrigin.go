// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PagingOriginPresentNon3gpp Enumerated = 0
)

type PagingOrigin struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
