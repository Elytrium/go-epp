package eppclient

import (
	"context"
	"net/http"

	"github.com/Elytrium/go-epp/requests"
	"github.com/Elytrium/go-epp/types"

	"github.com/imroc/req/v3"
	"github.com/nbio/xml"
)

type HTTPEPPSession struct {
	Client *req.Client
}

func NewHTTPSession(serverURL string) *HTTPEPPSession {
	return &HTTPEPPSession{
		Client: req.C().
			SetUserAgent("EPP Client /1.0").
			SetBaseURL(serverURL).
			SetXmlUnmarshal(xml.Unmarshal).
			SetXmlMarshal(func(v interface{}) ([]byte, error) {
				xmlBytes, err := xml.MarshalIndent(v, "", "  ")
				if err != nil {
					return nil, err
				}

				return []byte(xml.Header + string(xmlBytes)), nil
			}),
	}
}

func (s *HTTPEPPSession) WithKnownSessionID(sessionID string) *HTTPEPPSession {
	s.Client.Cookies = append(s.Client.Cookies, &http.Cookie{
		Name:  "EPPSESSIONID",
		Value: sessionID,
	})

	return s
}

func (s *HTTPEPPSession) WithClientCertificates(certPath, certKeyPath string) *HTTPEPPSession {
	s.Client = s.Client.
		SetCertFromFile(certPath, certKeyPath)
	return s
}

func (s *HTTPEPPSession) Do(request *requests.EPPContainer) (requests.EPPContainer, error) {
	response := requests.EPPContainer{}
	err := s.Client.Post().
		SetBodyXmlMarshal(request).
		Do(context.Background()).
		UnmarshalXml(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *HTTPEPPSession) DoCommand(request *requests.EPPCommandRequest) (*requests.EPPCommandResponse, error) {
	c, err := s.Do(&requests.EPPContainer{CommandRequest: request})
	return c.CommandResponse, err
}

func (s *HTTPEPPSession) Login(clientID, password string) (*requests.EPPCommandResponse, error) {
	request := requests.NewEPPRIPNLoginRequest(clientID, password)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) Logout() (*requests.EPPCommandResponse, error) {
	request := requests.NewEPPLogoutRequest()
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) Hello() (requests.EPPHelloResponse, error) {
	response, err := s.Do(&requests.EPPContainer{HelloRequest: &requests.EPPHelloRequest{}})
	return *response.HelloResponse, err
}

func (s *HTTPEPPSession) Poll() (*requests.EPPCommandResponse, *requests.EPPPollResponse, error) {
	request := requests.NewEPPPollRequest()
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, (*requests.EPPPollResponse)(response.Data.TransferResponse), err
}

func (s *HTTPEPPSession) PollAcknowledgement(id string) (*requests.EPPCommandResponse, error) {
	request := requests.NewEPPPollAcknowledgement(id)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) CreateOrganization(contactID string, organization types.EPPContactOrganization, verified bool, oidEsia string) (*requests.EPPCommandResponse, *requests.EPPCreateContactResponse, error) {
	request := requests.NewRIPNEPPCreateOrganizationRequest(contactID, organization, verified, oidEsia)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.CreateContactResponse, err
}

func (s *HTTPEPPSession) CreatePerson(contactID string, person types.EPPContactPerson, verified bool, oidEsia string) (*requests.EPPCommandResponse, *requests.EPPCreateContactResponse, error) {
	request := requests.NewRIPNEPPCreatePersonRequest(contactID, person, verified, oidEsia)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.CreateContactResponse, err
}

func (s *HTTPEPPSession) CreateHost(hostName string, addresses []types.EPPHostAddr) (*requests.EPPCommandResponse, *requests.EPPCreateHostResponse, error) {
	hostName, err := wrapDomainName(hostName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPCreateHostRequest(hostName, addresses)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.CreateHostResponse, err
}

func (s *HTTPEPPSession) CreateDomain(domainName string, period types.EPPDomainPeriod, registrant string,
	description []string, authInfo *types.EPPDomainAuthInfo) (*requests.EPPCommandResponse, *requests.EPPCreateDomainResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPCreateDomainRequest(domainName, period, registrant, description, authInfo)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.CreateDomainResponse, err
}

func (s *HTTPEPPSession) DeleteContact(contactID string) (*requests.EPPCommandResponse, error) {
	request := requests.NewRIPNEPPDeleteContactRequest(contactID)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) DeleteHost(hostName string) (*requests.EPPCommandResponse, error) {
	hostName, err := wrapDomainName(hostName)
	if err != nil {
		return nil, err
	}

	request := requests.NewRIPNEPPDeleteHostRequest(hostName)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) DeleteDomain(domainName string) (*requests.EPPCommandResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, err
	}

	request := requests.NewRIPNEPPDeleteDomainRequest(domainName)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) GetContactInfo(contactID, authinfo string) (*requests.EPPCommandResponse, *requests.EPPInfoContactResponse, error) {
	request := requests.NewRIPNEPPInfoContactRequest(contactID, authinfo)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.InfoContactResponse, err
}

func (s *HTTPEPPSession) GetHostInfo(hostName string) (*requests.EPPCommandResponse, *requests.EPPInfoHostResponse, error) {
	hostName, err := wrapDomainName(hostName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPInfoHostRequest(hostName)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.InfoHostResponse, err
}

func (s *HTTPEPPSession) GetDomainInfo(domainName, authinfo string) (*requests.EPPCommandResponse, *requests.EPPInfoDomainResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPInfoDomainRequest(domainName, authinfo)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.InfoDomainResponse, err
}

func (s *HTTPEPPSession) GetRegistrarInfo(registrarID string) (*requests.EPPCommandResponse, *requests.EPPInfoRegistrarResponse, error) {
	request := requests.NewRIPNEPPInfoRegistrarRequest(registrarID)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.InfoRegistrarResponse, err
}

func (s *HTTPEPPSession) RenewDomain(domainName string, curExpDate string, period types.EPPDomainPeriod) (*requests.EPPCommandResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, err
	}

	request := requests.NewRIPNEPPRenewDomainRequest(domainName, curExpDate, period)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) TransferRequest(domainName, authInfo string) (*requests.EPPCommandResponse, *requests.EPPTransferResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPTransferRequest(domainName, authInfo)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.TransferResponse, err
}

func (s *HTTPEPPSession) TransferApprove(domainName string) (*requests.EPPCommandResponse, *requests.EPPTransferResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPTransferApprove(domainName)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.TransferResponse, err
}

func (s *HTTPEPPSession) TransferReject(domainName string) (*requests.EPPCommandResponse, *requests.EPPTransferResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPTransferReject(domainName)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.TransferResponse, err
}

func (s *HTTPEPPSession) TransferCancel(domainName, authInfo string) (*requests.EPPCommandResponse, *requests.EPPTransferResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, nil, err
	}

	request := requests.NewRIPNEPPTransferCancel(domainName, authInfo)
	response, err := s.DoCommand(&request)
	if response.Data == nil {
		return response, nil, err
	}
	return response, response.Data.TransferResponse, err
}

