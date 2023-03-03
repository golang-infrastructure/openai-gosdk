package openai_gosdk

const editsURL = "https://api.openai.com/v1/edits"

type RequestEdits struct {
	// string Required
	// ID of the model to use. You can use the text-davinci-edit-001 or code-davinci-edit-001 model with this endpoint.
	Model *string `json:"model"`

	// string Optional Defaults to ''
	// The input text to use as a starting point for the edit.
	Input string `json:"input,omitempty"`

	// string Required
	// The instruction that tells the model how to edit the prompt.
	Instruction string `json:"instruction"`

	// integer Optional Defaults to 1
	// How many edits to generate for the input and instruction.
	N int `json:"n,omitempty"`

	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	//
	// We generally recommend altering this or top_p but not both.
	Temperature int `json:"temperature,omitempty"`

	// number Optional Defaults to 1
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or temperature but not both.
	TopP int `json:"top_p,omitempty"`
}

type ResponseEdits struct {
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Text  string `json:"text"`
		Index int    `json:"index"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func NewEdits(baseOpenAI BaseOpenAI) OpenAI[RequestEdits, ResponseEdits] {
	return OpenAI[RequestEdits, ResponseEdits]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  editsURL,
		Method:     POST,
	}
}
