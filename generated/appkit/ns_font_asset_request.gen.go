// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFontAssetRequest */


/* debug [class_header]: Header for NSFontAssetRequest */
// The class instance for the [FontAssetRequest] class.
var (
	FontAssetRequestClass     _FontAssetRequestClass
	FontAssetRequestClassOnce sync.Once
)

func getFontAssetRequestClass() _FontAssetRequestClass {
	FontAssetRequestClassOnce.Do(func() {
		FontAssetRequestClass = _FontAssetRequestClass{objc.GetClass("NSFontAssetRequest")}
	})
	return FontAssetRequestClass
}

type _FontAssetRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FontAssetRequest */
// An interface definition for the [FontAssetRequest] class.
type IFontAssetRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FontAssetRequest */
	// properties:
	DownloadedFontDescriptors() []FontDescriptor
	Progress() foundation.Progress
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FontAssetRequest */
	// methods:
	DownloadFontAssetsWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FontAssetRequest */
// Alloc allocates a new instance without initialization.
func (fc _FontAssetRequestClass) Alloc() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FontAssetRequestClass) New() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontAssetRequest) Init() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontAssetRequest) Autorelease() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontAssetRequest creates a new FontAssetRequest instance.
func NewFontAssetRequest() FontAssetRequest {
	return getFontAssetRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FontAssetRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest
type FontAssetRequest struct {
	objectivec.Object
}

// FontAssetRequestFrom constructs a [FontAssetRequest] from an unsafe.Pointer.
func FontAssetRequestFrom(ptr unsafe.Pointer) FontAssetRequest {
	return FontAssetRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FontAssetRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/init(fontDescriptors:options:)
func NewFontAssetRequestWithFontDescriptorsOptions(fontDescriptors []FontDescriptor, options FontAssetRequestOptions) FontAssetRequest {
	instance := getFontAssetRequestClass().Alloc()
	rv := objc.Send[FontAssetRequest](instance.ID, objc.Sel("initWithFontDescriptors:options:"), fontDescriptors, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFontAssetRequestWithFontDescriptorsOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FontAssetRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FontAssetRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FontAssetRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/download(withCompletionHandler:)
func (f_ FontAssetRequest) DownloadFontAssetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("downloadFontAssetsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: DownloadFontAssetsWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FontAssetRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/downloadedFontDescriptors
func (f_ FontAssetRequest) DownloadedFontDescriptors() []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("downloadedFontDescriptors"))
	return rv
}/* debug [instance_properties/getter]: downloadedFontDescriptors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/progress
func (f_ FontAssetRequest) Progress() foundation.Progress {
	rv := objc.Send[foundation.Progress](f_.ID, objc.Sel("progress"))
	return rv
}/* debug [instance_properties/getter]: progress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFontAssetRequest */


