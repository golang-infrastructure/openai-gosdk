package openai_gosdk

import "fmt"

const createFineTuneURL = "https://api.openai.com/v1/fine-tunes"

type RequestCreateFineTune struct {
	// string Required
	// The ID of an uploaded file that contains training data.
	//
	// See upload file for how to upload a file. https://platform.openai.com/docs/api-reference/files/upload
	//
	// Your dataset must be formatted as a JSONL file, where each training example is a JSON object with the keys "prompt" and "completion". Additionally, you must upload your file with the purpose .fine-tune
	//
	// See the fine-tuning guide for more details. https://platform.openai.com/docs/guides/fine-tuning/creating-training-data
	TrainingFile *string `json:"training_file"`

	// string Optional
	// The ID of an uploaded file that contains validation data.
	//
	// If you provide this file, the data is used to generate validation metrics periodically during fine-tuning. These metrics can be viewed in the fine-tuning results file https://platform.openai.com/docs/guides/fine-tuning/analyzing-your-fine-tuned-model. Your train and validation data should be mutually exclusive.
	//
	// Your dataset must be formatted as a JSONL file, where each validation example is a JSON object with the keys "prompt" and "completion". Additionally, you must upload your file with the purpose .fine-tune
	//
	// See the fine-tuning guide for more details. https://platform.openai.com/docs/guides/fine-tuning/creating-training-data
	ValidationFile string `json:"validation_file,omitempty"`

	// string Optional Defaults to curie
	// The name of the base model to fine-tune. You can select one of "ada", "babbage", "curie", "davinci", or a fine-tuned model created after 2022-04-21. To learn more about these models, see the Models documentation.
	Model string `json:"model,omitempty"`

	// integer Optional Defaults to 4
	// The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset.
	NEpochs int `json:"n_epochs,omitempty"`

	// integer Optional Defaults to null
	// TThe batch size to use for training. The batch size is the number of training examples used to train a single forward and backward pass.
	//
	// By default, the batch size will be dynamically configured to be ~0.2% of the number of examples in the training set, capped at 256 - in general, we've found that larger batch sizes tend to work better for larger datasets.
	BatchSize float64 `json:"batch_size,omitempty"`

	// number Optional Defaults to null
	// The learning rate multiplier to use for training. The fine-tuning learning rate is the original learning rate used for pretraining multiplied by this value.
	//
	// By default, the learning rate multiplier is the 0.05, 0.1, or 0.2 depending on final batch_size (larger learning rates tend to perform better with larger batch sizes). We recommend experimenting with values in the range 0.02 to 0.2 to see what produces the best results.
	LearningRateMultiplier float64 `json:"learning_rate_multiplier,omitempty"`

	// number Optional Defaults to 0.01
	// The weight to use for loss on the prompt tokens. This controls how much the model tries to learn to generate the prompt (as compared to the completion which always has a weight of 1.0), and can add a stabilizing effect to training when completions are short.
	//
	// If prompts are extremely long (relative to completions), it may make sense to reduce this weight so as to avoid over-prioritizing learning the prompt.
	PromptLossWeight float64 `json:"prompt_loss_weight,omitempty"`

	// boolean Optional Defaults to false
	// If set, we calculate classification-specific metrics such as accuracy and F-1 score using the validation set at the end of every epoch. These metrics can be viewed in the results file.
	//
	// In order to compute classification metrics, you must provide a validation_file. Additionally, you must specify classification_n_classes for multiclass classification or classification_positive_class for binary classification. https://platform.openai.com/docs/guides/fine-tuning/analyzing-your-fine-tuned-model
	ComputeClassificationMetrics bool `json:"compute_classification_metrics,omitempty"`

	// integer Optional Defaults to null
	// The number of classes in a classification task.
	//
	// This parameter is required for multiclass classification.
	ClassificationNClasses int `json:"classification_n_classes,omitempty"`

	// string Optional Defaults to null
	// The positive class in binary classification.
	//
	// This parameter is needed to generate precision, recall, and F1 metrics when doing binary classification.
	ClassificationPositiveClass string `json:"classification_positive_class,omitempty"`

	// array Optional Defaults to null
	// If this is provided, we calculate F-beta scores at the specified beta values. The F-beta score is a generalization of F-1 score. This is only used for binary classification.
	//
	// With a beta of 1 (i.e. the F-1 score), precision and recall are given the same weight. A larger beta score puts more weight on recall and less on precision. A smaller beta score puts more weight on precision and less on recall.
	ClassificationBetas []byte `json:"classification_betas,omitempty"`

	// string Optional Defaults to null
	// A string of up to 40 characters that will be added to your fine-tuned model name.
	//
	// For example, a suffix of "custom-model-name" would produce a model name like ada:ft-your-org:custom-model-name-2022-02-15-04-21-04.
	Suffix string `json:"suffix,omitempty"`
}