func (s *HTTPEPPSession) UpdateContact(contactID string, change *types.EPPContactData) (*requests.EPPCommandResponse, error) {
	request := requests.NewRIPNEPPUpdateContactRequest(contactID, change)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) UpdateHost(hostName string, change *types.EPPHostData) (*requests.EPPCommandResponse, error) {
	hostName, err := wrapDomainName(hostName)
	if err != nil {
		return nil, err
	}

	request := requests.NewRIPNEPPUpdateHostRequest(hostName, change)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) UpdateDomain(domainName string, change, add, remove *types.EPPDomainData) (*requests.EPPCommandResponse, error) {
	domainName, err := wrapDomainName(domainName)
	if err != nil {
		return nil, err
	}

	err = wrapDomainData(change)
	if err != nil {
		return nil, err
	}
	err = wrapDomainData(add)
	if err != nil {
		return nil, err
	}
	err = wrapDomainData(remove)
	if err != nil {
		return nil, err
	}

	request := requests.NewRIPNEPPUpdateDomainRequest(domainName, change, add, remove)
	response, err := s.DoCommand(&request)
	return response, err
}

func (s *HTTPEPPSession) UpdateRegistrar(registrarID string, change, add, remove *types.EPPRegistrarData) (*requests.EPPCommandResponse, error) {
	request := requests.NewRIPNEPPUpdateRegistrarRequest(registrarID, change, add, remove)
	response, err := s.DoCommand(&request)
	return response, err
}
