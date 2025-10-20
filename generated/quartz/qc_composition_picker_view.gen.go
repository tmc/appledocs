// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [QCCompositionPickerView] class.
var (
	QCCompositionPickerViewClass     _QCCompositionPickerViewClass
	QCCompositionPickerViewClassOnce sync.Once
)

func getQCCompositionPickerViewClass() _QCCompositionPickerViewClass {
	QCCompositionPickerViewClassOnce.Do(func() {
		QCCompositionPickerViewClass = _QCCompositionPickerViewClass{objc.GetClass("QCCompositionPickerView")}
	})
	return QCCompositionPickerViewClass
}

type _QCCompositionPickerViewClass struct {
	class objc.Class
}

// An interface definition for the [QCCompositionPickerView] class.
type IQCCompositionPickerView interface {
	appkit.IView
	SetSelectedComposition(composition unsafe.Pointer)
	ShowsCompositionNames() bool
}

// The class allows users to browse compositions that are in the Quartz Composer composition repository, and to preview them. You can set the default input parameters for a composition preview by using the method setDefaultValue:forInputKey:.
//
// Note that the composition picker view does not automatically refresh its content when the composition repository is updated. It’s your responsibility to perform any necessary updating.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerView
type QCCompositionPickerView struct {
	appkit.View
}

// QCCompositionPickerViewFrom constructs a [QCCompositionPickerView] from an unsafe.Pointer.
//
// The class allows users to browse compositions that are in the Quartz Composer composition repository, and to preview them. You can set the default input parameters for a composition preview by using the method setDefaultValue:forInputKey:.
func QCCompositionPickerViewFrom(ptr unsafe.Pointer) QCCompositionPickerView {
	return QCCompositionPickerView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QCCompositionPickerViewClass) Alloc() QCCompositionPickerView {
	rv := objc.Send[QCCompositionPickerView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCCompositionPickerViewClass) New() QCCompositionPickerView {
	rv := objc.Send[QCCompositionPickerView](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionPickerView) Init() QCCompositionPickerView {
	rv := objc.Send[QCCompositionPickerView](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionPickerView) Autorelease() QCCompositionPickerView {
	rv := objc.Send[QCCompositionPickerView](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionPickerView creates a new QCCompositionPickerView instance.
func NewQCCompositionPickerView() QCCompositionPickerView {
	return getQCCompositionPickerViewClass().New()
}


// Sets a composition as selected in the composition picker view.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerView/setSelectedComposition(_:)
func (q_ QCCompositionPickerView) SetSelectedComposition(composition unsafe.Pointer) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setSelectedComposition:"), composition)
}

// Retrieves whether composition names can be shown in the composition picker view.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerView/showsCompositionNames()
func (q_ QCCompositionPickerView) ShowsCompositionNames() bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("showsCompositionNames"))
	return rv
}



