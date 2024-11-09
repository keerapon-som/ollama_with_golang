package ollamasvc

type OllamaService interface {
	GetAllModelsName() ([]string, error)
}
