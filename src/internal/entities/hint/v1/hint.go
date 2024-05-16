package entities_hint_v1

type Hint struct {
	HintType uint32 `json:"hint_type"`
	Hint     string `json:"hint"`
}
