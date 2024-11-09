package ollama

import (
	"fmt"
	"testing"
)

func TestListLocalModels(t *testing.T) {
	o := NewOllama()
	resp, err := o.ListLocalModels()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Println(resp)
}

func TestPullModelStream(t *testing.T) {
	o := NewOllama()
	req := PullModelRequest{Name: "llama3.2:1b", Stream: true}
	resp, err := o.PullAModelStream(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	for r := range resp {
		fmt.Println(r)
	}
}

func TestGenerateACompletionStream(t *testing.T) {
	o := NewOllama()
	req := GenerateACompletionStreamRequest{Model: "llama3.2:1b", Prompt: "The quick brown fox", Stream: true}
	resp, err := o.GenerateACompletionStream(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	for r := range resp {
		fmt.Println(r)
	}
}

func TestShowModelInformation(t *testing.T) {
	o := NewOllama()
	req := ShowModelInformationRequest{Name: "llama3.2:1b"}
	resp, err := o.ShowModelInformation(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Println(resp)
}
