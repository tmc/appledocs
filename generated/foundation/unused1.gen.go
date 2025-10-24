// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class unused1 */


/* debug [class_header]: Header for unused1 */
// The class instance for the [unused1] class.
var (
	Unused1Class     _unused1Class
	Unused1ClassOnce sync.Once
)

func getunused1Class() _unused1Class {
	Unused1ClassOnce.Do(func() {
		Unused1Class = _unused1Class{objc.GetClass("unused1")}
	})
	return Unused1Class
}

type _unused1Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for unused1 */
// An interface definition for the [unused1] class.
type Iunused1 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for unused1 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for unused1 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for unused1 */
// Alloc allocates a new instance without initialization.
func (uc _unused1Class) Alloc() unused1 {
	rv := objc.Send[unused1](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _unused1Class) New() unused1 {
	rv := objc.Send[unused1](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ unused1) Init() unused1 {
	rv := objc.Send[unused1](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ unused1) Autorelease() unused1 {
	rv := objc.Send[unused1](u_.ID, objc.Sel("autorelease"))
	return rv
}

// Newunused1 creates a new unused1 instance.
func Newunused1() unused1 {
	return getunused1Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for unused1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unused1
type unused1 struct {
	objectivec.Object
}

// unused1From constructs a [unused1] from an unsafe.Pointer.
func unused1From(ptr unsafe.Pointer) unused1 {
	return unused1{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for unused1 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for unused1 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for unused1 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for unused1 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for unused1 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class unused1 */



