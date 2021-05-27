package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//搜索用户 https://bytedance.feishu.cn/docs/doccnizryz7NKuUmVfkRJWeZGVc
func (u User) SearchUser(query string, pageSize int, pageToken string) (*vo.SearchUserResp, error){
	queryParams := map[string]interface{}{
	}
	if query != ""{
		queryParams["query"] = query
	}
	if pageSize > 0{
		queryParams["page_size"] = pageSize
	}
	if pageToken != ""{
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiSearchUser, queryParams, http.BuildTokenHeaderOptions(u.UserAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SearchUserResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//使用手机号或邮箱获取用户 ID  https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN
func (t Tenant) BatchGetId(emails []string, mobiles []string) (*vo.BatchGetIdResp, error) {
	queryParams := make([]http.QueryParameter, 0)
	if emails != nil && len(emails) > 0 {
		for _, email := range emails {
			queryParams = append(queryParams, http.QueryParameter {
				Key: "emails",
				Value: email,
			})
		}
	}
	if mobiles != nil && len(mobiles) > 0 {
		for _, mobile := range mobiles {
			queryParams = append(queryParams, http.QueryParameter {
				Key: "emails",
				Value: mobile,
			})
		}
	}
	respBody, err := http.GetRepetition(consts.ApiBatchGetId, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.BatchGetIdResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
