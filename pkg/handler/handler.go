package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/medyagh/gopogh/pkg/db"
)

type HandlerDB struct {
	Database db.Datab
}

func (m *HandlerDB) ServeEnvironmentTestsAndTestCases(w http.ResponseWriter, r *http.Request) {
	data, err := m.Database.GetEnvironmentTestsAndTestCases()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON data", http.StatusInternalServerError)
		return
	}
}

// ServeTestCharts writes the individual test charts to a JSON HTTP response
func (m *HandlerDB) ServeTestCharts(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	env := queryValues.Get("env")
	if env == "" {
		http.Error(w, "missing environment name", http.StatusUnprocessableEntity)
		return
	}
	test := queryValues.Get("test")
	if test == "" {
		http.Error(w, "missing test name", http.StatusUnprocessableEntity)
		return
	}

	data, err := m.Database.GetTestCharts(env, test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON data", http.StatusInternalServerError)
		return
	}
}

// ServeEnvCharts writes the overall environment charts to a JSON HTTP response
func (m *HandlerDB) ServeEnvCharts(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	env := queryValues.Get("env")
	if env == "" {
		http.Error(w, "missing environment name", http.StatusUnprocessableEntity)
		return
	}
	testsInTopStr := queryValues.Get("tests_in_top")
	if testsInTopStr == "" {
		testsInTopStr = "10"
	}
	testsInTop, err := strconv.Atoi(testsInTopStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid number of top tests to use: %v", err), http.StatusUnprocessableEntity)
		return
	}
	data, err := m.Database.GetEnvCharts(env, testsInTop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON data", http.StatusInternalServerError)
		return
	}
}

// ServeOverview writes the overview chart for all of the environments to a JSON HTTP response
func (m *HandlerDB) ServeOverview(w http.ResponseWriter, _ *http.Request) {

	data, err := m.Database.GetOverview()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON data", http.StatusInternalServerError)
		return
	}
}
