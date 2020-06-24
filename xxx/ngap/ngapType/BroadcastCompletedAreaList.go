// Created By HaoDHH-245789 VHT2020
package ngapType

const (
	BroadcastCompletedAreaListPresentNothing int = iota /* No components present */
	BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA
	BroadcastCompletedAreaListPresentTAIBroadcastEUTRA
	BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA
	BroadcastCompletedAreaListPresentCellIDBroadcastNR
	BroadcastCompletedAreaListPresentTAIBroadcastNR
	BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR
	BroadcastCompletedAreaListPresentChoiceExtensions
)

type BroadcastCompletedAreaList struct {
	Present                       int
	CellIDBroadcastEUTRA          *CellIDBroadcastEUTRA
	TAIBroadcastEUTRA             *TAIBroadcastEUTRA
	EmergencyAreaIDBroadcastEUTRA *EmergencyAreaIDBroadcastEUTRA
	CellIDBroadcastNR             *CellIDBroadcastNR
	TAIBroadcastNR                *TAIBroadcastNR
	EmergencyAreaIDBroadcastNR    *EmergencyAreaIDBroadcastNR
	ChoiceExtensions              *ProtocolIESingleContainerBroadcastCompletedAreaListExtIEs
}
