package osmap

type OSMapData struct {
	PageProps struct {
		RouteID string `json:"routeId"`
		Route   struct {
			Status string `json:"status"`
			Data   struct {
				Metadata struct {
					Name        string `json:"name"`
					Description string `json:"description"`
					CreatedAt   string `json:"createdAt"`
				} `json:"metadata"`
				Characteristics struct {
					Distance         float64 `json:"distance"`
					ElevationAscent  float64 `json:"elevationAscent"`
					ElevationDescent float64 `json:"elevationDescent"`
					Activity         string  `json:"activity"`
					Looped           bool    `json:"looped"`
				} `json:"characteristics"`
				Features []Feature `json:"features"`
			} `json:"data"`
		} `json:"route"`
	} `json:"pageProps"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	Kind                 string                 `json:"kind"`
	DistanceFromPrevious float64                `json:"distanceFromPrevious"`
	Elevation            float64                `json:"elevation"`
	AscentFromPrevious   float64                `json:"ascentFromPrevious"`
	DescentFromPrevious  float64                `json:"descentFromPrevious"`
	PositionInSegment    int                    `json:"positionInSegment"`
	Style                map[string]interface{} `json:"style"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}
