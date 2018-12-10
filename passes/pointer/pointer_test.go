package pointer_test

import (
	"testing"

	"github.com/tenntenn/til/passes/pointer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, pointer.Analyzer, "a")
}
