// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class BADownload */


/* debug [class_header]: Header for BADownload */
// The class instance for the [BADownload] class.
var (
	BADownloadClass     _BADownloadClass
	BADownloadClassOnce sync.Once
)

func getBADownloadClass() _BADownloadClass {
	BADownloadClassOnce.Do(func() {
		BADownloadClass = _BADownloadClass{objc.GetClass("BADownload")}
	})
	return BADownloadClass
}

type _BADownloadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BADownload */
// An interface definition for the [BADownload] class.
type IBADownload interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BADownload */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	IsEssential() bool
	Priority() BADownloaderPriority /* typedef */
	State() BADownloadState
	UniqueIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BADownload */
	// methods:
	CopyAsNonEssential() unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BADownload */
// Alloc allocates a new instance without initialization.
func (bc _BADownloadClass) Alloc() BADownload {
	rv := objc.Send[BADownload](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BADownloadClass) New() BADownload {
	rv := objc.Send[BADownload](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BADownload) Init() BADownload {
	rv := objc.Send[BADownload](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BADownload) Autorelease() BADownload {
	rv := objc.Send[BADownload](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBADownload creates a new BADownload instance.
func NewBADownload() BADownload {
	return getBADownloadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BADownload */
// An object that represents an in-progress or concluded asset download.


// An object that represents an in-progress or concluded asset download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload
type BADownload struct {
	objectivec.Object
}

// BADownloadFrom constructs a [BADownload] from an unsafe.Pointer.
//
// An object that represents an in-progress or concluded asset download.
func BADownloadFrom(ptr unsafe.Pointer) BADownload {
	return BADownload{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BADownload *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BADownload */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BADownload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BADownload */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/removingEssential()
func (b_ BADownload) CopyAsNonEssential() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("copyAsNonEssential"))
	return rv
}/* debug [instance_methods/method]: CopyAsNonEssential */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BADownload */

// The app-specific string that uniquely identifies the downloadable asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/identifier
func (b_ BADownload) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/isEssential
func (b_ BADownload) IsEssential() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEssential"))
	return rv
}/* debug [instance_properties/getter]: isEssential */


// The download’s execution priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/priority-swift.property
func (b_ BADownload) Priority() BADownloaderPriority /* typedef */ {
	rv := objc.Send[int](b_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */


// The current state of the download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/state-swift.property
func (b_ BADownload) State() BADownloadState {
	rv := objc.Send[BADownloadState](b_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The system-provided string that uniquely identifies the download object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/uniqueIdentifier
func (b_ BADownload) UniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: uniqueIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class BADownload */



