package web

import "embed"

//go:embed all:build
var BuildDir embed.FS
