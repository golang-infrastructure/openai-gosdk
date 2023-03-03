# [OpenAI Doc](https://platform.openai.com/docs/api-reference/introduction)

# Quick start

```shell
go get github.com/songzhibin97/openai-gosdk
```

```go
package main

import (
	"encoding/json"
	"fmt"
	
	openai_gosdk "github.com/songzhibin97/openai-gosdk"
)

func main() {
	base := openai_gosdk.NewBaseOpenAI("sk-xxxxxx", "")
	resp, err := openai_gosdk.NewChat(base).DoRequest(openai_gosdk.RequestChat{
		Model: openai_gosdk.GPT3p5Turbo.ModelName(),
		Messages: []openai_gosdk.Message{
			{
				Role:    openai_gosdk.RoleUser,
				Content: "hello chatgpt!",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	v, _ := json.Marshal(resp)
	// {"id":"cha*****Vl","object":"chat.completion","created":1677828863,"choices":[{"index":0,"message":{"role":"assistant","content":"\n\nHello there, how can I assist you today?"},"finish_reason":"stop"}],"usage":{"prompt_tokens":12,"completion_tokens":12,"total_tokens":24}}
	fmt.Println(string(v))
}

```