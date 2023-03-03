package openai_gosdk

const (
	createImageURL = "https://api.openai.com/v1/images/generations"
)

type RequestCreateImage struct {
	// string Required
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt *string `json:"prompt"`

	// integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int `json:"n,omitempty"`

	// string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	Size string `json:"size,omitempty"`

	// string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of url or b64_json.
	ResponseFormat string `json:"response_format,omitempty"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more. https://platform.openai.com/docs/guides/safety-best-practices
	User string `json:"user,omitempty"`
}

type ResponseCreateImage struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewCreateImage(baseOpenAI BaseOpenAI) OpenAI[RequestCreateImage, ResponseCreateImage] {
	return OpenAI[RequestCreateImage, ResponseCreateImage]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  createImageURL,
		Method:     POST,
	}
}

const (
	editImageURL = "https://api.openai.com/v1/images/edits"
)

type RequestEditImage struct {
	// string Required
	// The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.
	Image *string `json:"image"`

	// string Optional
	// An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as image.
	Mask string `json:"mask,omitempty"`

	// string Required
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt *string `json:"prompt"`

	// integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int `json:"n,omitempty"`

	// string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	Size string `json:"size,omitempty"`

	// string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of url or b64_json.
	ResponseFormat string `json:"response_format,omitempty"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more. https://platform.openai.com/docs/guides/safety-best-practices
	User string `json:"user,omitempty"`
}

type ResponseEditImage struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewEditImage(baseOpenAI BaseOpenAI) OpenAI[RequestEditImage, ResponseEditImage] {
	return OpenAI[RequestEditImage, ResponseEditImage]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  editImageURL,
		Method:     POST,
	}
}

const (
	createVariationImageURL = "https://api.openai.com/v1/images/variations"
)

type RequestCreateVariationImage struct {
	// string Required
	// The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	Image *string `json:"image"`

	// integer Optional Defaults to 1
	// The number of images to generate. Must be between 1 and 10.
	N int `json:"n,omitempty"`

	// string Optional Defaults to 1024x1024
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	Size string `json:"size,omitempty"`

	// string Optional Defaults to url
	// The format in which the generated images are returned. Must be one of url or b64_json.
	ResponseFormat string `json:"response_format,omitempty"`

	// string Optional
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more. https://platform.openai.com/docs/guides/safety-best-practices
	User string `json:"user,omitempty"`
}

type ResponseCreateVariationImage struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

func NewCreateVariationImage(baseOpenAI BaseOpenAI) OpenAI[RequestCreateVariationImage, ResponseCreateVariationImage] {
	return OpenAI[RequestCreateVariationImage, ResponseCreateVariationImage]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  createVariationImageURL,
		Method:     POST,
	}
}
