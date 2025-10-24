// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mFolderListingFileRef */


/* debug [class_header]: Header for mFolderListingFileRef */
// The class instance for the [mFolderListingFileRef] class.
var (
	MFolderListingFileRefClass     _mFolderListingFileRefClass
	MFolderListingFileRefClassOnce sync.Once
)

func getmFolderListingFileRefClass() _mFolderListingFileRefClass {
	MFolderListingFileRefClassOnce.Do(func() {
		MFolderListingFileRefClass = _mFolderListingFileRefClass{objc.GetClass("mFolderListingFileRef")}
	})
	return MFolderListingFileRefClass
}

type _mFolderListingFileRefClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mFolderListingFileRef */
// An interface definition for the [mFolderListingFileRef] class.
type ImFolderListingFileRef interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mFolderListingFileRef */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mFolderListingFileRef */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mFolderListingFileRef */
// Alloc allocates a new instance without initialization.
func (mc _mFolderListingFileRefClass) Alloc() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mFolderListingFileRefClass) New() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mFolderListingFileRef) Init() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mFolderListingFileRef) Autorelease() mFolderListingFileRef {
	rv := objc.Send[mFolderListingFileRef](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmFolderListingFileRef creates a new mFolderListingFileRef instance.
func NewmFolderListingFileRef() mFolderListingFileRef {
	return getmFolderListingFileRefClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mFolderListingFileRef */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mFolderListingFileRef
type mFolderListingFileRef struct {
	objectivec.Object
}

// mFolderListingFileRefFrom constructs a [mFolderListingFileRef] from an unsafe.Pointer.
func mFolderListingFileRefFrom(ptr unsafe.Pointer) mFolderListingFileRef {
	return mFolderListingFileRef{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mFolderListingFileRef *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mFolderListingFileRef */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mFolderListingFileRef */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mFolderListingFileRef */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mFolderListingFileRef */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mFolderListingFileRef */



