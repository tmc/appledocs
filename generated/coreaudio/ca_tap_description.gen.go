// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TapDescription] class.
var (
	TapDescriptionClass     _TapDescriptionClass
	TapDescriptionClassOnce sync.Once
)

func getTapDescriptionClass() _TapDescriptionClass {
	TapDescriptionClassOnce.Do(func() {
		TapDescriptionClass = _TapDescriptionClass{objc.GetClass("CATapDescription")}
	})
	return TapDescriptionClass
}

type _TapDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [TapDescription] class.
type ITapDescription interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription
type TapDescription struct {
	objectivec.Object
}

// TapDescriptionFrom constructs a [TapDescription] from an unsafe.Pointer.
func TapDescriptionFrom(ptr unsafe.Pointer) TapDescription {
	return TapDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TapDescriptionClass) Alloc() TapDescription {
	rv := objc.Send[TapDescription](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TapDescriptionClass) New() TapDescription {
	rv := objc.Send[TapDescription](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TapDescription) Init() TapDescription {
	rv := objc.Send[TapDescription](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TapDescription) Autorelease() TapDescription {
	rv := objc.Send[TapDescription](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTapDescription creates a new TapDescription instance.
func NewTapDescription() TapDescription {
	return getTapDescriptionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initExcludingProcesses:andDeviceUID:withStream:
func NewTapDescriptionExcludingProcessesAndDeviceUIDWithStream(processesObjectIDsToExcludeFromTap unsafe.Pointer, deviceUID string, stream int) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initExcludingProcesses:andDeviceUID:withStream:"), processesObjectIDsToExcludeFromTap, objc.String(deviceUID), stream)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoGlobalTapButExcludeProcesses:
func NewTapDescriptionMonoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap unsafe.Pointer) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initMonoGlobalTapButExcludeProcesses:"), processesObjectIDsToExcludeFromTap)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoMixdownOfProcesses:
func NewTapDescriptionMonoMixdownOfProcesses(processesObjectIDsToIncludeInTap unsafe.Pointer) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initMonoMixdownOfProcesses:"), processesObjectIDsToIncludeInTap)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoGlobalTapButExcludeProcesses:
func NewTapDescriptionStereoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap unsafe.Pointer) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initStereoGlobalTapButExcludeProcesses:"), processesObjectIDsToExcludeFromTap)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoMixdownOfProcesses:
func NewTapDescriptionStereoMixdownOfProcesses(processesObjectIDsToIncludeInTap unsafe.Pointer) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initStereoMixdownOfProcesses:"), processesObjectIDsToIncludeInTap)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initWithProcesses:andDeviceUID:withStream:
func NewTapDescriptionWithProcessesAndDeviceUIDWithStream(processesObjectIDsToIncludeInTap unsafe.Pointer, deviceUID string, stream int) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initWithProcesses:andDeviceUID:withStream:"), processesObjectIDsToIncludeInTap, objc.String(deviceUID), stream)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/bundleIDs
func (t_ TapDescription) BundleIDs() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("bundleIDs"))
	return rv
}


// SetBundleIDs sets the value of the bundleIDs property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/bundleIDs
func (t_ TapDescription) SetBundleIDs(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setBundleIDs:"), nsArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/deviceUID
func (t_ TapDescription) DeviceUID() string {
	rv := objc.Send[string](t_.ID, objc.Sel("deviceUID"))
	return rv
}


// SetDeviceUID sets the value of the deviceUID property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/deviceUID
func (t_ TapDescription) SetDeviceUID(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDeviceUID:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isExclusive
func (t_ TapDescription) Exclusive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("exclusive"))
	return rv
}


// SetExclusive sets the value of the exclusive property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isExclusive
func (t_ TapDescription) SetExclusive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExclusive:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMixdown
func (t_ TapDescription) Mixdown() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("mixdown"))
	return rv
}


// SetMixdown sets the value of the mixdown property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMixdown
func (t_ TapDescription) SetMixdown(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMixdown:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMono
func (t_ TapDescription) Mono() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("mono"))
	return rv
}


// SetMono sets the value of the mono property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMono
func (t_ TapDescription) SetMono(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMono:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isPrivate
func (t_ TapDescription) PrivateTap() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("privateTap"))
	return rv
}


// SetPrivateTap sets the value of the privateTap property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isPrivate
func (t_ TapDescription) SetPrivateTap(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrivateTap:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isProcessRestoreEnabled
func (t_ TapDescription) ProcessRestoreEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("processRestoreEnabled"))
	return rv
}


// SetProcessRestoreEnabled sets the value of the processRestoreEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isProcessRestoreEnabled
func (t_ TapDescription) SetProcessRestoreEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setProcessRestoreEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/muteBehavior
func (t_ TapDescription) MuteBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("muteBehavior"))
	return rv
}


// SetMuteBehavior sets the value of the muteBehavior property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/muteBehavior
func (t_ TapDescription) SetMuteBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMuteBehavior:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t_ TapDescription) Name() string {
	rv := objc.Send[string](t_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t_ TapDescription) SetName(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/processes-3cdzw
func (t_ TapDescription) Processes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](t_.ID, objc.Sel("processes"))
	return rv
}


// SetProcesses sets the value of the processes property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/processes-3cdzw
func (t_ TapDescription) SetProcesses(value []foundation.Number) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setProcesses:"), nsArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/stream-u4ff
func (t_ TapDescription) Stream() foundation.Number {
	rv := objc.Send[foundation.Number](t_.ID, objc.Sel("stream"))
	return rv
}


// SetStream sets the value of the stream property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/stream-u4ff
func (t_ TapDescription) SetStream(value foundation.Number) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStream:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/uuid
func (t_ TapDescription) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("UUID"))
	return rv
}


// SetUUID sets the value of the UUID property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/uuid
func (t_ TapDescription) SetUUID(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUUID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isexclusive
func (t_ TapDescription) IsExclusive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isExclusive"))
	return rv
}


// SetIsExclusive sets the value of the isExclusive property.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isexclusive
func (t_ TapDescription) SetIsExclusive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsExclusive:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismixdown
func (t_ TapDescription) IsMixdown() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMixdown"))
	return rv
}


// SetIsMixdown sets the value of the isMixdown property.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismixdown
func (t_ TapDescription) SetIsMixdown(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMixdown:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismono
func (t_ TapDescription) IsMono() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMono"))
	return rv
}


// SetIsMono sets the value of the isMono property.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismono
func (t_ TapDescription) SetIsMono(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMono:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprivate
func (t_ TapDescription) IsPrivate() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isPrivate"))
	return rv
}


// SetIsPrivate sets the value of the isPrivate property.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprivate
func (t_ TapDescription) SetIsPrivate(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsPrivate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprocessrestoreenabled
func (t_ TapDescription) IsProcessRestoreEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isProcessRestoreEnabled"))
	return rv
}


// SetIsProcessRestoreEnabled sets the value of the isProcessRestoreEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprocessrestoreenabled
func (t_ TapDescription) SetIsProcessRestoreEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsProcessRestoreEnabled:"), value)
}


