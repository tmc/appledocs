// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDataAsset */


/* debug [class_header]: Header for NSDataAsset */
// The class instance for the [DataAsset] class.
var (
	DataAssetClass     _DataAssetClass
	DataAssetClassOnce sync.Once
)

func getDataAssetClass() _DataAssetClass {
	DataAssetClassOnce.Do(func() {
		DataAssetClass = _DataAssetClass{objc.GetClass("NSDataAsset")}
	})
	return DataAssetClass
}

type _DataAssetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DataAsset */
// An interface definition for the [DataAsset] class.
type IDataAsset interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DataAsset */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Name() DataAssetName /* typedef */
	TypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DataAsset */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DataAsset */
// Alloc allocates a new instance without initialization.
func (dc _DataAssetClass) Alloc() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DataAssetClass) New() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DataAsset) Init() DataAsset {
	rv := objc.Send[DataAsset](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DataAsset) Autorelease() DataAsset {
	rv := objc.Send[DataAsset](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDataAsset creates a new DataAsset instance.
func NewDataAsset() DataAsset {
	return getDataAssetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DataAsset */
// An object from a data set type stored in an asset catalog.
//
// The object’s content is stored as a set of one or more files with associated device attributes. These sets can also be tagged for use as on-demand resources.


// An object from a data set type stored in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset
type DataAsset struct {
	objectivec.Object
}

// DataAssetFrom constructs a [DataAsset] from an unsafe.Pointer.
//
// An object from a data set type stored in an asset catalog.
func DataAssetFrom(ptr unsafe.Pointer) DataAsset {
	return DataAsset{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DataAsset */

// Initializes and returns an object with a reference to the named data asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:)
func NewDataAssetWithName(name DataAssetName /* typedef */) DataAsset {
	instance := getDataAssetClass().Alloc()
	rv := objc.Send[DataAsset](instance.ID, objc.Sel("initWithName:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataAssetWithName */


// Initializes and returns an object with a reference to the named data asset that’s in an asset catalog in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:bundle:)
func NewDataAssetWithNameBundle(name DataAssetName /* typedef */, bundle foundation.Bundle) DataAsset {
	instance := getDataAssetClass().Alloc()
	rv := objc.Send[DataAsset](instance.ID, objc.Sel("initWithName:bundle:"), name, bundle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataAssetWithNameBundle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DataAsset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DataAsset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DataAsset */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DataAsset */

// The raw data values in the data asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/data
func (d_ DataAsset) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The name of the data set in the asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/name-swift.property
func (d_ DataAsset) Name() DataAssetName /* typedef */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The uniform type identifier for the data asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/typeIdentifier
func (d_ DataAsset) TypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("typeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: typeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDataAsset */


