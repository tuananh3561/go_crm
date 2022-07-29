package service

type MediaService interface {
	//Upload(pathFile string, pathUpload string, prefix string) (*DataMedia, error)
	//UploadFileWithOutSaveFileLocal(pathFile string, pathUpload string) (string, error)
	//UploadFile(context *gin.Context, FileUpload *multipart.FileHeader, pathStorage string, pathUpload string, HasNameOrigin bool, HasPrefix string, filename string) (string, error)
	//UploadFileImage(context *gin.Context, FileUpload *multipart.FileHeader, pathStorage string, pathUpload string, HasNameOrigin bool, HasPrefix string) (string, error)
	//SaveLocalStoragePlistFile(plistData interface{}, fileName string, pathStorage string) error
	//SaveLocalStorageJsonFile(jsonData interface{}, fileName string, pathStorage string) error
	//DownloadFileFromUrl(Url string, pathStorage string, fileName string) error
}

type mediaService struct {
}

func NewMediaService() MediaService {
	return &mediaService{}
}

type BodyResponse struct {
	Code    int        `json:"code"`
	Data    *DataMedia `json:"data"`
	Message string     `json:"message"`
	Status  string     `json:"status"`
}

type DataMedia struct {
	Id         int    `json:"id"`
	Link       string `json:"link"`
	LinkS3     string `json:"link_s3"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	FolderId   int    `json:"folder_id"`
	FormatFile string `json:"format_file"`
	IsPublic   int    `json:"is_public"`
	Size       int    `json:"size"`
}

const PrefixGenerate = "generate"

//func (m mediaService) UploadFile(context *gin.Context, FileUpload *multipart.FileHeader, pathStorage string, pathUpload string, HasNameOrigin bool, HasPrefix string, filename string) (string, error) {
//	prefix := ""
//	var pathStorageMkd = pathStorage
//	if HasPrefix == PrefixGenerate {
//		prefix = helper.GenerateString(1)
//		prefix = strings.ToLower(prefix)
//		pathStorageMkd += "/" + prefix
//	} else if HasPrefix != "" {
//		prefix = strings.ToLower(HasPrefix)
//		pathStorageMkd += "/" + prefix
//	}
//	err := helper.Mkdir(pathStorageMkd)
//	if err != nil {
//		return "", err
//	}
//
//	// generate filename
//	if filename == "" {
//		if HasNameOrigin == true {
//			filename = FileUpload.Filename
//		} else {
//			name := helper.GenerateString(22)
//			curTime := strconv.FormatUint(uint64(time.Now().UnixMilli()), 10)
//			filename = name + curTime + filepath.Ext(FileUpload.Filename)
//		}
//	}
//	if prefix != "" {
//		filename = prefix + "/" + filename
//	}
//	pathFile := pathStorage + "/" + filename
//
//	// save file on local storage
//	errSaveFile := context.SaveUploadedFile(FileUpload, pathFile)
//	if errSaveFile != nil {
//		return "", errSaveFile
//	}
//	// upload file media
//	_, errUploadMedia := m.Upload(pathFile, pathUpload, prefix)
//	if errUploadMedia != nil {
//		return "", errUploadMedia
//	}
//
//	// remove file on local storage
//	errRemove := helper.RemoveAllFileOrFolderLocalStorage(pathStorage)
//
//	if errRemove != nil {
//		return "", errRemove
//	}
//
//	return filename, nil
//}
//
//func (m mediaService) UploadFileImage(context *gin.Context, FileUpload *multipart.FileHeader, pathStorage string, pathUpload string, HasNameOrigin bool, HasPrefix string) (string, error) {
//	prefix := ""
//	if HasPrefix == PrefixGenerate {
//		prefix = helper.GenerateString(1)
//		prefix = strings.ToLower(prefix)
//	} else if HasPrefix != "" {
//		prefix = strings.ToLower(HasPrefix)
//	}
//	// generate filename
//	filename := ""
//	if HasNameOrigin == true {
//		filename = FileUpload.Filename
//	} else {
//		name := helper.GenerateString(22)
//		curTime := strconv.FormatUint(uint64(time.Now().UnixMilli()), 10)
//		filename = name + curTime + filepath.Ext(FileUpload.Filename)
//	}
//	fileNameHD, errUploadMediaHD := m.UploadFile(context, FileUpload, pathStorage+"/"+entity.PathHD, pathUpload+"/"+entity.PathHD, HasNameOrigin, prefix, filename)
//	if errUploadMediaHD != nil {
//		return "", errUploadMediaHD
//	}
//	_, errUploadMediaHDR := m.UploadFile(context, FileUpload, pathStorage+"/"+entity.PathHDR, pathUpload+"/"+entity.PathHDR, HasNameOrigin, prefix, filename)
//	if errUploadMediaHDR != nil {
//		return "", errUploadMediaHDR
//	}
//	return fileNameHD, nil
//}
//
//func (m mediaService) UploadFileWithOutSaveFileLocal(pathFile string, pathUpload string) (string, error) {
//	data, errUploadMedia := m.Upload(pathFile, pathUpload, "")
//
//	if errUploadMedia != nil {
//		return "", errUploadMedia
//	}
//
//	fmt.Println(data.Size)
//	return strconv.Itoa(data.Size), nil
//}
//
//func (m mediaService) Upload(pathFile string, pathUpload string, prefix string) (*DataMedia, error) {
//	var url = config.GetUrlByEnv("API_SERVICE_MEDIA") + "api/upload"
//	params := map[string]string{
//		"description":  "",
//		"folder_path":  pathUpload,
//		"is_mime_type": "1",
//	}
//	if prefix != "" {
//		params["prefix"] = prefix
//	}
//	request, errCurl := helper.CurlPostFile(url, params, "file", pathFile)
//	if errCurl != nil {
//		return nil, errCurl
//	}
//	client := &http.Client{}
//	response, err := client.Do(request)
//	if err != nil {
//		return nil, err
//	}
//	defer response.Body.Close()
//	var body BodyResponse
//	data3, _ := json.Marshal(response.Body)
//	fmt.Println(string(data3))
//	errorDecode := json.NewDecoder(response.Body).Decode(&body)
//	if errorDecode != nil {
//		return nil, errorDecode
//	}
//	if body.Status != "success" {
//		return nil, errors.New(body.Message)
//	}
//
//	return body.Data, nil
//}
//
//func (m mediaService) SaveLocalStoragePlistFile(plistData interface{}, fileName string, pathStorage string) error {
//	err := helper.Mkdir(pathStorage)
//	if err != nil {
//		return errors.New("folder creation failed")
//	}
//
//	// create file
//	pathFile := pathStorage + "/" + fileName
//	folder, err := os.Create(pathFile)
//	if err != nil {
//		return err
//	}
//	defer folder.Close()
//
//	// encode data interface to xml
//	encoder := plist.NewEncoder(os.Stdout)
//	encoder.Encode(plistData)
//
//	// convert data xml to byte
//	var plistBytes []byte
//	plistBytes, err = plist.MarshalIndent(plistData, plist.XMLFormat, "\t")
//	if err != nil {
//		return err
//	}
//
//	// write file plist
//	_, errWrite := folder.Write(plistBytes)
//	if errWrite != nil {
//		return errWrite
//	}
//
//	return nil
//}
//
//func (m mediaService) SaveLocalStorageJsonFile(jsonData interface{}, fileName string, pathStorage string) error {
//	err := helper.Mkdir(pathStorage)
//	if err != nil {
//		return errors.New("folder creation failed")
//	}
//
//	// create file
//	pathFile := pathStorage + "/" + fileName
//	folder, err := os.Create(pathFile)
//	if err != nil {
//		return err
//	}
//	defer folder.Close()
//
//	//convert json
//	dataFile, errMarshal := json.Marshal(jsonData)
//	if errMarshal != nil {
//		return errMarshal
//	}
//
//	// create file plist
//	_, errWrite := folder.Write(dataFile)
//	if errWrite != nil {
//		return errWrite
//	}
//
//	return nil
//}
//
//func (m mediaService) DownloadFileFromUrl(Url string, pathStorage string, fileName string) error {
//	//Mkdir
//	err := helper.Mkdir(pathStorage)
//	if err != nil {
//		return errors.New("folder creation failed")
//	}
//
//	//download file
//	response, err := http.Get(config.GetUrlByEnv("URL_DISPLAY_CDN") + Url)
//	if err != nil {
//		return err
//	}
//	defer response.Body.Close()
//
//	// Create the file
//	if fileName == "" {
//		fileName = path.Base(Url)
//	}
//	dir := path.Dir(fileName)
//	if dir != "" && dir != "." {
//		errMkdir := helper.Mkdir(pathStorage + "/" + dir)
//		if errMkdir != nil {
//			return errors.New("folder creation failed")
//		}
//	}
//	out, err := os.Create(pathStorage + "/" + fileName)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//
//	// Write the body to file
//	_, err = io.Copy(out, response.Body)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
