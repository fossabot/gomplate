package funcs

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	fs := afero.NewMemMapFs()
	ff := &FileFuncs{fs}

	_ = fs.Mkdir("/tmp", 0777)
	f, _ := fs.Create("/tmp/foo")
	_, _ = f.Write([]byte("foo"))

	assert.True(t, ff.Exists("/tmp/foo"))
	assert.False(t, ff.Exists("/tmp/bar"))
}

func TestFileIsDir(t *testing.T) {
	fs := afero.NewMemMapFs()
	ff := &FileFuncs{fs}

	_ = fs.Mkdir("/tmp", 0777)
	f, _ := fs.Create("/tmp/foo")
	_, _ = f.Write([]byte("foo"))

	assert.True(t, ff.IsDir("/tmp"))
	assert.False(t, ff.IsDir("/tmp/foo"))
}
