package openai_gosdk

const createTranscriptionURL = "https://api.openai.com/v1/audio/transcriptions"

type RequestTranscription struct {
	// string Required
	// The audio file to translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	File *string `json:"file"`

	// string Required
	// ID of the model to use. Only whisper-1 is currently available.
	Model *string `json:"model"`

	// string Optional
	// An optional text to guide the model's style or continue a previous audio segment. The prompt should match the audio language.. https://platform.openai.com/docs/guides/speech-to-text/prompting
	Prompt string `json:"prompt,omitempty"`

	// string Optional Defaults to json
	// The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	ResponseFormat string `json:"response_format,omitempty"`

	// number Optional Defaults to 0
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.
	Temperature int `json:"temperature,omitempty"`

	// string Optional
	// The language of the input audio. Supplying the input language in ISO-639-1 format will improve accuracy and latency.
	Language string `json:"language,omitempty"`
}

type ResponseTranscription struct {
	Text string `json:"text"`
}

func NewTranscription(baseOpenAI BaseOpenAI, request RequestTranscription) OpenAI[RequestTranscription, ResponseTranscription] {
	return OpenAI[RequestTranscription, ResponseTranscription]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  createTranscriptionURL,
		Method:     POST,
	}
}

const createTranslationURL = "https://api.openai.com/v1/audio/translations"

type RequestTranslation struct {
	// string Required
	// The audio file to translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	File *string `json:"file"`

	// string Required
	// ID of the model to use. Only is currently available.whisper-1
	Model *string `json:"model"`

	// string Optional
	// An optional text to guide the model's style or continue a previous audio segment. The prompt should be in English. https://platform.openai.com/docs/guides/speech-to-text/prompting
	Prompt string `json:"prompt,omitempty"`

	// string Optional Defaults to json
	// The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	ResponseFormat string `json:"response_format,omitempty"`

	// number Optional Defaults to 0
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.
	Temperature int `json:"temperature,omitempty"`
}

type ResponseTranslation struct {
	Text string `json:"text"`
}

func NewTranslation(baseOpenAI BaseOpenAI) OpenAI[RequestTranslation, ResponseTranslation] {
	return OpenAI[RequestTranslation, ResponseTranslation]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  createTranslationURL,
		Method:     POST,
	}
}
