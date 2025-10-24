// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [systemVersion] class.
var (
	SystemVersionClass     _systemVersionClass
	SystemVersionClassOnce sync.Once
)

func getsystemVersionClass() _systemVersionClass {
	SystemVersionClassOnce.Do(func() {
		SystemVersionClass = _systemVersionClass{objc.GetClass("systemVersion")}
	})
	return SystemVersionClass
}

type _systemVersionClass struct {
	class objc.Class
}





// An interface definition for the [systemVersion] class.
type IsystemVersion interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _systemVersionClass) Alloc() systemVersion {
	rv := objc.Send[systemVersion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _systemVersionClass) New() systemVersion {
	rv := objc.Send[systemVersion](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ systemVersion) Init() systemVersion {
	rv := objc.Send[systemVersion](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ systemVersion) Autorelease() systemVersion {
	rv := objc.Send[systemVersion](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsystemVersion creates a new systemVersion instance.
func NewsystemVersion() systemVersion {
	return getsystemVersionClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/systemVersion-c.ivar
type systemVersion struct {
	objectivec.Object
}

// systemVersionFrom constructs a [systemVersion] from an unsafe.Pointer.
func systemVersionFrom(ptr unsafe.Pointer) systemVersion {
	return systemVersion{objectivec.Object{objc.ID(ptr)}}
}































