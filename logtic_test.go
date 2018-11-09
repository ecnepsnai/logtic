package logtic

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"testing"
)

var file *File
var source *Source

func setupTest() {
	dir, err := ioutil.TempDir("", "logtic")
	if err != nil {
		fmt.Printf("Unable to create temporary directory: %s\n", err.Error())
		os.Exit(1)
	}
	f, s, err := New(path.Join(dir, "app.log"), LevelDebug, "logtic")
	if err != nil {
		fmt.Printf("Unable to create new logtic instance: %s\n", err.Error())
		os.Exit(1)
	}
	file = f
	source = s
}

func testdownTest() {
	file.Close()
}

func TestMain(m *testing.M) {
	setupTest()
	retCode := m.Run()
	testdownTest()
	os.Exit(retCode)
}

func TestParallelWrite(t *testing.T) {
	var wg sync.WaitGroup

	source.Info("TestParallelWrite")

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := Connect("goroutine1")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := Connect("goroutine2")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := Connect("goroutine3")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()

	wg.Wait()
}
