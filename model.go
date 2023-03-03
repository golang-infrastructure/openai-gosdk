package openai_gosdk

type ModelName struct {
	Name string `json:"name"`
}

func (m ModelName) ModelName() *string {
	return &m.Name
}

/*
GPT-3.5
GPT-3.5 models can understand and generate natural language or code. Our most capable and cost effective model is gpt-3.5-turbo which is optimized for chat but works well for traditional completions tasks as well.

LATEST MODEL	DESCRIPTION	MAX REQUEST	TRAINING DATA
gpt-3.5-turbo	Most capable GPT-3.5 model and optimized for chat at 1/10th the cost of text-davinci-003. Will be updated with our latest model iteration.	4,096 tokens	Up to Sep 2021
gpt-3.5-turbo-0301	Snapshot of gpt-3.5-turbo from March 1st 2023. Unlike gpt-3.5-turbo, this model will not receive updates, and will only be supported for a three month period ending on June 1st 2023.	4,096 tokens	Up to Sep 2021
text-davinci-003	Can do any language task with better quality, longer output, and consistent instruction-following than the curie, babbage, or ada models. Also supports inserting completions within text.	4,000 tokens	Up to Jun 2021
text-davinci-002	Similar capabilities to text-davinci-003 but trained with supervised fine-tuning instead of reinforcement learning	4,000 tokens	Up to Jun 2021
code-davinci-002	Optimized for code-completion tasks	4,000 tokens	Up to Jun 2021
We recommend using gpt-3.5-turbo while experimenting since it will yield the best results. Once you’ve got things working, we encourage trying the other models to see if you can get the same results with lower latency or cost.
*/

var (
	GPT3p5Turbo     = ModelName{Name: "gpt-3.5-turbo"}
	GPT3p5Turbo0301 = ModelName{Name: "gpt-3.5-turbo-0301"}
	TextDavinci003  = ModelName{Name: "text-davinci-003"}
	TextDavinci002  = ModelName{Name: "text-davinci-002"}
	CodeDavinci002  = ModelName{Name: "code-davinci-002"}
)

/*
GPT-3 models can understand and generate natural language. These models were superceded by the more powerful GPT-3.5 generation models. However, the original GPT-3 base models (, , , and ) are current the only models that are available to fine-tune.davincicurieadababbage

LATEST MODEL	DESCRIPTION	MAX REQUEST	TRAINING DATA
text-curie-001	Very capable, faster and lower cost than Davinci.	2,048 tokens	Up to Oct 2019
text-babbage-001	Capable of straightforward tasks, very fast, and lower cost.	2,048 tokens	Up to Oct 2019
text-ada-001	Capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.	2,048 tokens	Up to Oct 2019
davinci	Most capable GPT-3 model. Can do any task the other models can do, often with higher quality.	2,048 tokens	Up to Oct 2019
curie	Very capable, but faster and lower cost than Davinci.	2,048 tokens	Up to Oct 2019
babbage	Capable of straightforward tasks, very fast, and lower cost.	2,048 tokens	Up to Oct 2019
ada	Capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.	2,048 tokens	Up to Oct 2019
*/

var (
	TextCurie001   = ModelName{Name: "text-curie-001"}
	TextBabbage001 = ModelName{Name: "text-babbage-001"}
	TextAda001     = ModelName{Name: "text-ada-001"}
	Davinci        = ModelName{Name: "davinci"}
	Curie          = ModelName{Name: "curie"}
	Babbage        = ModelName{Name: "babbage"}
	Ada            = ModelName{Name: "ada"}
)
