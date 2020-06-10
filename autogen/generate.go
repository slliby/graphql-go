package autogen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GenDirectory(inputDir string, outPutDir string, packageName string) error {
	files, err := filepath.Glob(inputDir)
	if err != nil {
		return err
	}
	var schemaString string
	for _, file := range files {
		data, err := openFile(file)
		if err != nil {
			return err
		}
		schemaString += data
	}
	schema, err := ParseSchema(schemaString)
	if err != nil {
		return err
	}
	for _, file := range files {

	}
	return nil
}

func genFile(inputFile string, outputDir string, packageName string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	_, fileName := filepath.Split(file.Name())
	ext := filepath.Ext(file.Name())
	fileName = strings.TrimSuffix(fileName, ext)
	fileName += ".go"
	outputFile := filepath.Join(outputDir + fileName)
	wf, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer wf.Close()
	_, err = wf.WriteString("package " + packageName + "\n")
	if err != nil {
		return err
	}
	_, err = wf.WriteString(outputStr)
	return err
}

func openFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
