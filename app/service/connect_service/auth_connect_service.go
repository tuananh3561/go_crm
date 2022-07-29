package connect_service

type BodyResponseListUser struct {
	Code    int    `json:"code"`
	Data    Users  `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Users struct {
	//Data []entity.UserAuth
}

type BodyResponseUserIds struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

type AuthConnectService interface {
	//GetListUser(email string) ([]entity.UserAuth, error)
	//GetUserByIds(userIds string) (interface{}, error)
}

type authConnectService struct {
}

func NewCrmConnectService() AuthConnectService {
	return &authConnectService{}
}

//
//func (c authConnectService) GetListUser(email string) ([]entity.UserAuth, error) {
//	var url = config.GetUrlByEnv("API_SERVICE_AUTH") + "api/v2/get-list-user"
//
//	request, _ := helper.CurlGetClient(url, map[string]string{
//		"email": email,
//	})
//
//	client := &http.Client{}
//	response, _ := client.Do(request)
//
//	defer response.Body.Close()
//	var body BodyResponseListUser
//
//	errorDecode := json.NewDecoder(response.Body).Decode(&body)
//
//	if errorDecode != nil {
//		return nil, errorDecode
//	}
//
//	if body.Status != "success" {
//		return nil, errors.New(body.Message)
//	}
//
//	return body.Data.Data, nil
//}
//
//func (c authConnectService) GetUserByIds(userIds string) (interface{}, error) {
//	var url = config.GetUrlByEnv("API_SERVICE_AUTH") + "api/v2/get-name-by-list-user"
//
//	request, _ := helper.CurlGetClient(url, map[string]string{
//		"users": userIds,
//	})
//
//	client := &http.Client{}
//	response, _ := client.Do(request)
//
//	defer response.Body.Close()
//	var body BodyResponseUserIds
//
//	errorDecode := json.NewDecoder(response.Body).Decode(&body)
//
//	if errorDecode != nil {
//		return nil, errorDecode
//	}
//
//	if body.Status != "success" {
//		return nil, errors.New(body.Message)
//	}
//
//	return body.Data, nil
//}
