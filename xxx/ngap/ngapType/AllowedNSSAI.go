// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct AllowedNSSAI */
/* AllowedNSSAIItem */
type AllowedNSSAI struct {
	List []AllowedNSSAIItem `vht:"valueExt,sizeLB:1,sizeUB:8"`
}
