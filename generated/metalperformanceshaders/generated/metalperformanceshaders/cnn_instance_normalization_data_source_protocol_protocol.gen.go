// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNInstanceNormalizationDataSource is the MPSCNNInstanceNormalizationDataSource protocol interface.
//
// A protocol that defines methods that an instance normalization uses to initialize scale factors and bias terms.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 10.13.4+
//   - tvOS 11.3+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationDataSource
type PCNNInstanceNormalizationDataSource interface {
	// Required methods
	EncodeWithCoder(aCoder Coder /* not a class type */)/* debug [protocol_interface/required_method]: EncodeWithCoder */
	InitWithCoder(aDecoder Coder /* not a class type */) objectivec.IObject/* debug [protocol_interface/required_method]: InitWithCoder */
	Label()/* debug [protocol_interface/required_method]: Label */
	Beta()/* debug [protocol_interface/required_method]: Beta */
	Gamma()/* debug [protocol_interface/required_method]: Gamma */
	UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer unsafe.Pointer, instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState/* debug [protocol_interface/required_method]: UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch */
	UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) bool/* debug [protocol_interface/required_method]: UpdateGammaAndBetaWithInstanceNormalizationStateBatch */
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject/* debug [protocol_interface/required_method]: CopyWithZoneDevice */
	// Optional methods
	Encode()
	HasEncode() bool
	Epsilon()
	HasEpsilon() bool
	UpdateGammaAndBeta()
	HasUpdateGammaAndBeta() bool
	Copy()
	HasCopy() bool
	Load()
	HasLoad() bool
	Purge()
	HasPurge() bool
}

// CNNInstanceNormalizationDataSource is a delegate implementation builder for the PCNNInstanceNormalizationDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CNNInstanceNormalizationDataSource struct {
	_Encode func()
	_Epsilon func()
	_UpdateGammaAndBeta func()
	_Copy func()
	_Load func()
	_Purge func()
	_EncodeWithCoder func(aCoder Coder /* not a class type */)
	_InitWithCoder func(aDecoder Coder /* not a class type */) objectivec.IObject
	_Label func()
	_Beta func()
	_Gamma func()
	_UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch func(commandBuffer unsafe.Pointer, instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState
	_UpdateGammaAndBetaWithInstanceNormalizationStateBatch func(instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) bool
	_CopyWithZoneDevice func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
}

// SetEncode sets the handler for the Encode delegate method.
func (d *CNNInstanceNormalizationDataSource) SetEncode(f func()) {
	d._Encode = f
}

// SetEpsilon sets the handler for the Epsilon delegate method.
func (d *CNNInstanceNormalizationDataSource) SetEpsilon(f func()) {
	d._Epsilon = f
}

// SetUpdateGammaAndBeta sets the handler for the UpdateGammaAndBeta delegate method.
func (d *CNNInstanceNormalizationDataSource) SetUpdateGammaAndBeta(f func()) {
	d._UpdateGammaAndBeta = f
}

// SetCopy sets the handler for the Copy delegate method.
func (d *CNNInstanceNormalizationDataSource) SetCopy(f func()) {
	d._Copy = f
}

// SetLoad sets the handler for the Load delegate method.
func (d *CNNInstanceNormalizationDataSource) SetLoad(f func()) {
	d._Load = f
}

// SetPurge sets the handler for the Purge delegate method.
func (d *CNNInstanceNormalizationDataSource) SetPurge(f func()) {
	d._Purge = f
}

// SetEncodeWithCoder sets the handler for the EncodeWithCoder delegate method.
func (d *CNNInstanceNormalizationDataSource) SetEncodeWithCoder(f func(aCoder Coder /* not a class type */)) {
	d._EncodeWithCoder = f
}

// SetInitWithCoder sets the handler for the InitWithCoder delegate method.
func (d *CNNInstanceNormalizationDataSource) SetInitWithCoder(f func(aDecoder Coder /* not a class type */) objectivec.IObject) {
	d._InitWithCoder = f
}

// SetLabel sets the handler for the Label delegate method.
func (d *CNNInstanceNormalizationDataSource) SetLabel(f func()) {
	d._Label = f
}

// SetBeta sets the handler for the Beta delegate method.
func (d *CNNInstanceNormalizationDataSource) SetBeta(f func()) {
	d._Beta = f
}

// SetGamma sets the handler for the Gamma delegate method.
func (d *CNNInstanceNormalizationDataSource) SetGamma(f func()) {
	d._Gamma = f
}

// SetUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch sets the handler for the UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch delegate method.
func (d *CNNInstanceNormalizationDataSource) SetUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(f func(commandBuffer unsafe.Pointer, instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState) {
	d._UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch = f
}

// SetUpdateGammaAndBetaWithInstanceNormalizationStateBatch sets the handler for the UpdateGammaAndBetaWithInstanceNormalizationStateBatch delegate method.
func (d *CNNInstanceNormalizationDataSource) SetUpdateGammaAndBetaWithInstanceNormalizationStateBatch(f func(instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) bool) {
	d._UpdateGammaAndBetaWithInstanceNormalizationStateBatch = f
}

// SetCopyWithZoneDevice sets the handler for the CopyWithZoneDevice delegate method.
func (d *CNNInstanceNormalizationDataSource) SetCopyWithZoneDevice(f func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject) {
	d._CopyWithZoneDevice = f
}

// Encode implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Encode() {
	if d._Encode != nil {
		d._Encode()
	}
}

// HasEncode returns true if a handler for Encode has been set.
func (d *CNNInstanceNormalizationDataSource) HasEncode() bool {
	return d._Encode != nil
}

// Epsilon implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Epsilon() {
	if d._Epsilon != nil {
		d._Epsilon()
	}
}

// HasEpsilon returns true if a handler for Epsilon has been set.
func (d *CNNInstanceNormalizationDataSource) HasEpsilon() bool {
	return d._Epsilon != nil
}

// UpdateGammaAndBeta implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) UpdateGammaAndBeta() {
	if d._UpdateGammaAndBeta != nil {
		d._UpdateGammaAndBeta()
	}
}

// HasUpdateGammaAndBeta returns true if a handler for UpdateGammaAndBeta has been set.
func (d *CNNInstanceNormalizationDataSource) HasUpdateGammaAndBeta() bool {
	return d._UpdateGammaAndBeta != nil
}

// Copy implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Copy() {
	if d._Copy != nil {
		d._Copy()
	}
}

// HasCopy returns true if a handler for Copy has been set.
func (d *CNNInstanceNormalizationDataSource) HasCopy() bool {
	return d._Copy != nil
}

// Load implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Load() {
	if d._Load != nil {
		d._Load()
	}
}

// HasLoad returns true if a handler for Load has been set.
func (d *CNNInstanceNormalizationDataSource) HasLoad() bool {
	return d._Load != nil
}

// Purge implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Purge() {
	if d._Purge != nil {
		d._Purge()
	}
}

// HasPurge returns true if a handler for Purge has been set.
func (d *CNNInstanceNormalizationDataSource) HasPurge() bool {
	return d._Purge != nil
}

// EncodeWithCoder implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) EncodeWithCoder(aCoder Coder /* not a class type */) {
	if d._EncodeWithCoder != nil {
		d._EncodeWithCoder(aCoder)
	}
}

// HasEncodeWithCoder returns true if a handler for EncodeWithCoder has been set.
func (d *CNNInstanceNormalizationDataSource) HasEncodeWithCoder() bool {
	return d._EncodeWithCoder != nil
}

// InitWithCoder implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) InitWithCoder(aDecoder Coder /* not a class type */) objectivec.IObject {
	if d._InitWithCoder != nil {
		return d._InitWithCoder(aDecoder)
	}
	var zero objectivec.IObject
	return zero
}

// HasInitWithCoder returns true if a handler for InitWithCoder has been set.
func (d *CNNInstanceNormalizationDataSource) HasInitWithCoder() bool {
	return d._InitWithCoder != nil
}

// Label implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Label() {
	if d._Label != nil {
		d._Label()
	}
}

// HasLabel returns true if a handler for Label has been set.
func (d *CNNInstanceNormalizationDataSource) HasLabel() bool {
	return d._Label != nil
}

// Beta implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Beta() {
	if d._Beta != nil {
		d._Beta()
	}
}

// HasBeta returns true if a handler for Beta has been set.
func (d *CNNInstanceNormalizationDataSource) HasBeta() bool {
	return d._Beta != nil
}

// Gamma implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) Gamma() {
	if d._Gamma != nil {
		d._Gamma()
	}
}

// HasGamma returns true if a handler for Gamma has been set.
func (d *CNNInstanceNormalizationDataSource) HasGamma() bool {
	return d._Gamma != nil
}

// UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer unsafe.Pointer, instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState {
	if d._UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch != nil {
		return d._UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch(commandBuffer, instanceNormalizationStateBatch)
	}
	var zero CNNNormalizationGammaAndBetaState
	return zero
}

// HasUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch returns true if a handler for UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch has been set.
func (d *CNNInstanceNormalizationDataSource) HasUpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch() bool {
	return d._UpdateGammaAndBetaWithCommandBufferInstanceNormalizationStateBatch != nil
}

// UpdateGammaAndBetaWithInstanceNormalizationStateBatch implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch CNNInstanceNormalizationGradientStateBatch /* not a class type */) bool {
	if d._UpdateGammaAndBetaWithInstanceNormalizationStateBatch != nil {
		return d._UpdateGammaAndBetaWithInstanceNormalizationStateBatch(instanceNormalizationStateBatch)
	}
	var zero bool
	return zero
}

// HasUpdateGammaAndBetaWithInstanceNormalizationStateBatch returns true if a handler for UpdateGammaAndBetaWithInstanceNormalizationStateBatch has been set.
func (d *CNNInstanceNormalizationDataSource) HasUpdateGammaAndBetaWithInstanceNormalizationStateBatch() bool {
	return d._UpdateGammaAndBetaWithInstanceNormalizationStateBatch != nil
}

// CopyWithZoneDevice implements the PCNNInstanceNormalizationDataSource interface.
func (d *CNNInstanceNormalizationDataSource) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	if d._CopyWithZoneDevice != nil {
		return d._CopyWithZoneDevice(zone, device)
	}
	var zero objectivec.IObject
	return zero
}

// HasCopyWithZoneDevice returns true if a handler for CopyWithZoneDevice has been set.
func (d *CNNInstanceNormalizationDataSource) HasCopyWithZoneDevice() bool {
	return d._CopyWithZoneDevice != nil
}
