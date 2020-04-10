package returnfuncs

import (
	"io/ioutil"
	"os"
	"testing"
)

// setup returns teardown function so that caller can use it to perform cleanup
func setup(t *testing.T) (*os.File, func(), error) {
	teardown := func() {}

	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		return nil, teardown, err
	}

	teardown = func() {
		err := f.Close()
		if err != nil {
			t.Error("setup Close:", err)
		}

		err = os.RemoveAll(f.Name())
		if err != nil {
			t.Error("setup RemoveAll:", err)
		}
	}

	return f, teardown, nil
}

func testSomething(t *testing.T) {
	f, teardown, err := setup(t)
	defer teardown()
	if err != nil {
		t.Error("setup:", err)
	}

	if _, err = f.WriteString("hello"); err != nil {
		t.Error("file write:", err)
	}
}
