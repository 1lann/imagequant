//+build !windows

package imagequant

/*
#cgo CFLAGS: -O3 -fno-math-errno -fopenmp -funroll-loops -fomit-frame-pointer -Wall -Wno-attributes -std=c99 -DNDEBUG -fexcess-precision=fast
#cgo LDFLAGS: -lm -fopenmp
*/
import "C"
