// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct CriticalityDiagnostics_IE_List */
/* CriticalityDiagnosticsIEItem */
type CriticalityDiagnosticsIEList struct {
	List []CriticalityDiagnosticsIEItem `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
