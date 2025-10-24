// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */


/* debug [class_header]: Header for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
// The class instance for the [MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass     _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass() _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass {
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass = _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass
}

type _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
// An interface definition for the [MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent] class.
type IMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
	// properties:
	BytesDownloaded() objc.IObject /* cross-framework: NSNumber */
	SetBytesDownloaded(value objc.IObject /* cross-framework: NSNumber */)
	PlatformCode() objc.IObject /* cross-framework: NSNumber */
	SetPlatformCode(value objc.IObject /* cross-framework: NSNumber */)
	ProgressPercent() objc.IObject /* cross-framework: NSNumber */
	SetProgressPercent(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass) New() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) Init() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent creates a new MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent() MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	return getMTROTASoftwareUpdateRequestorClusterDownloadErrorEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw
type MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterDownloadErrorEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterDownloadErrorEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent {
	return MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/bytesDownloaded
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) BytesDownloaded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("bytesDownloaded"))
	return rv
}/* debug [instance_properties/getter]: bytesDownloaded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/bytesDownloaded
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetBytesDownloaded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBytesDownloaded:"), value)
}/* debug [instance_properties/setter]: bytesDownloaded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/platformCode
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) PlatformCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("platformCode"))
	return rv
}/* debug [instance_properties/getter]: platformCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/platformCode
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetPlatformCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlatformCode:"), value)
}/* debug [instance_properties/setter]: platformCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/progressPercent
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) ProgressPercent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("progressPercent"))
	return rv
}/* debug [instance_properties/getter]: progressPercent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/progressPercent
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetProgressPercent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressPercent:"), value)
}/* debug [instance_properties/setter]: progressPercent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/softwareVersion
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent-2w5rw/softwareVersion
func (m_ MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent */



