# OS Maps to GPX Converter

A Go tool that converts OS Maps route JSON format to GPX format compatible with Garmin watches and other GPS devices.

## Installation

### Download Pre-built Binaries

Download the latest release for your platform from the [releases page](https://github.com/owenrumney/os2gpx/releases):

- **Windows**: `os2gpx_Windows_x86_64.zip` or `os2gpx_Windows_arm64.zip`
- **macOS**: `os2gpx_Darwin_x86_64.tar.gz` or `os2gpx_Darwin_arm64.tar.gz`
- **Linux**: `os2gpx_Linux_x86_64.tar.gz` or `os2gpx_Linux_arm64.tar.gz`

### Build from Source

```bash
git clone https://github.com/owenrumney/os2gpx.git
cd /os2gpx
go build -o os2gpx main.go
```

## Features

- Converts OS Maps JSON route data to GPX 1.1 format
- Preserves elevation data from waypoints
- Includes route metadata (name, description, creation time)
- Garmin watch compatible output
- Handles both waypoints and track segments

## Usage

### Getting the JSON File

This is where it gets a bit faffy - you need to open up the Inspect Window in the developer tools and look at the Network traffic in the `Network` tab to get the JSON file... export the response to file and load that.

```bash
# Using pre-built binary
./os2gpx -input=routes/dovedale.json -output=dovedale.gpx

# Or build and run from source
go run main.go -input=routes/dovedale.json -output=dovedale.gpx

# Show version information
./github.com/owenrumney/os2gpx -version
```

### Parameters

- `-input`: Path to the OS Maps JSON file
- `-output`: Path for the output GPX file

## GPX Schema

The tool generates GPX files using the standard GPX 1.1 schema:
- **Namespace**: `http://www.topografix.com/GPX/1/1`
- **Structure**: Track (`<trk>`) with track segments (`<trkseg>`) containing track points (`<trkpt>`)
- **Data**: Latitude, longitude, elevation, and metadata

## Input Format

Expects OS Maps JSON export format with:
- Route metadata (name, description, creation date)
- Features array containing waypoints and track segments
- Coordinate data in [longitude, latitude] format
- Elevation data in waypoint properties

## Output Format

Generates standard GPX 1.1 XML with:
- Route name and description in metadata
- Single track with one track segment
- Track points with lat/lon coordinates and elevation
- Compatible with Garmin watches and other GPS devices

## Example

```bash
# Convert a route from OS Maps
./github.com/owenrumney/os2gpx -input=routes/my-route.json -output=my-route.gpx

# Upload the resulting GPX file to your Garmin watch
```

## Releases

This project uses GoReleaser to build cross-platform binaries. When you push a tag (e.g., `v1.0.0`), GitHub Actions will automatically:

- Build binaries for Windows, Linux, and macOS (both amd64 and arm64)
- Create a GitHub release with all binaries
- Generate checksums for verification

To create a release:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## Roadmap

In the next version you will be able to pass it the URL of the OS Maps route and it will do the work for you