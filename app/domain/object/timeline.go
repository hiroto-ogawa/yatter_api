package object

type Timeline struct {
	MaxID *int64 `json:"max_id,omitempty"`

	SinceID *int64 `json:"since_id,omitempty"`

	Limit *int64 `json:"limit,omitempty"`
}
