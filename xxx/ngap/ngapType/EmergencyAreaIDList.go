// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct EmergencyAreaIDList */
/* EmergencyAreaID */
type EmergencyAreaIDList struct {
	List []EmergencyAreaID `vht:"sizeLB:1,sizeUB:65535"`
}
