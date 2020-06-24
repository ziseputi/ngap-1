// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	PagingPriorityPresentPriolevel1 Enumerated = 0
	PagingPriorityPresentPriolevel2 Enumerated = 1
	PagingPriorityPresentPriolevel3 Enumerated = 2
	PagingPriorityPresentPriolevel4 Enumerated = 3
	PagingPriorityPresentPriolevel5 Enumerated = 4
	PagingPriorityPresentPriolevel6 Enumerated = 5
	PagingPriorityPresentPriolevel7 Enumerated = 6
	PagingPriorityPresentPriolevel8 Enumerated = 7
)

type PagingPriority struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:7"`
}
