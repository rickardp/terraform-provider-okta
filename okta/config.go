package okta

import (
	"fmt"

	"github.com/articulate/oktasdk-go/okta"
)

// Config is a struct containing our provider schema values
// plus the okta client object
type Config struct {
	orgName  string
	domain   string
	apiToken string

	oktaClient  *okta.Client
}

func (c *Config) loadAndValidate() error {

	client, err := okta.NewClientWithDomain(nil, c.orgName, c.domain, c.apiToken)
	if err != nil {
		return err
	}

	// quick test of our credentials by listing our authorization server(s)
	url := fmt.Sprintf("authorizationServers")
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req, nil)
	if err != nil {
		return err
	}

	// add our client object to Config
	c.oktaClient = client
	return nil
}
