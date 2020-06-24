// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	ReportAreaPresentCell Enumerated = 0
)

type ReportArea struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
