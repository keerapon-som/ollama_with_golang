package ollama

import "github.com/stretchr/testify/mock"

// type Ollama interface {
// 	GenerateACompletionStream(GenerateACompletionStreamRequest GenerateACompletionStreamRequest) (<-chan GenerateACompletionStreamResponse, error)
// 	ListLocalModels() (RespListLocalModels, error)
// 	ShowModelInformation(ShowModelInformationRequest ShowModelInformationRequest) (ShowModelInformationResponse, error)
// 	CopyAModel(CopyModelRequest CopyModelRequest) error
// 	DeleteAModel(modelName DeleteAModelRequest) error
// 	PullAModel(pullModelReq PullModelRequest) (PullModelResponse, error)
// 	PullAModelStream(pullModelReq PullModelRequest) (<-chan PullModelResponse, error)
// 	PushAModel(modelName PushAModelRequest) (PushAModelResponse, error)
// 	PushAModelStream(modelName PushAModelRequest) (<-chan PushAModelResponse, error)
// 	ListRunningModels() (ListRunningModelsResponse, error)
// }

type ollamaMock struct {
	mock.Mock
}

func NewOllamaMock() Ollama {
	return &ollamaMock{}
}

func (o *ollamaMock) GenerateACompletionStream(GenerateACompletionStreamRequest GenerateACompletionStreamRequest) (<-chan GenerateACompletionStreamResponse, error) {
	args := o.Called(GenerateACompletionStreamRequest)
	return args.Get(0).(<-chan GenerateACompletionStreamResponse), args.Error(1)
}

func (o *ollamaMock) ListLocalModels() (RespListLocalModels, error) {
	args := o.Called()
	return args.Get(0).(RespListLocalModels), args.Error(1)
}

func (o *ollamaMock) ShowModelInformation(ShowModelInformationRequest ShowModelInformationRequest) (ShowModelInformationResponse, error) {
	args := o.Called(ShowModelInformationRequest)
	return args.Get(0).(ShowModelInformationResponse), args.Error(1)
}

func (o *ollamaMock) CopyAModel(CopyModelRequest CopyModelRequest) error {
	args := o.Called(CopyModelRequest)
	return args.Error(0)
}

func (o *ollamaMock) DeleteAModel(modelName DeleteAModelRequest) error {
	args := o.Called(modelName)
	return args.Error(0)
}

func (o *ollamaMock) PullAModel(pullModelReq PullModelRequest) (PullModelResponse, error) {
	args := o.Called(pullModelReq)
	return args.Get(0).(PullModelResponse), args.Error(1)
}

func (o *ollamaMock) PullAModelStream(pullModelReq PullModelRequest) (<-chan PullModelResponse, error) {
	args := o.Called(pullModelReq)
	return args.Get(0).(<-chan PullModelResponse), args.Error(1)
}

func (o *ollamaMock) PushAModel(modelName PushAModelRequest) (PushAModelResponse, error) {
	args := o.Called(modelName)
	return args.Get(0).(PushAModelResponse), args.Error(1)
}

func (o *ollamaMock) PushAModelStream(modelName PushAModelRequest) (<-chan PushAModelResponse, error) {
	args := o.Called(modelName)
	return args.Get(0).(<-chan PushAModelResponse), args.Error(1)
}

func (o *ollamaMock) ListRunningModels() (ListRunningModelsResponse, error) {
	args := o.Called()
	return args.Get(0).(ListRunningModelsResponse), args.Error(1)
}
