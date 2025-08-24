package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"neuralops/api/proto/ai_engine/v1"
	"neuralops/llm/policies"
)

type OpenAIAdapter struct {
	apiKey   string
	endpoint string
	client   *http.Client
}

func NewOpenAIAdapter(apiKey, endpoint string) *OpenAIAdapter {
	return &OpenAIAdapter{
		apiKey:   apiKey,
		endpoint: endpoint,
		client:   &http.Client{},
	}
}

type openAIRequest struct {
	Model    string          `json:"model"`
	Messages []openAIMessage `json:"messages"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message openAIMessage `json:"message"`
	} `json:"choices"`
}

func (a *OpenAIAdapter) GeneratePlan(ctx context.Context, query string) (*ai_enginev1.PipelinePlan, error) {
	prompt := fmt.Sprintf(policies.SystemPrompt, query)

	reqPayload := openAIRequest{
		Model: "gpt-3.5-turbo", // Or any other model
		Messages: []openAIMessage{
			{Role: "system", Content: "You are a helpful assistant that generates JSON pipeline plans."},
			{Role: "user", Content: prompt},
		},
	}

	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	var openAIResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("no choices returned from OpenAI")
	}

	var plan ai_enginev1.PipelinePlan
	// The plan is expected to be a JSON string in the content
	err = json.Unmarshal([]byte(openAIResp.Choices[0].Message.Content), &plan)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal plan from LLM response: %w", err)
	}

	return &plan, nil
}
