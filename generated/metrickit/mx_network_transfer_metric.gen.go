// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXNetworkTransferMetric */


/* debug [class_header]: Header for MXNetworkTransferMetric */
// The class instance for the [MXNetworkTransferMetric] class.
var (
	MXNetworkTransferMetricClass     _MXNetworkTransferMetricClass
	MXNetworkTransferMetricClassOnce sync.Once
)

func getMXNetworkTransferMetricClass() _MXNetworkTransferMetricClass {
	MXNetworkTransferMetricClassOnce.Do(func() {
		MXNetworkTransferMetricClass = _MXNetworkTransferMetricClass{objc.GetClass("MXNetworkTransferMetric")}
	})
	return MXNetworkTransferMetricClass
}

type _MXNetworkTransferMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXNetworkTransferMetric */
// An interface definition for the [MXNetworkTransferMetric] class.
type IMXNetworkTransferMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXNetworkTransferMetric */
	// properties:
	CumulativeCellularDownload() unsafe.Pointer
	CumulativeCellularUpload() unsafe.Pointer
	CumulativeWifiDownload() unsafe.Pointer
	CumulativeWifiUpload() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXNetworkTransferMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXNetworkTransferMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXNetworkTransferMetricClass) Alloc() MXNetworkTransferMetric {
	rv := objc.Send[MXNetworkTransferMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXNetworkTransferMetricClass) New() MXNetworkTransferMetric {
	rv := objc.Send[MXNetworkTransferMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXNetworkTransferMetric) Init() MXNetworkTransferMetric {
	rv := objc.Send[MXNetworkTransferMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXNetworkTransferMetric) Autorelease() MXNetworkTransferMetric {
	rv := objc.Send[MXNetworkTransferMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXNetworkTransferMetric creates a new MXNetworkTransferMetric instance.
func NewMXNetworkTransferMetric() MXNetworkTransferMetric {
	return getMXNetworkTransferMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXNetworkTransferMetric */
// An object representing metrics about network transfers.


// An object representing metrics about network transfers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric
type MXNetworkTransferMetric struct {
	MXMetric
}

// MXNetworkTransferMetricFrom constructs a [MXNetworkTransferMetric] from an unsafe.Pointer.
//
// An object representing metrics about network transfers.
func MXNetworkTransferMetricFrom(ptr unsafe.Pointer) MXNetworkTransferMetric {
	return MXNetworkTransferMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXNetworkTransferMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXNetworkTransferMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXNetworkTransferMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXNetworkTransferMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXNetworkTransferMetric */

// The total amount of data downloaded over the cellular connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeCellularDownload
func (m_ MXNetworkTransferMetric) CumulativeCellularDownload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCellularDownload"))
	return rv
}/* debug [instance_properties/getter]: cumulativeCellularDownload */


// The total amount of data uploaded over the cellular connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeCellularUpload
func (m_ MXNetworkTransferMetric) CumulativeCellularUpload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCellularUpload"))
	return rv
}/* debug [instance_properties/getter]: cumulativeCellularUpload */


// The total amount of data downloaded over the WiFi connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeWifiDownload
func (m_ MXNetworkTransferMetric) CumulativeWifiDownload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeWifiDownload"))
	return rv
}/* debug [instance_properties/getter]: cumulativeWifiDownload */


// The total amount of data uploaded over the WiFi connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeWifiUpload
func (m_ MXNetworkTransferMetric) CumulativeWifiUpload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeWifiUpload"))
	return rv
}/* debug [instance_properties/getter]: cumulativeWifiUpload */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXNetworkTransferMetric */



