/*
Create: 2022/7/20
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"errors"
)

var (
	ErrMetaDir     = errors.New("metadata dir is not exist")
	ErrWalkMetaDir = errors.New("failed to walk metadata to load app configs")
	ErrLoadPart    = errors.New("failed to load part of apps")
	ErrAPP         = errors.New("failed to find app by name")
)
