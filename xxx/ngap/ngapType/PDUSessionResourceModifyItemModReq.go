// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            PDUSessionID
	NASPDU                                  *NASPDU `vht:"optional"`
	PDUSessionResourceModifyRequestTransfer OctetString
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs `vht:"optional"`
}
