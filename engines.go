package openai_gosdk

const enginesURL = "https://api.openai.com/v1/engines"

type RequestEngines struct{}

type ResponseEngines struct {
	Data []struct {
		Id     string `json:"id"`
		Object string `json:"object"`
		Owner  string `json:"owner"`
		Ready  bool   `json:"ready"`
	} `json:"data"`
	Object string `json:"object"`
}

func NewEngines(baseOpenAI BaseOpenAI) OpenAI[RequestEngines, ResponseEngines] {
	return OpenAI[RequestEngines, ResponseEngines]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  enginesURL,
		Method:     GET,
	}
}

const retrieveEngineURL = "https://api.openai.com/v1/engines/"

type RequestRetrieveEngine struct{}

type ResponseRetrieveEngine struct {
	Id     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

func NewRetrieveEngine(baseOpenAI BaseOpenAI, engineID string) OpenAI[RequestRetrieveEngine, ResponseRetrieveEngine] {
	return OpenAI[RequestRetrieveEngine, ResponseRetrieveEngine]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  retrieveEngineURL + engineID,
		Method:     GET,
	}
}
