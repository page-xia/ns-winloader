package winloader

import (
	"github.com/page-xia/go-winloader/internal/pe"
)

// NativeArch is a constant that will be equal to the PE machine type
// enumeration value that corresponds to the arch the binary is running as.
const NativeArch = pe.ImageFileMachinei386
