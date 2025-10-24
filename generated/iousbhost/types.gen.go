// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost


// C struct types
// IOUSBHostCIMessage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessage
type IOUSBHostCIMessage struct {
	Control uint32
	Data0 uint32
	Data1 uint64
}/* debug [types.gen.go/struct]: IOUSBHostCIMessage */

// IOUSBHostIOSourceDescriptors - The descriptors for a single endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSourceDescriptors
type IOUSBHostIOSourceDescriptors struct {
	BcdUSB uint16 // The USB version that the device supports.
	Descriptor USBEndpointDescriptor // The descriptor for a USB endpoint.
	SsCompanionDescriptor USBSuperSpeedEndpointCompanionDescriptor // The descriptor for a SuperSpeed USB endpoint companion.
	SspCompanionDescriptor USBSuperSpeedPlusIsochronousEndpointCompanionDescriptor // The descriptor for a SuperSpeedPlus isochronous USB endpoint companion.
}/* debug [types.gen.go/struct]: IOUSBHostIOSourceDescriptors */

// IOUSBHostIsochronousFrame - A structure that represents a single frame in an isochronous transfer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousFrame
type IOUSBHostIsochronousFrame struct {
	CompleteCount uint32 // The number of bytes that the system actually transferred for the frame.
	RequestCount uint32 // The number of requested bytes to transfer for the frame.
	Reserved uint32
	Status int // The completion status for an individual frame.
	TimeStamp USBHostTime // The observed time for the frame’s completion.
}/* debug [types.gen.go/struct]: IOUSBHostIsochronousFrame */

// IOUSBHostIsochronousTransaction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransaction
type IOUSBHostIsochronousTransaction struct {
	CompleteCount uint32
	Offset uint32
	Options USBHostIsochronousTransactionOptions
	RequestCount uint32
	Status int
	TimeStamp USBHostTime
}/* debug [types.gen.go/struct]: IOUSBHostIsochronousTransaction */





