// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class BAAssetPackManifest */


/* debug [class_header]: Header for BAAssetPackManifest */
// The class instance for the [BAAssetPackManifest] class.
var (
	BAAssetPackManifestClass     _BAAssetPackManifestClass
	BAAssetPackManifestClassOnce sync.Once
)

func getBAAssetPackManifestClass() _BAAssetPackManifestClass {
	BAAssetPackManifestClassOnce.Do(func() {
		BAAssetPackManifestClass = _BAAssetPackManifestClass{objc.GetClass("BAAssetPackManifest")}
	})
	return BAAssetPackManifestClass
}

type _BAAssetPackManifestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BAAssetPackManifest */
// An interface definition for the [BAAssetPackManifest] class.
type IBAAssetPackManifest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BAAssetPackManifest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BAAssetPackManifest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BAAssetPackManifest */
// Alloc allocates a new instance without initialization.
func (bc _BAAssetPackManifestClass) Alloc() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BAAssetPackManifestClass) New() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAssetPackManifest) Init() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAssetPackManifest) Autorelease() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAssetPackManifest creates a new BAAssetPackManifest instance.
func NewBAAssetPackManifest() BAAssetPackManifest {
	return getBAAssetPackManifestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BAAssetPackManifest */
// A representation of a manifest that lists asset packs that are available to download.
//
// This class applies only when you want to manage your asset packs manually. Don’t use this class if you want to opt in to automatic management of asset packs.


// A representation of a manifest that lists asset packs that are available to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest
type BAAssetPackManifest struct {
	objectivec.Object
}

// BAAssetPackManifestFrom constructs a [BAAssetPackManifest] from an unsafe.Pointer.
//
// A representation of a manifest that lists asset packs that are available to download.
func BAAssetPackManifestFrom(ptr unsafe.Pointer) BAAssetPackManifest {
	return BAAssetPackManifest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BAAssetPackManifest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BAAssetPackManifest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BAAssetPackManifest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BAAssetPackManifest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BAAssetPackManifest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class BAAssetPackManifest */



