package delivery

// RegisterEndpoints(e, db)

type Handler interface {
	GET(path string)
}
