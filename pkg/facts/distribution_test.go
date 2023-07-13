package facts

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDistributionCollector_collect(t *testing.T) {
	tests := []struct {
		description   string
		getFileOutput func(path string) ([]string, error)
		ok            DistributionFacts
		err           error
	}{
		{
			description: "Fedora 37",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/distribution-fc37-etc-osrelease")
			},
			ok:  DistributionFacts{Name: "Fedora Linux", Version: "37", ID: "Workstation Edition"},
			err: nil,
		},
		{
			description: "CentOS 7",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/distribution-c7-etc-osrelease")
			},
			ok:  DistributionFacts{Name: "CentOS Linux", Version: "7", ID: "Core"},
			err: nil,
		},
		{
			description: "CentOS Stream 8",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/distribution-cs8-etc-osrelease")
			},
			ok:  DistributionFacts{Name: "CentOS Stream", Version: "8", ID: "8"},
			err: nil,
		},
		{
			description: "CentOS Stream 9",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/distribution-cs9-etc-osrelease")
			},
			ok:  DistributionFacts{Name: "CentOS Stream", Version: "9", ID: "9"},
			err: nil,
		},
		{
			description: "RHEL 8.8",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/distribution-el88-etc-osrelease")
			},
			ok:  DistributionFacts{Name: "Red Hat Enterprise Linux", Version: "8.8", ID: "Ootpa"},
			err: nil,
		},
		{
			description: "RHEL 9.2",
			getFileOutput: func(path string) ([]string, error) {
				return getFileOutput("./test_data/distribution-el92-etc-osrelease")
			},
			ok:  DistributionFacts{Name: "Red Hat Enterprise Linux", Version: "9.2", ID: "Plow"},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			collector := NewDistributionCollector()
			collector.getFileOutput = test.getFileOutput

			facts, err := collector.GetData(true)

			if test.err == nil {
				// We expect valid results
				if facts.Name != test.ok.Name {
					t.Errorf("Name expected as %s, got %s", test.ok.Name, facts.Name)
				}
				if facts.Version != test.ok.Version {
					t.Errorf("Version expected as %s, got %s", test.ok.Version, facts.Version)
				}
				if facts.ID != test.ok.ID {
					t.Errorf("ID expected as %s, got %s", test.ok.ID, facts.ID)
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
