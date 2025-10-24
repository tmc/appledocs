// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio
import (
	"unsafe"

	"github.com/tmc/appledocs/generated/coreaudiotypes"
)


// C struct types
// AudioDriverPlugInHostInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInHostInfo
type AudioDriverPlugInHostInfo struct {
	MDeviceID AudioDeviceID
	MDevicePropertyChangedProc unsafe.Pointer
	MIOAudioDevice unsafe.Pointer
	MIOAudioEngine unsafe.Pointer
	MStreamPropertyChangedProc unsafe.Pointer
}/* debug [types.gen.go/struct]: AudioDriverPlugInHostInfo */

// AudioHardwareIOProcStreamUsage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareIOProcStreamUsage
type AudioHardwareIOProcStreamUsage struct {
	MIOProc unsafe.Pointer
	MNumberStreams unsafe.Pointer
	MStreamIsOn unsafe.Pointer
}/* debug [types.gen.go/struct]: AudioHardwareIOProcStreamUsage */

// AudioObjectPropertyAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyAddress
type AudioObjectPropertyAddress struct {
	MElement AudioObjectPropertyElement
	MScope AudioObjectPropertyScope
	MSelector AudioObjectPropertySelector
}/* debug [types.gen.go/struct]: AudioObjectPropertyAddress */

// AudioServerPlugInClientInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInClientInfo
type AudioServerPlugInClientInfo struct {
	MBundleID StringRef
	MClientID unsafe.Pointer
	MIsNativeEndian unsafe.Pointer
	MProcessID unsafe.Pointer
}/* debug [types.gen.go/struct]: AudioServerPlugInClientInfo */

// AudioServerPlugInCustomPropertyInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInCustomPropertyInfo
type AudioServerPlugInCustomPropertyInfo struct {
	MPropertyDataType unsafe.Pointer
	MQualifierDataType unsafe.Pointer
	MSelector AudioObjectPropertySelector
}/* debug [types.gen.go/struct]: AudioServerPlugInCustomPropertyInfo */

// AudioServerPlugInDriverInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInDriverInterface
type AudioServerPlugInDriverInterface struct {
	AbortDeviceConfigurationChange unsafe.Pointer
	AddDeviceClient unsafe.Pointer
	AddRef unsafe.Pointer
	BeginIOOperation unsafe.Pointer
	CreateDevice unsafe.Pointer
	DestroyDevice unsafe.Pointer
	DoIOOperation unsafe.Pointer
	EndIOOperation unsafe.Pointer
	GetPropertyData unsafe.Pointer
	GetPropertyDataSize unsafe.Pointer
	GetZeroTimeStamp unsafe.Pointer
	HasProperty unsafe.Pointer
	Initialize unsafe.Pointer
	IsPropertySettable unsafe.Pointer
	PerformDeviceConfigurationChange unsafe.Pointer
	QueryInterface unsafe.Pointer
	Release unsafe.Pointer
	RemoveDeviceClient unsafe.Pointer
	SetPropertyData unsafe.Pointer
	StartIO unsafe.Pointer
	StopIO unsafe.Pointer
	WillDoIOOperation unsafe.Pointer
}/* debug [types.gen.go/struct]: AudioServerPlugInDriverInterface */

// AudioServerPlugInHostInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInHostInterface
type AudioServerPlugInHostInterface struct {
	CopyFromStorage unsafe.Pointer
	DeleteFromStorage unsafe.Pointer
	PropertiesChanged unsafe.Pointer
	RequestDeviceConfigurationChange unsafe.Pointer
	WriteToStorage unsafe.Pointer
}/* debug [types.gen.go/struct]: AudioServerPlugInHostInterface */

// AudioServerPlugInIOCycleInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOCycleInfo
type AudioServerPlugInIOCycleInfo struct {
	MCurrentTime coreaudiotypes.AudioTimeStamp
	MDeviceHostTicksPerFrame unsafe.Pointer
	MInputTime coreaudiotypes.AudioTimeStamp
	MIOCycleCounter unsafe.Pointer
	MMainHostTicksPerFrame unsafe.Pointer
	MMasterHostTicksPerFrame unsafe.Pointer
	MNominalIOBufferFrameSize unsafe.Pointer
	MOutputTime coreaudiotypes.AudioTimeStamp
}/* debug [types.gen.go/struct]: AudioServerPlugInIOCycleInfo */

// AudioStreamRangedDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamRangedDescription
type AudioStreamRangedDescription struct {
	MFormat coreaudiotypes.AudioStreamBasicDescription
	MSampleRateRange coreaudiotypes.AudioValueRange
}/* debug [types.gen.go/struct]: AudioStreamRangedDescription */





