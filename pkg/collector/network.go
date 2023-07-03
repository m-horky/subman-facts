package collector

import (
	"fmt"
	"git.sr.ht/~spc/go-log"
	"net"
	"net/netip"
	"os"
	"strings"
)

type IPAddressType int

const (
	IPUnspecifiedAddress IPAddressType = iota
	IPGlobalAddress
	IPLocalAddress
	IPHostAddress
)

type IPWithMask struct {
	Address netip.Addr
	Mask    string
	Version int
	Type    IPAddressType
}

// NetworkCollector contains facts from 'ip' command
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
	_ = c.collectHostname()
	_ = c.collectFQDN()

	// Uncomment to collect network.ipv4_address and network.ipv6_address
	/*_ = c.collectAddresses()*/

	err := c.collectInterfaces()
	if err != nil {
		log.Errorf("Could not collect network interfaces: %s", err)
	}

	return nil
}

func (c *NetworkCollector) collectHostname() error {
	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("Could not collect hostname: %s", err)
		return err
	}
	c.data["network.hostname"] = hostname
	return nil
}

func (c *NetworkCollector) collectFQDN() error {
	fullName, err := getCommandOutput("/usr/bin/hostname", "--fqdn,")
	if err != nil {
		log.Errorf("Could not collect fully qualified domain name: %s", err)
		return err
	}
	c.data["network.fqdn"] = fullName
	return nil
}

// getInterfaceAddresses inspect an interface and returns its v4 and v6 addresses.
func getInterfaceAddresses(iface net.Interface) ([]IPWithMask, error) {
	var addresses []IPWithMask

	addrs, _ := iface.Addrs()
	for _, addr := range addrs {
		addrWithoutMask, maskStr, ok := strings.Cut(addr.String(), "/")
		if !ok {
			// TODO Log the error
		}
		address := netip.MustParseAddr(addrWithoutMask)

		// Categorize the address
		var addressType IPAddressType
		if address.IsGlobalUnicast() {
			addressType = IPGlobalAddress
		} else if address.IsLoopback() {
			addressType = IPHostAddress
		} else if address.IsLinkLocalUnicast() || address.IsLinkLocalMulticast() {
			addressType = IPLocalAddress
		}

		if address.Is4() {
			addresses = append(addresses, IPWithMask{address, maskStr, 4, addressType})
		} else {
			addresses = append(addresses, IPWithMask{address, maskStr, 6, addressType})
		}
	}

	return addresses, nil
}

// collectAddresses gathers 'network.ipv4_address' and 'network.ipv6_address' facts.
func (c *NetworkCollector) collectAddresses() error {
	var v4Addresses []string
	var v4Masks []string
	var v6Addresses []string
	var v6Masks []string

	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		// Do not include loopback addresses in the list
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := getInterfaceAddresses(iface)
		if err != nil {
			return err
		}

		for _, address := range addresses {
			if address.Version == 4 {
				v4Addresses = append(v4Addresses, address.Address.String())
				v4Masks = append(v4Masks, address.Mask)
			} else {
				// Do not include link local addresses
				if address.Type == IPLocalAddress {
					continue
				}
				v6Addresses = append(v6Addresses, address.Address.String())
				v6Masks = append(v6Masks, address.Mask)
			}
		}

	}

	c.data["network.ipv4_address"] = strings.Join(v4Addresses, ", ")
	c.data["network.ipv6_address"] = strings.Join(v6Addresses, ", ")
	return nil
}

// collectInterfaces gathers all 'net.' facts.
func (c *NetworkCollector) collectInterfaces() error {
	// TODO Rewrite this to be 'ip'-based
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		prefix := fmt.Sprintf("net.interface.%s", iface.Name)

		ifaceMac := iface.HardwareAddr.String()
		if ifaceMac != "" {
			c.data[fmt.Sprintf("%s.mac_address", prefix)] = ifaceMac
		}

		addresses, err := getInterfaceAddresses(iface)
		if err != nil {
			// TODO Log the error
		}

		var v4a []string
		var v4m []string
		var v6ga []string
		var v6gm []string
		var v6la []string
		var v6lm []string
		var v6ha []string
		var v6hm []string

		for _, address := range addresses {
			if address.Version == 4 {
				v4a = append(v4a, address.Address.String())
				v4m = append(v4m, address.Mask)
			} else if address.Type == IPGlobalAddress {
				v6ga = append(v6ga, address.Address.String())
				v6gm = append(v6gm, address.Mask)
			} else if address.Type == IPLocalAddress {
				v6la = append(v6la, address.Address.String())
				v6lm = append(v6lm, address.Mask)
			} else if address.Type == IPHostAddress {
				v6ha = append(v6ha, address.Address.String())
				v6hm = append(v6hm, address.Mask)
			}
		}

		if len(v4a) > 0 {
			c.data[fmt.Sprintf("%s.ipv4_address", prefix)] = v4a[0]
			c.data[fmt.Sprintf("%s.ipv4_address_list", prefix)] = strings.Join(v4a, ", ")
		}
		/*if len(v4m) > 0 {
			c.data[fmt.Sprintf("%s.ipv4_netmask", prefix)] = v4m[0]
			c.data[fmt.Sprintf("%s.ipv4_netmask_list", prefix)] = strings.Join(v4m, ", ")
		}*/
		if len(v6ga) > 0 {
			c.data[fmt.Sprintf("%s.ipv6_address.global", prefix)] = v6ga[0]
			c.data[fmt.Sprintf("%s.ipv6_address.global_list", prefix)] = strings.Join(v6ga, ", ")
		}
		/*if len(v6gm) > 0 {
			c.data[fmt.Sprintf("%s.ipv6_netmask.global", prefix)] = v6gm[0]
			c.data[fmt.Sprintf("%s.ipv6_netmask.global_list", prefix)] = strings.Join(v6gm, ", ")
		}*/
		if len(v6la) > 0 {
			c.data[fmt.Sprintf("%s.ipv6_address.link", prefix)] = v6la[0]
			c.data[fmt.Sprintf("%s.ipv6_address.link_list", prefix)] = strings.Join(v6la, ", ")
		}
		/*if len(v6lm) > 0 {
			c.data[fmt.Sprintf("%s.ipv6_netmask.link", prefix)] = v6lm[0]
			c.data[fmt.Sprintf("%s.ipv6_netmask.link_list", prefix)] = strings.Join(v6lm, ", ")
		}*/
		/*if len(v6ha) > 0 {
			c.data[fmt.Sprintf("%s.ipv6_address.host", prefix)] = v6ha[0]
			c.data[fmt.Sprintf("%s.ipv6_address.host_list", prefix)] = strings.Join(v6ha, ", ")
		}*/
		/*if len(v6hm) > 0 {
			c.data[fmt.Sprintf("%s.ipv6_netmask.host", prefix)] = v6hm[0]
			c.data[fmt.Sprintf("%s.ipv6_netmask.host_list", prefix)] = strings.Join(v6hm, ", ")
		}*/

	}

	return nil
}
