// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEFileInfo] class.
var (
	MEFileInfoClass     _MEFileInfoClass
	MEFileInfoClassOnce sync.Once
)

func getMEFileInfoClass() _MEFileInfoClass {
	MEFileInfoClassOnce.Do(func() {
		MEFileInfoClass = _MEFileInfoClass{objc.GetClass("MEFileInfo")}
	})
	return MEFileInfoClass
}

type _MEFileInfoClass struct {
	class objc.Class
}

// An interface definition for the [MEFileInfo] class.
type IMEFileInfo interface {
	objectivec.IObject
	// properties:
	FragmentsStatus() unsafe.Pointer
	SetFragmentsStatus(value unsafe.Pointer)
	Duration() objc.IObject /* cross-framework: Time */
	SetDuration(value objc.IObject /* cross-framework: Time */)
	SidecarFileName() objc.IObject /* cross-framework: NSString */
	SetSidecarFileName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An object that contains file properties from the media asset.


// An object that contains file properties from the media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo
type MEFileInfo struct {
	objectivec.Object
}

// MEFileInfoFrom constructs a [MEFileInfo] from an unsafe.Pointer.
//
// An object that contains file properties from the media asset.
func MEFileInfoFrom(ptr unsafe.Pointer) MEFileInfo {
	return MEFileInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEFileInfoClass) Alloc() MEFileInfo {
	rv := objc.Send[MEFileInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEFileInfoClass) New() MEFileInfo {
	rv := objc.Send[MEFileInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEFileInfo) Init() MEFileInfo {
	rv := objc.Send[MEFileInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEFileInfo) Autorelease() MEFileInfo {
	rv := objc.Send[MEFileInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEFileInfo creates a new MEFileInfo instance.
func NewMEFileInfo() MEFileInfo {
	return getMEFileInfoClass().New()
}



// Indicates if the media asset contains fragments or is extendable by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/fragmentsStatus-swift.property
func (m_ MEFileInfo) FragmentsStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentsStatus"))
	return rv
}


// Indicates if the media asset contains fragments or is extendable by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/fragmentsStatus-swift.property
func (m_ MEFileInfo) SetFragmentsStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentsStatus:"), value)
}


// The duration of the media asset, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mefileinfo/duration
func (m_ MEFileInfo) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](m_.ID, objc.Sel("duration"))
	return rv
}


// The duration of the media asset, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mefileinfo/duration
func (m_ MEFileInfo) SetDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mefileinfo/sidecarfilename
func (m_ MEFileInfo) SidecarFileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("sidecarFileName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mefileinfo/sidecarfilename
func (m_ MEFileInfo) SetSidecarFileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSidecarFileName:"), value)
}



