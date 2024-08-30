package types

import "encoding/xml"

type EPPDomainTransferAuthInfo struct {
	Password string `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:pw"`
}

type EPPDomainAuthInfo struct {
	Password string `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:pw"`
}

type EPPContactAuthInfo struct {
	Password string `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:pw"`
}

type EPPTransferOperation string

const (
	EPPTransferOperationRequest EPPTransferOperation = "request"
	EPPTransferOperationApprove EPPTransferOperation = "approve"
	EPPTransferOperationReject  EPPTransferOperation = "reject"
	EPPTransferOperationCancel  EPPTransferOperation = "cancel"
)

type EPPTransferStatus string

const (
	EPPTransferStatusClientApproved  EPPTransferStatus = "clientApproved"
	EPPTransferStatusClientCancelled EPPTransferStatus = "clientCancelled"
	EPPTransferStatusClientRejected  EPPTransferStatus = "clientRejected"
	EPPTransferStatusPending         EPPTransferStatus = "pending"
	EPPTransferStatusServerApproved  EPPTransferStatus = "serverApproved"
	EPPTransferStatusServerCancelled EPPTransferStatus = "serverCancelled"
)

type EPPDomainTransferInfo struct {
	XMLName    xml.Name                   `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:transfer"`
	DomainName string                     `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:name"`
	AuthInfo   *EPPDomainTransferAuthInfo `xml:"http://www.ripn.net/epp/ripn-domain-1.1 domain:authInfo,omitempty"`
}
