// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RegularExpression] class.
var (
	RegularExpressionClass     _RegularExpressionClass
	RegularExpressionClassOnce sync.Once
)

func getRegularExpressionClass() _RegularExpressionClass {
	RegularExpressionClassOnce.Do(func() {
		RegularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}
	})
	return RegularExpressionClass
}

type _RegularExpressionClass struct {
	class objc.Class
}

// An interface definition for the [RegularExpression] class.
type IRegularExpression interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type RegularExpression struct {
	objectivec.Object
}

// RegularExpressionFrom constructs a [RegularExpression] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func RegularExpressionFrom(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RegularExpressionClass) Alloc() RegularExpression {
	rv := objc.Send[RegularExpression](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RegularExpressionClass) New() RegularExpression {
	rv := objc.Send[RegularExpression](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RegularExpression) Init() RegularExpression {
	rv := objc.Send[RegularExpression](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RegularExpression) Autorelease() RegularExpression {
	rv := objc.Send[RegularExpression](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRegularExpression creates a new RegularExpression instance.
func NewRegularExpression() RegularExpression {
	return getRegularExpressionClass().New()
}




