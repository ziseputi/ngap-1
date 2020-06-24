// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	TNLAssociationUsagePresentUe    Enumerated = 0
	TNLAssociationUsagePresentNonUe Enumerated = 1
	TNLAssociationUsagePresentBoth  Enumerated = 2
)

type TNLAssociationUsage struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:2"`
}
