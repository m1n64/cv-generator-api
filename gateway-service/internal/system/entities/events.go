package entities

import "gateway-service/internal/system/enums"

type Event struct {
	Type    enums.NotificationType `json:"type"`
	Message string                 `json:"message"`
}

type NotificationEvent struct {
	Type    enums.NotificationType `json:"type"`
	UserID  string                 `json:"user_id"`
	Message string                 `json:"message"`
}
