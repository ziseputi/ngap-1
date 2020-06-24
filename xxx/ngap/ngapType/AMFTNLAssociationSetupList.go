// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct AMF_TNLAssociationSetupList */
/* AMFTNLAssociationSetupItem */
type AMFTNLAssociationSetupList struct {
	List []AMFTNLAssociationSetupItem `vht:"valueExt,sizeLB:1,sizeUB:32"`
}
