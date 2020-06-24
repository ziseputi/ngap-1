// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToSetupListPSReq */
/* PDUSessionResourceFailedToSetupItemPSReq */
type PDUSessionResourceFailedToSetupListPSReq struct {
	List []PDUSessionResourceFailedToSetupItemPSReq `vht:"valueExt,sizeLB:1,sizeUB:256"`
}
