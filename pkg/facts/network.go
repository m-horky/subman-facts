package facts

import (
	"encoding/json"
	"git.sr.ht/~spc/go-log"
	"os"
	"strings"
)

type ipRouteAddress struct {
	Family            string `json:"family"`
	Local             string `json:"local"`
	PrefixLength      int    `json:"prefixlen"`
	Scope             string `json:"scope"`
	Label             string `json:"label"`
	ValidLifeTime     int    `json:"valid_life_time"`
	PreferredLifeTime int    `json:"preferred_life_time"`
}

type ipRouteInterface struct {
	IfIndex             int              `json:"ifindex"`
	IfName              string           `json:"ifname"`
	Flags               []string         `json:"flags"`
	MTU                 int              `json:"mtu"`
	QDisc               string           `json:"qdisc"`
	OperationState      string           `json:"operstate"`
	Group               string           `json:"group"`
	TXQLen              int              `json:"txqlen"`
	LinkType            string           `json:"link_type"`
	MACAddress          string           `json:"address"`
	BroadcastMacAddress string           `json:"broadcast"`
	Addresses           []ipRouteAddress `json:"addr_info"`
}

type NetworkInterfaceFacts struct {
	MACAddress          string   `json:"mac_address"`
	IPv4Address         string   `json:"ipv4_address,omitempty"`
	IPv4Addresses       []string `json:"ipv4_address_list,omitempty"`
	IPv6GlobalAddress   string   `json:"ipv6_global_address,omitempty"`
	IPv6GlobalAddresses []string `json:"ipv6_global_address_list,omitempty"`
	IPv6LinkAddress     string   `json:"ipv6_link_address,omitempty"`
	IPv6LinkAddresses   []string `json:"ipv6_link_address_list,omitempty"`
}

type NetworkFacts struct {
	FQDN       string                           `json:"fqdn"`
	Hostname   string                           `json:"hostname"`
	Interfaces map[string]NetworkInterfaceFacts `json:"interfaces"`
}

// NetworkCollector contains facts from 'Network'
type NetworkCollector struct {
	data             NetworkFacts
	collected        bool
	getCommandOutput func(command string, args ...string) ([]string, []string, error)
}

func NewNetworkCollector() NetworkCollector {
	return NetworkCollector{
		data:             NetworkFacts{},
		collected:        false,
		getCommandOutput: getCommandOutput,
	}
}

// GetData collects 'ip' data and returns them as NetworkFacts.
func (c *NetworkCollector) GetData(rescan bool) (NetworkFacts, error) {
	if rescan || !c.collected {
		c.data = NetworkFacts{}
	}

	err := c.collect()
	if err != nil {
		return NetworkFacts{}, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *NetworkCollector) collect() error {
	_ = c.collectHostname()
	_ = c.collectFQDN()
	_ = c.collectIPRoute()
	c.collected = true
	return nil
}

func (c *NetworkCollector) collectHostname() error {
	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("could not collect hostname: %s", err)
		return err
	}
	c.data.Hostname = hostname
	return nil
}

func (c *NetworkCollector) collectFQDN() error {
	fullName, _, err := c.getCommandOutput("/usr/bin/hostname", "--fqdn")
	if err != nil {
		log.Errorf("could not collect fully qualified domain name: %s", err)
		return err
	}
	c.data.FQDN = fullName[0]
	return nil
}

func (c *NetworkCollector) collectIPRoute() error {
	stdout, _, err := c.getCommandOutput("/usr/sbin/ip", "--json", "address")
	if err != nil {
		log.Errorf("could not read output of ip: %s", err)
		return err
	}
	rawOutput := strings.Join(stdout, "\n")

	var ipOutput []ipRouteInterface
	err = json.Unmarshal([]byte(rawOutput), &ipOutput)
	if err != nil {
		log.Errorf("could not decode output of ip: %s", err)
		return err
	}

	ifaces := make(map[string]NetworkInterfaceFacts, 0)
	for _, ifaceData := range ipOutput {
		iface := NetworkInterfaceFacts{}
		iface.MACAddress = ifaceData.MACAddress

		var ipv4addresses []string
		var ipv6globalAddresses []string
		var ipv6linkAddresses []string
		for _, ifaceAddress := range ifaceData.Addresses {
			if ifaceAddress.Family == "inet" {
				ipv4addresses = append(ipv4addresses, ifaceAddress.Local)
			}
			if ifaceAddress.Family == "inet6" && ifaceAddress.Scope == "global" {
				ipv6globalAddresses = append(ipv6globalAddresses, ifaceAddress.Local)
			}
			if ifaceAddress.Family == "inet6" && ifaceAddress.Scope == "link" {
				ipv6linkAddresses = append(ipv6linkAddresses, ifaceAddress.Local)
			}
		}

		if len(ipv4addresses) > 0 {
			iface.IPv4Address = ipv4addresses[0]
			iface.IPv4Addresses = ipv4addresses
		}
		if len(ipv6globalAddresses) > 0 {
			iface.IPv6GlobalAddress = ipv6globalAddresses[0]
			iface.IPv6GlobalAddresses = ipv6globalAddresses
		}
		if len(ipv6linkAddresses) > 0 {
			iface.IPv6LinkAddress = ipv6linkAddresses[0]
			iface.IPv6LinkAddresses = ipv6linkAddresses
		}

		ifaces[ifaceData.IfName] = iface
	}
	c.data.Interfaces = ifaces
	return nil
}
