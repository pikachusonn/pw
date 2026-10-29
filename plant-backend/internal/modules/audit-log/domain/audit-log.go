package domain

import "time"

type AuditLog struct {
	Action          string    `json:"action"`
	Detail          string    `json:"detail"`
	RelatedObjectId string    `json:"relatedObjectId"`
	CreatedAt       time.Time `json:"createAt"`
}

func NewAuditLog(action string, detail string, relatedObject string) *AuditLog {
	return &AuditLog{
		Action:          action,
		Detail:          detail,
		CreatedAt:       time.Now(),
		RelatedObjectId: relatedObject,
	}
}
