// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MXNetworkTransferMetric] class.
type IMXNetworkTransferMetric interface {
	IMXMetric
	// properties:
	CumulativeCellularDownload() unsafe.Pointer
	CumulativeCellularUpload() unsafe.Pointer
	CumulativeWifiDownload() unsafe.Pointer
	CumulativeWifiUpload() unsafe.Pointer
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MXNetworkTransferMetricClass) Alloc() MXNetworkTransferMetric {
	rv := objc.Send[MXNetworkTransferMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The total amount of data downloaded over the cellular connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeCellularDownload
func (m_ MXNetworkTransferMetric) CumulativeCellularDownload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCellularDownload"))
	return rv
}


// The total amount of data uploaded over the cellular connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeCellularUpload
func (m_ MXNetworkTransferMetric) CumulativeCellularUpload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCellularUpload"))
	return rv
}


// The total amount of data downloaded over the WiFi connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeWifiDownload
func (m_ MXNetworkTransferMetric) CumulativeWifiDownload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeWifiDownload"))
	return rv
}


// The total amount of data uploaded over the WiFi connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXNetworkTransferMetric/cumulativeWifiUpload
func (m_ MXNetworkTransferMetric) CumulativeWifiUpload() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeWifiUpload"))
	return rv
}



