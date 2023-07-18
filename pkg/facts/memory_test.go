package facts

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMemoryCollector_collect(t *testing.T) {
	tests := []struct {
		description   string
		getFileOutput func(path string) ([]string, error)
		collected     bool
		ok            MemoryFacts
		err           error
	}{
		{
			description: "simplified input",
			getFileOutput: func(path string) ([]string, error) {
				return []string{"MemTotal:   4096 kB", "SwapTotal:   2048 kB"}, nil
			},
			collected: true,
			ok:        MemoryFacts{MemTotal: 4096, SwapTotal: 2048},
			err:       nil,
		},
		{
			description: "real life input",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/memory-laptop")
			},
			collected: true,
			ok:        MemoryFacts{MemTotal: 32604020, SwapTotal: 8388604},
			err:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			collector := NewMemoryCollector()
			collector.getFileOutput = test.getFileOutput

			facts, err := collector.GetData(true)

			if collector.collected != test.collected {
				t.Errorf("collector.collected expected as %t, got %t", test.collected, collector.collected)
				t.FailNow()
			}

			if test.err == nil {
				// We expect valid results
				if facts.MemTotal != test.ok.MemTotal {
					t.Errorf("MemTotal expected as %d, got %d", test.ok.MemTotal, facts.MemTotal)
				}
				if facts.SwapTotal != test.ok.SwapTotal {
					t.Errorf("SwapTotal expected as %d, got %d", test.ok.SwapTotal, facts.SwapTotal)
				}
			} else {
				// We expect failed results
				if !cmp.Equal(test.err, err) {
					t.Errorf("Error expected as %s, got %s", test.err, err)
				}
			}
		})
	}
}
