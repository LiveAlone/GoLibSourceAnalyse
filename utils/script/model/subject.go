/**
 * @brief
 * @file subject
 * @author zhangpeng
 * @version 1.0
 * @updateBy niuenlai@zuoyebang.com
 * @updateDate 2021/11/30 5:14 下午
 * @updateBrief 新增艺术类学科枚举值(59~65)
 */

package model

// 学科
const (
	SubjectYuWen           = 1
	SubjectShuXue          = 2
	SubjectKeXue           = 16
	SubjectYingYu          = 3
	SubjectShengWu         = 6
	SubjectDiLi            = 9
	SubjectLiShi           = 8
	SubjectWuLi            = 4
	SubjectHuaXue          = 5
	SubjectZhengZhi        = 7
	SubjectXingQu          = 10
	SubjectSiXiangPinDe    = 11
	SubjectJiangZuo        = 12
	SubjectLiZong          = 13
	SubjectWenZong         = 14
	SubjectAoShu           = 15
	SubjectDaoDeYuFaZhi    = 48
	SubjectYinYue          = 49
	SubjectMeiShu          = 50
	SubjectXinXiJiShu      = 51
	SubjectSiXiangZhengZhi = 52
	SubjectTiNengYunDong   = 59
	SubjectYiShuSuYang     = 60
	SubjectSiWeiLuoJi      = 61
	SubjectKeJiChuangXin   = 62
	SubjectYuYanWenXue     = 63
	SubjectChuanTongWenHua = 64
	SubjectSheHuiShiJian   = 65
	SubjectTiYuYuJianKang  = 72
	SubjectTongYongJiShu   = 73
	SubjectShuFa           = 74
	SubjectLaoDong         = 75
	JAPANESE               = 76
	SPANISH                = 77
)

var SubjectIDs = []int{
	SubjectYuWen,
	SubjectShuXue,
	SubjectKeXue,
	SubjectYingYu,
	SubjectShengWu,
	SubjectDiLi,
	SubjectLiShi,
	SubjectWuLi,
	SubjectHuaXue,
	SubjectZhengZhi,
	SubjectXingQu,
	SubjectSiXiangPinDe,
	SubjectJiangZuo,
	SubjectLiZong,
	SubjectWenZong,
	SubjectAoShu,
	SubjectDaoDeYuFaZhi,
	SubjectYinYue,
	SubjectMeiShu,
	SubjectXinXiJiShu,
	SubjectSiXiangZhengZhi,
	SubjectTiNengYunDong,
	SubjectYiShuSuYang,
	SubjectSiWeiLuoJi,
	SubjectKeJiChuangXin,
	SubjectYuYanWenXue,
	SubjectChuanTongWenHua,
	SubjectSheHuiShiJian,
	SubjectTiYuYuJianKang,
	SubjectTongYongJiShu,
	SubjectShuFa,
	SubjectLaoDong,
	JAPANESE,
	SPANISH,
}

var SubjectID2Name = map[int]string{
	SubjectYuWen:           "语文",
	SubjectShuXue:          "数学",
	SubjectKeXue:           "科学",
	SubjectYingYu:          "英语",
	SubjectShengWu:         "生物",
	SubjectDiLi:            "地理",
	SubjectLiShi:           "历史",
	SubjectWuLi:            "物理",
	SubjectHuaXue:          "化学",
	SubjectZhengZhi:        "政治",
	SubjectXingQu:          "兴趣课",
	SubjectSiXiangPinDe:    "思想品德",
	SubjectJiangZuo:        "讲座",
	SubjectLiZong:          "理综",
	SubjectWenZong:         "文综",
	SubjectAoShu:           "奥数",
	SubjectDaoDeYuFaZhi:    "道德与法治",
	SubjectYinYue:          "音乐",
	SubjectMeiShu:          "美术",
	SubjectXinXiJiShu:      "信息技术",
	SubjectSiXiangZhengZhi: "思想政治",
	SubjectTiNengYunDong:   "体能运动",
	SubjectYiShuSuYang:     "艺术素养",
	SubjectSiWeiLuoJi:      "思维逻辑",
	SubjectKeJiChuangXin:   "科技创新",
	SubjectYuYanWenXue:     "语言文学",
	SubjectChuanTongWenHua: "传统文化",
	SubjectSheHuiShiJian:   "社会实践",
	SubjectTiYuYuJianKang:  "体育与健康",
	SubjectTongYongJiShu:   "通用技术",
	SubjectShuFa:           "书法",
	SubjectLaoDong:         "劳动",
	JAPANESE:               "日语",
	SPANISH:                "西班牙语",
}

var SubjectID2Bit = map[int64]int64{
	SubjectYuWen:           1,
	SubjectShuXue:          2,
	SubjectKeXue:           4,
	SubjectYingYu:          8,
	SubjectShengWu:         16,
	SubjectDiLi:            32,
	SubjectLiShi:           64,
	SubjectWuLi:            128,
	SubjectHuaXue:          256,
	SubjectZhengZhi:        512,
	SubjectXingQu:          1024,
	SubjectSiXiangPinDe:    2048,
	SubjectJiangZuo:        4096,
	SubjectLiZong:          8192,
	SubjectWenZong:         16384,
	SubjectAoShu:           32768,
	SubjectDaoDeYuFaZhi:    65536,
	SubjectYinYue:          131072,
	SubjectMeiShu:          262144,
	SubjectXinXiJiShu:      524288,
	SubjectSiXiangZhengZhi: 1048576,
	SubjectTiNengYunDong:   2097152,
	SubjectYiShuSuYang:     4194304,
	SubjectSiWeiLuoJi:      8388608,
	SubjectKeJiChuangXin:   16777216,
	SubjectYuYanWenXue:     33554432,
	SubjectChuanTongWenHua: 67108864,
	SubjectSheHuiShiJian:   134217728,
	SubjectTiYuYuJianKang:  268435456,
	SubjectTongYongJiShu:   536870912,
	SubjectShuFa:           1073741824,
	SubjectLaoDong:         2147483648,
	JAPANESE:               4294967296,
	SPANISH:                8589934592,
}

func GetSubjectName(ID int) string {
	if value, ok := SubjectID2Name[ID]; ok {
		return value
	}
	return "未知"
}
