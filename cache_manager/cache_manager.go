package cachemanager

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/adamrdrew/mosh/config"
	"github.com/djherbis/atime"
)

//We use this in conversions
const MB = 1048576

//Get a new CacheManager
func MakeCacheManager() CacheManager {
	return CacheManager{
		conf: config.GetConfig(),
	}
}

type CacheManager struct {
	conf config.Config
}

//Delete old files and free up space if cache size is too large
func (c *CacheManager) PruneCache() {
	log.Print("CacheManager: Starting cache prune...")
	c.deleteOldFiles()
	c.deleteFilesIfCacheIsTooLarge()
}

func (c *CacheManager) cachePath(fileName string) string {
	return config.GetCacheDir() + "/" + fileName
}

func (c *CacheManager) fileIsAvailable(fileName string) bool {
	//Check if the file exists
	_, error := os.Stat(c.cachePath(fileName))
	if errors.Is(error, os.ErrNotExist) {
		return false
	}
	if errors.Is(error, os.ErrPermission) {
		return false
	}
	return true
}

func (c *CacheManager) deleteFilesIfCacheIsTooLarge() {
	log.Print("CacheManager: Checking if cache needs to shrink...")
	cacheSize := c.getCacheSizeMB()
	if cacheSize < int64(c.conf.CacheMaxSizeMB) {
		log.Printf("CacheManager: Total cache size %dMB is less than CacheMaxSizeMB %dMB.", cacheSize, c.conf.CacheMaxSizeMB)
		return
	}
	log.Printf("CacheManager: Total cache size %dMB is exceeds CacheMaxSizeMB %dMB. Cleaning up...", cacheSize, c.conf.CacheMaxSizeMB)
	deleteList := []string{}
	fileList := c.getFileList()
	spaceFreedUp := float64(0.0)
	spaceFreeUpTarget := float64(cacheSize) - (float64(c.conf.CacheMaxSizeMB) * .50)
	for _, fileInfo := range fileList {
		if !c.fileIsAvailable(fileInfo.Name()) {
			continue
		}
		size := fileInfo.Size()
		deleteList = append(deleteList, fileInfo.Name())
		spaceFreedUp += float64(size) / MB
		if spaceFreedUp >= spaceFreeUpTarget {
			break
		}
	}
	for _, fileName := range deleteList {
		c.deleteFile(fileName)
	}
	freedUpMB := int64(spaceFreedUp)
	log.Printf("CacheManager: Freed up %dMB of space.", freedUpMB)
}

func (c *CacheManager) getCacheSizeMB() int64 {
	out := int64(0)
	fileList := c.getFileList()
	for _, fileInfo := range fileList {
		out += fileInfo.Size() / MB
	}
	return out
}

func (c *CacheManager) deleteOldFiles() {
	log.Print("CacheManager: Checking for old files to delete...")
	deleted := 0
	files := c.getFileList()
	for _, fileInfo := range files {
		name := fileInfo.Name()
		if !c.fileIsAvailable(name) {
			continue
		}
		if c.isFileTooOld(name) {
			c.deleteFile(name)
			deleted += 1
		}
	}
	log.Printf("CacheManager: Deleted %d old files", deleted)
}

func (c *CacheManager) getFileList() []fs.FileInfo {
	var out []fs.FileInfo
	var err error
	out, err = ioutil.ReadDir(config.GetCacheDir())
	if err != nil {
		out = []fs.FileInfo{}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].ModTime().Before(out[j].ModTime())
	})
	return out
}

func (c *CacheManager) isFileTooOld(fileName string) bool {
	out := false
	at, err := atime.Stat(c.cachePath(fileName))
	if err != nil {
		return false
	}
	today := time.Now()
	oldDate := today.AddDate(0, 0, -1*c.conf.CacheMaxAgeDays)
	if at.Before(oldDate) {
		out = true
	}
	return out
}
func (c *CacheManager) deleteFile(fileName string) {
	if !c.fileIsAvailable(fileName) {
		return
	}
	os.Remove(c.cachePath(fileName))
}
