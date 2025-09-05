package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type GPXMetadata struct {
	Name   string `xml:"name"`
	Time   string `xml:"time"`
	Type   string `xml:"type"`
	Author struct {
		Name string `xml:"name"`
	} `xml:"author"`
	Link struct {
		Href string `xml:"href,attr"`
		Text string `xml:"text"`
	} `xml:"link"`
}

type GPXTrack struct {
	Name string `xml:"name"`
	Type string `xml:"type"`
}

type GPXRoute struct {
	Name string `xml:"name"`
}

type GPX struct {
	Metadata GPXMetadata `xml:"metadata"`
	Track    GPXTrack    `xml:"trk"`
	Route    GPXRoute    `xml:"rte"`
}

type RouteInfo struct {
	Filename string
	Name     string
	Type     string
	Source   string
	Link     string
	FileSize int64
}

func main() {
	files, err := filepath.Glob("*.gpx")
	if err != nil {
		panic(err)
	}

	var routes []RouteInfo

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		var gpx GPX
		err = xml.Unmarshal(data, &gpx)
		if err != nil {
			continue
		}

		route := RouteInfo{
			Filename: file,
			FileSize: info.Size(),
		}

		// Extract name (prefer track name, then metadata name, then filename)
		if gpx.Track.Name != "" {
			route.Name = gpx.Track.Name
		} else if gpx.Route.Name != "" && gpx.Route.Name != "Untitled Route" {
			route.Name = gpx.Route.Name
		} else if gpx.Metadata.Name != "" && gpx.Metadata.Name != "Untitled Route" {
			route.Name = gpx.Metadata.Name
		} else {
			route.Name = strings.TrimSuffix(file, ".gpx")
		}

		// Extract type (prefer track type, then metadata type, then default to "route")
		route.Type = gpx.Track.Type
		if route.Type == "" {
			route.Type = gpx.Metadata.Type
		}
		if route.Type == "" {
			route.Type = "route"
		}

		// Extract source and link
		if strings.Contains(string(data), "strava.com") {
			route.Source = "Strava"
			route.Link = gpx.Metadata.Link.Href
		} else if strings.Contains(string(data), "ridewithgps.com") {
			route.Source = "Ride with GPS"
			route.Link = gpx.Metadata.Link.Href
		} else if strings.Contains(string(data), "footpathapp.com") {
			route.Source = "Footpath"
			route.Link = gpx.Metadata.Link.Href
		}

		routes = append(routes, route)
	}

	// Generate markdown
	fmt.Println("# Awesome Routes")
	fmt.Println("")
	fmt.Println("This is a collection of some of my favorite routes for various activities. The list is not exhaustive and will be updated over time.")
	fmt.Println("")
	fmt.Printf("📊 **Total Routes:** %d\n", len(routes))
	fmt.Println("")
	fmt.Println("## Routes")
	fmt.Println("")
	fmt.Println("| Route Name | Type | Source | File Size | File |")
	fmt.Println("|------------|------|--------|-----------|------|")

	for _, route := range routes {
		fileSizeKB := route.FileSize / 1024
		sourceText := route.Source
		if route.Link != "" {
			sourceText = fmt.Sprintf("[%s](%s)", route.Source, route.Link)
		}
		
		fmt.Printf("| %s | %s | %s | %d KB | [📁](./%s) |\n", 
			route.Name, route.Type, sourceText, fileSizeKB, route.Filename)
	}

	fmt.Println("")
	fmt.Println("## Route Types")
	fmt.Println("")
	
	typeCount := make(map[string]int)
	for _, route := range routes {
		typeCount[route.Type]++
	}
	
	for routeType, count := range typeCount {
		fmt.Printf("- **%s**: %d routes\n", strings.Title(routeType), count)
	}
}