package matchers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type BeEqualToDirMatcher struct {
	Path string
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func sameFile(filePath, otherFilePath string) bool {
	content, _ := ioutil.ReadFile(filePath)
	otherContent, _ := ioutil.ReadFile(otherFilePath)
	return string(content) == string(otherContent)
}

func sameFolder(path, otherPath string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, fmt.Errorf("could not read directory %s", path)
	}
	for _, f := range files {
		srcPath := filepath.Join(path, f.Name())
		dstPath := filepath.Join(otherPath, f.Name())
		if !fileExists(dstPath) {
			return false, fmt.Errorf("file %s does not exists in the tree of %s", dstPath, otherPath)
		}
		// if f.IsDir() {
		// 	sameFolder(srcPath, dstPath)
		// 	continue
		// }
		if !sameFile(srcPath, dstPath) {
			return false, fmt.Errorf("file %s is not equal to %s", srcPath, dstPath)
		}

	}
	return true, nil
}

func (m *BeEqualToDirMatcher) Match(actual interface{}) (bool, error) {
	otherPath, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("%v must be a path to a directory", actual)
	}
	success, err := sameFolder(m.Path, otherPath)
	return success, err
}

func (m *BeEqualToDirMatcher) FailureMessage(actual interface{}) string {
	return "TODO"
}

func (m *BeEqualToDirMatcher) NegatedFailureMessage(actual interface{}) string {
	return "TODO"
}
