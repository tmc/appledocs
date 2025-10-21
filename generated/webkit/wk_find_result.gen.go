// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FindResult] class.
var (
	FindResultClass     _FindResultClass
	FindResultClassOnce sync.Once
)

func getFindResultClass() _FindResultClass {
	FindResultClassOnce.Do(func() {
		FindResultClass = _FindResultClass{objc.GetClass("WKFindResult")}
	})
	return FindResultClass
}

type _FindResultClass struct {
	class objc.Class
}

// An interface definition for the [FindResult] class.
type IFindResult interface {
	objectivec.IObject
}

// An object that contains the results of searching the web view’s contents.
//
// When you perform a search using the methods of , the web view creates a object and delivers it to your completion handler. You don’t create instances of this class directly. Use the objects that the web view provides to determine whether it found a match for the content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindResult
type FindResult struct {
	objectivec.Object
}

// FindResultFrom constructs a [FindResult] from an unsafe.Pointer.
//
// An object that contains the results of searching the web view’s contents.
func FindResultFrom(ptr unsafe.Pointer) FindResult {
	return FindResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FindResultClass) Alloc() FindResult {
	rv := objc.Send[FindResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FindResultClass) New() FindResult {
	rv := objc.Send[FindResult](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FindResult) Init() FindResult {
	rv := objc.Send[FindResult](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FindResult) Autorelease() FindResult {
	rv := objc.Send[FindResult](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFindResult creates a new FindResult instance.
func NewFindResult() FindResult {
	return getFindResultClass().New()
}


// A Boolean value that indicates whether the web view found a match during the search.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindresult/matchfound
func (f_ FindResult) MatchFound() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("matchFound"))
	return rv
}


// SetMatchFound sets the value of the matchFound property.
// A Boolean value that indicates whether the web view found a match during the search.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindresult/matchfound
func (f_ FindResult) SetMatchFound(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMatchFound:"), value)
}



