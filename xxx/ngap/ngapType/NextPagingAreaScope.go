// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	NextPagingAreaScopePresentSame    Enumerated = 0
	NextPagingAreaScopePresentChanged Enumerated = 1
)

type NextPagingAreaScope struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:1"`
}
