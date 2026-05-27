package service

import "errors"

// Validation errors returned by the service layer. They are sentinels so
// callers can branch on them with errors.Is instead of matching strings.
var (
	ErrEdgeRequired            = errors.New("edge is required")
	ErrNameRequired            = errors.New("name is required")
	ErrTier1NameRequired       = errors.New("db-shared name on tier 1 is required")
	ErrEdgeClusterNameRequired = errors.New("nsxt cluster name is empty")
)
