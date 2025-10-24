// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class BAURLDownload */


/* debug [class_header]: Header for BAURLDownload */
// The class instance for the [BAURLDownload] class.
var (
	BAURLDownloadClass     _BAURLDownloadClass
	BAURLDownloadClassOnce sync.Once
)

func getBAURLDownloadClass() _BAURLDownloadClass {
	BAURLDownloadClassOnce.Do(func() {
		BAURLDownloadClass = _BAURLDownloadClass{objc.GetClass("BAURLDownload")}
	})
	return BAURLDownloadClass
}

type _BAURLDownloadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BAURLDownload */
// An interface definition for the [BAURLDownload] class.
type IBAURLDownload interface {
	IBADownload
	
/* debug [class_interface_properties]: Properties for BAURLDownload */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BAURLDownload */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BAURLDownload */
// Alloc allocates a new instance without initialization.
func (bc _BAURLDownloadClass) Alloc() BAURLDownload {
	rv := objc.Send[BAURLDownload](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BAURLDownloadClass) New() BAURLDownload {
	rv := objc.Send[BAURLDownload](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAURLDownload) Init() BAURLDownload {
	rv := objc.Send[BAURLDownload](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAURLDownload) Autorelease() BAURLDownload {
	rv := objc.Send[BAURLDownload](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAURLDownload creates a new BAURLDownload instance.
func NewBAURLDownload() BAURLDownload {
	return getBAURLDownloadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BAURLDownload */
// An object that represents a remote asset to download.


// An object that represents a remote asset to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload
type BAURLDownload struct {
	BADownload
}

// BAURLDownloadFrom constructs a [BAURLDownload] from an unsafe.Pointer.
//
// An object that represents a remote asset to download.
func BAURLDownloadFrom(ptr unsafe.Pointer) BAURLDownload {
	return BAURLDownload{
		BADownload: BADownloadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BAURLDownload */

// Creates a download that uses the specified identifier and App Group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:applicationGroupIdentifier:)
func NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier(identifier objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:applicationGroupIdentifier:"), identifier, request, applicationGroupIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier */


// Creates a prioritized download that uses the specified identifier and App Group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:applicationGroupIdentifier:priority:)
func NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority(identifier objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */, priority BADownloaderPriority /* typedef */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:applicationGroupIdentifier:priority:"), identifier, request, applicationGroupIdentifier, priority)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:essential:fileSize:applicationGroupIdentifier:priority:)
func NewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority(identifier objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, essential bool, fileSize uint, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */, priority BADownloaderPriority /* typedef */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:essential:fileSize:applicationGroupIdentifier:priority:"), identifier, request, essential, fileSize, applicationGroupIdentifier, priority)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:fileSize:applicationGroupIdentifier:)
func NewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier(identifier objc.IObject /* cross-framework: NSString */, request foundation.URLRequest, fileSize uint, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:fileSize:applicationGroupIdentifier:"), identifier, request, fileSize, applicationGroupIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BAURLDownload */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BAURLDownload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BAURLDownload */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BAURLDownload */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class BAURLDownload */


