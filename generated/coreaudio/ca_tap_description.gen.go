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
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	BundleIDs() objc.IObject /* cross-framework: NSString */
	SetBundleIDs(value objc.IObject /* cross-framework: NSString */)
	DeviceUID() objc.IObject /* cross-framework: NSString */
	SetDeviceUID(value objc.IObject /* cross-framework: NSString */)
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
	MuteBehavior() TapMuteBehavior /* not a class type */
	SetMuteBehavior(value TapMuteBehavior /* not a class type */)
	Processes() unsafe.Pointer
	SetProcesses(value unsafe.Pointer)
	Stream() uint
	SetStream(value uint)
	Uuid() objc.IObject /* cross-framework: UUID */
	SetUuid(value objc.IObject /* cross-framework: UUID */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t_ TapDescription) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t_ TapDescription) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/bundleids
func (t_ TapDescription) BundleIDs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("bundleIDs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/bundleids
func (t_ TapDescription) SetBundleIDs(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBundleIDs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/deviceuid
func (t_ TapDescription) DeviceUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("deviceUID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/deviceuid
func (t_ TapDescription) SetDeviceUID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDeviceUID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isexclusive
func (t_ TapDescription) IsExclusive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isExclusive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isexclusive
func (t_ TapDescription) SetIsExclusive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsExclusive:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismixdown
func (t_ TapDescription) IsMixdown() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMixdown"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismixdown
func (t_ TapDescription) SetIsMixdown(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMixdown:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismono
func (t_ TapDescription) IsMono() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isMono"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/ismono
func (t_ TapDescription) SetIsMono(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsMono:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprivate
func (t_ TapDescription) IsPrivate() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isPrivate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprivate
func (t_ TapDescription) SetIsPrivate(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsPrivate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprocessrestoreenabled
func (t_ TapDescription) IsProcessRestoreEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isProcessRestoreEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/isprocessrestoreenabled
func (t_ TapDescription) SetIsProcessRestoreEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsProcessRestoreEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/mutebehavior
func (t_ TapDescription) MuteBehavior() TapMuteBehavior /* not a class type */ {
	rv := objc.Send[TapMuteBehavior](t_.ID, objc.Sel("muteBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/mutebehavior
func (t_ TapDescription) SetMuteBehavior(value TapMuteBehavior /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMuteBehavior:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/processes-1m4cr
func (t_ TapDescription) Processes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("processes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/processes-1m4cr
func (t_ TapDescription) SetProcesses(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setProcesses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/stream-ajk3
func (t_ TapDescription) Stream() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("stream"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/stream-ajk3
func (t_ TapDescription) SetStream(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStream:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/uuid
func (t_ TapDescription) Uuid() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](t_.ID, objc.Sel("uuid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudio/catapdescription/uuid
func (t_ TapDescription) SetUuid(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUuid:"), value)
}




