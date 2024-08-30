package requests

import "encoding/xml"

type EPPLoginOptions struct {
	Version  string `xml:"version"`
	Language string `xml:"lang"`
}

type EPPLoginServices struct {
	ObjectURIs []string `xml:"objURI"`
}

type EPPLoginCommand struct {
	XMLName     xml.Name         `xml:"login"`
	ClientID    string           `xml:"clID"`
	Password    string           `xml:"pw"`
	NewPassword string           `xml:"newPW,omitempty"`
	Options     EPPLoginOptions  `xml:"options"`
	Services    EPPLoginServices `xml:"svcs"`
}

var EPPLoginRIPNOptions = EPPLoginOptions{
	Version:  "1.0",
	Language: "ru",
}

var EPPLoginRIPNServies = []string{
	"http://www.ripn.net/epp/ripn-epp-1.0",
	"http://www.ripn.net/epp/ripn-eppcom-1.0",
	"http://www.ripn.net/epp/ripn-contact-1.0",
	"http://www.ripn.net/epp/ripn-domain-1.0",
	"http://www.ripn.net/epp/ripn-host-1.0",
	"http://www.ripn.net/epp/ripn-registrar-1.0",
	"http://www.ripn.net/epp/ripn-domain-1.1",
}

func NewEPPRIPNLoginRequest(clientID, password string) EPPCommandRequest {
	return NewEPPLoginRequest(clientID, password, EPPLoginRIPNOptions, EPPLoginRIPNServies)
}

func NewEPPLoginRequest(clientID, password string, options EPPLoginOptions, services []string) EPPCommandRequest {
	return WrapEPPCommand(EPPLoginCommand{
		ClientID: clientID,
		Password: password,
		Options:  options,
		Services: EPPLoginServices{
			ObjectURIs: services,
		},
	})
}
