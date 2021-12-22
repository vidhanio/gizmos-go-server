package models

type GizmoResponse struct {
	Message string `json:"message"`
	Gizmo   Gizmo  `json:"gizmo"`
}

type GizmosResponse struct {
	Message string  `json:"message"`
	Gizmos  []Gizmo `json:"gizmos"`
}

func NewGizmoResponse(message string, gizmo Gizmo) *GizmoResponse {
	return &GizmoResponse{
		Message: message,
		Gizmo:   gizmo,
	}
}

func NewGizmosResponse(message string, gizmos []Gizmo) *GizmosResponse {
	return &GizmosResponse{
		Message: message,
		Gizmos:  gizmos,
	}
}
