package requests

import (
	"encoding/xml"
)

type EPPDeleteRequest struct {
	XMLName xml.Name `xml:"delete"`
	Body    interface{}
}

type EPPDeleteContactRequest struct {
	XMLName   xml.Name `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:delete"`
	ContactID string   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:id"`
}

type EPPDeleteHostRequest struct {
	XMLName  xml.Name `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:delete"`
	HostName string   `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:name"`
}

type EPPDeleteDomainRequest struct {
	XMLName    xml.Name `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:delete"`
	DomainName string   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
}

func NewRIPNEPPDeleteContactRequest(contactID string) EPPCommandRequest {
	return WrapEPPCommand(EPPDeleteRequest{
		Body: EPPDeleteContactRequest{
			ContactID: contactID,
		},
	})
}

func NewRIPNEPPDeleteHostRequest(hostName string) EPPCommandRequest {
	return WrapEPPCommand(EPPDeleteRequest{
		Body: EPPDeleteHostRequest{
			HostName: hostName,
		},
	})
}

func NewRIPNEPPDeleteDomainRequest(domainName string) EPPCommandRequest {
	return WrapEPPCommand(EPPDeleteRequest{
		Body: EPPDeleteDomainRequest{
			DomainName: domainName,
		},
	})
}
