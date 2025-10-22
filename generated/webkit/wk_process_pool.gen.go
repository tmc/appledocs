// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ProcessPool] class.
var (
	ProcessPoolClass     _ProcessPoolClass
	ProcessPoolClassOnce sync.Once
)

func getProcessPoolClass() _ProcessPoolClass {
	ProcessPoolClassOnce.Do(func() {
		ProcessPoolClass = _ProcessPoolClass{objc.GetClass("WKProcessPool")}
	})
	return ProcessPoolClass
}

type _ProcessPoolClass struct {
	class objc.Class
}

// An interface definition for the [ProcessPool] class.
type IProcessPool interface {
	objectivec.IObject
	ProcessPool() WKProcessPool
	SetProcessPool(value IWKProcessPool)
}

// An opaque token that you use to run multiple web views in a single process.
//
// A object represents a single process that WebKit uses to manage web content. To provide a more secure and stable experience, WebKit renders the content of web views in separate processes, rather than in your app’s process space. By default, WebKit gives each web view its own process space until it reaches an implementation-defined process limit. After that, web views with the same object share the same web content process. If your app creates multiple web views, assign the same object to web views that may safely share a process space. Instantiate an instance of this class and assign it to the property of each web view’s object.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKProcessPool
type ProcessPool struct {
	objectivec.Object
}

// ProcessPoolFrom constructs a [ProcessPool] from an unsafe.Pointer.
//
// An opaque token that you use to run multiple web views in a single process.
func ProcessPoolFrom(ptr unsafe.Pointer) ProcessPool {
	return ProcessPool{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProcessPoolClass) Alloc() ProcessPool {
	rv := objc.Send[ProcessPool](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProcessPoolClass) New() ProcessPool {
	rv := objc.Send[ProcessPool](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProcessPool) Init() ProcessPool {
	rv := objc.Send[ProcessPool](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProcessPool) Autorelease() ProcessPool {
	rv := objc.Send[ProcessPool](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProcessPool creates a new ProcessPool instance.
func NewProcessPool() ProcessPool {
	return getProcessPoolClass().New()
}


// The object that coordinates the processes the web view uses to render its web content and execute scripts.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/processpool
func (p_ ProcessPool) ProcessPool() WKProcessPool {
	rv := objc.Send[WKProcessPool](p_.ID, objc.Sel("processPool"))
	return rv
}


// SetProcessPool sets the value of the processPool property.
// The object that coordinates the processes the web view uses to render its web content and execute scripts.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/processpool
func (p_ ProcessPool) SetProcessPool(value IWKProcessPool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessPool:"), value)
}



