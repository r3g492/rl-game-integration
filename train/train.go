package train

type Observation struct {
	CarX     float32 `json:"car_x"`
	CarY     float32 `json:"car_y"`
	CarZ     float32 `json:"car_z"`
	Velocity float32 `json:"velocity"`
	Yaw      float32 `json:"yaw"`
	GoalX    float32 `json:"goal_x"`
	GoalY    float32 `json:"goal_y"`
	GoalZ    float32 `json:"goal_z"`
}

type StepResponse struct {
	Observation Observation `json:"observation"`
	Reward      float32     `json:"reward"`
	Done        bool        `json:"done"`
	Truncated   bool        `json:"truncated"`  // timeout / step cap
	IsSuccess   bool        `json:"is_success"` // reached goal
}

type StepRequest struct {
	SpeedGradient    float32 `json:"speedGradient"`
	RotationGradient float32 `json:"rotationGradient"`
}
