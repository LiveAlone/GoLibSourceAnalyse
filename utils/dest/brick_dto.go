package dto

//ApiOrgGetOfficialAccountsReq  
type ApiOrgGetOfficialAccountsReq struct{
    OrgID    int  `json:"orgID" binding:"required"` // 机构ID 
}

//ApiOrgGetOfficialAccountsOfficialAccounts  
type ApiOrgGetOfficialAccountsOfficialAccounts struct{
    Name    string  `json:"name" binding:"required"`
    IconURL    string  `json:"iconURL" binding:"required"`
    BindId    int  `json:"bindId" binding:"required"`
    WxAppId    string  `json:"wxAppId" binding:"required"`
}

//ApiOrgGetOfficialAccountsRes  
type ApiOrgGetOfficialAccountsRes struct{
    OfficialAccounts    []ApiOrgGetOfficialAccountsOfficialAccounts  `json:"officialAccounts"`
}
