package llamasvc

type GenerateACompletionServiceRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Suffix string `json:"suffix,omitempty"`
	Image  string `json:"image,omitempty"`
	// Advanced parameters (optional):
	Format    string `json:"format,omitempty"`
	Options   string `json:"options,omitempty"`
	System    string `json:"system,omitempty"`
	Template  string `json:"template,omitempty"`
	Context   string `json:"context,omitempty"`
	Stream    bool   `json:"stream"`
	Raw       bool   `json:"raw,omitempty"`
	KeepAlive string `json:"keep_alive,omitempty"`
}

type GenerateACompletionStreamResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

type Llama3_2Service interface {
	GenerateACompletionFull(GenerateCompletionReq GenerateACompletionServiceRequest) ([]string, error)
	GenerateACompletionStream(GenerateCompletionReq GenerateACompletionServiceRequest) (<-chan []string, error)
}
