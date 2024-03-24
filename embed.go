package main

// Anonymous import for the new `embed` package.
import (
	"embed"
)

// Files can be made available through the embed package by adding a
// Go directive to the file. The directive must be a comment that starts
// with `//go:embed` followed by a space and the path to the file.
//
// You can also include a whole directory by using the embed.FS type.

//go:embed users_source.txt
var userSource string

//go:embed json_data/*
var modulesFs embed.FS
