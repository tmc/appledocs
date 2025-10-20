// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [LogicalTest] class.
type ILogicalTest interface {
	IScriptWhoseTest
}

// The logical combination of one or more specifier tests.
//
// Instances of this class perform logical operations of , , and on Boolean expressions represented by objects. These operators are equivalent to “ ”, “ ”, and “ ” in the C language. For and operations, an object is typically initialized with an array containing two or more objects. —inherited from —evaluates the array in a manner appropriate to the logical operation. For operations, an object is initialized with only one object; it simply reverses the Boolean outcome of the method. You don’t normally subclass .
//
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

// Alloc allocates a new instance without initialization.
func (lc _LogicalTestClass) Alloc() LogicalTest {
	rv := objc.Send[LogicalTest](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns an object initialized to perform an operation with the objects in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(orTestWith:)
func NewLogicalTestOrTestWithTests(subTests unsafe.Pointer) LogicalTest {
	instance := getLogicalTestClass().Alloc()
	rv := objc.Send[LogicalTest](instance.ID, objc.Sel("initOrTestWithTests:"), subTests)
	rv.Autorelease()
	return rv
}

// Returns an object initialized to perform an operation with the objects in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(andTestWith:)
func NewLogicalTestAndTestWithTests(subTests unsafe.Pointer) LogicalTest {
	instance := getLogicalTestClass().Alloc()
	rv := objc.Send[LogicalTest](instance.ID, objc.Sel("initAndTestWithTests:"), subTests)
	rv.Autorelease()
	return rv
}



