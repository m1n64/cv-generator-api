package entites

import "ai-service/internal/ai/enums"

type AiAnalyticsQueueMessage struct {
	Action   enums.AiAnalyticsAction `json:"action"`
	Prompt   string                  `json:"prompt"`
	Response string                  `json:"response"`
	Service  string                  `json:"service"`
	SendAt   string                  `json:"send_at"`
}
