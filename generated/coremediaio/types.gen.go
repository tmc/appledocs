// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio
import (
	"unsafe"
)


// C struct types
// CMIODeviceAVCCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceAVCCommand
type CMIODeviceAVCCommand struct {
	MCommand unsafe.Pointer
	MCommandLength unsafe.Pointer
	MResponse unsafe.Pointer
	MResponseLength unsafe.Pointer
	MResponseUsed unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIODeviceAVCCommand */

// CMIODeviceRS422Command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceRS422Command
type CMIODeviceRS422Command struct {
	MCommand unsafe.Pointer
	MCommandLength unsafe.Pointer
	MResponse unsafe.Pointer
	MResponseLength unsafe.Pointer
	MResponseUsed unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIODeviceRS422Command */

// CMIODeviceSMPTETimeCallback
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceSMPTETimeCallback
type CMIODeviceSMPTETimeCallback struct {
	MGetSMPTETimeProc IODeviceGetSMPTETimeProc
	MRefCon unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIODeviceSMPTETimeCallback */

// CMIODeviceStreamConfiguration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStreamConfiguration
type CMIODeviceStreamConfiguration struct {
	MNumberChannels unsafe.Pointer
	MNumberStreams unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIODeviceStreamConfiguration */

// CMIOHardwarePlugInInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOHardwarePlugInInterface
type CMIOHardwarePlugInInterface struct {
	AddRef unsafe.Pointer
	DeviceProcessAVCCommand unsafe.Pointer
	DeviceProcessRS422Command unsafe.Pointer
	DeviceResume unsafe.Pointer
	DeviceStartStream unsafe.Pointer
	DeviceStopStream unsafe.Pointer
	DeviceSuspend unsafe.Pointer
	Initialize unsafe.Pointer
	InitializeWithObjectID unsafe.Pointer
	ObjectGetPropertyData unsafe.Pointer
	ObjectGetPropertyDataSize unsafe.Pointer
	ObjectHasProperty unsafe.Pointer
	ObjectIsPropertySettable unsafe.Pointer
	ObjectSetPropertyData unsafe.Pointer
	ObjectShow unsafe.Pointer
	QueryInterface unsafe.Pointer
	Release unsafe.Pointer
	StreamCopyBufferQueue unsafe.Pointer
	StreamDeckCueTo unsafe.Pointer
	StreamDeckJog unsafe.Pointer
	StreamDeckPlay unsafe.Pointer
	StreamDeckStop unsafe.Pointer
	Teardown unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIOHardwarePlugInInterface */

// CMIOObjectPropertyAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyAddress
type CMIOObjectPropertyAddress struct {
	MElement IOObjectPropertyElement
	MScope IOObjectPropertyScope
	MSelector IOObjectPropertySelector
}/* debug [types.gen.go/struct]: CMIOObjectPropertyAddress */

// CMIOStreamDeck
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeck
type CMIOStreamDeck struct {
	MState unsafe.Pointer
	MState2 unsafe.Pointer
	MStatus unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIOStreamDeck */

// CMIOStreamScheduledOutputNotificationProcAndRefCon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamScheduledOutputNotificationProcAndRefCon
type CMIOStreamScheduledOutputNotificationProcAndRefCon struct {
	ScheduledOutputNotificationProc IOStreamScheduledOutputNotificationProc
	ScheduledOutputNotificationRefCon unsafe.Pointer
}/* debug [types.gen.go/struct]: CMIOStreamScheduledOutputNotificationProcAndRefCon */





