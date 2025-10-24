// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEFileInfo */


/* debug [class_header]: Header for MEFileInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEFileInfo */
// An interface definition for the [MEFileInfo] class.
type IMEFileInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEFileInfo */
	// properties:
	Duration() objc.IObject /* cross-framework: Time */
	SetDuration(value objc.IObject /* cross-framework: Time */)
	FragmentsStatus() MEFileInfoFragmentsStatus
	SetFragmentsStatus(value MEFileInfoFragmentsStatus)
	SidecarFileName() objc.IObject /* cross-framework: NSString */
	SetSidecarFileName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEFileInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEFileInfo */
// Alloc allocates a new instance without initialization.
func (mc _MEFileInfoClass) Alloc() MEFileInfo {
	rv := objc.Send[MEFileInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEFileInfo */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEFileInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEFileInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEFileInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEFileInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEFileInfo */

// The duration of the media asset, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/duration
func (m_ MEFileInfo) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration of the media asset, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/duration
func (m_ MEFileInfo) SetDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// Indicates if the media asset contains fragments or is extendable by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/fragmentsStatus-swift.property
func (m_ MEFileInfo) FragmentsStatus() MEFileInfoFragmentsStatus {
	rv := objc.Send[MEFileInfoFragmentsStatus](m_.ID, objc.Sel("fragmentsStatus"))
	return rv
}/* debug [instance_properties/getter]: fragmentsStatus */


// Indicates if the media asset contains fragments or is extendable by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/fragmentsStatus-swift.property
func (m_ MEFileInfo) SetFragmentsStatus(value MEFileInfoFragmentsStatus) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentsStatus:"), value)
}/* debug [instance_properties/setter]: fragmentsStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/sidecarFileName
func (m_ MEFileInfo) SidecarFileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("sidecarFileName"))
	return rv
}/* debug [instance_properties/getter]: sidecarFileName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/sidecarFileName
func (m_ MEFileInfo) SetSidecarFileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSidecarFileName:"), value)
}/* debug [instance_properties/setter]: sidecarFileName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEFileInfo */



