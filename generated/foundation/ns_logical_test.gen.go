// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLogicalTest */


/* debug [class_header]: Header for NSLogicalTest */
// The class instance for the [LogicalTest] class.
var (
	LogicalTestClass     _LogicalTestClass
	LogicalTestClassOnce sync.Once
)

func getLogicalTestClass() _LogicalTestClass {
	LogicalTestClassOnce.Do(func() {
		LogicalTestClass = _LogicalTestClass{objc.GetClass("NSLogicalTest")}
	})
	return LogicalTestClass
}

type _LogicalTestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LogicalTest */
// An interface definition for the [LogicalTest] class.
type ILogicalTest interface {
	IScriptWhoseTest
	
/* debug [class_interface_properties]: Properties for LogicalTest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LogicalTest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LogicalTest */
// Alloc allocates a new instance without initialization.
func (lc _LogicalTestClass) Alloc() LogicalTest {
	rv := objc.Send[LogicalTest](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LogicalTestClass) New() LogicalTest {
	rv := objc.Send[LogicalTest](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogicalTest) Init() LogicalTest {
	rv := objc.Send[LogicalTest](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogicalTest) Autorelease() LogicalTest {
	rv := objc.Send[LogicalTest](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogicalTest creates a new LogicalTest instance.
func NewLogicalTest() LogicalTest {
	return getLogicalTestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LogicalTest */
// The logical combination of one or more specifier tests.
//
// Instances of this class perform logical operations of , , and on Boolean expressions represented by objects. These operators are equivalent to “ ”, “ ”, and “ ” in the C language. For and operations, an object is typically initialized with an array containing two or more objects. —inherited from —evaluates the array in a manner appropriate to the logical operation. For operations, an object is initialized with only one object; it simply reverses the Boolean outcome of the method. You don’t normally subclass .


// The logical combination of one or more specifier tests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest
type LogicalTest struct {
	ScriptWhoseTest
}

// LogicalTestFrom constructs a [LogicalTest] from an unsafe.Pointer.
//
// The logical combination of one or more specifier tests.
func LogicalTestFrom(ptr unsafe.Pointer) LogicalTest {
	return LogicalTest{
		ScriptWhoseTest: ScriptWhoseTestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LogicalTest */

// Returns an object initialized to perform an operation with the objects in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(andTestWith:)
func NewLogicalTestAndTestWithTests(subTests []SpecifierTest) LogicalTest {
	instance := getLogicalTestClass().Alloc()
	rv := objc.Send[LogicalTest](instance.ID, objc.Sel("initAndTestWithTests:"), subTests)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLogicalTestAndTestWithTests */


// Returns an object initialized to perform a operation on the given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(notTestWith:)
func NewLogicalTestNotTestWithTest(subTest IScriptWhoseTest) LogicalTest {
	instance := getLogicalTestClass().Alloc()
	rv := objc.Send[LogicalTest](instance.ID, objc.Sel("initNotTestWithTest:"), subTest)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLogicalTestNotTestWithTest */


// Returns an object initialized to perform an operation with the objects in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(orTestWith:)
func NewLogicalTestOrTestWithTests(subTests []SpecifierTest) LogicalTest {
	instance := getLogicalTestClass().Alloc()
	rv := objc.Send[LogicalTest](instance.ID, objc.Sel("initOrTestWithTests:"), subTests)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLogicalTestOrTestWithTests */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LogicalTest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LogicalTest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LogicalTest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LogicalTest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLogicalTest */


