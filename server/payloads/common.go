package payloads

type ResponseCreate struct {
	ID   string `json:"id"`
	GOTO string `json:"goto,omitempty"`
	Data string `json:"data"`
}
