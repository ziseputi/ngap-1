// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct AMF_TNLAssociationToUpdateList */
/* AMFTNLAssociationToUpdateItem */
type AMFTNLAssociationToUpdateList struct {
	List []AMFTNLAssociationToUpdateItem `vht:"valueExt,sizeLB:1,sizeUB:32"`
}
