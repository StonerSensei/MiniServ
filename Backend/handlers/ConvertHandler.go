package handlers

import (
	"encoding/json"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	inputFormat := r.URL.Query().Get("input")
	outputFormat := r.URL.Query().Get("output")

	if inputFormat == "" || outputFormat == "" {
		http.Error(w, "Valid input please", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var intermediate map[string]interface{}
	switch inputFormat {
	case "json":
		err = json.Unmarshal(body, &intermediate)
	case "yaml":
		err = yaml.Unmarshal(body, &intermediate)
	case "toml":
		err = toml.Unmarshal(body, &intermediate)
	default:
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var output []byte
	switch outputFormat {
	case "json":
		output, err = json.MarshalIndent(intermediate, "", "  ")
		w.Header().Set("Content-Type", "application/json")
	case "yaml":
		output, err = yaml.Marshal(intermediate)
		w.Header().Set("Content-Type", "text/yaml")
	case "toml":
		output, err = toml.Marshal(intermediate)
		w.Header().Set("Content-Type", "text/plain")
	default:
		http.Error(w, "Invalid output format", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
