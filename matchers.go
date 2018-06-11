package gomegafs

import (
	"github.com/nodo/gomegafs/matchers"
	"github.com/onsi/gomega/types"
)

func BeEqualToDir(pathToDir string) types.GomegaMatcher {
	return &matchers.BeEqualToDirMatcher{
		Path: pathToDir,
	}
}
