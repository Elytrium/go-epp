# Go EPP

(RIPN-) EPP protocol Golang implementation.

## Usage 

```go
package main

import (
	"github.com/Elytrium/go-epp/eppclient"
	"github.com/Elytrium/go-epp/types"
)

func main() {
	session := eppclient.NewHTTPSession("https://test.ru/").
		WithClientCertificates("ru.pem", "rukey.pem") //.
		//WithKnownSessionID("1234567890")

	session.Login("ELYTRIUM-RU", "password")

	session.CreatePerson("testikov-1", types.EPPContactPerson{
		InternationalPostalInfo: &types.EPPContactPersonAddress{
			Name:    "Testov Test Testikovich",
			Address: []string{"Moskovskaya oblast, 123"},
		},
		LocalPostalInfo: &types.EPPContactPersonAddress{
			Name:    "Testov Test Testikovich",
			Address: []string{"Московская область, 123"},
		},
		Birthday: "2024-08-28",
		Passport: []string{"1234567890 TEST"},
		Voice:    []string{"+12345"},
		Email:    []string{"mail@example.ru"},
	}, true, "")

	session.CreateDomain("example.ru", types.EPPDomainPeriod{Period: 1, Unit: types.EPPDomainPeriodUnitYear}, "testikov-1", nil, nil)

	session.CreateHost("ns1.пример.рф", nil)
	session.CreateHost("ns2.пример.рф", nil)
	session.UpdateDomain("example.ru", nil, &types.EPPDomainData{
		NS: &types.EPPDomainNS{
			HostObjects: []string{"ns1.пример.рф", "ns2.пример.рф"},
		},
	}, nil)

	session.UpdateRegistrar("ELYTRIUM-RU", &types.EPPRegistrarData{
		WWW:   []string{"https://elytrium.ru"},
		Whois: []string{"whois.elytrium.ru"},
		Voice: []string{"+74993228228"},
		Fax:   []string{"+74993228228"},
	}, nil, nil)

	session.UpdateRegistrar("ELYTRIUM-RU", &types.EPPRegistrarData{
		AdminContact: []string{"https://whois.elytrium.ru/%domain_name%"},
	}, nil, nil)
}
```