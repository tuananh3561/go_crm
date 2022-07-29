package helper

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"unicode"
)

func GenerateString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func GetJSONString(obj interface{}, ignoreFields ...string) (string, error) {
	toJson, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	if len(ignoreFields) == 0 {
		return string(toJson), nil
	}

	toMap := map[string]interface{}{}
	json.Unmarshal([]byte(string(toJson)), &toMap)

	for _, field := range ignoreFields {
		delete(toMap, field)
	}

	toJson, err = json.Marshal(toMap)
	if err != nil {
		return "", err
	}

	return string(toJson), nil
}

func ConvertStructToMap(st interface{}, fillable string) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(t.Field(i).Name)
		typ := v.FieldByName(t.Field(i).Name).Kind().String()
		structTag := t.Field(i).Tag.Get("json")
		jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
		value := v.FieldByName(t.Field(i).Name)

		// skip attributes not fillable
		if !strings.Contains(fillable, jsonName) {
			continue
		}

		// if jsonName is not empty use it for the key
		if jsonName != "" && jsonName != "-" {
			key = jsonName
		}

		if typ == "string" {
			if !(value.String() == "" || strings.Contains(structTag, "omitempty")) {
				result[key] = value.String()
			}
		} else if typ == "int" {
			result[key] = value.Int()
		} else {
			result[key] = value.Interface()
		}

	}

	return result
}

func JsonDecode(data string) ([]string, error) {
	var newData []string
	err := json.Unmarshal([]byte(data), &newData)
	return newData, err
}

func JsonDecodeI(data string) (interface{}, error) {
	var newData interface{}
	err := json.Unmarshal([]byte(data), &newData)
	return newData, err
}

func ZipFileLocalStorage(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if source == path {
			return nil
		}
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)

		if err != nil {
			return err
		}

		return nil
	})
}

func RemoveAllFileOrFolderLocalStorage(Filepath string) error {
	arrPath := strings.Split(Filepath, "/")
	folderPath := arrPath[0]
	isNotExist := false

	for key, _ := range arrPath {
		path := folderPath

		if key < len(arrPath)-1 {
			folderPath = folderPath + "/" + arrPath[key+1]
		}

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			continue
		}

		isNotExist = true
	}

	if isNotExist {
		return nil
	}

	path := Filepath
	err := os.RemoveAll(path)

	if err != nil {
		return errors.New("Remove folder fail")
	}

	return nil
}

func Unzip(src string, destination string) ([]string, error) {
	filenames := []string{}
	r, err := zip.OpenReader(src)

	if err != nil {
		return filenames, err
	}

	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(destination, f.Name)

		if !strings.HasPrefix(fPath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s is an illegal filepath", fPath)
		}

		filenames = append(filenames, fPath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fPath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())

		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()

		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}

	return filenames, nil
}

func ConvertUnicode(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}
