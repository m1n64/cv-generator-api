package entities

import "cv-generator-service/internal/notifications/enums"

type NotificationEvent struct {
	Type    enums.NotificationType `json:"type"`
	UserID  string                 `json:"user_id"`
	Message string                 `json:"message"`
}
