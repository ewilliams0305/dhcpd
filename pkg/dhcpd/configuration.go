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