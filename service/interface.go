package service

type HitCounter interface {
	IncrementAndGetCounter() int
}

type DataWriter interface {
	Write(string, interface{}) error
	Read(string, interface{}) error
}
