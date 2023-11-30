package encoding

import (
	"os"
	"testing"

	"github.com/Yandex-Practicum/final-project-encoding-go/utils"
)

func TestEncodingJsonToYaml(t *testing.T) {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	jsonData := JSONData{FileInput: "jsonInput.json", FileOutput: "yamlOutput.yml"}

	err := jsonData.Encoding()

	if err != nil {
		t.Errorf("jsonData Encoding : %s", err)
	}

	expected, err := os.ReadFile("yamlInput.yml")

	if err != nil {
		t.Errorf("ReadFile yamlInput.yml : %s", err)
	}

	res, err := os.ReadFile(jsonData.FileOutput)

	if err != nil {
		t.Errorf("ReadFile json Output: %s", err)
	}

	if string(expected) != string(res) {
		t.Errorf("Files don't match")
	}

	os.Remove("jsonInput.json")
	os.Remove("yamlInput.yml")
	os.Remove("yamlOutput.yml")
}

func TestEncodingYamlToJson(t *testing.T) {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	yamlData := YAMLData{FileInput: "yamlInput.yml", FileOutput: "jsonOutput.json"}

	err := yamlData.Encoding()

	if err != nil {
		t.Errorf("yamlData Encoding : %s", err)
	}

	expected, err := os.ReadFile("jsonInput.json")

	if err != nil {
		t.Errorf("ReadFile jsonInput.json : %s", err)
	}

	res, err := os.ReadFile(yamlData.FileOutput)

	if err != nil {
		t.Errorf("ReadFile yaml Output: %s", err)
	}

	if string(expected) != string(res) {
		t.Errorf("Files don't match")
	}

	os.Remove("jsonInput.json")
	os.Remove("jsonOutput.json")
	os.Remove("yamlInput.yml")

}
