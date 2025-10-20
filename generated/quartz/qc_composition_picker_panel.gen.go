// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [QCCompositionPickerPanel] class.
var (
	QCCompositionPickerPanelClass     _QCCompositionPickerPanelClass
	QCCompositionPickerPanelClassOnce sync.Once
)

func getQCCompositionPickerPanelClass() _QCCompositionPickerPanelClass {
	QCCompositionPickerPanelClassOnce.Do(func() {
		QCCompositionPickerPanelClass = _QCCompositionPickerPanelClass{objc.GetClass("QCCompositionPickerPanel")}
	})
	return QCCompositionPickerPanelClass
}

type _QCCompositionPickerPanelClass struct {
	class objc.Class
}

// An interface definition for the [QCCompositionPickerPanel] class.
type IQCCompositionPickerPanel interface {
	appkit.IPanel
	CompositionPickerView() unsafe.Pointer
}

// The class represents a utility window that allows users to browse compositions that are in the Quartz Composer composition repository and, if supported, preview the composition. The class cannot be subclassed.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerPanel
type QCCompositionPickerPanel struct {
	appkit.Panel
}

// QCCompositionPickerPanelFrom constructs a [QCCompositionPickerPanel] from an unsafe.Pointer.
//
// The class represents a utility window that allows users to browse compositions that are in the Quartz Composer composition repository and, if supported, preview the composition. The class cannot be subclassed.
func QCCompositionPickerPanelFrom(ptr unsafe.Pointer) QCCompositionPickerPanel {
	return QCCompositionPickerPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QCCompositionPickerPanelClass) Alloc() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCCompositionPickerPanelClass) New() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionPickerPanel) Init() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionPickerPanel) Autorelease() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionPickerPanel creates a new QCCompositionPickerPanel instance.
func NewQCCompositionPickerPanel() QCCompositionPickerPanel {
	return getQCCompositionPickerPanelClass().New()
}


// Returns the composition picker view used by the panel so that it can be configured.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerPanel/compositionPickerView()
func (q_ QCCompositionPickerPanel) CompositionPickerView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](q_.ID, objc.Sel("compositionPickerView"))
	return rv
}



