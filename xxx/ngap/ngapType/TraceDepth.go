// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	TraceDepthPresentMinimum                               Enumerated = 0
	TraceDepthPresentMedium                                Enumerated = 1
	TraceDepthPresentMaximum                               Enumerated = 2
	TraceDepthPresentMinimumWithoutVendorSpecificExtension Enumerated = 3
	TraceDepthPresentMediumWithoutVendorSpecificExtension  Enumerated = 4
	TraceDepthPresentMaximumWithoutVendorSpecificExtension Enumerated = 5
)

type TraceDepth struct {
	Value Enumerated `vht:"valueExt,valueLB:0,valueUB:5"`
}
