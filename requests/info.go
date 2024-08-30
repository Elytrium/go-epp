package requests

import (
	"encoding/xml"

	"github.com/Elytrium/go-epp/types"
)

type EPPInfoRequest struct {
	XMLName xml.Name `xml:"info"`
	Body    interface{}
}

type EPPInfoResponse struct {
	InfoContactResponse   *EPPInfoContactResponse   `xml:",omitempty"`
	InfoHostResponse      *EPPInfoHostResponse      `xml:",omitempty"`
	InfoDomainResponse    *EPPInfoDomainResponse    `xml:",omitempty"`
	InfoRegistrarResponse *EPPInfoRegistrarResponse `xml:",omitempty"`
}

type EPPInfoContactRequest struct {
	XMLName   xml.Name                  `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:info"`
	ContactID string                    `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:id"`
	AuthInfo  *types.EPPContactAuthInfo `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:authInfo,omitempty"`
}

type EPPInfoContactResponse struct {
	XMLName     xml.Name `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:infData"`
	ContactID   string   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:id"`
	ContactROID string   `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:roid"`
	Status      []types.EPPContactStatus
	types.EPPContactData
	types.EPPContactExtendedData
}

type EPPInfoHostRequest struct {
	XMLName  xml.Name `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:info"`
	HostName string   `xml:"name"`
}

type EPPInfoHostResponse struct {
	XMLName  xml.Name `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:infData"`
	HostName string   `xml:"name"`
	HostROID string   `xml:"roid"`
	Status   []types.EPPHostStatus
	types.EPPHostData
	types.EPPHostExtendedData
}

type EPPInfoDomainRequest struct {
	XMLName    xml.Name                 `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:info"`
	DomainName string                   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
	AuthInfo   *types.EPPDomainAuthInfo `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:authInfo,omitempty"`
}

type EPPInfoDomainResponse struct {
	XMLName    xml.Name `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:infData"`
	DomainName string   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
	DomainROID string   `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:roid"`
	Status     []types.EPPDomainStatus
	types.EPPDomainData
	types.EPPDomainExtendedData
}

type EPPInfoRegistrarRequest struct {
	XMLName     xml.Name `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:info"`
	RegistrarID string   `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:id"`
}

type EPPInfoRegistrarResponse struct {
	XMLName     xml.Name `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:infData"`
	RegistrarID string   `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:id"`
	Status      []types.EPPRegistrarStatus
	types.EPPRegistrarData
	types.EPPRegistrarExtendedData
}

func NewRIPNEPPInfoContactRequest(contactID, authinfo string) EPPCommandRequest {
	return WrapEPPCommand(EPPInfoRequest{
		Body: EPPInfoContactRequest{
			ContactID: contactID,
			AuthInfo: &types.EPPContactAuthInfo{
				Password: authinfo,
			},
		},
	})
}

func NewRIPNEPPInfoHostRequest(hostName string) EPPCommandRequest {
	return WrapEPPCommand(EPPInfoRequest{
		Body: EPPInfoHostRequest{
			HostName: hostName,
		},
	})
}

func NewRIPNEPPInfoDomainRequest(domainName, authinfo string) EPPCommandRequest {
	return WrapEPPCommand(EPPInfoRequest{
		Body: EPPInfoDomainRequest{
			DomainName: domainName,
			AuthInfo: &types.EPPDomainAuthInfo{
				Password: authinfo,
			},
		},
	})
}

func NewRIPNEPPInfoRegistrarRequest(registrarID string) EPPCommandRequest {
	return WrapEPPCommand(EPPInfoRequest{
		Body: EPPInfoRegistrarRequest{
			RegistrarID: registrarID,
		},
	})
}
