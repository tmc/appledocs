// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class recordTypes */


/* debug [class_header]: Header for recordTypes */
// The class instance for the [recordTypes] class.
var (
	RecordTypesClass     _recordTypesClass
	RecordTypesClassOnce sync.Once
)

func getrecordTypesClass() _recordTypesClass {
	RecordTypesClassOnce.Do(func() {
		RecordTypesClass = _recordTypesClass{objc.GetClass("recordTypes")}
	})
	return RecordTypesClass
}

type _recordTypesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for recordTypes */
// An interface definition for the [recordTypes] class.
type IrecordTypes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for recordTypes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for recordTypes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for recordTypes */
// Alloc allocates a new instance without initialization.
func (rc _recordTypesClass) Alloc() recordTypes {
	rv := objc.Send[recordTypes](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _recordTypesClass) New() recordTypes {
	rv := objc.Send[recordTypes](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ recordTypes) Init() recordTypes {
	rv := objc.Send[recordTypes](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ recordTypes) Autorelease() recordTypes {
	rv := objc.Send[recordTypes](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrecordTypes creates a new recordTypes instance.
func NewrecordTypes() recordTypes {
	return getrecordTypesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for recordTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordTypes-c.ivar
type recordTypes struct {
	objectivec.Object
}

// recordTypesFrom constructs a [recordTypes] from an unsafe.Pointer.
func recordTypesFrom(ptr unsafe.Pointer) recordTypes {
	return recordTypes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for recordTypes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for recordTypes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for recordTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for recordTypes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for recordTypes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class recordTypes */



