// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNGroupNormalizationDataSource is the MPSCNNGroupNormalizationDataSource protocol interface.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationDataSource
type PCNNGroupNormalizationDataSource interface {
	// Required methods
	Beta()/* debug [protocol_interface/required_method]: Beta */
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject/* debug [protocol_interface/required_method]: CopyWithZoneDevice */
	EncodeWithCoder(aCoder foundation.Coder)/* debug [protocol_interface/required_method]: EncodeWithCoder */
	Gamma()/* debug [protocol_interface/required_method]: Gamma */
	InitWithCoder(aDecoder foundation.Coder) objectivec.IObject/* debug [protocol_interface/required_method]: InitWithCoder */
	Label()/* debug [protocol_interface/required_method]: Label */
	UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer unsafe.Pointer, groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState/* debug [protocol_interface/required_method]: UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch */
	UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) bool/* debug [protocol_interface/required_method]: UpdateGammaAndBetaWithGroupNormalizationStateBatch */
	// Optional methods
	Copy()
	HasCopy() bool
	Encode()
	HasEncode() bool
	Epsilon()
	HasEpsilon() bool
	UpdateGammaAndBeta()
	HasUpdateGammaAndBeta() bool
}

// CNNGroupNormalizationDataSource is a delegate implementation builder for the PCNNGroupNormalizationDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CNNGroupNormalizationDataSource struct {
	_Copy func()
	_Encode func()
	_Epsilon func()
	_UpdateGammaAndBeta func()
	_Beta func()
	_CopyWithZoneDevice func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	_EncodeWithCoder func(aCoder foundation.Coder)
	_Gamma func()
	_InitWithCoder func(aDecoder foundation.Coder) objectivec.IObject
	_Label func()
	_UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch func(commandBuffer unsafe.Pointer, groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState
	_UpdateGammaAndBetaWithGroupNormalizationStateBatch func(groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) bool
}

// SetCopy sets the handler for the Copy delegate method.
func (d *CNNGroupNormalizationDataSource) SetCopy(f func()) {
	d._Copy = f
}

// SetEncode sets the handler for the Encode delegate method.
func (d *CNNGroupNormalizationDataSource) SetEncode(f func()) {
	d._Encode = f
}

// SetEpsilon sets the handler for the Epsilon delegate method.
func (d *CNNGroupNormalizationDataSource) SetEpsilon(f func()) {
	d._Epsilon = f
}

// SetUpdateGammaAndBeta sets the handler for the UpdateGammaAndBeta delegate method.
func (d *CNNGroupNormalizationDataSource) SetUpdateGammaAndBeta(f func()) {
	d._UpdateGammaAndBeta = f
}

// SetBeta sets the handler for the Beta delegate method.
func (d *CNNGroupNormalizationDataSource) SetBeta(f func()) {
	d._Beta = f
}

// SetCopyWithZoneDevice sets the handler for the CopyWithZoneDevice delegate method.
func (d *CNNGroupNormalizationDataSource) SetCopyWithZoneDevice(f func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject) {
	d._CopyWithZoneDevice = f
}

// SetEncodeWithCoder sets the handler for the EncodeWithCoder delegate method.
func (d *CNNGroupNormalizationDataSource) SetEncodeWithCoder(f func(aCoder foundation.Coder)) {
	d._EncodeWithCoder = f
}

// SetGamma sets the handler for the Gamma delegate method.
func (d *CNNGroupNormalizationDataSource) SetGamma(f func()) {
	d._Gamma = f
}

// SetInitWithCoder sets the handler for the InitWithCoder delegate method.
func (d *CNNGroupNormalizationDataSource) SetInitWithCoder(f func(aDecoder foundation.Coder) objectivec.IObject) {
	d._InitWithCoder = f
}

// SetLabel sets the handler for the Label delegate method.
func (d *CNNGroupNormalizationDataSource) SetLabel(f func()) {
	d._Label = f
}

// SetUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch sets the handler for the UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch delegate method.
func (d *CNNGroupNormalizationDataSource) SetUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(f func(commandBuffer unsafe.Pointer, groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState) {
	d._UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch = f
}

// SetUpdateGammaAndBetaWithGroupNormalizationStateBatch sets the handler for the UpdateGammaAndBetaWithGroupNormalizationStateBatch delegate method.
func (d *CNNGroupNormalizationDataSource) SetUpdateGammaAndBetaWithGroupNormalizationStateBatch(f func(groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) bool) {
	d._UpdateGammaAndBetaWithGroupNormalizationStateBatch = f
}

// Copy implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) Copy() {
	if d._Copy != nil {
		d._Copy()
	}
}

