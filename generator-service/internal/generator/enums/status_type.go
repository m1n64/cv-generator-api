package enums

type StatusType string

const (
	StatusPending    StatusType = "pending"
	StatusInProgress StatusType = "in_progress"
	StatusCompleted  StatusType = "completed"
	StatusFailed     StatusType = "failed"
)

var validStatuses = map[StatusType]bool{
	StatusPending:    true,
	StatusInProgress: true,
	StatusCompleted:  true,
	StatusFailed:     true,
}

func (s StatusType) IsValid() bool {
	_, ok := validStatuses[s]
	return ok
}

func (s StatusType) String() string {
	return string(s)
}
