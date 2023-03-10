package slinter_test

import (
	"testing"

	"github.com/aoman-n/slinter"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	analyzer := slinter.Analyzer
	analyzer.Flags.Set("maxLines", "7")

	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, slinter.Analyzer, "a")
}