type ResponseCreateFineTune struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	CreatedAt int    `json:"created_at"`
	Events    []struct {
		Object    string `json:"object"`
		CreatedAt int    `json:"created_at"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"events"`
	FineTunedModel interface{} `json:"fine_tuned_model"`
	HyperParams    struct {
		BatchSize              int     `json:"batch_size"`
		LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
		NEpochs                int     `json:"n_epochs"`
		PromptLossWeight       float64 `json:"prompt_loss_weight"`
	} `json:"hyperparams"`
	OrganizationId  string        `json:"organization_id"`
	ResultFiles     []interface{} `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []interface{} `json:"validation_files"`
	TrainingFiles   []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt int `json:"updated_at"`
}

func NewCreateFineTune(baseOpenAI BaseOpenAI) OpenAI[RequestCreateFineTune, ResponseCreateFineTune] {
	return OpenAI[RequestCreateFineTune, ResponseCreateFineTune]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  createFineTuneURL,
		Method:     POST,
	}
}

const fineTunesURL = "https://api.openai.com/v1/fine-tunes"

type RequestFineTunes struct{}

type ResponseFineTunes struct {
	Object string `json:"object"`
	Data   []struct {
		Id             string      `json:"id"`
		Object         string      `json:"object"`
		Model          string      `json:"model"`
		CreatedAt      int         `json:"created_at"`
		FineTunedModel interface{} `json:"fine_tuned_model"`
		Hyperparams    struct {
		} `json:"hyperparams"`
		OrganizationId  string        `json:"organization_id"`
		ResultFiles     []interface{} `json:"result_files"`
		Status          string        `json:"status"`
		ValidationFiles []interface{} `json:"validation_files"`
		TrainingFiles   []struct {
		} `json:"training_files"`
		UpdatedAt int `json:"updated_at"`
	} `json:"data"`
}

func NewFineTunes(baseOpenAI BaseOpenAI) OpenAI[RequestFineTunes, ResponseFineTunes] {
	return OpenAI[RequestFineTunes, ResponseFineTunes]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  fineTunesURL,
		Method:     GET,
	}
}

const retrieveFineTuneURL = "https://api.openai.com/v1/fine-tunes/"

type RequestRetrieveFineTune struct {
	// string Required
	// The ID of the fine-tune job
	FineTurnID string `json:"fine_turn_id"`
}

type ResponseRetrieveFineTune struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	CreatedAt int    `json:"created_at"`
	Events    []struct {
		Object    string `json:"object"`
		CreatedAt int    `json:"created_at"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"events"`
	FineTunedModel string `json:"fine_tuned_model"`
	Hyperparams    struct {
		BatchSize              int     `json:"batch_size"`
		LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
		NEpochs                int     `json:"n_epochs"`
		PromptLossWeight       float64 `json:"prompt_loss_weight"`
	} `json:"hyperparams"`
	OrganizationId string `json:"organization_id"`
	ResultFiles    []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []interface{} `json:"validation_files"`
	TrainingFiles   []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt int `json:"updated_at"`
}

func NewRetrieveFineTune(baseOpenAI BaseOpenAI, fineTuneId string) OpenAI[RequestRetrieveFineTune, ResponseRetrieveFineTune] {
	return OpenAI[RequestRetrieveFineTune, ResponseRetrieveFineTune]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  retrieveFineTuneURL + fineTuneId,
		Method:     GET,
	}
}

const cancelFineTuneURL = "https://api.openai.com/v1/fine-tunes/%s/cancel"

type RequestCancelFineTune struct{}

type ResponseCancelFineTune struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	CreatedAt int    `json:"created_at"`
	Events    []struct {
	} `json:"events"`
	FineTunedModel interface{} `json:"fine_tuned_model"`
	Hyperparams    struct {
	} `json:"hyperparams"`
	OrganizationId  string        `json:"organization_id"`
	ResultFiles     []interface{} `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []interface{} `json:"validation_files"`
	TrainingFiles   []struct {
		Id        string `json:"id"`
		Object    string `json:"object"`
		Bytes     int    `json:"bytes"`
		CreatedAt int    `json:"created_at"`
		Filename  string `json:"filename"`
		Purpose   string `json:"purpose"`
	} `json:"training_files"`
	UpdatedAt int `json:"updated_at"`
}

func NewCancelFineTune(baseOpenAI BaseOpenAI, fineTuneId string) OpenAI[RequestCancelFineTune, ResponseCancelFineTune] {
	return OpenAI[RequestCancelFineTune, ResponseCancelFineTune]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  fmt.Sprintf(cancelFineTuneURL, fineTuneId),
		Method:     GET,
	}
}

const listFineTuneEventsURL = "https://api.openai.com/v1/fine-tunes/%s/events"

type RequestListFineTuneEvents struct{}

type ResponseListFineTuneEvents struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string `json:"object"`
		CreatedAt int    `json:"created_at"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"data"`
}

func NewListFineTuneEvents(baseOpenAI BaseOpenAI, fineTuneId string) OpenAI[RequestListFineTuneEvents, ResponseListFineTuneEvents] {
	return OpenAI[RequestListFineTuneEvents, ResponseListFineTuneEvents]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  fmt.Sprintf(listFineTuneEventsURL, fineTuneId),
		Method:     GET,
	}
}

const deleteFineTuneModelURL = "https://api.openai.com/v1/models/"

type RequestDeleteFineTuneModel struct{}

type ResponseDeleteFineTuneModel struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func NewDeleteFineTuneModel(baseOpenAI BaseOpenAI, modelId string) OpenAI[RequestDeleteFineTuneModel, ResponseDeleteFineTuneModel] {
	return OpenAI[RequestDeleteFineTuneModel, ResponseDeleteFineTuneModel]{
		BaseOpenAI: baseOpenAI,
		TargetURL:  deleteFineTuneModelURL + modelId,
		Method:     DELETE,
	}
}
