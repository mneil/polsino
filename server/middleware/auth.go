package middleware

type (
	AuthIFace interface {
		verify() func()
	}
)
