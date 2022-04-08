package types

type ResponseClip struct {
	ID        string `json:"id"`
	UserId    string `json:"userId"`
	DeviceId  string `json:"device_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}
