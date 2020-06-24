// Created By HaoDHH-245789 VHT2020
package ngapType

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs `vht:"optional"`
}
