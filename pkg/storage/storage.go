package storage

type Storage interface {
	Upload(content []byte) error
}
