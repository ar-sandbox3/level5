package level5

import (
	_ "embed"
)

//go:embed public.pem
var Public []byte

//go:embed private.pem
var Private []byte
