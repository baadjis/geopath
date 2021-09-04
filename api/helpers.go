package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	geojson "github.com/paulmach/go.geojson"
)

type Path struct {
	PathName string `json:"pathname"`

	Coordinates [][]float64 `json:"coordinates"`
}

//create a geojson feature collect with proprety pathname and coordinates
func NewGeoJSON(path Path) ([]byte, error) {
	featureCollection := geojson.NewFeatureCollection()
	feature := geojson.NewLineStringFeature(path.Coordinates)
	feature.SetProperty("pathname", path.PathName)

	featureCollection.AddFeature(feature)
	return featureCollection.MarshalJSON()
}

// add a new feature to feature collection
func AppendGeojson(path Path, featureCollection geojson.FeatureCollection) ([]byte, error) {

	feature := geojson.NewLineStringFeature(path.Coordinates)
	feature.SetProperty("pathname", path.PathName)

	featureCollection.AddFeature(feature)
	return featureCollection.MarshalJSON()
}

//read the geojson file aand return bytes
func ReadGeojsonFile() ([]byte, error) {
	data, err := ioutil.ReadFile("./public/path.geojson")

	if err != nil {
		fmt.Print(err)

	}
	return data, err
}

// write to the geojson file
func WriteGeojsonFile(data []byte) error {
	err := ioutil.WriteFile("./public/path.geojson", data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// calculate duration
func GetDurationWithCoordinates(cor [][]float64) float64 {
	return (cor[len(cor)-1][3] - cor[0][3])
}

//get duration of linestring feature
func GetDuration(feature geojson.Feature) float64 {

	var cor [][]float64

	cor = feature.Geometry.LineString

	duration := GetDurationWithCoordinates(cor)
	return duration
}

// calculate distance
func HaversianDistance(a, b []float64) float64 {

	Pa := math.Pi * a[0] / 180
	Pb := math.Pi * b[0] / 180
	La := math.Pi * a[1] / 180
	Lb := math.Pi * b[1] / 180

	arc := math.Pow(math.Sin((Pb-Pa)/2), 2) + math.Cos(Pa)*math.Cos(Pb)*math.Pow(math.Sin((Lb-La)/2), 2)
	cc := 2 * math.Atan(math.Sqrt(arc)/math.Sqrt(1-arc))
	return 6371 * cc
}

func GetDistanceWithCoordinates(cor [][]float64) float64 {
	dist := 0.0
	for i := range cor {
		if i > 0 {
			dist += HaversianDistance(cor[i], cor[i-1])
		}
	}
	return dist
}

//calculate distance of linestring
func GetDistance(feature geojson.Feature) float64 {

	var cor [][]float64

	cor = feature.Geometry.LineString

	dist := GetDistanceWithCoordinates(cor)
	return dist
}

//list all pathnames
func GetPathNames() []string {
	var names []string

	data, _ := ReadGeojsonFile()
	var featurecoll geojson.FeatureCollection
	json.Unmarshal(data, &featurecoll)
	for _, feature := range featurecoll.Features {
		names = append(names, feature.Properties["pathname"].(string))

	}
	return names

}

//search a feature with specific name
func GetFeatureByName(fname string, fcollection geojson.FeatureCollection) geojson.Feature {
	for _, feature := range fcollection.Features {
		if feature.Properties["pathname"] == fname {
			return *feature
		}
	}
	panic("not found")
}

// get all features of a collection
func GetFeatureCollection() geojson.FeatureCollection {
	data, _ := ReadGeojsonFile()
	var featurecoll geojson.FeatureCollection
	json.Unmarshal(data, &featurecoll)
	return featurecoll
}
