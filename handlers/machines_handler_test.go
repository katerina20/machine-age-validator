package handlers

import (
	"bytes"
	"encoding/json"
	"main/models"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestHandleMachineData(t *testing.T) {
	body, _ := os.ReadFile("../body_data.json")

	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleMachineData)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	var expectedOutput []models.Machine
	if err := readJSONFromFile("../result_data.json", &expectedOutput); err != nil {
		t.Fatal(err)
	}

	var gotOutliers []models.Machine
	if err := json.NewDecoder(recorder.Body).Decode(&gotOutliers); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if !reflect.DeepEqual(gotOutliers, expectedOutput) {
		t.Errorf("handler returned unexpected body: got %v want %v", gotOutliers, expectedOutput)
	}
}

func readJSONFromFile(filePath string, target interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}
