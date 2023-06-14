package openai_gosdk

const imageGenerationsURL = "https://api.openai.com/v1/images/generations"

type RequestImageGenerations struct {
	// string Required
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt *string `json:"prompt"`

	// integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int `json:"n,omitempty"`

	// string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of , , or .256x256 512x512 1024x1024
	Size string `json:"size,omitempty"`

	// string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of or .url b64_json
	ResponseFormat string `json:"response_format,omitempty"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids.
	User string `json:"user,omitempty"`
}

type ResponseImageGenerations struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewImageGenerations(baseOpenAI BaseOpenAI) OpenAI[RequestImageGenerations, ResponseImageGenerations] {
	return OpenAI[RequestImageGenerations, ResponseImageGenerations]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  imageGenerationsURL,
		Method:     POST,
	}
}

const imageEditURL = "https://api.openai.com/v1/images/edits"

type RequestImageEdit struct {
	// string Required
	// The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.
	Image *string `json:"image"`

	// string Optional
	// An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as . image image
	Mask string `json:"mask,omitempty"`

	// string Required
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt *string `json:"prompt"`

	// integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int `json:"n,omitempty"`

	// string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of , , or .256x256 512x512 1024x1024
	Size string `json:"size,omitempty"`

	// string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of or .url b64_json
	ResponseFormat string `json:"response_format,omitempty"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids.
	User string `json:"user,omitempty"`
}

type ResponseImageEdit struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewImageEdit(baseOpenAI BaseOpenAI) OpenAI[RequestImageEdit, ResponseImageEdit] {
	return OpenAI[RequestImageEdit, ResponseImageEdit]{
		BaseOpenAI: baseOpenAI,
		Method:     POST,
		TargetURL:  imageEditURL,
	}
}

const imageVariationURL = "https://api.openai.com/v1/images/variations"

type RequestImageVariation struct {
	// string Required
	// The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	Image *string `json:"image"`

	// integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int `json:"n,omitempty"`

	// string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of , , or .256x256 512x512 1024x1024
	Size string `json:"size,omitempty"`

	// string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of or .url b64_json
	ResponseFormat string `json:"response_format,omitempty"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids.
	User string `json:"user,omitempty"`
}

type ResponseImageVariation struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewImageVariation(baseOpenAI BaseOpenAI) OpenAI[RequestImageVariation, ResponseImageVariation] {
	return OpenAI[RequestImageVariation, ResponseImageVariation]{
		BaseOpenAI: baseOpenAI,
		Method:     POST,
		TargetURL:  imageVariationURL,
	}
}
