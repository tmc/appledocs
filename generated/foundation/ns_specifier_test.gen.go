// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSpecifierTest */


/* debug [class_header]: Header for NSSpecifierTest */
// The class instance for the [SpecifierTest] class.
var (
	SpecifierTestClass     _SpecifierTestClass
	SpecifierTestClassOnce sync.Once
)

func getSpecifierTestClass() _SpecifierTestClass {
	SpecifierTestClassOnce.Do(func() {
		SpecifierTestClass = _SpecifierTestClass{objc.GetClass("NSSpecifierTest")}
	})
	return SpecifierTestClass
}

type _SpecifierTestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpecifierTest */
// An interface definition for the [SpecifierTest] class.
type ISpecifierTest interface {
	IScriptWhoseTest
	
/* debug [class_interface_properties]: Properties for SpecifierTest */
	// properties:
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpecifierTest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpecifierTest */
// Alloc allocates a new instance without initialization.
func (sc _SpecifierTestClass) Alloc() SpecifierTest {
	rv := objc.Send[SpecifierTest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpecifierTestClass) New() SpecifierTest {
	rv := objc.Send[SpecifierTest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpecifierTest) Init() SpecifierTest {
	rv := objc.Send[SpecifierTest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpecifierTest) Autorelease() SpecifierTest {
	rv := objc.Send[SpecifierTest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpecifierTest creates a new SpecifierTest instance.
func NewSpecifierTest() SpecifierTest {
	return getSpecifierTestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpecifierTest */
// A comparison between an object specifier and a test object.
//
// Instances of this class represent a Boolean expression; they evaluate an object specifier and compare the resulting object to another object using a given comparison method. For more information on , see the method description for its sole public method, its initializer, . When an object is properly initialized, it holds two objects: A “value” or “test” object used as the basis of the comparison; this object can be a regular object or object specifier (such as “blue” in “words whose color is blue”). An object specifier evaluating to the container (“words”). The instance also encapsulates a selector identifying the method performing this comparison. The informal protocol defines a set of comparison methods useful for this purpose, while describes additional methods you may need to use for scripting. The test object is compared, using the selector, against each object in the container. Specifiers in these tests usually have invoked on their topmost container. You should rarely need to subclass .


// A comparison between an object specifier and a test object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSpecifierTest
type SpecifierTest struct {
	ScriptWhoseTest
}

// SpecifierTestFrom constructs a [SpecifierTest] from an unsafe.Pointer.
//
// A comparison between an object specifier and a test object.
func SpecifierTestFrom(ptr unsafe.Pointer) SpecifierTest {
	return SpecifierTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpecifierTest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpecifierTest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpecifierTest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpecifierTest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpecifierTest */

// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisobjectbeingtested
func (s_ SpecifierTest) ContainerIsObjectBeingTested() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containerIsObjectBeingTested"))
	return rv
}/* debug [instance_properties/getter]: containerIsObjectBeingTested */


// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisobjectbeingtested
func (s_ SpecifierTest) SetContainerIsObjectBeingTested(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContainerIsObjectBeingTested:"), value)
}/* debug [instance_properties/setter]: containerIsObjectBeingTested */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSpecifierTest */



