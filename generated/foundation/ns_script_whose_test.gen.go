// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScriptWhoseTest */


/* debug [class_header]: Header for NSScriptWhoseTest */
// The class instance for the [ScriptWhoseTest] class.
var (
	ScriptWhoseTestClass     _ScriptWhoseTestClass
	ScriptWhoseTestClassOnce sync.Once
)

func getScriptWhoseTestClass() _ScriptWhoseTestClass {
	ScriptWhoseTestClassOnce.Do(func() {
		ScriptWhoseTestClass = _ScriptWhoseTestClass{objc.GetClass("NSScriptWhoseTest")}
	})
	return ScriptWhoseTestClass
}

type _ScriptWhoseTestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScriptWhoseTest */
// An interface definition for the [ScriptWhoseTest] class.
type IScriptWhoseTest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScriptWhoseTest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScriptWhoseTest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScriptWhoseTest */
// Alloc allocates a new instance without initialization.
func (sc _ScriptWhoseTestClass) Alloc() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScriptWhoseTestClass) New() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptWhoseTest) Init() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptWhoseTest) Autorelease() ScriptWhoseTest {
	rv := objc.Send[ScriptWhoseTest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptWhoseTest creates a new ScriptWhoseTest instance.
func NewScriptWhoseTest() ScriptWhoseTest {
	return getScriptWhoseTestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScriptWhoseTest */
// An abstract class that provides the basis for testing specifiers one at a time or in groups.
//
// is an abstract class whose sole method is . Two concrete subclasses of generate objects representing Boolean expressions comparing one object with another and objects representing multiple Boolean expressions connected by logical operators ( , , ). These classes are, respectively, and . In evaluating itself, an invokes the method of its “test” object. You shouldn’t need to subclass , and you should rarely need to subclass one of its subclasses.


// An abstract class that provides the basis for testing specifiers one at a time or in groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest
type ScriptWhoseTest struct {
	objectivec.Object
}

// ScriptWhoseTestFrom constructs a [ScriptWhoseTest] from an unsafe.Pointer.
//
// An abstract class that provides the basis for testing specifiers one at a time or in groups.
func ScriptWhoseTestFrom(ptr unsafe.Pointer) ScriptWhoseTest {
	return ScriptWhoseTest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScriptWhoseTest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest/init(coder:)
func NewScriptWhoseTestWithCoder(inCoder ICoder) ScriptWhoseTest {
	instance := getScriptWhoseTestClass().Alloc()
	rv := objc.Send[ScriptWhoseTest](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScriptWhoseTestWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScriptWhoseTest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScriptWhoseTest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScriptWhoseTest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScriptWhoseTest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScriptWhoseTest */


