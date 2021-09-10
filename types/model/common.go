package model

// EntityTimestamp 实体追踪时间
type EntityTimestamp struct {
	CreatedAt int64 `json:"createdAt" query:"createdAt" xml:"created_at" bson:"createdAt"`
	UpdatedAt int64 `json:"updatedAt" query:"updatedAt" xml:"updated_at" bson:"updatedAt"`
}