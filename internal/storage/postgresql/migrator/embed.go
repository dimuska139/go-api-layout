package migrator

import "embed"

//go:embed *.sql
var migrations embed.FS
