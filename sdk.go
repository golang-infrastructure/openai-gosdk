package openai_gosdk

import (
	"encoding/json"
	"errors"

	resty "github.com/go-resty/resty/v2"
)

type Method string

var (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
)

var (
	ErrorUnsupportedMethod = errors.New("unsupported method")
)

type config struct {
	proxy string
}

type BaseOpenAI struct {
	APIKey       string `json:"api_key"`
	Organization string `json:"organization"`
	config
}

type OpenAI[Request, Response any] struct {
	BaseOpenAI
	Method    Method `json:"method"`
	TargetURL string `json:"target_url"`
}

type AbnormalReturn map[string]interface{}

func (a AbnormalReturn) Error() string {
	v, _ := json.Marshal(a)
	return string(v)
}

func (o OpenAI[Request, Response]) DoRequest(request Request) (Response, error) {
	var zeroResponse Response
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return zeroResponse, err
	}

	cli := resty.
		New()
	if o.proxy != "" {
		cli.SetProxy(o.proxy)
	}

	query := cli.
		R().
		SetHeader("Authorization", "Bearer "+o.APIKey).
		SetHeader("OpenAI-Organization", o.Organization).
		SetHeader("Content-Type", "application/json").
		SetBody(string(requestBytes))

	var (
		response *resty.Response
	)
	switch o.Method {
	case GET:
		response, err = query.Get(o.TargetURL)
	case POST:
		response, err = query.Post(o.TargetURL)
	case DELETE:
		response, err = query.Delete(o.TargetURL)
	default:
		err = ErrorUnsupportedMethod
	}

	if err != nil {
		return zeroResponse, err
	}

	if response.StatusCode() != 200 {
		var abnormalReturn AbnormalReturn = make(map[string]interface{})
		_ = json.Unmarshal(response.Body(), &abnormalReturn)
		return zeroResponse, abnormalReturn
	}

	err = json.Unmarshal(response.Body(), &zeroResponse)
	if err != nil {
		return zeroResponse, err
	}
	return zeroResponse, nil
}

type Option func(*config)

func NewBaseOpenAI(apiKey, organization string, options ...Option) BaseOpenAI {
	base := BaseOpenAI{
		APIKey:       apiKey,
		Organization: organization,
	}
	for _, option := range options {
		option(&base.config)
	}
	return base
}

func SetProxy(url string) Option {
	return func(c *config) {
		c.proxy = url
	}
}
