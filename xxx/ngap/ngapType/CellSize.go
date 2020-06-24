// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	CellSizePresentVerysmall Enumerated = 0
	CellSizePresentSmall     Enumerated = 1
	CellSizePresentMedium    Enumerated = 2
	CellSizePresentLarge     Enumerated = 3
)

type CellSize struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:3"`
}
