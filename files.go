package openai_gosdk

import "fmt"

const filesURL = "https://api.openai.com/v1/files"

type RequestFiles struct {
}

type ResponseFiles struct {
	Data []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"data"`
	Object string `json:"object"`
}

func NewFiles(baseOpenAI BaseOpenAI) OpenAI[RequestFiles, ResponseFiles] {
	return OpenAI[RequestFiles, ResponseFiles]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  filesURL,
		Method:     GET,
	}
}

const uploadFileURL = "https://api.openai.com/v1/files"

type RequestUploadFile struct {
	// string Required Name of the JSON Lines file to be uploaded.
	//
	// Name of the JSON Lines file to be uploaded.
	//
	// If the purpose is set to "fine-tune", each line is a JSON record with "prompt" and "completion" fields representing your training examples.
	File *string `json:"file"`

	// string Required
	// The intended purpose of the uploaded documents.
	//
	// Use "fine-tune" for Fine-tuning. This allows us to validate the format of the uploaded file.
	Purpose *string `json:"purpose"`
}

type ResponseUploadFile struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

func NewUploadFile(baseOpenAI BaseOpenAI) OpenAI[RequestUploadFile, ResponseUploadFile] {
	return OpenAI[RequestUploadFile, ResponseUploadFile]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  uploadFileURL,
		Method:     POST,
	}
}

const deleteFileURL = "https://api.openai.com/v1/files/"

type RequestDeleteFile struct{}

type ResponseDeleteFile struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func NewDeleteFile(baseOpenAI BaseOpenAI, fileID string) OpenAI[RequestDeleteFile, ResponseDeleteFile] {
	return OpenAI[RequestDeleteFile, ResponseDeleteFile]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  deleteFileURL + fileID,
		Method:     DELETE,
	}
}

const retrieveFileURL = "https://api.openai.com/v1/files/"

type RequestRetrieveFile struct{}

type ResponseRetrieveFile struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

func NewRetrieveFile(baseOpenAI BaseOpenAI, fileID string) OpenAI[RequestRetrieveFile, ResponseRetrieveFile] {
	return OpenAI[RequestRetrieveFile, ResponseRetrieveFile]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  retrieveFileURL + fileID,
		Method:     GET,
	}
}

const retrieveFileContentURL = "https://api.openai.com/v1/files/%s/content"

type RequestRetrieveFileContent struct{}

type ResponseRetrieveFileContent struct {
	// TODO
}

func NewRetrieveFileContent(baseOpenAI BaseOpenAI, fileID string) OpenAI[RequestRetrieveFileContent, ResponseRetrieveFileContent] {
	return OpenAI[RequestRetrieveFileContent, ResponseRetrieveFileContent]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  fmt.Sprintf(retrieveFileContentURL, fileID),
		Method:     GET,
	}
}
