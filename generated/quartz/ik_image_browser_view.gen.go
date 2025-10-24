// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IKImageBrowserView] class.
var (
	IKImageBrowserViewClass     _IKImageBrowserViewClass
	IKImageBrowserViewClassOnce sync.Once
)

func getIKImageBrowserViewClass() _IKImageBrowserViewClass {
	IKImageBrowserViewClassOnce.Do(func() {
		IKImageBrowserViewClass = _IKImageBrowserViewClass{objc.GetClass("IKImageBrowserView")}
	})
	return IKImageBrowserViewClass
}

type _IKImageBrowserViewClass struct {
	class objc.Class
}

// An interface definition for the [IKImageBrowserView] class.
type IIKImageBrowserView interface {
	appkit.IView
	// properties:
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	// methods:
	SetDropIndexDropOperation(index int, operation unsafe.Pointer)
}

// A view for displaying and browsing a large collection of images and movies.
//
// The class is a view for displaying and browsing a large amount of images and movies efficiently. This class will be deprecated in a future release. Please switch to instead. You must set a datasource for the view and implement, at a minimum, the and described in . The items must conform to the IKImageBrowserItem Protocol protocol. The class’s delegate object must conform to IKImageBrowserDelegate Protocol protocol. It receives notification of changes in selection, as well as mouse events in the cells.


// A view for displaying and browsing a large collection of images and movies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView
type IKImageBrowserView struct {
	appkit.View
}

// IKImageBrowserViewFrom constructs a [IKImageBrowserView] from an unsafe.Pointer.
//
// A view for displaying and browsing a large collection of images and movies.
func IKImageBrowserViewFrom(ptr unsafe.Pointer) IKImageBrowserView {
	return IKImageBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IKImageBrowserViewClass) Alloc() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IKImageBrowserViewClass) New() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageBrowserView) Init() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageBrowserView) Autorelease() IKImageBrowserView {
	rv := objc.Send[IKImageBrowserView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageBrowserView creates a new IKImageBrowserView instance.
func NewIKImageBrowserView() IKImageBrowserView {
	return getIKImageBrowserViewClass().New()
}



// Allows the class to retarget the drop action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserView/setDrop(_:dropOperation:)
func (i_ IKImageBrowserView) SetDropIndexDropOperation(index int, operation unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDropIndex:dropOperation:"), index, operation)
}


// Returns the data source of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowserview/datasource
func (i_ IKImageBrowserView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dataSource"))
	return rv
}


// Returns the data source of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowserview/datasource
func (i_ IKImageBrowserView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataSource:"), value)
}


// Returns the delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowserview/delegate
func (i_ IKImageBrowserView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}


// Returns the delegate of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowserview/delegate
func (i_ IKImageBrowserView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}



