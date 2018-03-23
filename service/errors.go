package service

import (
	"github.com/go-errors/errors"
)

// ErrPathNotExist indicates a PathNotExist error
var ErrPathNotExist = errors.New("path does not exist")

// ErrUnexpected indicates an unexpected error
var ErrUnexpected = errors.New("unexpected error")

// ErrInvalidSiteID indicates an invalid SiteID
// Typically a SiteID <= 0
var ErrInvalidSiteID = errors.New("invalid siteID")
