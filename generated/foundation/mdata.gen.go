// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mdata */


/* debug [class_header]: Header for mdata */
// The class instance for the [mdata] class.
var (
	MdataClass     _mdataClass
	MdataClassOnce sync.Once
)

func getmdataClass() _mdataClass {
	MdataClassOnce.Do(func() {
		MdataClass = _mdataClass{objc.GetClass("mdata")}
	})
	return MdataClass
}

type _mdataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mdata */
// An interface definition for the [mdata] class.
type Imdata interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mdata */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mdata */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mdata */
// Alloc allocates a new instance without initialization.
func (mc _mdataClass) Alloc() mdata {
	rv := objc.Send[mdata](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mdataClass) New() mdata {
	rv := objc.Send[mdata](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mdata) Init() mdata {
	rv := objc.Send[mdata](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mdata) Autorelease() mdata {
	rv := objc.Send[mdata](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmdata creates a new mdata instance.
func Newmdata() mdata {
	return getmdataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mdata */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/mdata
type mdata struct {
	objectivec.Object
}

// mdataFrom constructs a [mdata] from an unsafe.Pointer.
func mdataFrom(ptr unsafe.Pointer) mdata {
	return mdata{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mdata *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mdata */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mdata */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mdata */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mdata */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mdata */



