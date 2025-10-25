// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class classInfoImported */


/* debug [class_header]: Header for classInfoImported */
// The class instance for the [classInfoImported] class.
var (
	ClassInfoImportedClass     _classInfoImportedClass
	ClassInfoImportedClassOnce sync.Once
)

func getclassInfoImportedClass() _classInfoImportedClass {
	ClassInfoImportedClassOnce.Do(func() {
		ClassInfoImportedClass = _classInfoImportedClass{objc.GetClass("classInfoImported")}
	})
	return ClassInfoImportedClass
}

type _classInfoImportedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for classInfoImported */
// An interface definition for the [classInfoImported] class.
type IclassInfoImported interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for classInfoImported */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for classInfoImported */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for classInfoImported */
// Alloc allocates a new instance without initialization.
func (cc _classInfoImportedClass) Alloc() classInfoImported {
	rv := objc.Send[classInfoImported](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _classInfoImportedClass) New() classInfoImported {
	rv := objc.Send[classInfoImported](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ classInfoImported) Init() classInfoImported {
	rv := objc.Send[classInfoImported](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ classInfoImported) Autorelease() classInfoImported {
	rv := objc.Send[classInfoImported](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewclassInfoImported creates a new classInfoImported instance.
func NewclassInfoImported() classInfoImported {
	return getclassInfoImportedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for classInfoImported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/classInfoImported
type classInfoImported struct {
	objectivec.Object
}

// classInfoImportedFrom constructs a [classInfoImported] from an unsafe.Pointer.
func classInfoImportedFrom(ptr unsafe.Pointer) classInfoImported {
	return classInfoImported{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for classInfoImported *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for classInfoImported */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for classInfoImported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for classInfoImported */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for classInfoImported */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class classInfoImported */



