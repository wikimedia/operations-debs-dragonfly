// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskRegisterRequest task register request
// swagger:model TaskRegisterRequest
type TaskRegisterRequest struct {

	// IP address which peer client carries
	// Format: ipv4
	IP strfmt.IPv4 `json:"IP,omitempty"`

	// This attribute represents the node as a seed node for the taskURL.
	//
	AsSeed bool `json:"asSeed,omitempty"`

	// CID means the client ID. It maps to the specific dfget process.
	// When user wishes to download an image/file, user would start a dfget process to do this.
	// This dfget is treated a client and carries a client ID.
	// Thus, multiple dfget processes on the same peer have different CIDs.
	//
	CID string `json:"cID,omitempty"`

	// This attribute represents where the dfget requests come from. Dfget will pass
	// this field to supernode and supernode can do some checking and filtering via
	// black/white list mechanism to guarantee security, or some other purposes like debugging.
	//
	// Min Length: 1
	CallSystem string `json:"callSystem,omitempty"`

	// tells whether it is a call from dfdaemon. dfdaemon is a long running
	// process which works for container engines. It translates the image
	// pulling request into raw requests into those dfget recognizes.
	//
	Dfdaemon bool `json:"dfdaemon,omitempty"`

	// This attribute represents the length of resource, dfdaemon or dfget catches and calculates
	// this parameter from the headers of request URL. If fileLength is vaild, the supernode need
	// not get the length of resource by accessing the rawURL.
	//
	FileLength int64 `json:"fileLength,omitempty"`

	// extra HTTP headers sent to the rawURL.
	// This field is carried with the request to supernode.
	// Supernode will extract these HTTP headers, and set them in HTTP downloading requests
	// from source server as user's wish.
	//
	Headers []string `json:"headers"`

	// host name of peer client node.
	// Min Length: 1
	HostName string `json:"hostName,omitempty"`

	// special attribute of remote source file. This field is used with taskURL to generate new taskID to
	// identify different downloading task of remote source file. For example, if user A and user B uses
	// the same taskURL and taskID to download file, A and B will share the same peer network to distribute files.
	// If user A additionally adds an identifier with taskURL, while user B still carries only taskURL, then A's
	// generated taskID is different from B, and the result is that two users use different peer networks.
	//
	Identifier string `json:"identifier,omitempty"`

	// tells whether skip secure verify when supernode download the remote source file.
	//
	Insecure bool `json:"insecure,omitempty"`

	// md5 checksum for the resource to distribute. dfget catches this parameter from dfget's CLI
	// and passes it to supernode. When supernode finishes downloading file/image from the source location,
	// it will validate the source file with this md5 value to check whether this is a valid file.
	//
	Md5 string `json:"md5,omitempty"`

	// path is used in one peer A for uploading functionality. When peer B hopes
	// to get piece C from peer A, B must provide a URL for piece C.
	// Then when creating a task in supernode, peer A must provide this URL in request.
	//
	Path string `json:"path,omitempty"`

	// when registering, dfget will setup one uploader process.
	// This one acts as a server for peer pulling tasks.
	// This port is which this server listens on.
	//
	// Maximum: 65000
	// Minimum: 15000
	Port int32 `json:"port,omitempty"`

	// The is the resource's URL which user uses dfget to download. The location of URL can be anywhere, LAN or WAN.
	// For image distribution, this is image layer's URL in image registry.
	// The resource url is provided by command line parameter.
	//
	RawURL string `json:"rawURL,omitempty"`

	// The root ca cert from client used to download the remote source file.
	//
	RootCAs []strfmt.Base64 `json:"rootCAs"`

	// The address of supernode that the client can connect to
	SuperNodeIP string `json:"superNodeIp,omitempty"`

	// Dfdaemon or dfget could specific the taskID which will represents the key of this resource
	// in supernode.
	//
	TaskID string `json:"taskId,omitempty"`

	// taskURL is generated from rawURL. rawURL may contains some queries or parameter, dfget will filter some queries via
	// --filter parameter of dfget. The usage of it is that different rawURL may generate the same taskID.
	//
	TaskURL string `json:"taskURL,omitempty"`

	// version number of dfget binary.
	Version string `json:"version,omitempty"`
}

// Validate validates this task register request
func (m *TaskRegisterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCAs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskRegisterRequest) validateIP(formats strfmt.Registry) error {

	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if err := validate.FormatOf("IP", "body", "ipv4", m.IP.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskRegisterRequest) validateCallSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.CallSystem) { // not required
		return nil
	}

	if err := validate.MinLength("callSystem", "body", string(m.CallSystem), 1); err != nil {
		return err
	}

	return nil
}

func (m *TaskRegisterRequest) validateHostName(formats strfmt.Registry) error {

	if swag.IsZero(m.HostName) { // not required
		return nil
	}

	if err := validate.MinLength("hostName", "body", string(m.HostName), 1); err != nil {
		return err
	}

	return nil
}

func (m *TaskRegisterRequest) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 15000, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65000, false); err != nil {
		return err
	}

	return nil
}

func (m *TaskRegisterRequest) validateRootCAs(formats strfmt.Registry) error {

	if swag.IsZero(m.RootCAs) { // not required
		return nil
	}

	for i := 0; i < len(m.RootCAs); i++ {

		// Format "byte" (base64 string) is already validated when unmarshalled

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskRegisterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskRegisterRequest) UnmarshalBinary(b []byte) error {
	var res TaskRegisterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
