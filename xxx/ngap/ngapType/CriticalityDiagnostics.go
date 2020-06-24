// Created By HaoDHH-245789 VHT2020
package ngapType

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                                          `vht:"optional"`
	TriggeringMessage         *TriggeringMessage                                      `vht:"optional"`
	ProcedureCriticality      *Criticality                                            `vht:"optional"`
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList                           `vht:"optional"`
	IEExtensions              *ProtocolExtensionContainerCriticalityDiagnosticsExtIEs `vht:"optional"`
}
