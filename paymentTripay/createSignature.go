package paymenttripay

import "github.com/valyala/fasthttp"

type Environment int

const (
	Development Environment = iota
	Production
)

type Tripay struct {
	ApiKey     []byte
	PrivateKey string

	/*Environment to decide what this program running on

	use Development var for sandbox
	use Production var for Production node
	*/
	f            *fasthttp.Client
	Host         string
	MerchantCode string
	MerchantName string
}

func New(ApiKey, PrivateKey, MerchantCode string, environment Environment) *Tripay {
	host := ""
	switch environment {
	case Development:
		host = "https://tripay.co.id/api-sandbox"
	case Production:
		host = "https://tripay.co.id/api"
	}
	return &Tripay{ApiKey: []byte("Bearer " + ApiKey), PrivateKey: PrivateKey, Host: host, f: &fasthttp.Client{}, MerchantCode: MerchantCode}
}

/*SetHttpClient used for set fasthttp.Client
Default fasthttp is used if no set
*/
func (t *Tripay) SetHttpClient(f *fasthttp.Client) {
	t.f = f
}
