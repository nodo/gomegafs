package matchers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	. "github.com/nodo/gomegafs"
	"github.com/nodo/gomegafs/matchers"
)

func TestEqualityBewtweenTwoDirectories(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect("fixtures/equal/one").To(BeEqualToDir("fixtures/equal/two"))
}

func TestDirectoryWithDifferentContent(t *testing.T) {
	g := NewGomegaWithT(t)
	success, err := (&matchers.BeEqualToDirMatcher{
		Path: "fixtures/different/one",
	}).Match("fixtures/different/two")
	g.Expect(success).To(BeFalse(), "Expected matcher to fail, but it succeeded")
	g.Expect(err).To(MatchError("file fixtures/different/one/hola.txt is not equal to fixtures/different/two/hola.txt"))
}

func TestDirectoryWithDifferentNames(t *testing.T) {
	g := NewGomegaWithT(t)
	success, err := (&matchers.BeEqualToDirMatcher{
		Path: "fixtures/different/one",
	}).Match("fixtures/different/three")
	g.Expect(success).To(BeFalse(), "Expected matcher to fail, but it succeeded")
	g.Expect(err).To(MatchError("file fixtures/different/three/hola.txt does not exists in the tree of fixtures/different/three"))
}

func TestInvalidInput(t *testing.T) {
	g := NewGomegaWithT(t)
	success, err := (&matchers.BeEqualToDirMatcher{}).Match(nil)
	g.Expect(success).To(BeFalse(), "Expected matcher to fail, but it succeeded")
	g.Expect(err).To(MatchError("<nil> must be a path to a directory"))
}

func TestUnreadableDirectory(t *testing.T) {
	g := NewGomegaWithT(t)
	success, err := (&matchers.BeEqualToDirMatcher{
		Path: "fixtures/unreadable",
	}).Match("fixtures/equal")
	g.Expect(success).To(BeFalse(), "Expected matcher to fail, but it succeeded")
	g.Expect(err).To(MatchError("could not read directory fixtures/unreadable"))
}
