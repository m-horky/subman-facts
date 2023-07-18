package facts

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"os"
	"strings"
	"testing"
)

func TestVirtCollector_collect(t *testing.T) {
	tests := []struct {
		description      string
		getFileOutput    func(path string) ([]string, error)
		getCommandOutput func(cmd string, args ...string) ([]string, []string, error)
		collected        bool
		ok               VirtFacts
		err              error
	}{
		{
			description: "bare metal",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				return []string{}, []string{}, nil
			},
			getFileOutput: func(path string) ([]string, error) {
				return []string{}, fmt.Errorf("getFileOutput is not mocked")
			},
			collected: true,
			ok:        VirtFacts{IsGuest: false, HostType: "Not Applicable", UUID: ""},
			err:       nil,
		},
		{
			description: "non-zero error code",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				return []string{"virt-what: unrecognized option 'foo'"}, []string{}, fmt.Errorf("exit status 1")
			},
			getFileOutput: func(path string) ([]string, error) {
				return []string{}, fmt.Errorf("getFileOutput is not mocked")
			},
			collected: false,
			ok:        VirtFacts{}, // FIXME sub-man reports 'is_guest: "Unknown"'
			err:       nil,
		},
		{
			description: "made-up virtualization",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				return []string{"made-up"}, []string{}, nil
			},
			getFileOutput: func(path string) ([]string, error) {
				return []string{}, fmt.Errorf("mock is missing for path %s", path)
			},
			collected: true,
			ok:        VirtFacts{IsGuest: true, HostType: "made-up", UUID: ""},
			err:       nil,
		},
		{
			description: "kvm",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/virt-what":
					return []string{"kvm"}, []string{}, nil
				case "/usr/sbin/dmidecode":
					lines, _ := getFileOutput("./test_data/virt-kvm-dmidecode")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			getFileOutput: func(path string) ([]string, error) {
				return []string{}, fmt.Errorf("mock is missing for path %s", path)
			},
			collected: true,
			ok:        VirtFacts{IsGuest: true, HostType: "kvm", UUID: "fake-uuid"},
			err:       nil,
		},
		{
			description: "ppc64le (vm,uuid)",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/virt-what":
					return []string{"ibm_power-lpar_dedicated"}, []string{}, nil
				case "/usr/sbin/dmidecode":
					return []string{}, []string{}, os.ErrNotExist
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			getFileOutput: func(path string) ([]string, error) {
				switch path {
				case "/proc/device-tree/vm,uuid":
					return []string{"fake-uuid"}, nil
				case "/proc/device-tree/ibm,partition-uuid":
					return []string{}, os.ErrNotExist
				case "/sys/hypervisor/uuid":
					return []string{}, os.ErrNotExist
				default:
					return []string{}, fmt.Errorf("mock is missing for %s", path)
				}
			},
			collected: true,
			ok:        VirtFacts{IsGuest: true, HostType: "ibm_power-lpar_dedicated", UUID: "fake-uuid"},
			err:       nil,
		},
		{
			description: "ppc64le (ibm,partition-uuid)",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/virt-what":
					return []string{"ibm_power-lpar_dedicated"}, []string{}, nil
				case "/usr/sbin/dmidecode":
					return []string{}, []string{}, os.ErrNotExist
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			getFileOutput: func(path string) ([]string, error) {
				switch path {
				case "/proc/device-tree/vm,uuid":
					return []string{}, os.ErrNotExist
				case "/proc/device-tree/ibm,partition-uuid":
					return []string{"fake-uuid"}, nil
				case "/sys/hypervisor/uuid":
					return []string{}, os.ErrNotExist
				default:
					return []string{}, fmt.Errorf("mock is missing for %s", path)
				}
			},
			collected: true,
			ok:        VirtFacts{IsGuest: true, HostType: "ibm_power-lpar_dedicated", UUID: "fake-uuid"},
			err:       nil,
		},
		{
			description: "xen",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/virt-what":
					return []string{"xen"}, []string{}, nil
				case "/usr/sbin/dmidecode":
					return []string{}, []string{}, os.ErrNotExist
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			getFileOutput: func(path string) ([]string, error) {
				switch path {
				case "/proc/device-tree/vm,uuid":
					return []string{}, os.ErrNotExist
				case "/proc/device-tree/ibm,partition-uuid":
					return []string{}, os.ErrNotExist
				case "/sys/hypervisor/uuid":
					return []string{"fake-uuid"}, nil
				default:
					return []string{}, fmt.Errorf("mock is missing for %s", path)
				}
			},
			collected: true,
			ok:        VirtFacts{IsGuest: true, HostType: "xen", UUID: "fake-uuid"},
			err:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			collector := NewVirtCollector()
			collector.getFileOutput = test.getFileOutput
			collector.getCommandOutput = test.getCommandOutput

			facts, err := collector.GetData(true)

			if collector.collected != test.collected {
				t.Errorf("collector.collected expected as %t, got %t", test.collected, collector.collected)
				t.FailNow()
			}

			if test.err == nil {
				// We expect valid results
				if facts.IsGuest != test.ok.IsGuest {
					t.Errorf("IsGuest expected as %t, got %t", test.ok.IsGuest, facts.IsGuest)
				}
				if facts.HostType != test.ok.HostType {
					t.Errorf("HostType expected as %s, got %s", test.ok.HostType, facts.HostType)
				}
				if facts.UUID != test.ok.UUID {
					t.Errorf("UUID expected as %s, got %s", test.ok.UUID, facts.UUID)
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
