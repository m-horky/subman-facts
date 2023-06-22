package collector

import (
	"github.com/Showmax/go-fqdn"
	"net"
	"net/netip"
	"os"
	"os/exec"
	"strings"
)

type NetworkCollector struct {
	data map[string]string
}

func (c *NetworkCollector) String() string {
	return "Network collector"
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c *NetworkCollector) Flush() {
	c.data = make(map[string]string)
}

// GetData collects network interface data.
func (c *NetworkCollector) GetData() (map[string]string, error) {
	if c.data == nil {
		c.Flush()
	}
	if len(c.data) == 0 {
		err := c.collect()
		if err != nil {
			return nil, err
		}
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *NetworkCollector) collect() error {
	err := c.collectHostname()
	if err != nil {
		// TODO Log the error
	}
	err = c.collectFQDN_lib()
	if err != nil {
		// TODO Log the error
	}
	err = c.collectFQDN_shell()
	if err != nil {
		// TODO Log the error
	}
	err = c.collectAddresses()
	if err != nil {
		// TODO Log the error
	}

	return nil
}

func (c *NetworkCollector) collectHostname() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	c.data["network.hostname"] = hostname
	return nil
}

func (c *NetworkCollector) collectFQDN_lib() error {
	fullName, err := fqdn.FqdnHostname()
	if err != nil {
		return err
	}
	c.data["network.fqdn(lib)"] = fullName
	return nil
}

func (c *NetworkCollector) collectFQDN_shell() error {
	fullRawName, err := exec.Command("/usr/bin/hostname", "-f").Output()
	if err != nil {
		return err
	}
	fullName := strings.TrimRight(string(fullRawName), "\n")
	c.data["network.fqdn(shell)"] = fullName
	return nil
}

func (c *NetworkCollector) collectAddresses() error {
	var ipv4s = []string{}
	var ipv6s = []string{}

	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		// Do not include loopback addresses in the list
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			addrWithoutMask := strings.Split(addr.String(), "/")[0]
			address := netip.MustParseAddr(addrWithoutMask)

			if address.Is4() {
				ipv4s = append(ipv4s, address.String())
			} else {
				// Do not include link local addresses
				if address.IsLinkLocalUnicast() || address.IsLinkLocalMulticast() {
					continue
				}
				ipv6s = append(ipv6s, address.String())
			}
		}
	}

	c.data["network.ipv4_address"] = strings.Join(ipv4s, ", ")
	c.data["network.ipv6_address"] = strings.Join(ipv6s, ", ")
	return nil
}
