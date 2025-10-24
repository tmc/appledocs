// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CATapDescription */


/* debug [class_header]: Header for CATapDescription */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TapDescription */
// An interface definition for the [TapDescription] class.
type ITapDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TapDescription */
	// properties:
	BundleIDs() []string
	SetBundleIDs(value []string)
	DeviceUID() objc.IObject /* cross-framework: NSString */
	SetDeviceUID(value objc.IObject /* cross-framework: NSString */)
	Exclusive() bool
	SetExclusive(value bool)
	Mixdown() bool
	SetMixdown(value bool)
	Mono() bool
	SetMono(value bool)
	PrivateTap() bool
	SetPrivateTap(value bool)
	ProcessRestoreEnabled() bool
	SetProcessRestoreEnabled(value bool)
	MuteBehavior() TapMuteBehavior
	SetMuteBehavior(value TapMuteBehavior)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Processes() []foundation.Number
	SetProcesses(value []foundation.Number)
	Stream() objc.IObject /* cross-framework: NSNumber */
	SetStream(value objc.IObject /* cross-framework: NSNumber */)
	UUID() foundation.UUID
	SetUUID(value foundation.UUID)
	IsExclusive() bool
	SetIsExclusive(value bool)
	IsMixdown() bool
	SetIsMixdown(value bool)
	IsMono() bool
	SetIsMono(value bool)
	IsPrivate() bool
	SetIsPrivate(value bool)
	IsProcessRestoreEnabled() bool
	SetIsProcessRestoreEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TapDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TapDescription */
