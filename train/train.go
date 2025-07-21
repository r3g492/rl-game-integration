package train

type Observation struct {
	CarX  float32 `json:"car_x"`
	CarY  float32 `json:"car_y"`
	CarZ  float32 `json:"car_z"`
	Yaw   float32 `json:"yaw"`
	GoalX float32 `json:"goal_x"`
	GoalY float32 `json:"goal_y"`
	GoalZ float32 `json:"goal_z"`
}

type StepResponse struct {
	Observation Observation `json:"observation"`
	Reward      float32     `json:"reward"`
	Done        bool        `json:"done"`
	Info        interface{} `json:"info,omitempty"` // Optional: debugging/extra info
}

type StepInfo struct {
	DistanceToGoal float32 `json:"distance_to_goal"`
	Collision      bool    `json:"collision"`
}
