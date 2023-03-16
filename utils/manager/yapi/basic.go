package yapi

func QueryProjectInfo(token string, apiList string) {
	//projectBaseInfo := GetProjectInfo(token)

	//	var apiIds []string
	//	var apiList []*ProjectApiInfo
	//	if yapiAllApi {
	//		page, size := 1, 20
	//		for true {
	//			pageApiInfo := new(PageApiInfo)
	//			err = common.GetWithErrorCodeResp("https://yapi.zuoyebang.cc/api/interface/list", map[string]string{
	//				"token":      token,
	//				"project_id": strconv.Itoa(projectBaseInfo.ID),
	//				"page":       strconv.Itoa(page),
	//				"size":       strconv.Itoa(size),
	//			}, pageApiInfo)
	//			if err != nil {
	//				log.Fatalf("page api info err:%v", err)
	//			}
	//			if len(pageApiInfo.List) == 0 {
	//				break
	//			}
	//			for _, info := range pageApiInfo.List {
	//				apiIds = append(apiIds, strconv.FormatInt(info.Id, 10))
	//			}
	//			page += 1
	//		}
	//	}
	//
	//	if !yapiAllApi {
	//		apiIds = strings.Split(strings.TrimSpace(api), ",")
	//	}
	//	for _, apiId := range apiIds {
	//		apiInfo := new(ProjectApiInfo)
	//		err = common.GetWithErrorCodeResp("https://yapi.zuoyebang.cc/api/interface/get", map[string]string{
	//			"token": token,
	//			"id":    apiId,
	//		}, apiInfo)
	//		if err != nil {
	//			log.Fatalf("single api info err:%v", err)
	//		}
	//		apiList = append(apiList, apiInfo)
	//	}
	//
	//	if len(apiList) == 0 {
	//		return nil
	//	}
	//
	//	return &ProjectInfo{
	//		BaseInfo: projectBaseInfo,
	//		ApiList:  apiList,
	//	}
}
