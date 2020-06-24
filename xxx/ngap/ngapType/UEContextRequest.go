// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	UEContextRequestPresentRequested Enumerated = 0
)

type UEContextRequest struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:0"`
}
