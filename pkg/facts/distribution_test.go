package facts

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDistributionCollector_collect(t *testing.T) {
	tests := []struct {
		description string
		read        func(path string) ([]string, error)
		wants       DistributionFacts
		wantsErr    error
	}{
		{
			description: "Fedora 37",
			read: func(path string) ([]string, error) {
				return read("./test_data/distribution-fc37-etc-osrelease")
			},
			wants:    DistributionFacts{Name: "Fedora Linux", Version: "37", ID: "Workstation Edition"},
			wantsErr: nil,
		},
		{
			description: "CentOS 7",
			read: func(path string) ([]string, error) {
				return read("./test_data/distribution-c7-etc-osrelease")
			},
			wants:    DistributionFacts{Name: "CentOS Linux", Version: "7", ID: "Core"},
			wantsErr: nil,
		},
		{
			description: "CentOS Stream 8",
			read: func(path string) ([]string, error) {
				return read("./test_data/distribution-cs8-etc-osrelease")
			},
			wants:    DistributionFacts{Name: "CentOS Stream", Version: "8", ID: "8"},
			wantsErr: nil,
		},
		{
			description: "CentOS Stream 9",
			read: func(path string) ([]string, error) {
				return read("./test_data/distribution-cs9-etc-osrelease")
			},
			wants:    DistributionFacts{Name: "CentOS Stream", Version: "9", ID: "9"},
			wantsErr: nil,
		},
		{
			description: "RHEL 8.8",
			read: func(path string) ([]string, error) {
				return read("./test_data/distribution-el88-etc-osrelease")
			},
			wants:    DistributionFacts{Name: "Red Hat Enterprise Linux", Version: "8.8", ID: "Ootpa"},
			wantsErr: nil,
		},
		{
			description: "RHEL 9.2",
			read: func(path string) ([]string, error) {
				return read("./test_data/distribution-el92-etc-osrelease")
			},
			wants:    DistributionFacts{Name: "Red Hat Enterprise Linux", Version: "9.2", ID: "Plow"},
			wantsErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			OS.Read = test.read

			collector := NewDistributionCollector()
			facts, err := collector.GetData(true)

			if test.wantsErr == nil {
				// We expect valid results
				if facts.Name != test.wants.Name {
					t.Errorf("Name expected as %s, got %s", test.wants.Name, facts.Name)
				}
				if facts.Version != test.wants.Version {
					t.Errorf("Version expected as %s, got %s", test.wants.Version, facts.Version)
				}
				if facts.ID != test.wants.ID {
					t.Errorf("ID expected as %s, got %s", test.wants.ID, facts.ID)
				}
			} else {
				// We expect failed results
				if !cmp.Equal(test.wantsErr.Error(), err.Error()) {
					t.Errorf("Error expected as %s, got %s", test.wantsErr, err)
				}
			}
		})
	}
}
