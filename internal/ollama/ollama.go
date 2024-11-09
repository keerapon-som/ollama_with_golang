package ollama

type GenerateACompletionStreamRequest struct {
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

type ShowModelInformationRequest struct {
	Name    string `json:"name"`
	Verbose bool   `json:"verbose"`
}

type ShowModelInformationResponse struct {
	Modelfile  string `json:"modelfile"`
	Parameters string `json:"parameters"`
	Template   string `json:"template"`
	Details    struct {
		ParentModel       string   `json:"parent_model"`
		Format            string   `json:"format"`
		Family            string   `json:"family"`
		Families          []string `json:"families"`
		ParameterSize     string   `json:"parameter_size"`
		QuantizationLevel string   `json:"quantization_level"`
	} `json:"details"`
	ModelInfo struct {
		GeneralArchitecture               string        `json:"general.architecture"`
		GeneralFileType                   int           `json:"general.file_type"`
		GeneralParameterCount             int64         `json:"general.parameter_count"`
		GeneralQuantizationVersion        int           `json:"general.quantization_version"`
		LlamaAttentionHeadCount           int           `json:"llama.attention.head_count"`
		LlamaAttentionHeadCountKv         int           `json:"llama.attention.head_count_kv"`
		LlamaAttentionLayerNormRmsEpsilon float64       `json:"llama.attention.layer_norm_rms_epsilon"`
		LlamaBlockCount                   int           `json:"llama.block_count"`
		LlamaContextLength                int           `json:"llama.context_length"`
		LlamaEmbeddingLength              int           `json:"llama.embedding_length"`
		LlamaFeedForwardLength            int           `json:"llama.feed_forward_length"`
		LlamaRopeDimensionCount           int           `json:"llama.rope.dimension_count"`
		LlamaRopeFreqBase                 int           `json:"llama.rope.freq_base"`
		LlamaVocabSize                    int           `json:"llama.vocab_size"`
		TokenizerGgmlBosTokenID           int           `json:"tokenizer.ggml.bos_token_id"`
		TokenizerGgmlEosTokenID           int           `json:"tokenizer.ggml.eos_token_id"`
		TokenizerGgmlMerges               []interface{} `json:"tokenizer.ggml.merges"`
		TokenizerGgmlModel                string        `json:"tokenizer.ggml.model"`
		TokenizerGgmlPre                  string        `json:"tokenizer.ggml.pre"`
		TokenizerGgmlTokenType            []interface{} `json:"tokenizer.ggml.token_type"`
		TokenizerGgmlTokens               []interface{} `json:"tokenizer.ggml.tokens"`
	} `json:"model_info"`
}

type CopyModelRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type DeleteAModelRequest struct {
	Name string `json:"name"`
}

type RespListLocalModels struct {
	Models []struct {
		Name       string `json:"name"`
		ModifiedAt string `json:"modified_at"`
		Size       int64  `json:"size"`
		Digest     string `json:"digest"`
		Details    struct {
			Format            string      `json:"format"`
			Family            string      `json:"family"`
			Families          interface{} `json:"families"`
			ParameterSize     string      `json:"parameter_size"`
			QuantizationLevel string      `json:"quantization_level"`
		} `json:"details"`
	} `json:"models"`
}

type PullModelRequest struct {
	Name     string `json:"name"`
	Insecure bool   `json:"insecure,omitempty"`
	Stream   bool   `json:"stream,omitempty"`
}

type PullModelResponse struct {
	Status    string `json:"status"`
	Digest    string `json:"digest,omitempty"`
	Total     int64  `json:"total,omitempty"`
	Completed int64  `json:"completed,omitempty"`
}

type PushAModelRequest struct {
	Name     string `json:"name"`
	Insecure bool   `json:"insecure,omitempty"`
	Stream   bool   `json:"stream,omitempty"`
}

type PushAModelResponse struct {
	Status string `json:"status"`
	Digest string `json:"digest"`
	Total  int    `json:"total"`
}

type ListRunningModelsResponse struct {
	Models []struct {
		Name    string `json:"name"`
		Model   string `json:"model"`
		Size    int64  `json:"size"`
		Digest  string `json:"digest"`
		Details struct {
			ParentModel       string   `json:"parent_model"`
			Format            string   `json:"format"`
			Family            string   `json:"family"`
			Families          []string `json:"families"`
			ParameterSize     string   `json:"parameter_size"`
			QuantizationLevel string   `json:"quantization_level"`
		} `json:"details"`
		ExpiresAt string `json:"expires_at"`
		SizeVram  int64  `json:"size_vram"`
	} `json:"models"`
}

type Ollama interface {
	GenerateACompletionStream(GenerateACompletionStreamRequest GenerateACompletionStreamRequest) (<-chan GenerateACompletionStreamResponse, error)
	ListLocalModels() (RespListLocalModels, error)
	ShowModelInformation(ShowModelInformationRequest ShowModelInformationRequest) (ShowModelInformationResponse, error)
	CopyAModel(CopyModelRequest CopyModelRequest) error
	DeleteAModel(modelName DeleteAModelRequest) error
	PullAModel(pullModelReq PullModelRequest) (PullModelResponse, error)
	PullAModelStream(pullModelReq PullModelRequest) (<-chan PullModelResponse, error)
	PushAModel(modelName PushAModelRequest) (PushAModelResponse, error)
	PushAModelStream(modelName PushAModelRequest) (<-chan PushAModelResponse, error)
	ListRunningModels() (ListRunningModelsResponse, error)
}
