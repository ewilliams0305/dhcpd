package dhcpd

const (
	// Starts the dhcpd service
	start Action = iota
	// Stops the dhcpd service
	stop Action = 1
	// Restarts the dhcpd service
	restart Action = 2
	// Enables the dhcpd service
	enable Action = 3
	// Disables the dhcpd service
	disable Action = 4
	// Returns the status of the dhcpd service
	status Action = 5
)

// The type of action to invoke on the systemd dhcpd service
type Action int

// The result of an action invoked on the systemd dhcp service
type ActionResult struct {
	// Action invoked
	action Action
	// Possibly err returned.
	err error
}

// Creates a new action response from the action
func newActionResult(action Action, err error) ActionResult {
	return ActionResult{
		action: action,
		err:    err,
	}
}

// Creates a result from the action send to the service
func (a *Action) Result(err error) ActionResult {
	return ActionResult{
		action: *a,
		err:    err,
	}
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
