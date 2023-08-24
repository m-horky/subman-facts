package facts

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestNetworkCollector_collectIPRoute(t *testing.T) {
	tests := []struct {
		description string
		run         func(cmd string, args ...string) ([]string, []string, error)
		wants       map[string]NetworkInterfaceFacts
		wantsErr    error
	}{
		{
			description: "no iproute",
			run: func(cmd string, args ...string) ([]string, []string, error) {
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
			run: func(cmd string, args ...string) ([]string, []string, error) {
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
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-loopback.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"lo": {
					MACAddress:    "00:00:00:00:00:00",
					IPv4Address:   "127.0.0.1",
					IPv4Addresses: []string{"127.0.0.1"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "more IPv4 addresses on one interface",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-more-ipv4s.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"fake0": {
					MACAddress:    "00:11:22:33:44:55",
					IPv4Address:   "10.0.1.10",
					IPv4Addresses: []string{"10.0.1.10", "10.0.2.247"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "more IPv6 addresses on one interface",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-more-ipv6s.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"fake0": {
					MACAddress:          "00:11:22:33:44:55",
					IPv6GlobalAddress:   "2620:52:0:0:0:1:2:3",
					IPv6GlobalAddresses: []string{"2620:52:0:0:0:1:2:3", "2620:52:0:0:0:1:2:4"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "system with no internet connection",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-no-connection.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"lo": {
					MACAddress:    "00:00:00:00:00:00",
					IPv4Address:   "127.0.0.1",
					IPv4Addresses: []string{"127.0.0.1"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "system with 🍉 emoji as an interface name",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-emoji.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"🍉": {
					MACAddress:        "00:11:22:33:44:55",
					IPv4Address:       "10.0.0.2",
					IPv4Addresses:     []string{"10.0.0.2"},
					IPv6LinkAddress:   "fe80::1:2",
					IPv6LinkAddresses: []string{"fe80::1:2"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "laptop data",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-laptop.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"lo": {
					MACAddress:    "00:00:00:00:00:00",
					IPv4Address:   "127.0.0.1",
					IPv4Addresses: []string{"127.0.0.1"},
				},
				"enp0s20f0u2u1": {
					MACAddress:          "a4:ae:12:01:02:03",
					IPv4Address:         "10.0.0.56",
					IPv4Addresses:       []string{"10.0.0.56"},
					IPv6GlobalAddress:   "2620:52:0:0:0:1:2:3",
					IPv6GlobalAddresses: []string{"2620:52:0:0:0:1:2:3"},
					IPv6LinkAddress:     "fe80::0:1:2:3",
					IPv6LinkAddresses:   []string{"fe80::0:1:2:3"},
				},
				"enp0s31f6": {
					MACAddress: "38:f3:ab:01:02:03",
				},
				"wlp0s20f3": {
					MACAddress:          "d6:73:ed:01:02:03",
					PermanentMACAddress: "28:d0:ea:01:02:03",
				},
				"virbr0": {
					MACAddress:          "52:54:00:01:02:03",
					IPv4Address:         "192.168.122.1",
					IPv4Addresses:       []string{"192.168.122.1"},
					IPv6GlobalAddress:   "2038:dead:beef::1",
					IPv6GlobalAddresses: []string{"2038:dead:beef::1"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "bonding of two interfaces",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-bonding.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"lo": {
					MACAddress:    "00:00:00:00:00:00",
					IPv4Address:   "127.0.0.1",
					IPv4Addresses: []string{"127.0.0.1"},
				},
				"enp1s0": {
					PermanentMACAddress: "52:54:00:00:00:01",
					MACAddress:          "72:4b:f3:aa:aa:aa",
				},
				"enp2s0": {
					PermanentMACAddress: "52:54:00:00:00:02",
					MACAddress:          "72:4b:f3:aa:aa:aa",
				},
				"bond0": {
					MACAddress:          "72:4b:f3:aa:aa:aa",
					IPv4Address:         "192.168.122.35",
					IPv4Addresses:       []string{"192.168.122.35"},
					IPv6GlobalAddress:   "2038:dead:beef::1",
					IPv6GlobalAddresses: []string{"2038:dead:beef::1"},
					IPv6LinkAddress:     "fe80::1",
					IPv6LinkAddresses:   []string{"fe80::1"},
				},
			},
			wantsErr: nil,
		},
		{
			description: "teaming of two interfaces",
			run: func(cmd string, args ...string) ([]string, []string, error) {
				fullCmd := strings.TrimRight(fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")), " ")
				switch fullCmd {
				case "/usr/sbin/ip --json address":
					lines, _ := read("./test_data/network-teaming.json")
					return lines, []string{}, nil
				default:
					return []string{}, []string{}, fmt.Errorf("mock is missing for cmd %s", fullCmd)
				}
			},
			wants: map[string]NetworkInterfaceFacts{
				"lo": {
					MACAddress:    "00:00:00:00:00:00",
					IPv4Address:   "127.0.0.1",
					IPv4Addresses: []string{"127.0.0.1"},
				},
				"enp1s0": {
					MACAddress:          "52:54:00:dd:19:07",
					PermanentMACAddress: "",
					IPv4Address:         "192.168.122.204",
					IPv4Addresses:       []string{"192.168.122.204"},
					IPv6GlobalAddress:   "2038:dead:beef::1a9",
					IPv6GlobalAddresses: []string{"2038:dead:beef::1a9"},
					IPv6LinkAddress:     "fe80::5054:ff:fedd:1907",
					IPv6LinkAddresses:   []string{"fe80::5054:ff:fedd:1907"},
				},
				"enp7s0": {
					MACAddress: "52:54:00:00:00:01",
				},
				"enp8s0": {
					PermanentMACAddress: "52:54:00:00:00:02",
					MACAddress:          "52:54:00:00:00:01",
				},
				"team0": {
					MACAddress:          "52:54:00:00:00:01",
					IPv4Address:         "192.168.122.35",
					IPv4Addresses:       []string{"192.168.122.35"},
					IPv6GlobalAddress:   "2038:dead:beef::1",
					IPv6GlobalAddresses: []string{"2038:dead:beef::1"},
					IPv6LinkAddress:     "fe80::1",
					IPv6LinkAddresses:   []string{"fe80::1"},
				},
			},
			wantsErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			OS.Run = test.run

			collector := NewNetworkCollector()
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
