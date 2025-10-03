package gpx

import (
	"encoding/xml"
	"time"
)

type GPX struct {
	XMLName  xml.Name `xml:"gpx"`
	Version  string   `xml:"version,attr"`
	Creator  string   `xml:"creator,attr"`
	Xmlns    string   `xml:"xmlns,attr"`
	Metadata Metadata `xml:"metadata"`
	Tracks   []Track  `xml:"trk"`
}

type Metadata struct {
	Name        string    `xml:"name"`
	Description string    `xml:"desc,omitempty"`
	Time        time.Time `xml:"time"`
}

type Track struct {
	Name        string         `xml:"name"`
	Description string         `xml:"desc,omitempty"`
	Segments    []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	Points []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	Lat       float64 `xml:"lat,attr"`
	Lon       float64 `xml:"lon,attr"`
	Elevation float64 `xml:"ele,omitempty"`
	Time      string  `xml:"time,omitempty"`
}
