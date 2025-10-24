// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class errorCorrectionLevel */


/* debug [class_header]: Header for errorCorrectionLevel */
// The class instance for the [errorCorrectionLevel] class.
var (
	ErrorCorrectionLevelClass     _errorCorrectionLevelClass
	ErrorCorrectionLevelClassOnce sync.Once
)

func geterrorCorrectionLevelClass() _errorCorrectionLevelClass {
	ErrorCorrectionLevelClassOnce.Do(func() {
		ErrorCorrectionLevelClass = _errorCorrectionLevelClass{objc.GetClass("errorCorrectionLevel")}
	})
	return ErrorCorrectionLevelClass
}

type _errorCorrectionLevelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for errorCorrectionLevel */
// An interface definition for the [errorCorrectionLevel] class.
type IerrorCorrectionLevel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for errorCorrectionLevel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for errorCorrectionLevel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for errorCorrectionLevel */
// Alloc allocates a new instance without initialization.
func (ec _errorCorrectionLevelClass) Alloc() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _errorCorrectionLevelClass) New() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ errorCorrectionLevel) Init() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ errorCorrectionLevel) Autorelease() errorCorrectionLevel {
	rv := objc.Send[errorCorrectionLevel](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewerrorCorrectionLevel creates a new errorCorrectionLevel instance.
func NewerrorCorrectionLevel() errorCorrectionLevel {
	return geterrorCorrectionLevelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for errorCorrectionLevel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-c.ivar
type errorCorrectionLevel struct {
	objectivec.Object
}

// errorCorrectionLevelFrom constructs a [errorCorrectionLevel] from an unsafe.Pointer.
func errorCorrectionLevelFrom(ptr unsafe.Pointer) errorCorrectionLevel {
	return errorCorrectionLevel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for errorCorrectionLevel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for errorCorrectionLevel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for errorCorrectionLevel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for errorCorrectionLevel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for errorCorrectionLevel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class errorCorrectionLevel */