// Alloc allocates a new instance without initialization.
func (tc _TapDescriptionClass) Alloc() TapDescription {
	rv := objc.Send[TapDescription](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TapDescription */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription
type TapDescription struct {
	objectivec.Object
}

// TapDescriptionFrom constructs a [TapDescription] from an unsafe.Pointer.
func TapDescriptionFrom(ptr unsafe.Pointer) TapDescription {
	return TapDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TapDescription */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initExcludingProcesses:andDeviceUID:withStream:
func NewTapDescriptionExcludingProcessesAndDeviceUIDWithStream(processesObjectIDsToExcludeFromTap []foundation.Number, deviceUID objc.IObject /* cross-framework: NSString */, stream int) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initExcludingProcesses:andDeviceUID:withStream:"), processesObjectIDsToExcludeFromTap, deviceUID, stream)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTapDescriptionExcludingProcessesAndDeviceUIDWithStream */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoGlobalTapButExcludeProcesses:
func NewTapDescriptionMonoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.Number) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initMonoGlobalTapButExcludeProcesses:"), processesObjectIDsToExcludeFromTap)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTapDescriptionMonoGlobalTapButExcludeProcesses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoMixdownOfProcesses:
func NewTapDescriptionMonoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.Number) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initMonoMixdownOfProcesses:"), processesObjectIDsToIncludeInTap)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTapDescriptionMonoMixdownOfProcesses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoGlobalTapButExcludeProcesses:
func NewTapDescriptionStereoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.Number) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initStereoGlobalTapButExcludeProcesses:"), processesObjectIDsToExcludeFromTap)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTapDescriptionStereoGlobalTapButExcludeProcesses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoMixdownOfProcesses:
func NewTapDescriptionStereoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.Number) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initStereoMixdownOfProcesses:"), processesObjectIDsToIncludeInTap)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTapDescriptionStereoMixdownOfProcesses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initWithProcesses:andDeviceUID:withStream:
func NewTapDescriptionWithProcessesAndDeviceUIDWithStream(processesObjectIDsToIncludeInTap []foundation.Number, deviceUID objc.IObject /* cross-framework: NSString */, stream int) TapDescription {
	instance := getTapDescriptionClass().Alloc()
	rv := objc.Send[TapDescription](instance.ID, objc.Sel("initWithProcesses:andDeviceUID:withStream:"), processesObjectIDsToIncludeInTap, deviceUID, stream)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTapDescriptionWithProcessesAndDeviceUIDWithStream */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TapDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TapDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TapDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TapDescription */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/bundleIDs
func (t_ TapDescription) BundleIDs() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("bundleIDs"))
	return rv
}/* debug [instance_properties/getter]: bundleIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/bundleIDs
func (t_ TapDescription) SetBundleIDs(value []string) {
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
}/* debug [instance_properties/setter]: bundleIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/deviceUID
func (t_ TapDescription) DeviceUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("deviceUID"))
	return rv
}/* debug [instance_properties/getter]: deviceUID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/deviceUID
func (t_ TapDescription) SetDeviceUID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDeviceUID:"), value)
}/* debug [instance_properties/setter]: deviceUID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isExclusive
func (t_ TapDescription) Exclusive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("exclusive"))
	return rv
}/* debug [instance_properties/getter]: exclusive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isExclusive
func (t_ TapDescription) SetExclusive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExclusive:"), value)
}/* debug [instance_properties/setter]: exclusive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMixdown
func (t_ TapDescription) Mixdown() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("mixdown"))
	return rv
}/* debug [instance_properties/getter]: mixdown */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMixdown
func (t_ TapDescription) SetMixdown(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMixdown:"), value)
}/* debug [instance_properties/setter]: mixdown */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMono
func (t_ TapDescription) Mono() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("mono"))
	return rv
}/* debug [instance_properties/getter]: mono */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMono
func (t_ TapDescription) SetMono(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMono:"), value)
}/* debug [instance_properties/setter]: mono */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isPrivate
func (t_ TapDescription) PrivateTap() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("privateTap"))
	return rv
}/* debug [instance_properties/getter]: privateTap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isPrivate
func (t_ TapDescription) SetPrivateTap(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrivateTap:"), value)
}/* debug [instance_properties/setter]: privateTap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isProcessRestoreEnabled
func (t_ TapDescription) ProcessRestoreEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("processRestoreEnabled"))
	return rv
}/* debug [instance_properties/getter]: processRestoreEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isProcessRestoreEnabled
func (t_ TapDescription) SetProcessRestoreEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setProcessRestoreEnabled:"), value)
}/* debug [instance_properties/setter]: processRestoreEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/muteBehavior
func (t_ TapDescription) MuteBehavior() TapMuteBehavior {
	rv := objc.Send[TapMuteBehavior](t_.ID, objc.Sel("muteBehavior"))
	return rv
}/* debug [instance_properties/getter]: muteBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/muteBehavior
func (t_ TapDescription) SetMuteBehavior(value TapMuteBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMuteBehavior:"), value)
}/* debug [instance_properties/setter]: muteBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t_ TapDescription) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t_ TapDescription) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/processes-3cdzw
func (t_ TapDescription) Processes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](t_.ID, objc.Sel("processes"))
	return rv
}/* debug [instance_properties/getter]: processes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/processes-3cdzw
func (t_ TapDescription) SetProcesses(value []foundation.Number) {
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
}/* debug [instance_properties/setter]: processes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/stream-u4ff
func (t_ TapDescription) Stream() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](t_.ID, objc.Sel("stream"))
	return rv
}/* debug [instance_properties/getter]: stream */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/stream-u4ff
func (t_ TapDescription) SetStream(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStream:"), value)
}/* debug [instance_properties/setter]: stream */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/uuid
func (t_ TapDescription) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](t_.ID, objc.Sel("UUID"))
	return rv
}/* debug [instance_properties/getter]: UUID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/uuid
func (t_ TapDescription) SetUUID(value foundation.UUID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUUID:"), value)
}/* debug [instance_properties/setter]: UUID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isexclusive
func (t_ TapDescription) IsExclusive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isExclusive"))
	return rv
}/* debug [instance_properties/getter]: isExclusive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isexclusive
func (t_ TapDescription) SetIsExclusive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsExclusive:"), value)
}/* debug [instance_properties/setter]: isExclusive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismixdown
func (t_ TapDescription) IsMixdown() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMixdown"))
	return rv
}/* debug [instance_properties/getter]: isMixdown */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismixdown
func (t_ TapDescription) SetIsMixdown(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMixdown:"), value)
}/* debug [instance_properties/setter]: isMixdown */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismono
func (t_ TapDescription) IsMono() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMono"))
	return rv
}/* debug [instance_properties/getter]: isMono */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismono
func (t_ TapDescription) SetIsMono(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMono:"), value)
}/* debug [instance_properties/setter]: isMono */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprivate
func (t_ TapDescription) IsPrivate() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isPrivate"))
	return rv
}/* debug [instance_properties/getter]: isPrivate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprivate
func (t_ TapDescription) SetIsPrivate(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsPrivate:"), value)
}/* debug [instance_properties/setter]: isPrivate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprocessrestoreenabled
func (t_ TapDescription) IsProcessRestoreEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isProcessRestoreEnabled"))
	return rv
}/* debug [instance_properties/getter]: isProcessRestoreEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprocessrestoreenabled
func (t_ TapDescription) SetIsProcessRestoreEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsProcessRestoreEnabled:"), value)
}/* debug [instance_properties/setter]: isProcessRestoreEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CATapDescription */


