// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CauseNasPresentNormalRelease         Enumerated = 0
	CauseNasPresentAuthenticationFailure Enumerated = 1
	CauseNasPresentDeregister            Enumerated = 2
	CauseNasPresentUnspecified           Enumerated = 3
)

type CauseNas struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:3"`
}