// HasCopy returns true if a handler for Copy has been set.
func (d *CNNGroupNormalizationDataSource) HasCopy() bool {
	return d._Copy != nil
}

// Encode implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) Encode() {
	if d._Encode != nil {
		d._Encode()
	}
}

// HasEncode returns true if a handler for Encode has been set.
func (d *CNNGroupNormalizationDataSource) HasEncode() bool {
	return d._Encode != nil
}

// Epsilon implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) Epsilon() {
	if d._Epsilon != nil {
		d._Epsilon()
	}
}

// HasEpsilon returns true if a handler for Epsilon has been set.
func (d *CNNGroupNormalizationDataSource) HasEpsilon() bool {
	return d._Epsilon != nil
}

// UpdateGammaAndBeta implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) UpdateGammaAndBeta() {
	if d._UpdateGammaAndBeta != nil {
		d._UpdateGammaAndBeta()
	}
}

// HasUpdateGammaAndBeta returns true if a handler for UpdateGammaAndBeta has been set.
func (d *CNNGroupNormalizationDataSource) HasUpdateGammaAndBeta() bool {
	return d._UpdateGammaAndBeta != nil
}

// Beta implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) Beta() {
	if d._Beta != nil {
		d._Beta()
	}
}

// HasBeta returns true if a handler for Beta has been set.
func (d *CNNGroupNormalizationDataSource) HasBeta() bool {
	return d._Beta != nil
}

// CopyWithZoneDevice implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	if d._CopyWithZoneDevice != nil {
		return d._CopyWithZoneDevice(zone, device)
	}
	var zero objectivec.IObject
	return zero
}

// HasCopyWithZoneDevice returns true if a handler for CopyWithZoneDevice has been set.
func (d *CNNGroupNormalizationDataSource) HasCopyWithZoneDevice() bool {
	return d._CopyWithZoneDevice != nil
}

// EncodeWithCoder implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) EncodeWithCoder(aCoder foundation.Coder) {
	if d._EncodeWithCoder != nil {
		d._EncodeWithCoder(aCoder)
	}
}

// HasEncodeWithCoder returns true if a handler for EncodeWithCoder has been set.
func (d *CNNGroupNormalizationDataSource) HasEncodeWithCoder() bool {
	return d._EncodeWithCoder != nil
}

// Gamma implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) Gamma() {
	if d._Gamma != nil {
		d._Gamma()
	}
}

// HasGamma returns true if a handler for Gamma has been set.
func (d *CNNGroupNormalizationDataSource) HasGamma() bool {
	return d._Gamma != nil
}

// InitWithCoder implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) InitWithCoder(aDecoder foundation.Coder) objectivec.IObject {
	if d._InitWithCoder != nil {
		return d._InitWithCoder(aDecoder)
	}
	var zero objectivec.IObject
	return zero
}

// HasInitWithCoder returns true if a handler for InitWithCoder has been set.
func (d *CNNGroupNormalizationDataSource) HasInitWithCoder() bool {
	return d._InitWithCoder != nil
}

// Label implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) Label() {
	if d._Label != nil {
		d._Label()
	}
}

// HasLabel returns true if a handler for Label has been set.
func (d *CNNGroupNormalizationDataSource) HasLabel() bool {
	return d._Label != nil
}

// UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer unsafe.Pointer, groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) CNNNormalizationGammaAndBetaState {
	if d._UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch != nil {
		return d._UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch(commandBuffer, groupNormalizationStateBatch)
	}
	var zero CNNNormalizationGammaAndBetaState
	return zero
}

// HasUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch returns true if a handler for UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch has been set.
func (d *CNNGroupNormalizationDataSource) HasUpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch() bool {
	return d._UpdateGammaAndBetaWithCommandBufferGroupNormalizationStateBatch != nil
}

// UpdateGammaAndBetaWithGroupNormalizationStateBatch implements the PCNNGroupNormalizationDataSource interface.
func (d *CNNGroupNormalizationDataSource) UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch CNNGroupNormalizationGradientStateBatch /* not a class type */) bool {
	if d._UpdateGammaAndBetaWithGroupNormalizationStateBatch != nil {
		return d._UpdateGammaAndBetaWithGroupNormalizationStateBatch(groupNormalizationStateBatch)
	}
	var zero bool
	return zero
}

// HasUpdateGammaAndBetaWithGroupNormalizationStateBatch returns true if a handler for UpdateGammaAndBetaWithGroupNormalizationStateBatch has been set.
func (d *CNNGroupNormalizationDataSource) HasUpdateGammaAndBetaWithGroupNormalizationStateBatch() bool {
	return d._UpdateGammaAndBetaWithGroupNormalizationStateBatch != nil
}
