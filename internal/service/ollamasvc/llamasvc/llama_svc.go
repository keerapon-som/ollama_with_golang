package llamasvc

import "ollamawithgo/internal/ollama"

type llama3_2Service struct {
	ollama ollama.Ollama
}

func NewLlama3_2Service(ollama ollama.Ollama) Llama3_2Service {
	return &llama3_2Service{
		ollama: ollama,
	}
}

func (o *llama3_2Service) GenerateACompletionFull(GenerateCompletionReq GenerateACompletionServiceRequest) (Resp []string, err error) {
	completionReq := ollama.GenerateACompletionStreamRequest{
		Model:     GenerateCompletionReq.Model,
		Prompt:    GenerateCompletionReq.Prompt,
		Suffix:    GenerateCompletionReq.Suffix,
		Image:     GenerateCompletionReq.Image,
		Format:    GenerateCompletionReq.Format,
		Options:   GenerateCompletionReq.Options,
		System:    GenerateCompletionReq.System,
		Template:  GenerateCompletionReq.Template,
		Context:   GenerateCompletionReq.Context,
		Stream:    GenerateCompletionReq.Stream,
		Raw:       GenerateCompletionReq.Raw,
		KeepAlive: GenerateCompletionReq.KeepAlive,
	}
	ch, err := o.ollama.GenerateACompletionStream(completionReq)
	if err != nil {
		return nil, err
	}
	for r := range ch {
		Resp = append(Resp, r.Response)
	}
	return Resp, nil
}

func (o *llama3_2Service) GenerateACompletionStream(GenerateCompletionReq GenerateACompletionServiceRequest) (<-chan []string, error) {
	completionReq := ollama.GenerateACompletionStreamRequest{
		Model:     GenerateCompletionReq.Model,
		Prompt:    GenerateCompletionReq.Prompt,
		Suffix:    GenerateCompletionReq.Suffix,
		Image:     GenerateCompletionReq.Image,
		Format:    GenerateCompletionReq.Format,
		Options:   GenerateCompletionReq.Options,
		System:    GenerateCompletionReq.System,
		Template:  GenerateCompletionReq.Template,
		Context:   GenerateCompletionReq.Context,
		Stream:    GenerateCompletionReq.Stream,
		Raw:       GenerateCompletionReq.Raw,
		KeepAlive: GenerateCompletionReq.KeepAlive,
	}
	ch, err := o.ollama.GenerateACompletionStream(completionReq)
	if err != nil {
		return nil, err
	}

	var respChan = make(chan []string)
	go func() {
		defer close(respChan)
		for r := range ch {
			respChan <- []string{r.Response}
		}
	}()
	return respChan, nil
}
