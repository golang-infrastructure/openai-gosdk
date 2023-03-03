package openai_gosdk

const createModerationURL = "https://api.openai.com/v1/moderations"

type RequestCreateModeration struct {
	// string or array Required
	// The input text to classify
	Input StrongOrArray `json:"input"`

	// string Optional Defaults to text-moderation-latest
	// Two content moderations models are available: text-moderation-stable and text-moderation-latest.
	//
	// The default is text-moderation-latest which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use text-moderation-stable, we will provide advanced notice before updating the model. Accuracy of text-moderation-stable may be slightly lower than for text-moderation-latest.
	Model string `json:"model,omitempty"`
}

type ResponseCreateModeration struct {
	Id      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Categories struct {
			Hate            bool `json:"hate"`
			HateThreatening bool `json:"hate/threatening"`
			SelfHarm        bool `json:"self-harm"`
			Sexual          bool `json:"sexual"`
			SexualMinors    bool `json:"sexual/minors"`
			Violence        bool `json:"violence"`
			ViolenceGraphic bool `json:"violence/graphic"`
		} `json:"categories"`
		CategoryScores struct {
			Hate            float64 `json:"hate"`
			HateThreatening float64 `json:"hate/threatening"`
			SelfHarm        float64 `json:"self-harm"`
			Sexual          float64 `json:"sexual"`
			SexualMinors    float64 `json:"sexual/minors"`
			Violence        float64 `json:"violence"`
			ViolenceGraphic float64 `json:"violence/graphic"`
		} `json:"category_scores"`
		Flagged bool `json:"flagged"`
	} `json:"results"`
}

func NewCreateModeration(baseOpenAI BaseOpenAI) OpenAI[RequestCreateModeration, ResponseCreateModeration] {
	return OpenAI[RequestCreateModeration, ResponseCreateModeration]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  createModerationURL,
		Method:     POST,
	}
}
