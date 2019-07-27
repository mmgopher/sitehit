package service

import (
	"site-hit/config"
	"site-hit/util"
	"sync"
	"time"
)

type HitCount struct {
	hitsMap map[int64]int
	windowTime int64
	storageFile string
	sync.Mutex
}

func (hc *HitCount) IncrementAndGetCounter() int {
	hc.Mutex.Lock()
	defer hc.Mutex.Unlock()
	now := time.Now().Unix()
	count := hc.hitsMap[now]
	hc.hitsMap[now] = count + 1
	counter := 0
	newHitsMap := make(map[int64]int)
	for i := now; i > now-hc.windowTime; i-- {
		counter = counter + hc.hitsMap[i]
		newHitsMap[i] = hc.hitsMap[i]
	}
	hc.hitsMap = newHitsMap
	util.WriteGob(hc.storageFile,hc.hitsMap)
	return counter
}

func NewHitCounter() *HitCount {
	return &HitCount{hitsMap: make(map[int64]int),windowTime: int64(config.GetConfiguration().WindowTime),storageFile:config.GetConfiguration().StorageFile}
}
