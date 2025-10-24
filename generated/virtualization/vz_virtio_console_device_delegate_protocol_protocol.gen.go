// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

// PVZVirtioConsoleDeviceDelegate is the VZVirtioConsoleDeviceDelegate protocol interface.
//
// Optional methods that you use to respond when a console port opens or closes in the virtual machine.
//
// Availability:
//   - macOS 13.0+
//
// See: doc://com.apple.virtualization/documentation/Virtualization/VZVirtioConsoleDeviceDelegate
type PVZVirtioConsoleDeviceDelegate interface {
	// Optional methods
	ConsoleDeviceDidClosePort(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort)
	HasConsoleDeviceDidClosePort() bool
	ConsoleDeviceDidOpenPort(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort)
	HasConsoleDeviceDidOpenPort() bool
}

// VZVirtioConsoleDeviceDelegate is a delegate implementation builder for the PVZVirtioConsoleDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type VZVirtioConsoleDeviceDelegate struct {
	_ConsoleDeviceDidClosePort func(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort)
	_ConsoleDeviceDidOpenPort  func(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort)
}

// SetConsoleDeviceDidClosePort sets the handler for the ConsoleDeviceDidClosePort delegate method.
//
// Tells the delegate that the framework closed a console port.
func (d *VZVirtioConsoleDeviceDelegate) SetConsoleDeviceDidClosePort(f func(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort)) {
	d._ConsoleDeviceDidClosePort = f
}

// SetConsoleDeviceDidOpenPort sets the handler for the ConsoleDeviceDidOpenPort delegate method.
//
// Tells the delegate that the framework opened a console port.
func (d *VZVirtioConsoleDeviceDelegate) SetConsoleDeviceDidOpenPort(f func(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort)) {
	d._ConsoleDeviceDidOpenPort = f
}

// ConsoleDeviceDidClosePort implements the PVZVirtioConsoleDeviceDelegate interface.
func (d *VZVirtioConsoleDeviceDelegate) ConsoleDeviceDidClosePort(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort) {
	if d._ConsoleDeviceDidClosePort != nil {
		d._ConsoleDeviceDidClosePort(consoleDevice, consolePort)
	}
}

// HasConsoleDeviceDidClosePort returns true if a handler for ConsoleDeviceDidClosePort has been set.
func (d *VZVirtioConsoleDeviceDelegate) HasConsoleDeviceDidClosePort() bool {
	return d._ConsoleDeviceDidClosePort != nil
}

// ConsoleDeviceDidOpenPort implements the PVZVirtioConsoleDeviceDelegate interface.
func (d *VZVirtioConsoleDeviceDelegate) ConsoleDeviceDidOpenPort(consoleDevice IVZVirtioConsoleDevice, consolePort IVZVirtioConsolePort) {
	if d._ConsoleDeviceDidOpenPort != nil {
		d._ConsoleDeviceDidOpenPort(consoleDevice, consolePort)
	}
}

// HasConsoleDeviceDidOpenPort returns true if a handler for ConsoleDeviceDidOpenPort has been set.
func (d *VZVirtioConsoleDeviceDelegate) HasConsoleDeviceDidOpenPort() bool {
	return d._ConsoleDeviceDidOpenPort != nil
}
