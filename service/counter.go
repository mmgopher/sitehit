package service

import (
	"fmt"
	"site-hit/config"
	"site-hit/logger"
	"site-hit/util"
	"sync"
	"time"
)

type HitCount struct {
	hitsMap     map[int64]int
	windowTime  int64
	storageFile string
	dataWriter  DataWriter
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
	err := hc.dataWriter.Write(hc.storageFile, hc.hitsMap)
	if err != nil {
		logger.Error("Can't persist counter in storage file", err)
	}
	return counter
}

func NewHitCounter(dataWriter DataWriter) *HitCount {
	storageFile := config.GetConfiguration().StorageFile
	logger.Info(fmt.Sprintf("Initialize hitCounter. Storage file: %s", storageFile))
	hitMap := make(map[int64]int)
	if util.FileExist(storageFile) {
		err := dataWriter.Read(storageFile, &hitMap)
		if err != nil {
			logger.Error("Can't read storage file", err)
		}
	}
	return &HitCount{hitsMap: hitMap, windowTime: int64(config.GetConfiguration().WindowTime), storageFile: storageFile, dataWriter: dataWriter}
}
