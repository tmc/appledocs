// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Hypervisor Functions (125 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_hv_capability func(unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_config_create func() unsafe.Pointer
	_hv_gic_config_set_distributor_base func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_gic_config_set_msi_interrupt_range func(unsafe.Pointer, uint32, uint32) unsafe.Pointer
	_hv_gic_config_set_msi_region_base func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_gic_config_set_redistributor_base func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_gic_create func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_distributor_base_alignment func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_distributor_reg func(unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_get_distributor_size func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_icc_reg func(unsafe.Pointer, unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_get_ich_reg func(unsafe.Pointer, unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_get_icv_reg func(unsafe.Pointer, unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_get_msi_reg func(unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_get_msi_region_base_alignment func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_msi_region_size func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_redistributor_base func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_redistributor_base_alignment func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_redistributor_reg func(unsafe.Pointer, unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_gic_get_redistributor_region_size func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_redistributor_size func(unsafe.Pointer) unsafe.Pointer
	_hv_gic_get_spi_interrupt_range func([]uint32, []uint32) unsafe.Pointer
	_hv_gic_reset func() unsafe.Pointer
	_hv_gic_send_msi func(unsafe.Pointer, uint32) unsafe.Pointer
	_hv_gic_set_distributor_reg func(unsafe.Pointer, uint64) unsafe.Pointer
	_hv_gic_set_icc_reg func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	_hv_gic_set_ich_reg func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	_hv_gic_set_icv_reg func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	_hv_gic_set_msi_reg func(unsafe.Pointer, uint64) unsafe.Pointer
	_hv_gic_set_redistributor_reg func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	_hv_gic_set_spi func(uint32, bool) unsafe.Pointer
	_hv_gic_set_state func(unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_gic_state_create func() unsafe.Pointer
	_hv_gic_state_get_data func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_gic_state_get_size func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_sme_config_get_max_svl_bytes func(unsafe.Pointer) unsafe.Pointer
	_hv_tsc_clock func() uint64
	_hv_vcpu_apic_ctrl func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_get_state func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_lsc_enter_imm32 func(unsafe.Pointer, uint64, unsafe.Pointer, unsafe.Pointer, uint32, uint64, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_lsc_enter_r32 func(unsafe.Pointer, bool, uint64, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint64, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_lsc_invalidate func(unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_put_state func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_read func(unsafe.Pointer, uint32, []uint32) unsafe.Pointer
	_hv_vcpu_apic_trigger_lvt func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_apic_write func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_destroy func(unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_exit_apic_access_read func(unsafe.Pointer, []uint32) unsafe.Pointer
	_hv_vcpu_exit_info func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_exit_init_ap func(unsafe.Pointer, bool, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_exit_inject_excp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []uint32, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_exit_ioapic_eoi func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_exit_startup_ap func(unsafe.Pointer, bool, unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_vcpu_flush func(unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_get_exec_time func(unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_vcpu_get_idle_time func(unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_vcpu_get_pending_interrupt func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_get_sme_p_reg func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vcpu_get_sme_state func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_get_sme_z_reg func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vcpu_get_sme_za_reg func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vcpu_get_sme_zt0_reg func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_get_vtimer_mask func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_get_vtimer_offset func(unsafe.Pointer, []uint64) unsafe.Pointer
	_hv_vcpu_inject_extint func(unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_interrupt func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_invalidate_tlb func(unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_run func(unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_run_until func(unsafe.Pointer, uint64) unsafe.Pointer
	_hv_vcpu_set_pending_interrupt func(unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_hv_vcpu_set_sme_p_reg func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vcpu_set_sme_state func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_set_sme_z_reg func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vcpu_set_sme_za_reg func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vcpu_set_sme_zt0_reg func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_set_tsc_relative func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vcpu_set_vtimer_mask func(unsafe.Pointer, bool) unsafe.Pointer
	_hv_vcpu_set_vtimer_offset func(unsafe.Pointer, uint64) unsafe.Pointer
	_hv_vcpu_vmx_status func(unsafe.Pointer, []uint32) unsafe.Pointer
	_hv_vcpus_exit func(unsafe.Pointer, uint32) unsafe.Pointer
	_hv_vm_add_pio_notifier func(unsafe.Pointer, uintptr, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vm_allocate func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_hv_vm_atpic_assert_irq func(int) unsafe.Pointer
	_hv_vm_atpic_deassert_irq func(int) unsafe.Pointer
	_hv_vm_atpic_get_state func(unsafe.Pointer, bool) unsafe.Pointer
	_hv_vm_atpic_port_read func(int, unsafe.Pointer) unsafe.Pointer
	_hv_vm_atpic_port_write func(int, unsafe.Pointer) unsafe.Pointer
	_hv_vm_atpic_put_state func(unsafe.Pointer, bool) unsafe.Pointer
	_hv_vm_config_create func() unsafe.Pointer
	_hv_vm_config_get_default_ipa_granule func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_config_get_default_ipa_size func([]uint32) unsafe.Pointer
	_hv_vm_config_get_el2_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vm_config_get_el2_supported func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_config_get_ipa_granule func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vm_config_get_ipa_size func(unsafe.Pointer, []uint32) unsafe.Pointer
	_hv_vm_config_get_max_ipa_size func([]uint32) unsafe.Pointer
	_hv_vm_config_set_el2_enabled func(unsafe.Pointer, bool) unsafe.Pointer
	_hv_vm_config_set_ipa_granule func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vm_config_set_ipa_size func(unsafe.Pointer, uint32) unsafe.Pointer
	_hv_vm_create func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_deallocate func(unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vm_destroy func() unsafe.Pointer
	_hv_vm_get_max_vcpu_count func([]uint32) unsafe.Pointer
	_hv_vm_ioapic_assert_irq func(int) unsafe.Pointer
	_hv_vm_ioapic_deassert_irq func(int) unsafe.Pointer
	_hv_vm_ioapic_get_state func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_ioapic_pulse_irq func(int) unsafe.Pointer
	_hv_vm_ioapic_put_state func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_ioapic_read func(unsafe.Pointer, []uint32) unsafe.Pointer
	_hv_vm_ioapic_write func(unsafe.Pointer, uint32) unsafe.Pointer
	_hv_vm_lapic_msi func(uint64, uint64) unsafe.Pointer
	_hv_vm_lapic_set_intr func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vm_map func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_hv_vm_map_space func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_hv_vm_protect func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_hv_vm_protect_space func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_hv_vm_remove_pio_notifier func(unsafe.Pointer, uintptr, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_hv_vm_send_ioapic_intr func(uint64) unsafe.Pointer
	_hv_vm_set_apic_bus_freq func(uint64) unsafe.Pointer
	_hv_vm_space_create func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_space_destroy func(unsafe.Pointer) unsafe.Pointer
	_hv_vm_unmap func(unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vm_unmap_space func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_hv_vmx_vcpu_set_apic_address_space func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_hv_capability, lib, "hv_capability")
	tryRegister(&_hv_gic_config_create, lib, "hv_gic_config_create")
	tryRegister(&_hv_gic_config_set_distributor_base, lib, "hv_gic_config_set_distributor_base")
	tryRegister(&_hv_gic_config_set_msi_interrupt_range, lib, "hv_gic_config_set_msi_interrupt_range")
	tryRegister(&_hv_gic_config_set_msi_region_base, lib, "hv_gic_config_set_msi_region_base")
	tryRegister(&_hv_gic_config_set_redistributor_base, lib, "hv_gic_config_set_redistributor_base")
	tryRegister(&_hv_gic_create, lib, "hv_gic_create")
	tryRegister(&_hv_gic_get_distributor_base_alignment, lib, "hv_gic_get_distributor_base_alignment")
	tryRegister(&_hv_gic_get_distributor_reg, lib, "hv_gic_get_distributor_reg")
	tryRegister(&_hv_gic_get_distributor_size, lib, "hv_gic_get_distributor_size")
	tryRegister(&_hv_gic_get_icc_reg, lib, "hv_gic_get_icc_reg")
	tryRegister(&_hv_gic_get_ich_reg, lib, "hv_gic_get_ich_reg")
	tryRegister(&_hv_gic_get_icv_reg, lib, "hv_gic_get_icv_reg")
	tryRegister(&_hv_gic_get_msi_reg, lib, "hv_gic_get_msi_reg")
	tryRegister(&_hv_gic_get_msi_region_base_alignment, lib, "hv_gic_get_msi_region_base_alignment")
	tryRegister(&_hv_gic_get_msi_region_size, lib, "hv_gic_get_msi_region_size")
	tryRegister(&_hv_gic_get_redistributor_base, lib, "hv_gic_get_redistributor_base")
	tryRegister(&_hv_gic_get_redistributor_base_alignment, lib, "hv_gic_get_redistributor_base_alignment")
	tryRegister(&_hv_gic_get_redistributor_reg, lib, "hv_gic_get_redistributor_reg")
	tryRegister(&_hv_gic_get_redistributor_region_size, lib, "hv_gic_get_redistributor_region_size")
	tryRegister(&_hv_gic_get_redistributor_size, lib, "hv_gic_get_redistributor_size")
	tryRegister(&_hv_gic_get_spi_interrupt_range, lib, "hv_gic_get_spi_interrupt_range")
	tryRegister(&_hv_gic_reset, lib, "hv_gic_reset")
	tryRegister(&_hv_gic_send_msi, lib, "hv_gic_send_msi")
	tryRegister(&_hv_gic_set_distributor_reg, lib, "hv_gic_set_distributor_reg")
	tryRegister(&_hv_gic_set_icc_reg, lib, "hv_gic_set_icc_reg")
	tryRegister(&_hv_gic_set_ich_reg, lib, "hv_gic_set_ich_reg")
	tryRegister(&_hv_gic_set_icv_reg, lib, "hv_gic_set_icv_reg")
	tryRegister(&_hv_gic_set_msi_reg, lib, "hv_gic_set_msi_reg")
	tryRegister(&_hv_gic_set_redistributor_reg, lib, "hv_gic_set_redistributor_reg")
	tryRegister(&_hv_gic_set_spi, lib, "hv_gic_set_spi")
	tryRegister(&_hv_gic_set_state, lib, "hv_gic_set_state")
	tryRegister(&_hv_gic_state_create, lib, "hv_gic_state_create")
	tryRegister(&_hv_gic_state_get_data, lib, "hv_gic_state_get_data")
	tryRegister(&_hv_gic_state_get_size, lib, "hv_gic_state_get_size")
	tryRegister(&_hv_sme_config_get_max_svl_bytes, lib, "hv_sme_config_get_max_svl_bytes")
	tryRegister(&_hv_tsc_clock, lib, "hv_tsc_clock")
	tryRegister(&_hv_vcpu_apic_ctrl, lib, "hv_vcpu_apic_ctrl")
	tryRegister(&_hv_vcpu_apic_get_state, lib, "hv_vcpu_apic_get_state")
	tryRegister(&_hv_vcpu_apic_lsc_enter_imm32, lib, "hv_vcpu_apic_lsc_enter_imm32")
	tryRegister(&_hv_vcpu_apic_lsc_enter_r32, lib, "hv_vcpu_apic_lsc_enter_r32")
	tryRegister(&_hv_vcpu_apic_lsc_invalidate, lib, "hv_vcpu_apic_lsc_invalidate")
	tryRegister(&_hv_vcpu_apic_put_state, lib, "hv_vcpu_apic_put_state")
	tryRegister(&_hv_vcpu_apic_read, lib, "hv_vcpu_apic_read")
	tryRegister(&_hv_vcpu_apic_trigger_lvt, lib, "hv_vcpu_apic_trigger_lvt")
	tryRegister(&_hv_vcpu_apic_write, lib, "hv_vcpu_apic_write")
	tryRegister(&_hv_vcpu_create, lib, "hv_vcpu_create")
	tryRegister(&_hv_vcpu_destroy, lib, "hv_vcpu_destroy")
	tryRegister(&_hv_vcpu_exit_apic_access_read, lib, "hv_vcpu_exit_apic_access_read")
	tryRegister(&_hv_vcpu_exit_info, lib, "hv_vcpu_exit_info")
	tryRegister(&_hv_vcpu_exit_init_ap, lib, "hv_vcpu_exit_init_ap")
	tryRegister(&_hv_vcpu_exit_inject_excp, lib, "hv_vcpu_exit_inject_excp")
	tryRegister(&_hv_vcpu_exit_ioapic_eoi, lib, "hv_vcpu_exit_ioapic_eoi")
	tryRegister(&_hv_vcpu_exit_startup_ap, lib, "hv_vcpu_exit_startup_ap")
	tryRegister(&_hv_vcpu_flush, lib, "hv_vcpu_flush")
	tryRegister(&_hv_vcpu_get_exec_time, lib, "hv_vcpu_get_exec_time")
	tryRegister(&_hv_vcpu_get_idle_time, lib, "hv_vcpu_get_idle_time")
	tryRegister(&_hv_vcpu_get_pending_interrupt, lib, "hv_vcpu_get_pending_interrupt")
	tryRegister(&_hv_vcpu_get_sme_p_reg, lib, "hv_vcpu_get_sme_p_reg")
	tryRegister(&_hv_vcpu_get_sme_state, lib, "hv_vcpu_get_sme_state")
	tryRegister(&_hv_vcpu_get_sme_z_reg, lib, "hv_vcpu_get_sme_z_reg")
	tryRegister(&_hv_vcpu_get_sme_za_reg, lib, "hv_vcpu_get_sme_za_reg")
	tryRegister(&_hv_vcpu_get_sme_zt0_reg, lib, "hv_vcpu_get_sme_zt0_reg")
	tryRegister(&_hv_vcpu_get_vtimer_mask, lib, "hv_vcpu_get_vtimer_mask")
	tryRegister(&_hv_vcpu_get_vtimer_offset, lib, "hv_vcpu_get_vtimer_offset")
	tryRegister(&_hv_vcpu_inject_extint, lib, "hv_vcpu_inject_extint")
	tryRegister(&_hv_vcpu_interrupt, lib, "hv_vcpu_interrupt")
	tryRegister(&_hv_vcpu_invalidate_tlb, lib, "hv_vcpu_invalidate_tlb")
	tryRegister(&_hv_vcpu_run, lib, "hv_vcpu_run")
	tryRegister(&_hv_vcpu_run_until, lib, "hv_vcpu_run_until")
	tryRegister(&_hv_vcpu_set_pending_interrupt, lib, "hv_vcpu_set_pending_interrupt")
	tryRegister(&_hv_vcpu_set_sme_p_reg, lib, "hv_vcpu_set_sme_p_reg")
	tryRegister(&_hv_vcpu_set_sme_state, lib, "hv_vcpu_set_sme_state")
	tryRegister(&_hv_vcpu_set_sme_z_reg, lib, "hv_vcpu_set_sme_z_reg")
	tryRegister(&_hv_vcpu_set_sme_za_reg, lib, "hv_vcpu_set_sme_za_reg")
	tryRegister(&_hv_vcpu_set_sme_zt0_reg, lib, "hv_vcpu_set_sme_zt0_reg")
	tryRegister(&_hv_vcpu_set_tsc_relative, lib, "hv_vcpu_set_tsc_relative")
	tryRegister(&_hv_vcpu_set_vtimer_mask, lib, "hv_vcpu_set_vtimer_mask")
	tryRegister(&_hv_vcpu_set_vtimer_offset, lib, "hv_vcpu_set_vtimer_offset")
	tryRegister(&_hv_vcpu_vmx_status, lib, "hv_vcpu_vmx_status")
	tryRegister(&_hv_vcpus_exit, lib, "hv_vcpus_exit")
	tryRegister(&_hv_vm_add_pio_notifier, lib, "hv_vm_add_pio_notifier")
	tryRegister(&_hv_vm_allocate, lib, "hv_vm_allocate")
	tryRegister(&_hv_vm_atpic_assert_irq, lib, "hv_vm_atpic_assert_irq")
	tryRegister(&_hv_vm_atpic_deassert_irq, lib, "hv_vm_atpic_deassert_irq")
	tryRegister(&_hv_vm_atpic_get_state, lib, "hv_vm_atpic_get_state")
	tryRegister(&_hv_vm_atpic_port_read, lib, "hv_vm_atpic_port_read")
	tryRegister(&_hv_vm_atpic_port_write, lib, "hv_vm_atpic_port_write")
	tryRegister(&_hv_vm_atpic_put_state, lib, "hv_vm_atpic_put_state")
	tryRegister(&_hv_vm_config_create, lib, "hv_vm_config_create")
	tryRegister(&_hv_vm_config_get_default_ipa_granule, lib, "hv_vm_config_get_default_ipa_granule")
	tryRegister(&_hv_vm_config_get_default_ipa_size, lib, "hv_vm_config_get_default_ipa_size")
	tryRegister(&_hv_vm_config_get_el2_enabled, lib, "hv_vm_config_get_el2_enabled")
	tryRegister(&_hv_vm_config_get_el2_supported, lib, "hv_vm_config_get_el2_supported")
	tryRegister(&_hv_vm_config_get_ipa_granule, lib, "hv_vm_config_get_ipa_granule")
	tryRegister(&_hv_vm_config_get_ipa_size, lib, "hv_vm_config_get_ipa_size")
	tryRegister(&_hv_vm_config_get_max_ipa_size, lib, "hv_vm_config_get_max_ipa_size")
	tryRegister(&_hv_vm_config_set_el2_enabled, lib, "hv_vm_config_set_el2_enabled")
	tryRegister(&_hv_vm_config_set_ipa_granule, lib, "hv_vm_config_set_ipa_granule")
	tryRegister(&_hv_vm_config_set_ipa_size, lib, "hv_vm_config_set_ipa_size")
	tryRegister(&_hv_vm_create, lib, "hv_vm_create")
	tryRegister(&_hv_vm_deallocate, lib, "hv_vm_deallocate")
	tryRegister(&_hv_vm_destroy, lib, "hv_vm_destroy")
	tryRegister(&_hv_vm_get_max_vcpu_count, lib, "hv_vm_get_max_vcpu_count")
	tryRegister(&_hv_vm_ioapic_assert_irq, lib, "hv_vm_ioapic_assert_irq")
	tryRegister(&_hv_vm_ioapic_deassert_irq, lib, "hv_vm_ioapic_deassert_irq")
	tryRegister(&_hv_vm_ioapic_get_state, lib, "hv_vm_ioapic_get_state")
	tryRegister(&_hv_vm_ioapic_pulse_irq, lib, "hv_vm_ioapic_pulse_irq")
	tryRegister(&_hv_vm_ioapic_put_state, lib, "hv_vm_ioapic_put_state")
	tryRegister(&_hv_vm_ioapic_read, lib, "hv_vm_ioapic_read")
	tryRegister(&_hv_vm_ioapic_write, lib, "hv_vm_ioapic_write")
	tryRegister(&_hv_vm_lapic_msi, lib, "hv_vm_lapic_msi")
	tryRegister(&_hv_vm_lapic_set_intr, lib, "hv_vm_lapic_set_intr")
	tryRegister(&_hv_vm_map, lib, "hv_vm_map")
	tryRegister(&_hv_vm_map_space, lib, "hv_vm_map_space")
	tryRegister(&_hv_vm_protect, lib, "hv_vm_protect")
	tryRegister(&_hv_vm_protect_space, lib, "hv_vm_protect_space")
	tryRegister(&_hv_vm_remove_pio_notifier, lib, "hv_vm_remove_pio_notifier")
	tryRegister(&_hv_vm_send_ioapic_intr, lib, "hv_vm_send_ioapic_intr")
	tryRegister(&_hv_vm_set_apic_bus_freq, lib, "hv_vm_set_apic_bus_freq")
	tryRegister(&_hv_vm_space_create, lib, "hv_vm_space_create")
	tryRegister(&_hv_vm_space_destroy, lib, "hv_vm_space_destroy")
	tryRegister(&_hv_vm_unmap, lib, "hv_vm_unmap")
	tryRegister(&_hv_vm_unmap_space, lib, "hv_vm_unmap_space")
	tryRegister(&_hv_vmx_vcpu_set_apic_address_space, lib, "hv_vmx_vcpu_set_apic_address_space")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Gets the value of capabilities of the system.
//
// Added in macOS 10.15.
// Gets the value of capabilities of the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_capability(_:_:)
func hv_capability(capability unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_capability(capability, value)
}

// Creates a generic interrupt controller (GIC) configuration object.
//
// Added in macOS 15.0.
// Creates a generic interrupt controller (GIC) configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_create()
func hv_gic_config_create() unsafe.Pointer {
	return _hv_gic_config_create()
}

// Sets the generic interrupt controller (GIC) distributor region’s base address.
//
// Added in macOS 15.0.
// Sets the generic interrupt controller (GIC) distributor region’s base address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_distributor_base(_:_:)
func hv_gic_config_set_distributor_base(config unsafe.Pointer, distributor_base_address unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_config_set_distributor_base(config, distributor_base_address)
}

// Sets the range of message signaled interrupts (MSIs) the generic interrupt controller supports.
//
// Added in macOS 15.0.
// Sets the range of message signaled interrupts (MSIs) the generic interrupt controller supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_msi_interrupt_range(_:_:_:)
func hv_gic_config_set_msi_interrupt_range(config unsafe.Pointer, msi_intid_base uint32, msi_intid_count uint32) unsafe.Pointer {
	return _hv_gic_config_set_msi_interrupt_range(config, msi_intid_base, msi_intid_count)
}

// Sets the generic interrupt controllers message signaled interrupts (MSIs) region base address.
//
// Added in macOS 15.0.
// Sets the generic interrupt controllers message signaled interrupts (MSIs) region base address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_msi_region_base(_:_:)
func hv_gic_config_set_msi_region_base(config unsafe.Pointer, msi_region_base_address unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_config_set_msi_region_base(config, msi_region_base_address)
}

// Sets the generic interrupt controller (GIC) redistributor region base address.
//
// Added in macOS 15.0.
// Sets the generic interrupt controller (GIC) redistributor region base address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_redistributor_base(_:_:)
func hv_gic_config_set_redistributor_base(config unsafe.Pointer, redistributor_base_address unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_config_set_redistributor_base(config, redistributor_base_address)
}

// Creates a generic interrupt controller (GIC) v3 device for a VM configuration.
//
// Added in macOS 15.0.
// Creates a generic interrupt controller (GIC) v3 device for a VM configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_create(_:)
func hv_gic_create(gic_config unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_create(gic_config)
}

// Gets the alignment for the base address of the generic interrupt controller (GIC) distributor region, in bytes.
//
// Added in macOS 15.0.
// Gets the alignment for the base address of the generic interrupt controller (GIC) distributor region, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_distributor_base_alignment(_:)
func hv_gic_get_distributor_base_alignment(distributor_base_alignment unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_distributor_base_alignment(distributor_base_alignment)
}

// Reads a generic interrupt controller (GIC) distributor register.
//
// Added in macOS 15.0.
// Reads a generic interrupt controller (GIC) distributor register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_distributor_reg(_:_:)
func hv_gic_get_distributor_reg(reg unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_gic_get_distributor_reg(reg, value)
}

// Gets the size of the generic interrupt controller (GIC) distributor region, in bytes.
//
// Added in macOS 15.0.
// Gets the size of the generic interrupt controller (GIC) distributor region, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_distributor_size(_:)
func hv_gic_get_distributor_size(distributor_size unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_distributor_size(distributor_size)
}

// Reads a generic interrupt controller’s ICC CPU system register.
//
// Added in macOS 15.0.
// Reads a generic interrupt controller’s ICC CPU system register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_icc_reg(_:_:_:)
func hv_gic_get_icc_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_gic_get_icc_reg(vcpu, reg, value)
}

// Reads a generic interrupt controller’s (GIC) ICH virtualization control system register.
//
// Added in macOS 15.0.
// Reads a generic interrupt controller’s (GIC) ICH virtualization control system register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_ich_reg(_:_:_:)
func hv_gic_get_ich_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_gic_get_ich_reg(vcpu, reg, value)
}

// Writes a generic interrupt controller’s (GIC) ICV system register.
//
// Added in macOS 15.0.
// Writes a generic interrupt controller’s (GIC) ICV system register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_icv_reg(_:_:_:)
func hv_gic_get_icv_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_gic_get_icv_reg(vcpu, reg, value)
}

// Reads a generic interrupt controller (GIC) distributor message signaled interrupt (MSI) register.
//
// Added in macOS 15.0.
// Reads a generic interrupt controller (GIC) distributor message signaled interrupt (MSI) register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_msi_reg(_:_:)
func hv_gic_get_msi_reg(reg unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_gic_get_msi_reg(reg, value)
}

// Gets the alignment, in bytes, for the base address of the generic interrupt controller’s message signaled interrupts (MSI) region.
//
// Added in macOS 15.0.
// Gets the alignment, in bytes, for the base address of the generic interrupt controller’s message signaled interrupts (MSI) region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_msi_region_base_alignment(_:)
func hv_gic_get_msi_region_base_alignment(msi_region_base_alignment unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_msi_region_base_alignment(msi_region_base_alignment)
}

// Gets the size in bytes of the generic interrupt controller’s (GIC) message signaled interrupts (MSI) region.
//
// Added in macOS 15.0.
// Gets the size in bytes of the generic interrupt controller’s (GIC) message signaled interrupts (MSI) region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_msi_region_size(_:)
func hv_gic_get_msi_region_size(msi_region_size unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_msi_region_size(msi_region_size)
}

// Gets the redistributor base guest physical address for the given vCPU.
//
// Added in macOS 15.0.
// Gets the redistributor base guest physical address for the given vCPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_base(_:_:)
func hv_gic_get_redistributor_base(vcpu unsafe.Pointer, redistributor_base_address unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_redistributor_base(vcpu, redistributor_base_address)
}

// Gets the alignment for the base address of the generic interrupt controller (GIC) redistributor region, in bytes.
//
// Added in macOS 15.0.
// Gets the alignment for the base address of the generic interrupt controller (GIC) redistributor region, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_base_alignment(_:)
func hv_gic_get_redistributor_base_alignment(redistributor_base_alignment unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_redistributor_base_alignment(redistributor_base_alignment)
}

// Read a generic interrupt controller (GIC) redistributor register.
//
// Added in macOS 15.0.
// Read a generic interrupt controller (GIC) redistributor register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_reg(_:_:_:)
func hv_gic_get_redistributor_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value []uint64) unsafe.Pointer {
	return _hv_gic_get_redistributor_reg(vcpu, reg, value)
}

// Gets the total size in bytes of the generic interrupt controller (GIC) redistributor region.
//
// Added in macOS 15.0.
// Gets the total size in bytes of the generic interrupt controller (GIC) redistributor region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_region_size(_:)
func hv_gic_get_redistributor_region_size(redistributor_region_size unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_redistributor_region_size(redistributor_region_size)
}

// Gets the size in bytes of a single generic interrupt controller (GIC) redistributor.
//
// Added in macOS 15.0.
// Gets the size in bytes of a single generic interrupt controller (GIC) redistributor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_size(_:)
func hv_gic_get_redistributor_size(redistributor_size unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_get_redistributor_size(redistributor_size)
}

// Gets the range of shared peripheral interrupts (SPIs) the generic interrupt controller supports.
//
// Added in macOS 15.0.
// Gets the range of shared peripheral interrupts (SPIs) the generic interrupt controller supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_spi_interrupt_range(_:_:)
func hv_gic_get_spi_interrupt_range(spi_intid_base []uint32, spi_intid_count []uint32) unsafe.Pointer {
	return _hv_gic_get_spi_interrupt_range(spi_intid_base, spi_intid_count)
}

// Resets the generic interrupt controller (GIC) device.
//
// Added in macOS 15.0.
// Resets the generic interrupt controller (GIC) device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_reset()
func hv_gic_reset() unsafe.Pointer {
	return _hv_gic_reset()
}

// Sends a message signaled interrupt (MSI).
//
// Added in macOS 15.0.
// Sends a message signaled interrupt (MSI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_send_msi(_:_:)
func hv_gic_send_msi(address unsafe.Pointer, intid uint32) unsafe.Pointer {
	return _hv_gic_send_msi(address, intid)
}

// Writes the provided value to a generic interrupt controller (GIC) distributor register you specify.
//
// Added in macOS 15.0.
// Writes the provided value to a generic interrupt controller (GIC) distributor register you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_distributor_reg(_:_:)
func hv_gic_set_distributor_reg(reg unsafe.Pointer, value uint64) unsafe.Pointer {
	return _hv_gic_set_distributor_reg(reg, value)
}

// Writes to a generic interrupt controller (GIC) ICC cpu system register.
//
// Added in macOS 15.0.
// Writes to a generic interrupt controller (GIC) ICC cpu system register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_icc_reg(_:_:_:)
func hv_gic_set_icc_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value uint64) unsafe.Pointer {
	return _hv_gic_set_icc_reg(vcpu, reg, value)
}

// Writes to a generic interrupt controller (GIC) ICH virtualization control system register.
//
// Added in macOS 15.0.
// Writes to a generic interrupt controller (GIC) ICH virtualization control system register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_ich_reg(_:_:_:)
func hv_gic_set_ich_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value uint64) unsafe.Pointer {
	return _hv_gic_set_ich_reg(vcpu, reg, value)
}

// Writes to a generic interrupt controller (GIC) ICV system register.
//
// Added in macOS 15.0.
// Writes to a generic interrupt controller (GIC) ICV system register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_icv_reg(_:_:_:)
func hv_gic_set_icv_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value uint64) unsafe.Pointer {
	return _hv_gic_set_icv_reg(vcpu, reg, value)
}

// Writes to a generic interrupt controller distributor message signaled interrupt (MSI) register.
//
// Added in macOS 15.0.
// Writes to a generic interrupt controller distributor message signaled interrupt (MSI) register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_msi_reg(_:_:)
func hv_gic_set_msi_reg(reg unsafe.Pointer, value uint64) unsafe.Pointer {
	return _hv_gic_set_msi_reg(reg, value)
}

// Writes to a GIC redistributor register.
//
// Added in macOS 15.0.
// Writes to a GIC redistributor register.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_redistributor_reg(_:_:_:)
func hv_gic_set_redistributor_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value uint64) unsafe.Pointer {
	return _hv_gic_set_redistributor_reg(vcpu, reg, value)
}

// Triggers a shared peripheral interrupt (SPI).
//
// Added in macOS 15.0.
// Triggers a shared peripheral interrupt (SPI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_spi(_:_:)
func hv_gic_set_spi(intid uint32, level bool) unsafe.Pointer {
	return _hv_gic_set_spi(intid, level)
}

// Sets the state of a generic interrupt controller (GIC) device.
//
// Added in macOS 15.0.
// Sets the state of a generic interrupt controller (GIC) device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_state(_:_:)
func hv_gic_set_state(gic_state_data unsafe.Pointer, gic_state_size uintptr) unsafe.Pointer {
	return _hv_gic_set_state(gic_state_data, gic_state_size)
}

// Create a generic interrupt controller (GIC) state object.
//
// Added in macOS 15.0.
// Create a generic interrupt controller (GIC) state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_create()
func hv_gic_state_create() unsafe.Pointer {
	return _hv_gic_state_create()
}

// Gets the state data for generic interrupt controller (GIC).
//
// Added in macOS 15.0.
// Gets the state data for generic interrupt controller (GIC).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_get_data(_:_:)
func hv_gic_state_get_data(state unsafe.Pointer, gic_state_data unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_state_get_data(state, gic_state_data)
}

// Gets the size of the buffer required for generic interrupt controller (GIC) state.
//
// Added in macOS 15.0.
// Gets the size of the buffer required for generic interrupt controller (GIC) state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_get_size(_:_:)
func hv_gic_state_get_size(state unsafe.Pointer, gic_state_size unsafe.Pointer) unsafe.Pointer {
	return _hv_gic_state_get_size(state, gic_state_size)
}

// hv_sme_config_get_max_svl_bytes is a Hypervisor function.
//
// Added in macOS 15.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_sme_config_get_max_svl_bytes(_:)
func hv_sme_config_get_max_svl_bytes(value unsafe.Pointer) unsafe.Pointer {
	return _hv_sme_config_get_max_svl_bytes(value)
}

// Returns the value of an abstract clock.
//
// Added in macOS 11.0.
// Returns the value of an abstract clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_tsc_clock()
func hv_tsc_clock() uint64 {
	return _hv_tsc_clock()
}

// hv_vcpu_apic_ctrl is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_ctrl(_:_:)
func hv_vcpu_apic_ctrl(vcpu unsafe.Pointer, ctrls unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_ctrl(vcpu, ctrls)
}

// hv_vcpu_apic_get_state is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_get_state(_:_:)
func hv_vcpu_apic_get_state(vcpu unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_get_state(vcpu, state)
}

// hv_vcpu_apic_lsc_enter_imm32 is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_lsc_enter_imm32(_:_:_:_:_:_:_:)
func hv_vcpu_apic_lsc_enter_imm32(vcpu unsafe.Pointer, rip uint64, ilen unsafe.Pointer, cs unsafe.Pointer, imm32 uint32, uva uint64, count unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_lsc_enter_imm32(vcpu, rip, ilen, cs, imm32, uva, count)
}

// hv_vcpu_apic_lsc_enter_r32 is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_lsc_enter_r32(_:_:_:_:_:_:_:_:)
func hv_vcpu_apic_lsc_enter_r32(vcpu unsafe.Pointer, is_load bool, rip uint64, ilen unsafe.Pointer, cs unsafe.Pointer, reg unsafe.Pointer, uva uint64, count unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_lsc_enter_r32(vcpu, is_load, rip, ilen, cs, reg, uva, count)
}

// hv_vcpu_apic_lsc_invalidate is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_lsc_invalidate(_:)
func hv_vcpu_apic_lsc_invalidate(vcpu unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_lsc_invalidate(vcpu)
}

// hv_vcpu_apic_put_state is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_put_state(_:_:)
func hv_vcpu_apic_put_state(vcpu unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_put_state(vcpu, state)
}

// hv_vcpu_apic_read is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_read(_:_:_:)
func hv_vcpu_apic_read(vcpu unsafe.Pointer, offset uint32, data []uint32) unsafe.Pointer {
	return _hv_vcpu_apic_read(vcpu, offset, data)
}

// hv_vcpu_apic_trigger_lvt is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_trigger_lvt(_:_:)
func hv_vcpu_apic_trigger_lvt(vcpu unsafe.Pointer, flavor unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_trigger_lvt(vcpu, flavor)
}

// hv_vcpu_apic_write is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_apic_write(_:_:_:_:)
func hv_vcpu_apic_write(vcpu unsafe.Pointer, offset uint32, data uint32, no_side_effect unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_apic_write(vcpu, offset, data, no_side_effect)
}

// Creates a vCPU instance for the current thread.
//
// Added in macOS 10.10.
// Creates a vCPU instance for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_create(_:_:)
func hv_vcpu_create(vcpu unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_create(vcpu, flags)
}

// Destroys the vCPU instance associated with the current thread.
//
// Added in macOS 11.0.
// Destroys the vCPU instance associated with the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_destroy(_:)
func hv_vcpu_destroy(vcpu unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_destroy(vcpu)
}

// hv_vcpu_exit_apic_access_read is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_apic_access_read(_:_:)
func hv_vcpu_exit_apic_access_read(vcpu unsafe.Pointer, value []uint32) unsafe.Pointer {
	return _hv_vcpu_exit_apic_access_read(vcpu, value)
}

// hv_vcpu_exit_info is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_info(_:_:)
func hv_vcpu_exit_info(vcpu unsafe.Pointer, code unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_exit_info(vcpu, code)
}

// hv_vcpu_exit_init_ap is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_init_ap(_:_:_:)
func hv_vcpu_exit_init_ap(vcpu unsafe.Pointer, is_actv bool, count unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_exit_init_ap(vcpu, is_actv, count)
}

// hv_vcpu_exit_inject_excp is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_inject_excp(_:_:_:_:_:)
func hv_vcpu_exit_inject_excp(vcpu unsafe.Pointer, vec unsafe.Pointer, valid unsafe.Pointer, code []uint32, restart unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_exit_inject_excp(vcpu, vec, valid, code, restart)
}

// hv_vcpu_exit_ioapic_eoi is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_ioapic_eoi(_:_:)
func hv_vcpu_exit_ioapic_eoi(vcpu unsafe.Pointer, vec unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_exit_ioapic_eoi(vcpu, vec)
}

// hv_vcpu_exit_startup_ap is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_startup_ap(_:_:_:_:)
func hv_vcpu_exit_startup_ap(vcpu unsafe.Pointer, is_actv bool, count unsafe.Pointer, ap_rip []uint64) unsafe.Pointer {
	return _hv_vcpu_exit_startup_ap(vcpu, is_actv, count, ap_rip)
}

// Flushes the cached state of a vCPU.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.10.
// Flushes the cached state of a vCPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_flush(_:)
func hv_vcpu_flush(vcpu unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_flush(vcpu)
}

// Returns, by reference, the cumulative execution time of a vCPU, in nanoseconds.
//
// Added in macOS 11.0.
// Returns, by reference, the cumulative execution time of a vCPU, in nanoseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_exec_time(_:_:)
func hv_vcpu_get_exec_time(vcpu unsafe.Pointer, time []uint64) unsafe.Pointer {
	return _hv_vcpu_get_exec_time(vcpu, time)
}

// hv_vcpu_get_idle_time is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_idle_time(_:_:)
func hv_vcpu_get_idle_time(vcpu unsafe.Pointer, time []uint64) unsafe.Pointer {
	return _hv_vcpu_get_idle_time(vcpu, time)
}

// Gets pending interrupts for a vCPU.
//
// Added in macOS 11.0.
// Gets pending interrupts for a vCPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_pending_interrupt(_:_:_:)
func hv_vcpu_get_pending_interrupt(vcpu unsafe.Pointer, type_ unsafe.Pointer, pending unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_get_pending_interrupt(vcpu, type_, pending)
}

// Returns the value of a vCPU P predicate register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Returns the value of a vCPU P predicate register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_p_reg(_:_:_:_:)
func hv_vcpu_get_sme_p_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _hv_vcpu_get_sme_p_reg(vcpu, reg, value, length)
}

// Gets the current Scalable Matrix Extension (SME) state.
//
// Added in macOS 15.2.
// Gets the current Scalable Matrix Extension (SME) state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_state(_:_:)
func hv_vcpu_get_sme_state(vcpu unsafe.Pointer, sme_state unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_get_sme_state(vcpu, sme_state)
}

// Returns the value of a vCPU Z vector register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Returns the value of a vCPU Z vector register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_z_reg(_:_:_:_:)
func hv_vcpu_get_sme_z_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _hv_vcpu_get_sme_z_reg(vcpu, reg, value, length)
}

// Returns the value of the vCPU ZA matrix register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Returns the value of the vCPU ZA matrix register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_za_reg(_:_:_:)
func hv_vcpu_get_sme_za_reg(vcpu unsafe.Pointer, value unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _hv_vcpu_get_sme_za_reg(vcpu, value, length)
}

// Returns the current value of the vCPU ZT0 register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Returns the current value of the vCPU ZT0 register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_zt0_reg(_:_:)
func hv_vcpu_get_sme_zt0_reg(vcpu unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_get_sme_zt0_reg(vcpu, value)
}

// Gets the virtual timer mask.
//
// Added in macOS 11.0.
// Gets the virtual timer mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_vtimer_mask(_:_:)
func hv_vcpu_get_vtimer_mask(vcpu unsafe.Pointer, vtimer_is_masked unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_get_vtimer_mask(vcpu, vtimer_is_masked)
}

// Returns the vTimer offset for the vCPU ID you specify.
//
// Added in macOS 11.0.
// Returns the vTimer offset for the vCPU ID you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_vtimer_offset(_:_:)
func hv_vcpu_get_vtimer_offset(vcpu unsafe.Pointer, vtimer_offset []uint64) unsafe.Pointer {
	return _hv_vcpu_get_vtimer_offset(vcpu, vtimer_offset)
}

// hv_vcpu_inject_extint is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_inject_extint(_:)
func hv_vcpu_inject_extint(vcpu unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_inject_extint(vcpu)
}

// Forces the vCPU instances you provide to immediately exit the VM.
//
// Added in macOS 10.10.
// Forces the vCPU instances you provide to immediately exit the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_interrupt(_:_:)
func hv_vcpu_interrupt(vcpus unsafe.Pointer, vcpu_count unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_interrupt(vcpus, vcpu_count)
}

// Invalidates the translation look-aside buffer (TLB) of a vCPU.
//
// Added in macOS 10.10.
// Invalidates the translation look-aside buffer (TLB) of a vCPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_invalidate_tlb(_:)
func hv_vcpu_invalidate_tlb(vcpu unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_invalidate_tlb(vcpu)
}

// Starts the execution of a vCPU.
//
// Added in macOS 11.0.
// Starts the execution of a vCPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_run(_:)
func hv_vcpu_run(vcpu unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_run(vcpu)
}

// Executes a vCPU until it reaches the deadline defined in absolute time units you provide.
//
// Added in macOS 10.15.
// Executes a vCPU until it reaches the deadline defined in absolute time units you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_run_until(_:_:)
func hv_vcpu_run_until(vcpu unsafe.Pointer, deadline uint64) unsafe.Pointer {
	return _hv_vcpu_run_until(vcpu, deadline)
}

// Sets pending interrupts for a vCPU.
//
// Added in macOS 11.0.
// Sets pending interrupts for a vCPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_pending_interrupt(_:_:_:)
func hv_vcpu_set_pending_interrupt(vcpu unsafe.Pointer, type_ unsafe.Pointer, pending bool) unsafe.Pointer {
	return _hv_vcpu_set_pending_interrupt(vcpu, type_, pending)
}

// Sets the value of a vCPU P predicate register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Sets the value of a vCPU P predicate register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_p_reg(_:_:_:_:)
func hv_vcpu_set_sme_p_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _hv_vcpu_set_sme_p_reg(vcpu, reg, value, length)
}

// Sets the SME state consisting of the streaming Scalable Vector Extension (SVE) mode and ZA storage enable.
//
// Added in macOS 15.2.
// Sets the SME state consisting of the streaming Scalable Vector Extension (SVE) mode and ZA storage enable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_state(_:_:)
func hv_vcpu_set_sme_state(vcpu unsafe.Pointer, sme_state unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_set_sme_state(vcpu, sme_state)
}

// Sets the value of a vCPU Z vector register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Sets the value of a vCPU Z vector register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_z_reg(_:_:_:_:)
func hv_vcpu_set_sme_z_reg(vcpu unsafe.Pointer, reg unsafe.Pointer, value unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _hv_vcpu_set_sme_z_reg(vcpu, reg, value, length)
}

// Sets the value of the vCPU ZA matrix register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Sets the value of the vCPU ZA matrix register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_za_reg(_:_:_:)
func hv_vcpu_set_sme_za_reg(vcpu unsafe.Pointer, value unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _hv_vcpu_set_sme_za_reg(vcpu, value, length)
}

// Sets the value of the vCPU ZT0 register in streaming Scalable Vector Extension (SVE) mode.
//
// Added in macOS 15.2.
// Sets the value of the vCPU ZT0 register in streaming Scalable Vector Extension (SVE) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_zt0_reg(_:_:)
func hv_vcpu_set_sme_zt0_reg(vcpu unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_set_sme_zt0_reg(vcpu, value)
}

// Sets the offset of the guest timestamp-counter (TSC) relative to the Hypervisor’s TSC clock.
//
// Added in macOS 11.0.
// Sets the offset of the guest timestamp-counter (TSC) relative to the Hypervisor’s TSC clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_tsc_relative(_:_:)
func hv_vcpu_set_tsc_relative(vcpu unsafe.Pointer, offset unsafe.Pointer) unsafe.Pointer {
	return _hv_vcpu_set_tsc_relative(vcpu, offset)
}

// Sets or clears the virtual timer mask.
//
// Added in macOS 11.0.
// Sets or clears the virtual timer mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_vtimer_mask(_:_:)
func hv_vcpu_set_vtimer_mask(vcpu unsafe.Pointer, vtimer_is_masked bool) unsafe.Pointer {
	return _hv_vcpu_set_vtimer_mask(vcpu, vtimer_is_masked)
}

// Sets the vTimer offset to a value that you provide.
//
// Added in macOS 11.0.
// Sets the vTimer offset to a value that you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_vtimer_offset(_:_:)
func hv_vcpu_set_vtimer_offset(vcpu unsafe.Pointer, vtimer_offset uint64) unsafe.Pointer {
	return _hv_vcpu_set_vtimer_offset(vcpu, vtimer_offset)
}

// hv_vcpu_vmx_status is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_vmx_status(_:_:)
func hv_vcpu_vmx_status(vcpu unsafe.Pointer, status []uint32) unsafe.Pointer {
	return _hv_vcpu_vmx_status(vcpu, status)
}

// Forces an immediate exit of a set of vCPUs of the VM.
//
// Added in macOS 11.0.
// Forces an immediate exit of a set of vCPUs of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpus_exit(_:_:)
func hv_vcpus_exit(vcpus unsafe.Pointer, vcpu_count uint32) unsafe.Pointer {
	return _hv_vcpus_exit(vcpus, vcpu_count)
}

// Generate a notification when the Hypervisor issues a matching guest port I/O.
//
// Added in macOS 11.0.
// Generate a notification when the Hypervisor issues a matching guest port I/O.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_add_pio_notifier(_:_:_:_:_:)
func hv_vm_add_pio_notifier(addr unsafe.Pointer, size uintptr, value uint32, mach_port unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_add_pio_notifier(addr, size, value, mach_port, flags)
}

// hv_vm_allocate is a Hypervisor function.
//
// Added in macOS 12.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_allocate(_:_:_:)
func hv_vm_allocate(uvap unsafe.Pointer, size uintptr, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_allocate(uvap, size, flags)
}

// hv_vm_atpic_assert_irq is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_atpic_assert_irq(_:)
func hv_vm_atpic_assert_irq(irq int) unsafe.Pointer {
	return _hv_vm_atpic_assert_irq(irq)
}

// hv_vm_atpic_deassert_irq is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_atpic_deassert_irq(_:)
func hv_vm_atpic_deassert_irq(irq int) unsafe.Pointer {
	return _hv_vm_atpic_deassert_irq(irq)
}

// hv_vm_atpic_get_state is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_atpic_get_state(_:_:)
func hv_vm_atpic_get_state(state unsafe.Pointer, is_primary bool) unsafe.Pointer {
	return _hv_vm_atpic_get_state(state, is_primary)
}

// hv_vm_atpic_port_read is a Hypervisor function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_atpic_port_read(_:_:)
func hv_vm_atpic_port_read(port int, valuep unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_atpic_port_read(port, valuep)
}

// hv_vm_atpic_port_write is a Hypervisor function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_atpic_port_write(_:_:)
func hv_vm_atpic_port_write(port int, value unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_atpic_port_write(port, value)
}

// hv_vm_atpic_put_state is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_atpic_put_state(_:_:)
func hv_vm_atpic_put_state(state unsafe.Pointer, is_primary bool) unsafe.Pointer {
	return _hv_vm_atpic_put_state(state, is_primary)
}

// Creates a virtual machine configuration object.
//
// Added in macOS 13.0.
// Creates a virtual machine configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_create()
func hv_vm_config_create() unsafe.Pointer {
	return _hv_vm_config_create()
}

// hv_vm_config_get_default_ipa_granule is a Hypervisor function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_default_ipa_granule(_:)
func hv_vm_config_get_default_ipa_granule(granule unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_config_get_default_ipa_granule(granule)
}

// hv_vm_config_get_default_ipa_size is a Hypervisor function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_default_ipa_size(_:)
func hv_vm_config_get_default_ipa_size(ipa_bit_length []uint32) unsafe.Pointer {
	return _hv_vm_config_get_default_ipa_size(ipa_bit_length)
}

// Return a status value that indicates whether the VM configuration enables support for Exception Level 2 (EL2).
//
// Added in macOS 15.0.
// Return a status value that indicates whether the VM configuration enables support for Exception Level 2 (EL2).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_el2_enabled(_:_:)
func hv_vm_config_get_el2_enabled(config unsafe.Pointer, el2_enabled unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_config_get_el2_enabled(config, el2_enabled)
}

// Returns a status value that indicates whether the current platform supports Exception Level 2 (EL2).
//
// Added in macOS 15.0.
// Returns a status value that indicates whether the current platform supports Exception Level 2 (EL2).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_el2_supported(_:)
func hv_vm_config_get_el2_supported(el2_supported unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_config_get_el2_supported(el2_supported)
}

// hv_vm_config_get_ipa_granule is a Hypervisor function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_ipa_granule(_:_:)
func hv_vm_config_get_ipa_granule(config unsafe.Pointer, granule unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_config_get_ipa_granule(config, granule)
}

// hv_vm_config_get_ipa_size is a Hypervisor function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_ipa_size(_:_:)
func hv_vm_config_get_ipa_size(config unsafe.Pointer, ipa_bit_length []uint32) unsafe.Pointer {
	return _hv_vm_config_get_ipa_size(config, ipa_bit_length)
}

// hv_vm_config_get_max_ipa_size is a Hypervisor function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_max_ipa_size(_:)
func hv_vm_config_get_max_ipa_size(ipa_bit_length []uint32) unsafe.Pointer {
	return _hv_vm_config_get_max_ipa_size(ipa_bit_length)
}

// Sets whether the specified VM configuration enables support for Exception Level 2 (EL2).
//
// Added in macOS 15.0.
// Sets whether the specified VM configuration enables support for Exception Level 2 (EL2).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_set_el2_enabled(_:_:)
func hv_vm_config_set_el2_enabled(config unsafe.Pointer, el2_enabled bool) unsafe.Pointer {
	return _hv_vm_config_set_el2_enabled(config, el2_enabled)
}

// hv_vm_config_set_ipa_granule is a Hypervisor function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_set_ipa_granule(_:_:)
func hv_vm_config_set_ipa_granule(config unsafe.Pointer, granule unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_config_set_ipa_granule(config, granule)
}

// hv_vm_config_set_ipa_size is a Hypervisor function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_set_ipa_size(_:_:)
func hv_vm_config_set_ipa_size(config unsafe.Pointer, ipa_bit_length uint32) unsafe.Pointer {
	return _hv_vm_config_set_ipa_size(config, ipa_bit_length)
}

// Creates a VM instance for the current process.
//
// Added in macOS 11.0.
// Creates a VM instance for the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_create(_:)
func hv_vm_create(config unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_create(config)
}

// hv_vm_deallocate is a Hypervisor function.
//
// Added in macOS 12.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_deallocate(_:_:)
func hv_vm_deallocate(uva unsafe.Pointer, size uintptr) unsafe.Pointer {
	return _hv_vm_deallocate(uva, size)
}

// Destroys the VM instance associated with the current process.
//
// Added in macOS 11.0.
// Destroys the VM instance associated with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_destroy()
func hv_vm_destroy() unsafe.Pointer {
	return _hv_vm_destroy()
}

// Returns the maximum number of vCPUs that the hypervisor supports.
//
// Added in macOS 11.0.
// Returns the maximum number of vCPUs that the hypervisor supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_get_max_vcpu_count(_:)
func hv_vm_get_max_vcpu_count(max_vcpu_count []uint32) unsafe.Pointer {
	return _hv_vm_get_max_vcpu_count(max_vcpu_count)
}

// hv_vm_ioapic_assert_irq is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_assert_irq(_:)
func hv_vm_ioapic_assert_irq(intin int) unsafe.Pointer {
	return _hv_vm_ioapic_assert_irq(intin)
}

// hv_vm_ioapic_deassert_irq is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_deassert_irq(_:)
func hv_vm_ioapic_deassert_irq(intin int) unsafe.Pointer {
	return _hv_vm_ioapic_deassert_irq(intin)
}

// hv_vm_ioapic_get_state is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_get_state(_:)
func hv_vm_ioapic_get_state(state unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_ioapic_get_state(state)
}

// hv_vm_ioapic_pulse_irq is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_pulse_irq(_:)
func hv_vm_ioapic_pulse_irq(intin int) unsafe.Pointer {
	return _hv_vm_ioapic_pulse_irq(intin)
}

// hv_vm_ioapic_put_state is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_put_state(_:)
func hv_vm_ioapic_put_state(state unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_ioapic_put_state(state)
}

// hv_vm_ioapic_read is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_read(_:_:)
func hv_vm_ioapic_read(gpa unsafe.Pointer, datap []uint32) unsafe.Pointer {
	return _hv_vm_ioapic_read(gpa, datap)
}

// hv_vm_ioapic_write is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_ioapic_write(_:_:)
func hv_vm_ioapic_write(gpa unsafe.Pointer, data uint32) unsafe.Pointer {
	return _hv_vm_ioapic_write(gpa, data)
}

// hv_vm_lapic_msi is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_lapic_msi(_:_:)
func hv_vm_lapic_msi(addr uint64, data uint64) unsafe.Pointer {
	return _hv_vm_lapic_msi(addr, data)
}

// hv_vm_lapic_set_intr is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_lapic_set_intr(_:_:_:)
func hv_vm_lapic_set_intr(vcpu unsafe.Pointer, vector unsafe.Pointer, trig unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_lapic_set_intr(vcpu, vector, trig)
}

// Maps a region in the virtual address space of the current process into the guest physical address space of the VM.
//
// Added in macOS 11.0.
// Maps a region in the virtual address space of the current process into the guest physical address space of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_map(_:_:_:_:)
func hv_vm_map(addr unsafe.Pointer, ipa unsafe.Pointer, size uintptr, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_map(addr, ipa, size, flags)
}

// Maps a region in the virtual address space of the current task into a guest physical address space of the VM.
//
// Added in macOS 10.15.
// Maps a region in the virtual address space of the current task into a guest physical address space of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_map_space(_:_:_:_:_:)
func hv_vm_map_space(asid unsafe.Pointer, uva unsafe.Pointer, gpa unsafe.Pointer, size uintptr, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_map_space(asid, uva, gpa, size, flags)
}

// Modifies the permissions of a region in the guest physical address space of the VM.
//
// Added in macOS 11.0.
// Modifies the permissions of a region in the guest physical address space of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_protect(_:_:_:)
func hv_vm_protect(ipa unsafe.Pointer, size uintptr, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_protect(ipa, size, flags)
}

// Modifies the permissions of a region in a guest physical address space of the VM.
//
// Added in macOS 10.15.
// Modifies the permissions of a region in a guest physical address space of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_protect_space(_:_:_:_:)
func hv_vm_protect_space(asid unsafe.Pointer, gpa unsafe.Pointer, size uintptr, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_protect_space(asid, gpa, size, flags)
}

// Removes an existing I/O notifier that matches the specifications you provide.
//
// Added in macOS 11.0.
// Removes an existing I/O notifier that matches the specifications you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_remove_pio_notifier(_:_:_:_:_:)
func hv_vm_remove_pio_notifier(addr unsafe.Pointer, size uintptr, value uint32, mach_port unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_remove_pio_notifier(addr, size, value, mach_port, flags)
}

// hv_vm_send_ioapic_intr is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_send_ioapic_intr(_:)
func hv_vm_send_ioapic_intr(data uint64) unsafe.Pointer {
	return _hv_vm_send_ioapic_intr(data)
}

// hv_vm_set_apic_bus_freq is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_set_apic_bus_freq(_:)
func hv_vm_set_apic_bus_freq(freq uint64) unsafe.Pointer {
	return _hv_vm_set_apic_bus_freq(freq)
}

// Creates an additional guest address space for the current task.
//
// Added in macOS 10.15.
// Creates an additional guest address space for the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_space_create(_:)
func hv_vm_space_create(asid unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_space_create(asid)
}

// Destroys the address space instance associated with the current task.
//
// Added in macOS 10.15.
// Destroys the address space instance associated with the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_space_destroy(_:)
func hv_vm_space_destroy(asid unsafe.Pointer) unsafe.Pointer {
	return _hv_vm_space_destroy(asid)
}

// Unmaps a region in the guest physical address space of the VM.
//
// Added in macOS 11.0.
// Unmaps a region in the guest physical address space of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_unmap(_:_:)
func hv_vm_unmap(ipa unsafe.Pointer, size uintptr) unsafe.Pointer {
	return _hv_vm_unmap(ipa, size)
}

// Umaps a region in a guest physical address space of the VM.
//
// Added in macOS 10.15.
// Umaps a region in a guest physical address space of the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_unmap_space(_:_:_:)
func hv_vm_unmap_space(asid unsafe.Pointer, gpa unsafe.Pointer, size uintptr) unsafe.Pointer {
	return _hv_vm_unmap_space(asid, gpa, size)
}

// hv_vmx_vcpu_set_apic_address_space is a Hypervisor function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vmx_vcpu_set_apic_address_space(_:_:_:)
func hv_vmx_vcpu_set_apic_address_space(vcpu unsafe.Pointer, asid unsafe.Pointer, gpa unsafe.Pointer) unsafe.Pointer {
	return _hv_vmx_vcpu_set_apic_address_space(vcpu, asid, gpa)
}



