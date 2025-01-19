package repositories

type ColorRepository interface {
	GetColors() (map[string]string, error)
}
