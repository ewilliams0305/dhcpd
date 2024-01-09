package dhcpd

// The devices hardware ID used by the dhcp lease
type MacAddress string

// The ip address assigned to the device from the dhcp server
type IpAddress string

// Te ip netmask used for the dhcp server
type NetMask string

// The fully qualified hostname of the device.
type Hostname string

type DhcpOptions

// Duration for the lease
type LeaseTime int

// Represents a lease on the dhcp server.
type Lease struct {
	MacAddress
	Hostname
	IpAddress
}

// Represents a collection of static leases used by the server.
type StaticLeases []Lease

// Reresents a subnet declaration for an internal DHCP server LAN
type SubnetDeclaration struct {
	Subnet  IpAddress
	Netmask NetMask
	Start   IpAddress
	End     IpAddress
	Gateway IpAddress

	OptionDomainName       string
	OptionDnsServers       string
	OptionBroadcastAddress string

	DefaultLeaseTime LeaseTime
	MaxLeaseTime     LeaseTime
}

type DhcpConfig struct {
	subnet      *SubnetDeclaration
	staticLease *StaticLeases
}
