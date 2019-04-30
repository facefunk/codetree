package codetree_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aerogo/codetree"
)

func TestCodeTree(t *testing.T) {
	bytes, _ := ioutil.ReadFile("test/example.txt")
	code := string(bytes)
	tree, err := codetree.New(code)

	assert.NoError(t, err)
	defer tree.Close()

	assert.Equal(t, -1, tree.Indent)
	assert.Equal(t, 6, len(tree.Children))
	assert.Equal(t, "child1", tree.Children[5].Children[0].Line)
}

func TestBadIndentation(t *testing.T) {
	bytes, _ := ioutil.ReadFile("test/bad-indentation.txt")
	code := string(bytes)
	tree, err := codetree.New(code)

	assert.Nil(t, tree)
	assert.Error(t, err)
}

func BenchmarkCodeTree(b *testing.B) {
	bytes, _ := ioutil.ReadFile("example.txt")
	code := string(bytes)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree, err := codetree.New(code)

		if err != nil {
			b.Fail()
		}

		tree.Close()
	}
}
