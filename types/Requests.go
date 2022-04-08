package types

type RequestClip struct {
	UserId    string `json:"user_id"`
	DeviceId  string `json:"device_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
