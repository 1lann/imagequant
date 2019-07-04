# imagequant
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2F1lann%2Fimagequant.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2F1lann%2Fimagequant?ref=badge_shield)


This library was taken from https://code.ivysaur.me/go-imagequant/.

The only code changes made are
- Change import paths to `github.com/1lann/imagequant`.
- Use `[]byte` as `NewImage`'s `rgba32data` argument instead of `string`.
- `GoImageToRgba32` and `Rgb8PaletteToGoImage` are taken from `cmd/gopngquant/main.go` and replicated in `Image.go` such that it's available in the package.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2F1lann%2Fimagequant.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2F1lann%2Fimagequant?ref=badge_large)