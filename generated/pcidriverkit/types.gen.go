// Code generated from Apple documentation for PCIDriverKit. DO NOT EDIT.

package pcidriverkit

// C struct types
// IOPCIDevice - A DriverKit provider object that manages access to your custom PCI hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice
type IOPCIDevice struct {
} /* debug [types.gen.go/struct]: IOPCIDevice */

// ClientCrashed
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ClientCrashed
type ClientCrashed struct {
} /* debug [types.gen.go/struct]: ClientCrashed */

// Close - Closes the session to the PCI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/Close
type Close struct {
} /* debug [types.gen.go/struct]: Close */

// ConfigurationRead16 - Reads a 16-bit data value synchronously from the device’s configuration space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigurationRead16
type ConfigurationRead16 struct {
} /* debug [types.gen.go/struct]: ConfigurationRead16 */

// ConfigurationRead32 - Reads a 32-bit data value synchronously from the device’s configuration space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigurationRead32
type ConfigurationRead32 struct {
} /* debug [types.gen.go/struct]: ConfigurationRead32 */

// ConfigurationRead8 - Reads an 8-bit data value synchronously from the device’s configuration space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigurationRead8
type ConfigurationRead8 struct {
} /* debug [types.gen.go/struct]: ConfigurationRead8 */

// ConfigurationWrite16 - Writes an 16-bit data value to the device’s configuration space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigurationWrite16
type ConfigurationWrite16 struct {
} /* debug [types.gen.go/struct]: ConfigurationWrite16 */

// ConfigurationWrite32 - Writes an 32-bit data value to the device’s configuration space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigurationWrite32
type ConfigurationWrite32 struct {
} /* debug [types.gen.go/struct]: ConfigurationWrite32 */

// ConfigurationWrite8 - Writes an 8-bit data value to the device’s configuration space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigurationWrite8
type ConfigurationWrite8 struct {
} /* debug [types.gen.go/struct]: ConfigurationWrite8 */

// ConfigureInterrupts
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/ConfigureInterrupts
type ConfigureInterrupts struct {
} /* debug [types.gen.go/struct]: ConfigureInterrupts */

// EnablePCIPowerManagement - Configures the device’s PCI bus power management capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/EnablePCIPowerManagement
type EnablePCIPowerManagement struct {
} /* debug [types.gen.go/struct]: EnablePCIPowerManagement */

// FindPCICapability - Search the configuration space for a PCI capability register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/FindPCICapability
type FindPCICapability struct {
} /* debug [types.gen.go/struct]: FindPCICapability */

// free - Performs any final cleanup for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/free
type free struct {
} /* debug [types.gen.go/struct]: free */

// GetBARInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/GetBARInfo
type GetBARInfo struct {
} /* debug [types.gen.go/struct]: GetBARInfo */

// GetBusDeviceFunction - Returns the device’s bus, device, and function numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/GetBusDeviceFunction
type GetBusDeviceFunction struct {
} /* debug [types.gen.go/struct]: GetBusDeviceFunction */

// GetLinkSpeed
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/GetLinkSpeed
type GetLinkSpeed struct {
} /* debug [types.gen.go/struct]: GetLinkSpeed */

// HasPCIPowerManagement - Determines whether the device has the specified PCI bus power management capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/HasPCIPowerManagement
type HasPCIPowerManagement struct {
} /* debug [types.gen.go/struct]: HasPCIPowerManagement */

// init - Initializes the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/init
type init struct {
} /* debug [types.gen.go/struct]: init */

// MemoryRead16 - Reads a 16-bit value synchronously from the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryRead16-50bq8
type MemoryRead16 struct {
} /* debug [types.gen.go/struct]: MemoryRead16 */

// MemoryRead32 - Reads a 32-bit value synchronously from the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryRead32-60hg9
type MemoryRead32 struct {
} /* debug [types.gen.go/struct]: MemoryRead32 */

// MemoryRead64 - Reads a 64-bit value synchronously from the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryRead64-37uob
type MemoryRead64 struct {
} /* debug [types.gen.go/struct]: MemoryRead64 */

// MemoryRead8 - Reads a 8-bit value synchronously from the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryRead8-1edw0
type MemoryRead8 struct {
} /* debug [types.gen.go/struct]: MemoryRead8 */

// MemoryWrite16 - Writes an 16-bit value to the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryWrite16-534yk
type MemoryWrite16 struct {
} /* debug [types.gen.go/struct]: MemoryWrite16 */

// MemoryWrite32 - Writes an 32-bit value to the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryWrite32-4pmh
type MemoryWrite32 struct {
} /* debug [types.gen.go/struct]: MemoryWrite32 */

// MemoryWrite64 - Writes an 64-bit value to the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryWrite64-8qyob
type MemoryWrite64 struct {
} /* debug [types.gen.go/struct]: MemoryWrite64 */

// MemoryWrite8 - Writes an 8-bit value to the PCI device’s aperture at the specified memory index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/MemoryWrite8-1uey6
type MemoryWrite8 struct {
} /* debug [types.gen.go/struct]: MemoryWrite8 */

// Open - Opens a session to the PCI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/Open
type Open struct {
} /* debug [types.gen.go/struct]: Open */

// Reset
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/Reset
type Reset struct {
} /* debug [types.gen.go/struct]: Reset */

// RestoreDeviceState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/RestoreDeviceState
type RestoreDeviceState struct {
} /* debug [types.gen.go/struct]: RestoreDeviceState */

// SaveDeviceState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/SaveDeviceState
type SaveDeviceState struct {
} /* debug [types.gen.go/struct]: SaveDeviceState */

// SetASPMState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/SetASPMState
type SetASPMState struct {
} /* debug [types.gen.go/struct]: SetASPMState */

// SetLinkSpeed
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/SetLinkSpeed
type SetLinkSpeed struct {
} /* debug [types.gen.go/struct]: SetLinkSpeed */

// SetProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIDevice/SetProperties
type SetProperties struct {
} /* debug [types.gen.go/struct]: SetProperties */
