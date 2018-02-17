/*
Copyright (c) 2016, The go-imagequant author(s)

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND ISC DISCLAIMS ALL WARRANTIES WITH REGARD TO
THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS.
IN NO EVENT SHALL ISC BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR
CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA
OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS
ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS
SOFTWARE.
*/

package imagequant

import (
	"errors"
	"image"
	"image/color"
	"unsafe"
)

/*
#include "libimagequant.h"
*/
import "C"

type Image struct {
	p        *C.struct_liq_image
	w, h     int
	released bool
}

// Callers MUST call Release() on the returned object to free memory.
func NewImage(attr *Attributes, rgba32data []byte, width, height int, gamma float64) (*Image, error) {
	pImg := C.liq_image_create_rgba(attr.p, unsafe.Pointer(&rgba32data[0]), C.int(width), C.int(height), C.double(gamma))
	if pImg == nil {
		return nil, errors.New("Failed to create image (invalid argument)")
	}

	return &Image{
		p:        pImg,
		w:        width,
		h:        height,
		released: false,
	}, nil
}

// Free memory. Callers must not use this object after Release has been called.
func (this *Image) Release() {
	C.liq_image_destroy(this.p)
	this.released = true
}

func (this *Image) Quantize(attr *Attributes) (*Result, error) {
	res := Result{
		im: this,
	}
	liqerr := C.liq_image_quantize(this.p, attr.p, &res.p)
	if liqerr != C.LIQ_OK {
		return nil, translateError(liqerr)
	}

	return &res, nil
}

func GoImageToRgba32(im image.Image) []byte {
	w := im.Bounds().Max.X
	h := im.Bounds().Max.Y
	ret := make([]byte, w*h*4)

	p := 0

	for y := 0; y < h; y += 1 {
		for x := 0; x < w; x += 1 {
			r16, g16, b16, a16 := im.At(x, y).RGBA() // Each value ranges within [0, 0xffff]

			ret[p+0] = uint8(r16 >> 8)
			ret[p+1] = uint8(g16 >> 8)
			ret[p+2] = uint8(b16 >> 8)
			ret[p+3] = uint8(a16 >> 8)
			p += 4
		}
	}

	return ret
}

func Rgb8PaletteToGoImage(w, h int, rgb8data []byte, pal color.Palette) image.Image {
	rect := image.Rectangle{
		Max: image.Point{
			X: w,
			Y: h,
		},
	}

	ret := image.NewPaletted(rect, pal)

	for y := 0; y < h; y += 1 {
		for x := 0; x < w; x += 1 {
			ret.SetColorIndex(x, y, rgb8data[y*w+x])
		}
	}

	return ret
}
