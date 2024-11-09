package ollamasvc

import "ollamawithgo/internal/ollama"

type ollamaService struct {
	ollama ollama.Ollama
}

func NewOllamaService(ollama ollama.Ollama) OllamaService {
	return &ollamaService{
		ollama: ollama,
	}
}

func (s *ollamaService) GetAllModelsName() ([]string, error) {
	resp, err := s.ollama.ListLocalModels()
	if err != nil {
		return nil, err
	}
	var models []string
	for _, model := range resp.Models {
		models = append(models, model.Name)
	}
	return models, nil
}
