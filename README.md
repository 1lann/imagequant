# imagequant

This library was taken from https://code.ivysaur.me/go-imagequant/.

The only code changes made are
- Change import paths to `github.com/1lann/imagequant`.
- Use `[]byte` as `NewImage`'s `rgba32data` argument instead of `string`.
- `GoImageToRgba32` and `Rgb8PaletteToGoImage` are taken from `cmd/gopngquant/main.go` and replicated in `Image.go` such that it's available in the package.
