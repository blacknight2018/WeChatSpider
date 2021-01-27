package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var urlHome string
var urlGetMsg string
var urlGetMsgEx string
var offset int
var count int
var cookie string
var curMpInfo ArticleInfo

func getArticleDetail(mpUrl string) ArticleEx {
	var ret ArticleEx
	respText := HttpsGet(mpUrl, nil)
	biz := strings.Index(respText, "__biz")
	if biz != -1 {
		chksm := strings.Index(respText, "chksm")
		if chksm != -1 {
			param := respText[biz:chksm]
			param = strings.ReplaceAll(param, "&amp;", "&") + "is_only_read=1"
			respText = HttpsPost(urlGetMsgEx, nil, &param)
			json.Unmarshal([]byte(respText), &ret)
		}
	}
	fmt.Println(respText)
	time.Sleep(time.Second * 5)
	return ret
}

func getNotLoginFiled(mpUrl string) ArticleNotLoginFiled {
	var ret ArticleNotLoginFiled

	respText := HttpsGet(mpUrl, nil)

	//这几个的算法都是有BUG的，待修正
	ret.WriteSelf = strings.Index(respText, "原创") != -1
	ret.ContainVideo = strings.Index(respText, `vid=“wxv`) != -1
	ret.ContainAudio = strings.Index(respText, `voice_encode_fileid=`) != -1

	p1 := strings.Index(respText, `i="2`)
	p2 := strings.Index(respText, `e(t,n,i,document.getElementById`)
	if p1 != -1 && p2 != -1 {
		ret.CreateTime = respText[p1+3 : p2-3]
	}
	return ret
}

func main() {

	g := gin.Default()
	g = g.Delims("[[", "]]")
	g.LoadHTMLFiles("./static/index.html")
	g.GET("/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"head_img": curMpInfo.Bizinfo.Headimg, "nick_name": curMpInfo.Bizinfo.Nickname, "signature": curMpInfo.Bizinfo.Signature})
	})
	g.GET("/", func(context *gin.Context) {
		urlData, _ := url.QueryUnescape(context.Query("url"))
		params := strings.Split(urlData, "&")
		offset = 0
		count = 10
		urlHome = "https://mp.weixin.qq.com/mp/profile_ext?action=home" + "&" + params[1] + "&" + params[3] + "&" + params[4] + "&" + "f=json"
		urlGetMsg = "https://mp.weixin.qq.com/mp/profile_ext?action=getmsg" + "&" + params[1] + "&" + params[3] + "&" + params[4] + "&" + "f=json"
		urlGetMsgEx = "https://mp.weixin.qq.com/mp/getappmsgext?" + params[3] + "&" + params[4]
		cookie, _ = url.QueryUnescape(context.Query("cookie"))
		jsonText := HttpsGet(urlHome, &cookie)
		var t ArticleInfo
		json.Unmarshal([]byte(jsonText), &t)
		curMpInfo = t
	})
	g.GET("/mp", func(context *gin.Context) {
		localOffset := MustA2I(context.Query("offset"))
		localCount := MustA2I(context.Query("count"))
		offset = localOffset
		count = localCount
		tmpUrlGetMsg := urlGetMsg + "&" + "offset=" + strconv.Itoa(offset) + "&count=" + strconv.Itoa(count)
		jsonText := HttpsGet(tmpUrlGetMsg, &cookie)
		var t Msg
		var d GeneralMsgList
		json.Unmarshal([]byte(jsonText), &t)
		json.Unmarshal([]byte(t.GeneralMsgList), &d)
		fmt.Println(d)
		var res []ArticleFinal
		for _, v := range d.List {
			var tmp ArticleFinal
			//fmt.Println(v.AppMsgExtInfo.Title, v.AppMsgExtInfo.ContentURL)
			tmp.Title = v.AppMsgExtInfo.Title
			tmp.Author = v.AppMsgExtInfo.Author
			if len(tmp.Author) == 0 {
				tmp.Author = curMpInfo.Bizinfo.Nickname
			}
			tmp.ArticleNotLoginFiled = getNotLoginFiled(v.AppMsgExtInfo.ContentURL)
			if len(tmp.Title) != 0 {
				detail := getArticleDetail(v.AppMsgExtInfo.ContentURL)
				tmp.ReadNums = int(detail.Appmsgstat.ReadNum)
				tmp.LikeNums = int(detail.Appmsgstat.LikeNum)
				tmp.GoodNums = int(detail.Appmsgstat.OldLikeNum)
				res = append(res, tmp)
			}
			for _, v2 := range v.AppMsgExtInfo.MultiAppMsgItemList {
				//fmt.Println(v2.Title, v2.ContentURL)
				tmp.Title = v2.Title
				tmp.Author = v2.Author
				if len(tmp.Author) == 0 {
					tmp.Author = curMpInfo.Bizinfo.Nickname
				}
				tmp.ArticleNotLoginFiled = getNotLoginFiled(v2.ContentURL)
				if len(tmp.Title) != 0 {
					detail := getArticleDetail(v2.ContentURL)
					tmp.ReadNums = int(detail.Appmsgstat.ReadNum)
					tmp.LikeNums = int(detail.Appmsgstat.LikeNum)
					tmp.GoodNums = int(detail.Appmsgstat.OldLikeNum)
					res = append(res, tmp)
				}
			}
		}
		context.JSON(http.StatusOK, res)

	})
	g.GET("/test", func(context *gin.Context) {
		url := context.Query("url")
		context.JSON(http.StatusOK, getArticleDetail(url))
	})
	g.Run(":8887")
}
