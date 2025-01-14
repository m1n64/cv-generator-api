package enums

type NotificationType string

const (
	TypeSuccess NotificationType = "success"
	TypeError   NotificationType = "error"
)

func (t NotificationType) String() string {
	return string(t)
}
