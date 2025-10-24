// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */


/* debug [class_header]: Header for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
// The class instance for the [MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass     _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass
	MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass() _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass {
	MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass = _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent interface {
	IMTROTASoftwareUpdateRequestorClusterDownloadErrorEvent
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
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

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass) New() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) Init() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent creates a new MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent() MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	return getMTROtaSoftwareUpdateRequestorClusterDownloadErrorEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t
type MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent struct {
	MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent
}

// MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterDownloadErrorEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent {
	return MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent{
		MTROTASoftwareUpdateRequestorClusterDownloadErrorEvent: MTROTASoftwareUpdateRequestorClusterDownloadErrorEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/bytesDownloaded
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) BytesDownloaded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("bytesDownloaded"))
	return rv
}/* debug [instance_properties/getter]: bytesDownloaded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/bytesDownloaded
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetBytesDownloaded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBytesDownloaded:"), value)
}/* debug [instance_properties/setter]: bytesDownloaded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/platformCode
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) PlatformCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("platformCode"))
	return rv
}/* debug [instance_properties/getter]: platformCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/platformCode
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetPlatformCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlatformCode:"), value)
}/* debug [instance_properties/setter]: platformCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/progressPercent
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) ProgressPercent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("progressPercent"))
	return rv
}/* debug [instance_properties/getter]: progressPercent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/progressPercent
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetProgressPercent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgressPercent:"), value)
}/* debug [instance_properties/setter]: progressPercent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/softwareVersion
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent-73h3t/softwareVersion
func (m_ MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateRequestorClusterDownloadErrorEvent */



