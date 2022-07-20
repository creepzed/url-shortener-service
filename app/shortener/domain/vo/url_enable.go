package vo

const (
	Enabled  = true
	Disabled = false
)

type UrlEnabled struct {
	value bool
}

func NewUrlEnabled(value bool) UrlEnabled {
	return UrlEnabled{
		value: value,
	}
}

func (u UrlEnabled) Value() bool {
	return u.value
}
