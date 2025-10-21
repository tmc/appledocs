// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QCCompositionParameterView] class.
var (
	QCCompositionParameterViewClass     _QCCompositionParameterViewClass
	QCCompositionParameterViewClassOnce sync.Once
)

func getQCCompositionParameterViewClass() _QCCompositionParameterViewClass {
	QCCompositionParameterViewClassOnce.Do(func() {
		QCCompositionParameterViewClass = _QCCompositionParameterViewClass{objc.GetClass("QCCompositionParameterView")}
	})
	return QCCompositionParameterViewClass
}

type _QCCompositionParameterViewClass struct {
	class objc.Class
}

// An interface definition for the [QCCompositionParameterView] class.
type IQCCompositionParameterView interface {
	appkit.IView
	SetDelegate(delegate objectivec.IObject)
}

// A class that allows users to edit the input parameters of a composition in real time. The composition can be rendering in any of the following objects: , , or .
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionParameterView
type QCCompositionParameterView struct {
	appkit.View
}

// QCCompositionParameterViewFrom constructs a [QCCompositionParameterView] from an unsafe.Pointer.
//
// A class that allows users to edit the input parameters of a composition in real time. The composition can be rendering in any of the following objects: , , or .
func QCCompositionParameterViewFrom(ptr unsafe.Pointer) QCCompositionParameterView {
	return QCCompositionParameterView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QCCompositionParameterViewClass) Alloc() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCCompositionParameterViewClass) New() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionParameterView) Init() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionParameterView) Autorelease() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionParameterView creates a new QCCompositionParameterView instance.
func NewQCCompositionParameterView() QCCompositionParameterView {
	return getQCCompositionParameterViewClass().New()
}


// Sets the composition parameter view delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionParameterView/setDelegate(_:)
func (q_ QCCompositionParameterView) SetDelegate(delegate objectivec.IObject) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setDelegate:"), delegate)
}



