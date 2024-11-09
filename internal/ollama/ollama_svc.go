package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ollamawithgo/config"
	"strings"
)

type ollama struct {
}

func NewOllama() Ollama {
	return &ollama{}
}

func (s *ollama) ListLocalModels() (response RespListLocalModels, err error) {
	config := config.GetConfig()
	url := config.Ollama.BASEURL + "api/tags"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return RespListLocalModels{}, ErrRequestFailed
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return RespListLocalModels{}, ErrDoFailed
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return RespListLocalModels{}, ErrReadbody
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return RespListLocalModels{}, ErrUnmarshal
	}

	return response, nil

}

func (s *ollama) PullAModel(request PullModelRequest) (PullModelResponse, error) {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/pull"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return PullModelResponse{}, ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return PullModelResponse{}, ErrRequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PullModelResponse{}, ErrStatuscode
	}

	var response PullModelResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return PullModelResponse{}, ErrUnmarshal
	}

	return response, nil
}

func (s *ollama) PullAModelStream(request PullModelRequest) (<-chan PullModelResponse, error) {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/pull"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, ErrRequestFailed
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrStatuscode
	}

	ch := make(chan PullModelResponse)
	go func() {
		defer close(ch)
		decoder := json.NewDecoder(resp.Body)
		for {
			var response PullModelResponse
			if err := decoder.Decode(&response); err == io.EOF {
				break
			} else if err != nil {
				break
			}
			ch <- response
		}
	}()

	return ch, nil
}

func (s *ollama) GenerateACompletionStream(request GenerateACompletionStreamRequest) (<-chan GenerateACompletionStreamResponse, error) {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/generate"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, ErrRequestFailed
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrStatuscode
	}

	ch := make(chan GenerateACompletionStreamResponse)
	go func() {
		defer close(ch)
		decoder := json.NewDecoder(resp.Body)
		for {
			var response GenerateACompletionStreamResponse
			if err := decoder.Decode(&response); err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				break
			}
			ch <- response
		}
	}()

	return ch, nil
}

func (s *ollama) ShowModelInformation(request ShowModelInformationRequest) (ShowModelInformationResponse, error) {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/show"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return ShowModelInformationResponse{}, ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return ShowModelInformationResponse{}, ErrRequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ShowModelInformationResponse{}, ErrStatuscode
	}

	var response ShowModelInformationResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return ShowModelInformationResponse{}, ErrUnmarshal
	}

	return response, nil
}

func (s *ollama) CopyAModel(request CopyModelRequest) error {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/copy"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return ErrRequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrStatuscode
	}

	return nil
}

func (s *ollama) DeleteAModel(request DeleteAModelRequest) error {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/delete"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return ErrRequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrStatuscode
	}

	return nil
}

func (s *ollama) PushAModel(request PushAModelRequest) (PushAModelResponse, error) {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/push"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return PushAModelResponse{}, ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return PushAModelResponse{}, ErrRequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PushAModelResponse{}, ErrStatuscode
	}

	var response PushAModelResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return PushAModelResponse{}, ErrUnmarshal
	}

	return response, nil
}

func (s *ollama) PushAModelStream(request PushAModelRequest) (<-chan PushAModelResponse, error) {
	config := config.GetConfig()

	url := config.Ollama.BASEURL + "api/push"
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, ErrMarshal
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, ErrRequestFailed
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrStatuscode
	}

	ch := make(chan PushAModelResponse)
	go func() {
		defer close(ch)
		decoder := json.NewDecoder(resp.Body)
		for {
			var response PushAModelResponse
			if err := decoder.Decode(&response); err == io.EOF {
				break
			} else if err != nil {
				break
			}
			ch <- response
		}
	}()

	return ch, nil
}

func (s *ollama) ListRunningModels() (ListRunningModelsResponse, error) {
	config := config.GetConfig()
	url := config.Ollama.BASEURL + "api/ps"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return ListRunningModelsResponse{}, ErrRequestFailed
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ListRunningModelsResponse{}, ErrDoFailed
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ListRunningModelsResponse{}, ErrReadbody
	}

	var response ListRunningModelsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return ListRunningModelsResponse{}, ErrUnmarshal
	}

	return response, nil
}
