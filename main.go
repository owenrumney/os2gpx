package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/owenrumney/os2gpx/pkg/gpx"
	"github.com/owenrumney/os2gpx/pkg/osmap"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	inputFile := flag.String("input", "", "Input OS Map JSON file")
	outputFile := flag.String("output", "", "Output GPX file")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("github.com/owenrumney/os2gpx %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Usage: github.com/owenrumney/os2gpx -input=dovedale.json -output=dovedale.gpx")
		fmt.Println("       github.com/owenrumney/os2gpx -version")
		os.Exit(1)
	}

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var osMapData osmap.OSMapData
	if err := json.Unmarshal(data, &osMapData); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	gpxData := convertToGPX(osMapData)

	output, err := xml.MarshalIndent(gpxData, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling GPX: %v", err)
	}

	xmlHeader := `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	fullOutput := xmlHeader + string(output)

	if err := os.WriteFile(*outputFile, []byte(fullOutput), 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", *inputFile, *outputFile)
}

func convertToGPX(osMapData osmap.OSMapData) gpx.GPX {
	routeData := osMapData.PageProps.Route.Data

	createdTime, _ := time.Parse(time.RFC3339, routeData.Metadata.CreatedAt)
	if createdTime.IsZero() {
		createdTime = time.Now()
	}

	gpxData := gpx.GPX{
		Version: "1.1",
		Creator: "osToGPX Converter",
		Xmlns:   "http://www.topografix.com/GPX/1/1",
		Metadata: gpx.Metadata{
			Name:        routeData.Metadata.Name,
			Description: routeData.Metadata.Description,
			Time:        createdTime,
		},
	}

	track := gpx.Track{
		Name:        routeData.Metadata.Name,
		Description: routeData.Metadata.Description,
	}

	var trackSegment gpx.TrackSegment

	// Only process TRACK_SEGMENT features for the actual route path
	// Waypoints are summary/milestone points and would duplicate the route
	for _, feature := range routeData.Features {
		if feature.Properties.Kind == "TRACK_SEGMENT" && feature.Geometry.Type == "LineString" {
			coords, ok := feature.Geometry.Coordinates.([]interface{})
			if !ok {
				continue
			}

			for _, coord := range coords {
				coordArray, ok := coord.([]interface{})
				if !ok || len(coordArray) < 2 {
					continue
				}

				lon, ok1 := coordArray[0].(float64)
				lat, ok2 := coordArray[1].(float64)
				if !ok1 || !ok2 {
					continue
				}

				trackPoint := gpx.TrackPoint{
					Lat: lat,
					Lon: lon,
				}

				trackSegment.Points = append(trackSegment.Points, trackPoint)
			}
		}
	}

	if len(trackSegment.Points) > 0 {
		track.Segments = append(track.Segments, trackSegment)
	}

	gpxData.Tracks = append(gpxData.Tracks, track)

	return gpxData
}
