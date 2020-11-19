package codes

var (
	ArchiveTypeForbidAdd = 10001
	ArchiveExist = 10002
	ArchiveNotExist = 10003
	ArchiveAlreadyDel = 10004
)

type alertCode int

// access a value of an unexported identifier
func New(value int) alertCode {
	return alertCode(value)
}

// unexported fields from an exported struct
type Status struct {
	Name string
	message string
}
