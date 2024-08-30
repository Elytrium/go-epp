package requests

import (
	"encoding/xml"

	"github.com/Elytrium/go-epp/types"
)

type EPPUpdateRequest struct {
	XMLName xml.Name `xml:"update"`
	Body    interface{}
}

type EPPUpdateContactRequest struct {
	XMLName   xml.Name              `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:update"`
	ContactID string                `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:id"`
	Change    *types.EPPContactData `xml:"http://www.ripn.net/epp/ripn-contact-1.0 contact:chg"`
}

type EPPUpdateHostRequest struct {
	XMLName  xml.Name           `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:update"`
	HostName string             `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:name"`
	Change   *types.EPPHostData `xml:"http://www.ripn.net/epp/ripn-host-1.0 host:chg"`
}

type EPPUpdateDomainRequest struct {
	XMLName    xml.Name             `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:update"`
	DomainName string               `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:name"`
	Change     *types.EPPDomainData `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:chg"`
	Add        *types.EPPDomainData `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:add"`
	Remove     *types.EPPDomainData `xml:"http://www.ripn.net/epp/ripn-domain-1.0 domain:rem"`
}

type EPPUpdateRegistrarRequest struct {
	XMLName     xml.Name                `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:update"`
	RegistrarID string                  `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:id"`
	Change      *types.EPPRegistrarData `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:chg"`
	Add         *types.EPPRegistrarData `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:add"`
	Remove      *types.EPPRegistrarData `xml:"http://www.ripn.net/epp/ripn-registrar-1.0 registrar:rem"`
}

func NewRIPNEPPUpdateContactRequest(contactID string, change *types.EPPContactData) EPPCommandRequest {
	return WrapEPPCommand(EPPUpdateRequest{
		Body: EPPUpdateContactRequest{
			ContactID: contactID,
			Change:    change,
		},
	})
}

func NewRIPNEPPUpdateHostRequest(hostName string, change *types.EPPHostData) EPPCommandRequest {
	return WrapEPPCommand(EPPUpdateRequest{
		Body: EPPUpdateHostRequest{
			HostName: hostName,
			Change:   change,
		},
	})
}

func NewRIPNEPPUpdateDomainRequest(domainName string, change, add, remove *types.EPPDomainData) EPPCommandRequest {
	return WrapEPPCommand(EPPUpdateRequest{
		Body: EPPUpdateDomainRequest{
			DomainName: domainName,
			Change:     change,
			Add:        add,
			Remove:     remove,
		},
	})
}

func NewRIPNEPPUpdateRegistrarRequest(registrarID string, change, add, remove *types.EPPRegistrarData) EPPCommandRequest {
	return WrapEPPCommand(EPPUpdateRequest{
		Body: EPPUpdateRegistrarRequest{
			RegistrarID: registrarID,
			Change:      change,
			Add:         add,
			Remove:      remove,
		},
	})
}
