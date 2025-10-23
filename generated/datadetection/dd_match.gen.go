// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DDMatch] class.
var (
	DDMatchClass     _DDMatchClass
	DDMatchClassOnce sync.Once
)

func getDDMatchClass() _DDMatchClass {
	DDMatchClassOnce.Do(func() {
		DDMatchClass = _DDMatchClass{objc.GetClass("DDMatch")}
	})
	return DDMatchClass
}

type _DDMatchClass struct {
	class objc.Class
}

// An interface definition for the [DDMatch] class.
type IDDMatch interface {
	objectivec.IObject
	// properties:
	MatchedString() string /* primitive/slice/pointer. */
	// methods:
}

// A base class for common types of data that the data detection system matches.
//
// The DataDetection framework returns results in objects that are subclasses of , which are specific to the type of matching data. Each object contains the matched string.


// A base class for common types of data that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatch
type DDMatch struct {
	objectivec.Object
}

// DDMatchFrom constructs a [DDMatch] from an unsafe.Pointer.
//
// A base class for common types of data that the data detection system matches.
func DDMatchFrom(ptr unsafe.Pointer) DDMatch {
	return DDMatch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchClass) Alloc() DDMatch {
	rv := objc.Send[DDMatch](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchClass) New() DDMatch {
	rv := objc.Send[DDMatch](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatch) Init() DDMatch {
	rv := objc.Send[DDMatch](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatch) Autorelease() DDMatch {
	rv := objc.Send[DDMatch](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatch creates a new DDMatch instance.
func NewDDMatch() DDMatch {
	return getDDMatchClass().New()
}



// A substring that the data detection system identifies from an original string as a common type of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatch/matchedString
func (d_ DDMatch) MatchedString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("matchedString"))
	return rv
}



