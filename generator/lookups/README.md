# Lookups Generator

This tool generates string-to-color and color-to-string lookup maps and functions for the `colors` package and its subpackages (`crayola`, `pantone`, `web`, `x11`).

## Usage

From the root of the repository, run:

```bash
go run generator/lookups/main.go
```

This will parse the `*.go` files in the root and subpackages, find all variables of type `color.RGBA`, and generate:
- `lookup.go`: Contains `FromString` and `ToString` lookup functions.
- `name_to_color.go`: Contains the `nameToColor` map mapping lowercase names to `color.RGBA`.
- `color_to_name.go`: Contains the `colorToName` map mapping `color.RGBA` to string names.
- `maps_sync_test.go`: A basic sanity test to ensure the maps populate correctly.
