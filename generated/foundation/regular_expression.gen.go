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
	regularExpressionClass     _RegularExpressionClass
	regularExpressionClassOnce sync.Once
)

func getRegularExpressionClass() _RegularExpressionClass {
	regularExpressionClassOnce.Do(func() {
		regularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}
	})
	return regularExpressionClass
}

type _RegularExpressionClass struct {
	class objc.Class
}

// An interface definition for the [RegularExpression] class.
type IRegularExpression interface {
	objectivec.IObject
}

// An immutable representation of a compiled regular expression that you apply to Unicode strings.
//
// The fundamental matching method for is a Block iterator method that allows clients to supply a Block object which will be invoked each time the regular expression matches a portion of the target string. There are additional convenience methods for returning all the matches as an array, the total number of matches, the first match, and the range of the first match. An individual match is represented by an instance of the class, which carries information about the overall matched range (via its property), and the range of each individual capture group (via the method). For basic objects, these match results will be of type , but subclasses may use other types.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression
type RegularExpression struct {
	objectivec.Object
}

// RegularExpressionFrom constructs a [RegularExpression] from an unsafe.Pointer.
//
// An immutable representation of a compiled regular expression that you apply to Unicode strings.
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




