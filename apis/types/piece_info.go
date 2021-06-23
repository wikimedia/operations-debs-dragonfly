// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PieceInfo Peer's detailed information in supernode.
// swagger:model PieceInfo
type PieceInfo struct {

	// the peerID that dfget task should download from
	PID string `json:"pID,omitempty"`

	// The URL path to download the specific piece from the target peer's uploader.
	//
	Path string `json:"path,omitempty"`

	// When dfget needs to download a piece from another peer. Supernode will return a PieceInfo
	// that contains a peerIP. This peerIP represents the IP of this dfget's target peer.
	//
	PeerIP string `json:"peerIP,omitempty"`

	// When dfget needs to download a piece from another peer. Supernode will return a PieceInfo
	// that contains a peerPort. This peerPort represents the port of this dfget's target peer's uploader.
	//
	PeerPort int32 `json:"peerPort,omitempty"`

	// the MD5 information of piece which is generated by supernode when doing CDN cache.
	// This value will be returned to dfget in order to validate the piece's completeness.
	//
	PieceMD5 string `json:"pieceMD5,omitempty"`

	// the range of specific piece in the task, example "0-45565".
	//
	PieceRange string `json:"pieceRange,omitempty"`

	// The size of pieces which is calculated as per the following strategy
	// 1. If file's total size is less than 200MB, then the piece size is 4MB by default.
	// 2. Otherwise, it equals to the smaller value between totalSize/100MB + 2 MB and 15MB.
	//
	PieceSize int32 `json:"pieceSize,omitempty"`
}

// Validate validates this piece info
func (m *PieceInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PieceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PieceInfo) UnmarshalBinary(b []byte) error {
	var res PieceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}