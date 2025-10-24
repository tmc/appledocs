// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRContentLauncherClusterContentSearch] class.
var (
	MTRContentLauncherClusterContentSearchClass     _MTRContentLauncherClusterContentSearchClass
	MTRContentLauncherClusterContentSearchClassOnce sync.Once
)

func getMTRContentLauncherClusterContentSearchClass() _MTRContentLauncherClusterContentSearchClass {
	MTRContentLauncherClusterContentSearchClassOnce.Do(func() {
		MTRContentLauncherClusterContentSearchClass = _MTRContentLauncherClusterContentSearchClass{objc.GetClass("MTRContentLauncherClusterContentSearch")}
	})
	return MTRContentLauncherClusterContentSearchClass
}

type _MTRContentLauncherClusterContentSearchClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterContentSearch] class.
type IMTRContentLauncherClusterContentSearch interface {
	IMTRContentLauncherClusterContentSearchStruct
	// properties:
	ParameterList() unsafe.Pointer
	SetParameterList(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearch
type MTRContentLauncherClusterContentSearch struct {
	MTRContentLauncherClusterContentSearchStruct
}

// MTRContentLauncherClusterContentSearchFrom constructs a [MTRContentLauncherClusterContentSearch] from an unsafe.Pointer.
func MTRContentLauncherClusterContentSearchFrom(ptr unsafe.Pointer) MTRContentLauncherClusterContentSearch {
	return MTRContentLauncherClusterContentSearch{
		MTRContentLauncherClusterContentSearchStruct: MTRContentLauncherClusterContentSearchStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterContentSearchClass) Alloc() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterContentSearchClass) New() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterContentSearch) Init() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterContentSearch) Autorelease() MTRContentLauncherClusterContentSearch {
	rv := objc.Send[MTRContentLauncherClusterContentSearch](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterContentSearch creates a new MTRContentLauncherClusterContentSearch instance.
func NewMTRContentLauncherClusterContentSearch() MTRContentLauncherClusterContentSearch {
	return getMTRContentLauncherClusterContentSearchClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclustercontentsearch/parameterlist
func (m_ MTRContentLauncherClusterContentSearch) ParameterList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("parameterList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclustercontentsearch/parameterlist
func (m_ MTRContentLauncherClusterContentSearch) SetParameterList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameterList:"), value)
}



