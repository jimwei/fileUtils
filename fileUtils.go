// os.RemoveAll not work correctly on windows platform,
// if there a file which attribute is read only, the RemoveAll function will fail.
// so i wrote a RemoveAllEx function to fix it.
package fileUtils

import (
	"log"
	"os"
	"path/filepath"
)

// remove the given path, if the path is a file
// remove sub folders and all files in the path, if the path is a directory
func RemoveAllEx(path string) error {
	err := resetReadOnlyFlagAll(path)
	if err != nil {
		return err
	}
	return os.RemoveAll(path)
}

// check the given file exists or not
func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}
func resetReadOnlyFlagAll(path string) error {

	fi, err := os.Stat(path)
	if err != nil {
		//the directory not exists
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if !fi.IsDir() {
		err := os.Chmod(path, 0666)
		if err != nil {
			return err
		}
	}
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	names, _ := fd.Readdirnames(-1)
	for _, name := range names {
		newNames := filepath.Join(path, name)
		log.Println("the sub name is", newNames)
		resetReadOnlyFlagAll(newNames)
	}

	return nil
}
