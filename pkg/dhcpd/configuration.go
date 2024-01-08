package dhcpd

const (
	filePath string = "/etc/dhcp/dhcpd.conf"
)

var config *DhcpConfig

func saveFile() error {
	return nil
}

func readFile() (*DhcpConfig, error) {
	config = &DhcpConfig{}
	return config, nil
}

func (scope *SubnetDeclaration) createScopeSection() (string, error) {
	return `
  subnet 10.5.5.0 netmask 255.255.255.224 {
  range 10.5.5.26 10.5.5.30;
  option domain-name-servers ns1.internal.example.org;
  option domain-name "internal.example.org";
  option routers 10.5.5.1;
  option broadcast-address 10.5.5.31;
  default-lease-time 600;
  max-lease-time 7200;
  }`, nil
}
