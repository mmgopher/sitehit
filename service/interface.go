package service

type HitCounter interface {
	IncrementAndGetCounter() int
}

type PersistenceManager interface {
	Write(string, interface{}) error
	Read(string, interface{}) error
}
