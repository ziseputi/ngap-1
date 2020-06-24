// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	DLNGUTNLInformationReusedPresentTrue Enumerated = 0
)

type DLNGUTNLInformationReused struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
