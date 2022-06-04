package cachemanager

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/adamrdrew/mosh/config"
	"github.com/stretchr/testify/assert"
)

var ORIG_CACHE_DIR = os.Getenv(ENV)

const DIR = "CACHE_MANAGER_TEST"
const MAX_SIZE = 25
const MAX_AGE = 2
const ENV = "MOSH_CACHE_DIR"

func getCacheManager() CacheManager {
	return CacheManager{
		conf: config.Config{
			CacheMaxSizeMB:  MAX_SIZE,
			CacheMaxAgeDays: MAX_AGE,
		},
	}
}

func makeDir() {
	os.Mkdir(DIR, 0777)
	os.Setenv(ENV, DIR)
}

func makeFiles(count int, size int, dayDif int) {
	for i := 0; i < count; i++ {
		makeFile(size, dayDif)
	}
}

func makeFile(sizeMB int, dayDif int) {
	size := int64(sizeMB * 1024 * 1024)
	min := 1000000
	max := 9999999
	rNum := rand.Intn(max-min) + min
	fname := DIR + "/" + fmt.Sprint(rNum)
	fd, err := os.Create(fname)
	if err != nil {
		log.Fatal("Failed to create output")
	}
	_, err = fd.Seek(size-1, 0)
	if err != nil {
		log.Fatal("Failed to seek")
	}
	_, err = fd.Write([]byte{0})
	if err != nil {
		log.Fatal("Write failed")
	}
	err = fd.Close()
	if err != nil {
		log.Fatal("Failed to close file")
	}
	spoofDate := time.Now().AddDate(0, 0, dayDif)
	os.Chtimes(fname, spoofDate, spoofDate)
}

func rmDir() {
	os.RemoveAll(DIR)
	os.Setenv(ENV, ORIG_CACHE_DIR)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestNoOldFilesCacheNotTooLarge(t *testing.T) {
	makeDir()
	cm := getCacheManager()
	makeFiles(4, MAX_SIZE/8, 0)
	sizeBeforePrune := cm.getCacheSizeMB()
	listBeforePrune := cm.getFileList()
	cm.PruneCache()
	assert.Equal(t, sizeBeforePrune, cm.getCacheSizeMB())
	assert.Equal(t, listBeforePrune, cm.getFileList())
	rmDir()
}

func TestNoOldFilesCacheTooLarge(t *testing.T) {
	makeDir()
	cm := getCacheManager()
	makeFiles(4, MAX_SIZE, 0)
	sizeBeforePrune := cm.getCacheSizeMB()
	listBeforePrune := cm.getFileList()
	cm.PruneCache()
	assert.Greater(t, sizeBeforePrune, cm.getCacheSizeMB())
	assert.Greater(t, len(listBeforePrune), len(cm.getFileList()))
	rmDir()
}

func TestOldFilesCacheNotTooLarge(t *testing.T) {
	makeDir()
	cm := getCacheManager()
	makeFiles(2, MAX_SIZE/8, -10)
	makeFiles(2, MAX_SIZE/8, 0)
	sizeBeforePrune := cm.getCacheSizeMB()
	listBeforePrune := cm.getFileList()
	cm.PruneCache()
	sizeAfterPrune := cm.getCacheSizeMB()
	listAfterPrune := cm.getFileList()
	assert.Equal(t, sizeBeforePrune/2, sizeAfterPrune)
	assert.Equal(t, len(listBeforePrune)/2, len(listAfterPrune))
	rmDir()
}

func TestOldFilesCacheTooLarge(t *testing.T) {
	makeDir()
	cm := getCacheManager()
	makeFiles(2, MAX_SIZE, -10)
	makeFiles(2, MAX_SIZE, 0)
	//sizeBeforePrune := cm.getCacheSizeMB()
	//listBeforePrune := cm.getFileList()
	cm.PruneCache()
	sizeAfterPrune := cm.getCacheSizeMB()
	listAfterPrune := cm.getFileList()
	assert.Equal(t, int64(0), sizeAfterPrune)
	assert.Equal(t, 0, len(listAfterPrune))
	rmDir()
}
