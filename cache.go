package cache

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type KeyMap struct {
	FolderName string
}

type Cache struct {
	FolderName string
	Map        map[string]string
}

const (
	DEFAULT_FOLDER_NAME = "cache"
)

var c *Cache

// Creates the root folder
func Default() *Cache {
	c = &Cache{FolderName: DEFAULT_FOLDER_NAME}
	createFolderIfNotExists(c.FolderName)
	return c
}

// Creates the cache folder with new name
func New(folder string) *Cache {
	c = &Cache{FolderName: folder}
	createFolderIfNotExists(c.FolderName)
	return c
}

// Saving a file in a particular folder
// key: folder name
// filename: the name of file to be saved
// data: the content of the the file
func (c *Cache) Save(key string, filename string, data []byte) {
	// /<cache>/<key>
	keyPath := generateKeyPath(c, key)
	filePath := generateFilePath(keyPath, filename)

	// Make sure key path exists
	createFolderIfNotExists(keyPath)

	// Only write new file cache has not been found
	if !isFileExists(filePath) {
		err := ioutil.WriteFile(filePath, data, 0755)
		panic(err)
	}
}

func generateKeyPath(c *Cache, key string) string {
	return filepath.Join(c.FolderName, key)
}

func generateFilePath(root, file string) string {
	return filepath.Join(root, file)
}

func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// Creates a folder in current directory if it not exists
func createFolderIfNotExists(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.Mkdir(name, 0755)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
