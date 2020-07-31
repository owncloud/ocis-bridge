package middleware

import (
	"net/http"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	acc "github.com/owncloud/ocis-accounts/pkg/proto/v0"
	"github.com/owncloud/ocis-pkg/v2/log"
	"github.com/owncloud/ocis-proxy/pkg/config"
	storepb "github.com/owncloud/ocis-store/pkg/proto/v0"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	// Logger to use for logging, must be set
	Logger log.Logger
	// TokenManagerConfig for communicating with the reva token manager
	TokenManagerConfig config.TokenManager
	// HTTPClient to use for communication with the oidc provider
	HTTPClient *http.Client
	// AccountsClient for resolving accounts
	AccountsClient acc.AccountsService
	// OIDCProviderFunc to lazily initialize a provider, must be set for the oidcProvider middleware
	OIDCProviderFunc func() (OIDCProvider, error)
	// OIDCIss is the oidc-issuer
	OIDCIss string
	// RevaGatewayClient to send requests to the reva gateway
	RevaGatewayClient gateway.GatewayAPIClient
	// Store for persisting data
	Store storepb.StoreService
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Logger provides a function to set the logger option.
func Logger(l log.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// TokenManagerConfig provides a function to set the token manger config option.
func TokenManagerConfig(cfg config.TokenManager) Option {
	return func(o *Options) {
		o.TokenManagerConfig = cfg
	}
}

// HTTPClient provides a function to set the http client config option.
func HTTPClient(c *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = c
	}
}

// AccountsClient provides a function to set the accounts client config option.
func AccountsClient(ac acc.AccountsService) Option {
	return func(o *Options) {
		o.AccountsClient = ac
	}
}

// OIDCProviderFunc provides a function to set the the oidc provider function option.
func OIDCProviderFunc(f func() (OIDCProvider, error)) Option {
	return func(o *Options) {
		o.OIDCProviderFunc = f
	}
}

// OIDCIss sets the oidc issuer url
func OIDCIss(iss string) Option {
	return func(o *Options) {
		o.OIDCIss = iss
	}
}

// RevaGatewayClient provides a function to set the the reva gateway service client option.
func RevaGatewayClient(gc gateway.GatewayAPIClient) Option {
	return func(o *Options) {
		o.RevaGatewayClient = gc
	}
}

// Store provides a function to set the store option.
func Store(sc storepb.StoreService) Option {
	return func(o *Options) {
		o.Store = sc
	}
}
