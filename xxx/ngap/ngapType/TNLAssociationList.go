// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct TNLAssociationList */
/* TNLAssociationItem */
type TNLAssociationList struct {
	List []TNLAssociationItem `vht:"valueExt,sizeLB:1,sizeUB:32"`
}
