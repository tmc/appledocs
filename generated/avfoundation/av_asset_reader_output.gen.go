// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetReaderOutput */


/* debug [class_header]: Header for AVAssetReaderOutput */
// The class instance for the [AssetReaderOutput] class.
var (
	AssetReaderOutputClass     _AssetReaderOutputClass
	AssetReaderOutputClassOnce sync.Once
)

func getAssetReaderOutputClass() _AssetReaderOutputClass {
	AssetReaderOutputClassOnce.Do(func() {
		AssetReaderOutputClass = _AssetReaderOutputClass{objc.GetClass("AVAssetReaderOutput")}
	})
	return AssetReaderOutputClass
}

type _AssetReaderOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetReaderOutput */
// An interface definition for the [AssetReaderOutput] class.
type IAssetReaderOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetReaderOutput */
	// properties:
	AlwaysCopiesSampleData() bool
	SetAlwaysCopiesSampleData(value bool)
	MediaType() MediaType /* typedef */
	SupportsRandomAccess() bool
	SetSupportsRandomAccess(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetReaderOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetReaderOutput */
// Alloc allocates a new instance without initialization.
func (ac _AssetReaderOutputClass) Alloc() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderOutputClass) New() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderOutput) Init() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderOutput) Autorelease() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderOutput creates a new AssetReaderOutput instance.
func NewAssetReaderOutput() AssetReaderOutput {
	return getAssetReaderOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetReaderOutput */
// An abstract class that defines the interface to read media samples from an asset reader.
//
// You add concrete output instances, such as or , to an asset reader to perform specific tasks.


// An abstract class that defines the interface to read media samples from an asset reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput
type AssetReaderOutput struct {
	objectivec.Object
}

// AssetReaderOutputFrom constructs a [AssetReaderOutput] from an unsafe.Pointer.
//
// An abstract class that defines the interface to read media samples from an asset reader.
func AssetReaderOutputFrom(ptr unsafe.Pointer) AssetReaderOutput {
	return AssetReaderOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetReaderOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetReaderOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetReaderOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetReaderOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetReaderOutput */

// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/alwaysCopiesSampleData
func (a_ AssetReaderOutput) AlwaysCopiesSampleData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysCopiesSampleData"))
	return rv
}/* debug [instance_properties/getter]: alwaysCopiesSampleData */


// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/alwaysCopiesSampleData
func (a_ AssetReaderOutput) SetAlwaysCopiesSampleData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysCopiesSampleData:"), value)
}/* debug [instance_properties/setter]: alwaysCopiesSampleData */


// The media type of samples that the output reads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/mediaType
func (a_ AssetReaderOutput) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// A Boolean value that indicates whether the output supports reconfiguring the time ranges it reads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/supportsRandomAccess
func (a_ AssetReaderOutput) SupportsRandomAccess() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsRandomAccess"))
	return rv
}/* debug [instance_properties/getter]: supportsRandomAccess */


// A Boolean value that indicates whether the output supports reconfiguring the time ranges it reads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/supportsRandomAccess
func (a_ AssetReaderOutput) SetSupportsRandomAccess(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportsRandomAccess:"), value)
}/* debug [instance_properties/setter]: supportsRandomAccess */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetReaderOutput */



