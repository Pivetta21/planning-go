package enum

type SessionOrigin int

const (
	Web    SessionOrigin = 1
	Mobile SessionOrigin = 2
)

func (e SessionOrigin) String() string {
	switch e {
	case Web:
		return "WEB"
	case Mobile:
		return "MOBILE"
	default:
		return ""
	}
}

func (e SessionOrigin) Index() int {
	return int(e)
}

func (e SessionOrigin) IsDefined() bool {
	return e.String() != ""
}
