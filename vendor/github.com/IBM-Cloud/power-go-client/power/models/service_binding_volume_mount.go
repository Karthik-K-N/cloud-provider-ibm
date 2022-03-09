// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceBindingVolumeMount service binding volume mount
//
// swagger:model ServiceBindingVolumeMount
type ServiceBindingVolumeMount struct {

	// container dir
	// Required: true
	ContainerDir *string `json:"container_dir"`

	// device
	// Required: true
	Device *ServiceBindingVolumeMountDevice `json:"device"`

	// device type
	// Required: true
	// Enum: [shared]
	DeviceType *string `json:"device_type"`

	// driver
	// Required: true
	Driver *string `json:"driver"`

	// mode
	// Required: true
	// Enum: [r rw]
	Mode *string `json:"mode"`
}

// Validate validates this service binding volume mount
func (m *ServiceBindingVolumeMount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBindingVolumeMount) validateContainerDir(formats strfmt.Registry) error {

	if err := validate.Required("container_dir", "body", m.ContainerDir); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBindingVolumeMount) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

var serviceBindingVolumeMountTypeDeviceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["shared"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceBindingVolumeMountTypeDeviceTypePropEnum = append(serviceBindingVolumeMountTypeDeviceTypePropEnum, v)
	}
}

const (

	// ServiceBindingVolumeMountDeviceTypeShared captures enum value "shared"
	ServiceBindingVolumeMountDeviceTypeShared string = "shared"
)

// prop value enum
func (m *ServiceBindingVolumeMount) validateDeviceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceBindingVolumeMountTypeDeviceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceBindingVolumeMount) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeviceTypeEnum("device_type", "body", *m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBindingVolumeMount) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("driver", "body", m.Driver); err != nil {
		return err
	}

	return nil
}

var serviceBindingVolumeMountTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["r","rw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceBindingVolumeMountTypeModePropEnum = append(serviceBindingVolumeMountTypeModePropEnum, v)
	}
}

const (

	// ServiceBindingVolumeMountModeR captures enum value "r"
	ServiceBindingVolumeMountModeR string = "r"

	// ServiceBindingVolumeMountModeRw captures enum value "rw"
	ServiceBindingVolumeMountModeRw string = "rw"
)

// prop value enum
func (m *ServiceBindingVolumeMount) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceBindingVolumeMountTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceBindingVolumeMount) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service binding volume mount based on the context it is used
func (m *ServiceBindingVolumeMount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBindingVolumeMount) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBindingVolumeMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBindingVolumeMount) UnmarshalBinary(b []byte) error {
	var res ServiceBindingVolumeMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
