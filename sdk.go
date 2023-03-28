package openai_gosdk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	resty "github.com/go-resty/resty/v2"
	"io"
	"os"
	"sync/atomic"
)

type Method string

var (
	GET            Method = "GET"
	POST           Method = "POST"
	DELETE       Method = "DELETE"
	OpenaiApiKeyEnv        = "OPENAI_API_KEY"
)

var (
	ErrorUnsupportedMethod = errors.New("unsupported method")
)

type config struct {
	client *resty.Client
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

type OpenAIWithStream[Request any, Response any] struct {
	BaseOpenAI
	Method    Method `json:"method"`
	TargetURL string `json:"target_url"`
}

type Stream[Response any] struct {
	isFinished atomic.Bool

	reader   *bufio.Reader
	response *resty.Response
}

func (stream *Stream[Response]) Recv() (Response, error) {

	var zeroResponse Response
	if stream.isFinished.Load() {
		return zeroResponse, io.EOF
	}

waitForData:
	line, err := stream.reader.ReadBytes('\n')
	if err != nil {
		return zeroResponse, err
	}

	var headerData = []byte("data: ")
	line = bytes.TrimSpace(line)
	if !bytes.HasPrefix(line, headerData) {
		goto waitForData
	}
	line = bytes.TrimPrefix(line, headerData)
	if string(line) == "[DONE]" {
		stream.isFinished.Store(true)
		return zeroResponse, io.EOF
	}

	err = json.Unmarshal(line, &zeroResponse)
	return zeroResponse, err
}

func (stream *Stream[Response]) Close() {
	if stream == nil {
		return
	}
	if stream.response == nil {
		return
	}
	if stream.response.RawBody() != nil {
		return
	}
	stream.response.RawBody().Close()
}

func (o OpenAIWithStream[Request, Response]) DoRequestByStream(request Request) (*Stream[Response], error) {

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	if o.client == nil {
		o.client = resty.New()
	}

	query := o.client.
		R().
		SetHeader("Authorization", "Bearer "+o.APIKey).
		SetHeader("OpenAI-Organization", o.Organization).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "text/event-stream").
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Connection", "keep-alive").
		SetDoNotParseResponse(true).
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
	if response.StatusCode() != 200 {
		var abnormalReturn AbnormalReturn = make(map[string]interface{})
		_ = json.Unmarshal(response.Body(), &abnormalReturn)
		return nil, abnormalReturn
	}
	return &Stream[Response]{
		reader:   bufio.NewReader(response.RawBody()),
		response: response,
	}, nil
}

func (o OpenAI[Request, Response]) DoRequest(request Request) (Response, error) {
	var zeroResponse Response
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return zeroResponse, err
	}

	if o.client == nil {
		o.client = resty.New()
	}

	query := o.client.
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
	if apiKey == "" {
		apiKey = os.Getenv(OpenaiApiKeyEnv)
	}
	base := BaseOpenAI{
		APIKey:       apiKey,
		Organization: organization,
	}
	for _, option := range options {
		option(&base.config)
	}
	return base
}

func SetRestyClient(client *resty.Client) Option {
	return func(c *config) {
		c.client = client
	}
}
