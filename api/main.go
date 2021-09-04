package api

import (
	"encoding/json"
	"io"

	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func getDistance(w http.ResponseWriter, r *http.Request) {
	path_name := r.URL.Query().Get("pathname")
	featurecoll := GetFeatureCollection()
	if path_name != "" {

		feature := GetFeatureByName(path_name, featurecoll)

		dist := GetDistance(feature)

		distance, err := json.Marshal(dist)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), r.Response.StatusCode)
		} else {
			w.Write(distance)
		}

	} else {
		distances := make(map[string]float64)
		for _, feature := range featurecoll.Features {
			path := feature.Properties["pathname"].(string)
			distances[path] = GetDistance(*feature)
		}
		distances_bytes, err := json.Marshal(distances)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), r.Response.StatusCode)
		} else {
			w.Write(distances_bytes)
		}

	}

}
func getDuration(w http.ResponseWriter, r *http.Request) {
	path_name := r.URL.Query().Get("pathname")

	featurecoll := GetFeatureCollection()
	if path_name != "" {
		feature := GetFeatureByName(path_name, featurecoll)
		dur := GetDuration(feature)
		duration, err := json.Marshal(dur)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), r.Response.StatusCode)
		} else {
			w.Write(duration)
		}
	} else {
		durations := make(map[string]float64)
		for _, feature := range featurecoll.Features {
			path := feature.Properties["pathname"].(string)
			durations[path] = GetDuration(*feature)
		}
		durations_bytes, err := json.Marshal(durations)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), r.Response.StatusCode)
		} else {
			w.Write(durations_bytes)
		}

	}

}

func getPath(w http.ResponseWriter, r *http.Request) {

	data, _ := ReadGeojsonFile()

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func postPath(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)

	var path Path
	err := json.Unmarshal(body, &path)

	if err != nil {
		http.Error(w, err.Error(), r.Response.StatusCode)
	} else {

		featurecoll := GetFeatureCollection()

		bytes, e := AppendGeojson(path, featurecoll)

		if e != nil {
			log.Fatal(e)
		}
		WriteGeojsonFile(bytes)

	}

}

func getPathNames(w http.ResponseWriter, r *http.Request) {
	names := GetPathNames()
	data, _ := json.Marshal(names)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/getPath", getPath).Methods("GET")
	myRouter.HandleFunc("/getPath", postPath).Methods("POST")

	myRouter.HandleFunc("/getPathNames", getPathNames)
	myRouter.HandleFunc("/getDuration", getDuration)
	myRouter.HandleFunc("/getDistance", getDistance)

	//serve statics file
	staticFileDirectory := http.Dir("./statics/")

	staticFileHandler := http.StripPrefix("/statics/", http.FileServer(staticFileDirectory))

	myRouter.PathPrefix("/statics/").Handler(staticFileHandler).Methods("GET")

	// serve html file and others assets
	myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
