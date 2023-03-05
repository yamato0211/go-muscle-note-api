package schemas

type NewMenuInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Target      string  `json:"target"`
	Weight      float64 `json:"weight"`
	IsJoint     bool    `json:"is_joint"`
	Link        string  `json:"link"`
}
