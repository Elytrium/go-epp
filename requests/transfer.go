package requests

import (
	"encoding/xml"
	"time"

	"github.com/Elytrium/go-epp/types"
)

type EPPTransferRequest struct {
	XMLName   xml.Name                   `xml:"transfer"`
	Operation types.EPPTransferOperation `xml:"op,attr"`
	Body      interface{}
}

type EPPTransferResponse struct {
	XMLName    xml.Name                `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:trnData"`
	DomainName string                  `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:name"`
	Status     types.EPPTransferStatus `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:trStatus"`
	ResultID   string                  `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:reID"`
	ResultDate time.Time               `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:reDate"`
	ActionID   string                  `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:acID"`
	ActionDate time.Time               `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:acDate"`
	ExpiryDate time.Time               `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:exDate,omitempty"`
}

func NewRIPNEPPTransferDomainInfoWithAuth(domainName, authInfo string) types.EPPDomainTransferInfo {
	return types.EPPDomainTransferInfo{
		DomainName: domainName,
		AuthInfo: &types.EPPDomainTransferAuthInfo{
			Password: authInfo,
		},
	}
}

func NewRIPNEPPTransferDomainInfo(domainName string) types.EPPDomainTransferInfo {
	return types.EPPDomainTransferInfo{
		DomainName: domainName,
	}
}

func NewRIPNEPPTransferRequest(domainName, authInfo string) EPPCommandRequest {
	return WrapEPPCommand(EPPTransferRequest{
		Operation: types.EPPTransferOperationRequest,
		Body:      NewRIPNEPPTransferDomainInfoWithAuth(domainName, authInfo),
	})
}

func NewRIPNEPPTransferReject(domainName string) EPPCommandRequest {
	return WrapEPPCommand(EPPTransferRequest{
		Operation: types.EPPTransferOperationReject,
		Body:      NewRIPNEPPTransferDomainInfo(domainName),
	})
}

func NewRIPNEPPTransferApprove(domainName string) EPPCommandRequest {
	return WrapEPPCommand(EPPTransferRequest{
		Operation: types.EPPTransferOperationApprove,
		Body:      NewRIPNEPPTransferDomainInfo(domainName),
	})
}

func NewRIPNEPPTransferCancel(domainName, authInfo string) EPPCommandRequest {
	return WrapEPPCommand(EPPTransferRequest{
		Operation: types.EPPTransferOperationCancel,
		Body:      NewRIPNEPPTransferDomainInfoWithAuth(domainName, authInfo),
	})
}
