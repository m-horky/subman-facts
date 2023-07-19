package facts

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestNetworkCollector_collectIPRoute(t *testing.T) {
	tests := []struct {
		description      string
		getCommandOutput func(cmd string, args ...string) ([]string, []string, error)
		wants            map[string]NetworkInterfaceFacts
		wantsErr         error
	}{
		{
			description: "no iproute",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					return []string{}, []string{}, fmt.Errorf("no such file or directory")
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants:    map[string]NetworkInterfaceFacts(nil),
			wantsErr: nil,
		},
		{
			description: "empty",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					return []string{"[]"}, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants:    map[string]NetworkInterfaceFacts{},
			wantsErr: nil,
		},
		{
			description: "only loopback",
			getCommandOutput: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := getFileOutput("./test_data/network-loopback.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"lo": {
					MACAddress:    "",
					IPv4Address:   "127.0.0.1",
					IPv4Addresses: []string{"127.0.0.1"},
				},
			},
			wantsErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			collector := NewNetworkCollector()
			collector.getCommandOutput = test.getCommandOutput

			err := collector.collectIPRoute()
			interfaces := collector.data.Interfaces

			if test.wantsErr == nil {
				// We expect valid results
				if !cmp.Equal(test.wants, interfaces) {
					t.Errorf("Expected %#v, got %#v", test.wants, interfaces)
				}
			} else {
				// We expect failed results
				if !cmp.Equal(test.wantsErr, err) {
					t.Errorf("Error expected as %s, got %s", test.wantsErr, err)
				}
			}
		})
	}
}
