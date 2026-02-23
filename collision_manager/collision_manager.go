package collision_manager

type Controller interface {
	Status() ControllerResponse
	SetActive(bool)
}
