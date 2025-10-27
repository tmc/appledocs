// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [localProxyCount] class.
var (
	LocalProxyCountClass     _localProxyCountClass
	LocalProxyCountClassOnce sync.Once
)

func getlocalProxyCountClass() _localProxyCountClass {
	LocalProxyCountClassOnce.Do(func() {
		LocalProxyCountClass = _localProxyCountClass{objc.GetClass("localProxyCount")}
	})
	return LocalProxyCountClass
}

type _localProxyCountClass struct {
	class objc.Class
}





// An interface definition for the [localProxyCount] class.
type IlocalProxyCount interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _localProxyCountClass) Alloc() localProxyCount {
	rv := objc.Send[localProxyCount](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _localProxyCountClass) New() localProxyCount {
	rv := objc.Send[localProxyCount](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ localProxyCount) Init() localProxyCount {
	rv := objc.Send[localProxyCount](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ localProxyCount) Autorelease() localProxyCount {
	rv := objc.Send[localProxyCount](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlocalProxyCount creates a new localProxyCount instance.
func NewlocalProxyCount() localProxyCount {
	return getlocalProxyCountClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/localProxyCount
type localProxyCount struct {
	objectivec.Object
}

// localProxyCountFrom constructs a [localProxyCount] from an unsafe.Pointer.
func localProxyCountFrom(ptr unsafe.Pointer) localProxyCount {
	return localProxyCount{objectivec.Object{objc.ID(ptr)}}
}































