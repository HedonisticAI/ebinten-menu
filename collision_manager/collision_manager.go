package collision_manager

type Collision_manager interface {
	IsCollision() (bool, error)
}
