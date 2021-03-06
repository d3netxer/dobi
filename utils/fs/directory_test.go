package fs

import (
	"os"
	"testing"
	"time"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/assert/cmp"
	"github.com/gotestyourself/gotestyourself/fs"
)

func TestLastModified(t *testing.T) {
	tmpdir := fs.NewDir(t, "test-directory-last-modified",
		fs.WithDir("a"),
		fs.WithDir("b",
			fs.WithDir("c")))
	defer tmpdir.Remove()

	for index, dir := range []string{"a", "b", "b/c"} {
		mtime := time.Now().AddDate(0, 0, index+10)
		assert.Assert(t, cmp.Nil(touch(tmpdir.Join(dir, "file"), mtime)))

		actual, err := LastModified(tmpdir.Path())
		assert.NilError(t, err)
		assert.Equal(t, actual, mtime)
	}
}

func touch(name string, mtime time.Time) error {
	w, err := os.Create(name)
	if err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return os.Chtimes(name, mtime, mtime)
}
