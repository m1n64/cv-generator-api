package enums

type AiAnalyticsAction string

const (
	AiAnalyticsActionGenerate AiAnalyticsAction = "ai_generate"
)

func (a AiAnalyticsAction) String() string {
	return string(a)
}
