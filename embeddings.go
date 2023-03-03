package openai_gosdk

const embeddingsURL = "https://api.openai.com/v1/embeddings"

type RequestEmbeddings struct {
	// string Required
	// ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them. https://platform.openai.com/docs/api-reference/models/list
	Model *string `json:"model"`

	// string or array Required
	// Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings for multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed 8192 tokens in length.
	Input StrongOrArray `json:"input"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more. https://platform.openai.com/docs/guides/safety-best-practices https://platform.openai.com/docs/api-reference/models/list
	User string `json:"user,omitempty"`
}

type ResponseEmbeddings struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func NewEmbeddings(baseOpenAI BaseOpenAI) OpenAI[RequestEmbeddings, ResponseEmbeddings] {
	return OpenAI[RequestEmbeddings, ResponseEmbeddings]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  embeddingsURL,
		Method:     POST,
	}
}
