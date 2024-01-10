package dhcpd

const (
	start   Action = iota
	stop    Action = 1
	restart Action = 2
	enable  Action = 3
	disable Action = 4
	status  Action = 5
)

// The type of action to invoke on the systemd dhcpd service
type Action int

// The result of an action invoked on the systemd dhcp service
type ActionResult struct {
	// Action invoked
	action Action
	// Possibly error returned.
	error error
}

// Converts the action to a string value used by the .exec command
func (a *Action) String() string {
	switch *a {
	case start:
		return "start"
	case stop:
		return "stop"
	case restart:
		return "restart"
	case enable:
		return "enable"
	case disable:
		return "disable"
	case status:
		return "status"
	}
	return "error"
}
