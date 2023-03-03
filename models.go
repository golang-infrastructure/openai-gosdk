package openai_gosdk

const modelsUrl = "https://api.openai.com/v1/models"

type Model interface {
	ModelName() *string
}

type RequestModels struct {
}

type ResponseModels struct {
	Data []struct {
		Id         string        `json:"id"`
		Object     string        `json:"object"`
		OwnedBy    string        `json:"owned_by"`
		Permission []interface{} `json:"permission"`
	} `json:"data"`
	Object string `json:"object"`
}

func NewModels(baseOpenAI BaseOpenAI) OpenAI[RequestModels, ResponseModels] {
	return OpenAI[RequestModels, ResponseModels]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  modelsUrl,
		Method:     GET,
	}
}

type RequestModel struct {
}

type ResponseModel struct {
	Id         string        `json:"id"`
	Object     string        `json:"object"`
	OwnedBy    string        `json:"owned_by"`
	Permission []interface{} `json:"permission"`
}

func NewModelList(baseOpenAI BaseOpenAI, modelID string) OpenAI[RequestModel, ResponseModel] {
	return OpenAI[RequestModel, ResponseModel]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  modelsUrl + "/" + modelID,
		Method:     GET,
	}
}
