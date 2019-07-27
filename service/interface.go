package service

type HitCounter interface {
	IncrementAndGetCounter() int
}
