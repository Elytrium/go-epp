package eppclient

import (
	"strings"

	"github.com/Elytrium/go-epp/requests"
	"github.com/Elytrium/go-epp/types"
	"golang.org/x/net/idna"
)

var idnaProfile = idna.New()

type EPPSession interface {
	DoEPPContainer(request, response *requests.EPPContainer) error
	Do(request, response interface{}) error
}

func wrapDomainName(domainName string) (string, error) {
	return idnaProfile.ToASCII(strings.ToLower(domainName))
}

func wrapDomainData(domainData *types.EPPDomainData) error {
	if domainData == nil {
		return nil
	}

	if domainData.NS == nil {
		return nil
	}

	for i, o := range domainData.NS.HostObjects {
		str, err := wrapDomainName(o)
		if err != nil {
			return err
		}

		domainData.NS.HostObjects[i] = str
	}

	return nil
}
