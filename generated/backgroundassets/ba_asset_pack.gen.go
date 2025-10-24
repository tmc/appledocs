// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class BAAssetPack */


/* debug [class_header]: Header for BAAssetPack */
// The class instance for the [BAAssetPack] class.
var (
	BAAssetPackClass     _BAAssetPackClass
	BAAssetPackClassOnce sync.Once
)

func getBAAssetPackClass() _BAAssetPackClass {
	BAAssetPackClassOnce.Do(func() {
		BAAssetPackClass = _BAAssetPackClass{objc.GetClass("BAAssetPack")}
	})
	return BAAssetPackClass
}

type _BAAssetPackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BAAssetPack */
// An interface definition for the [BAAssetPack] class.
type IBAAssetPack interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BAAssetPack */
	// properties:
	DownloadSize() int
	Identifier() objc.IObject /* cross-framework: NSString */
	UserInfo() objc.IObject /* cross-framework: NSData */
	Version() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BAAssetPack */
	// methods:
	Download() IBADownload
	DownloadForContentRequest(contentRequest BAContentRequest) IBADownload
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BAAssetPack */
// Alloc allocates a new instance without initialization.
func (bc _BAAssetPackClass) Alloc() BAAssetPack {
	rv := objc.Send[BAAssetPack](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BAAssetPackClass) New() BAAssetPack {
	rv := objc.Send[BAAssetPack](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAssetPack) Init() BAAssetPack {
	rv := objc.Send[BAAssetPack](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAssetPack) Autorelease() BAAssetPack {
	rv := objc.Send[BAAssetPack](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAssetPack creates a new BAAssetPack instance.
func NewBAAssetPack() BAAssetPack {
	return getBAAssetPackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BAAssetPack */
// An archive of assets that the system downloads together.
//
// An instance of this class can be invalidated when the asset pack that it represents is updated on the server.


// An archive of assets that the system downloads together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack
type BAAssetPack struct {
	objectivec.Object
}

// BAAssetPackFrom constructs a [BAAssetPack] from an unsafe.Pointer.
//
// An archive of assets that the system downloads together.
func BAAssetPackFrom(ptr unsafe.Pointer) BAAssetPack {
	return BAAssetPack{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BAAssetPack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BAAssetPack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BAAssetPack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BAAssetPack */

// Creates a download object for the asset pack that you schedule using a download manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/download
func (b_ BAAssetPack) Download() IBADownload {
	rv := objc.Send[BADownload](b_.ID, objc.Sel("download"))
	return rv
}/* debug [instance_methods/method]: Download */


// Creates a download object for the asset pack that you schedule using a download manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/downloadForContentRequest:
func (b_ BAAssetPack) DownloadForContentRequest(contentRequest BAContentRequest) IBADownload {
	rv := objc.Send[BADownload](b_.ID, objc.Sel("downloadForContentRequest:"), contentRequest)
	return rv
}/* debug [instance_methods/method]: DownloadForContentRequest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BAAssetPack */

// The size of the download file containing the asset pack in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/downloadSize
func (b_ BAAssetPack) DownloadSize() int {
	rv := objc.Send[int](b_.ID, objc.Sel("downloadSize"))
	return rv
}/* debug [instance_properties/getter]: downloadSize */


// A unique identifier for the asset pack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/identifier
func (b_ BAAssetPack) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// JSON-encoded custom information that’s associated with the asset pack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/userInfo
func (b_ BAAssetPack) UserInfo() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](b_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// The asset pack’s version number
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/version
func (b_ BAAssetPack) Version() int {
	rv := objc.Send[int](b_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class BAAssetPack */



