package types

type ResponseClip struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	DeviceId  string `json:"device_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
