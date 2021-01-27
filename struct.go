package main

type ArticleFinal struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	LikeNums int    `json:"like_nums"`
	ReadNums int    `json:"read_nums"`
	GoodNums int    `json:"good_nums"`
	ArticleNotLoginFiled
}

type ArticleNotLoginFiled struct {
	WriteSelf    bool   `json:"write_self"`
	ContainVideo bool   `json:"contain_video"`
	ContainAudio bool   `json:"contain_audio"`
	CreateTime   string `json:"create_time"`
}

type ArticleEx struct {
	AdvertisementInfo []interface{} `json:"advertisement_info"`
	AppmsgAlbumVideos []interface{} `json:"appmsg_album_videos"`
	Appmsgact         struct {
		FavoriteBefore float64 `json:"favorite_before"`
		OldLikedBefore float64 `json:"old_liked_before"`
		SeenBefore     float64 `json:"seen_before"`
		ShareBefore    float64 `json:"share_before"`
	} `json:"appmsgact"`
	Appmsgstat struct {
		FriendLikeNum   float64 `json:"friend_like_num"`
		IsLogin         bool    `json:"is_login"`
		LikeDisabled    bool    `json:"like_disabled"`
		LikeNum         float64 `json:"like_num"`
		Liked           bool    `json:"liked"`
		OldLikeNum      float64 `json:"old_like_num"`
		OldLiked        bool    `json:"old_liked"`
		OldLikedBefore  float64 `json:"old_liked_before"`
		Prompted        float64 `json:"prompted"`
		ReadNum         float64 `json:"read_num"`
		RealReadNum     float64 `json:"real_read_num"`
		Ret             float64 `json:"ret"`
		Show            bool    `json:"show"`
		ShowGray        float64 `json:"show_gray"`
		ShowLike        float64 `json:"show_like"`
		ShowLikeGray    float64 `json:"show_like_gray"`
		ShowOldLike     float64 `json:"show_old_like"`
		ShowOldLikeGray float64 `json:"show_old_like_gray"`
		ShowRead        float64 `json:"show_read"`
		Style           float64 `json:"style"`
		Version         float64 `json:"version"`
		VideoPv         float64 `json:"video_pv"`
		VideoUv         float64 `json:"video_uv"`
	} `json:"appmsgstat"`
	BaseResp struct {
		Wxtoken float64 `json:"wxtoken"`
	} `json:"base_resp"`
	FriendSubscribeCount float64       `json:"friend_subscribe_count"`
	IsFans               float64       `json:"is_fans"`
	MoreReadList         []interface{} `json:"more_read_list"`
	OriginalArticleCount float64       `json:"original_article_count"`
	PublicTagInfo        struct {
		Tags []interface{} `json:"tags"`
	} `json:"public_tag_info"`
	RelatedTagArticle []interface{} `json:"related_tag_article"`
	RelatedTagVideo   []interface{} `json:"related_tag_video"`
	RewardHeadImgs    []interface{} `json:"reward_head_imgs"`
	ShareFlag         struct {
		Show     float64 `json:"show"`
		ShowGray float64 `json:"show_gray"`
	} `json:"share_flag"`
	TestFlag          float64       `json:"test_flag"`
	VideoContinueFlag float64       `json:"video_continue_flag"`
	VideoSharePageTag []interface{} `json:"video_share_page_tag"`
}

type Msg struct {
	CanMsgContinue float64       `json:"can_msg_continue"`
	Errmsg         string        `json:"errmsg"`
	GeneralMsgList string        `json:"general_msg_list"`
	HomePageList   []interface{} `json:"home_page_list"`
	MsgCount       float64       `json:"msg_count"`
	NextOffset     float64       `json:"next_offset"`
	RealType       float64       `json:"real_type"`
	Ret            float64       `json:"ret"`
	UseVideoTab    float64       `json:"use_video_tab"`
	VideoCount     float64       `json:"video_count"`
}

type GeneralMsgList struct {
	List []struct {
		AppMsgExtInfo struct {
			AudioFileid            float64 `json:"audio_fileid"`
			Author                 string  `json:"author"`
			Content                string  `json:"content"`
			ContentURL             string  `json:"content_url"`
			CopyrightStat          float64 `json:"copyright_stat"`
			Cover                  string  `json:"cover"`
			DelFlag                float64 `json:"del_flag"`
			Digest                 string  `json:"digest"`
			Duration               float64 `json:"duration"`
			Fileid                 float64 `json:"fileid"`
			IsMulti                float64 `json:"is_multi"`
			ItemShowType           float64 `json:"item_show_type"`
			MaliciousContentType   float64 `json:"malicious_content_type"`
			MaliciousTitleReasonID float64 `json:"malicious_title_reason_id"`
			MultiAppMsgItemList    []struct {
				AudioFileid            float64 `json:"audio_fileid"`
				Author                 string  `json:"author"`
				Content                string  `json:"content"`
				ContentURL             string  `json:"content_url"`
				CopyrightStat          float64 `json:"copyright_stat"`
				Cover                  string  `json:"cover"`
				DelFlag                float64 `json:"del_flag"`
				Digest                 string  `json:"digest"`
				Duration               float64 `json:"duration"`
				Fileid                 float64 `json:"fileid"`
				ItemShowType           float64 `json:"item_show_type"`
				MaliciousContentType   float64 `json:"malicious_content_type"`
				MaliciousTitleReasonID float64 `json:"malicious_title_reason_id"`
				PlayURL                string  `json:"play_url"`
				SourceURL              string  `json:"source_url"`
				Title                  string  `json:"title"`
			} `json:"multi_app_msg_item_list"`
			PlayURL   string  `json:"play_url"`
			SourceURL string  `json:"source_url"`
			Subtype   float64 `json:"subtype"`
			Title     string  `json:"title"`
		} `json:"app_msg_ext_info"`
		CommMsgInfo struct {
			Content  string  `json:"content"`
			Datetime float64 `json:"datetime"`
			Fakeid   string  `json:"fakeid"`
			ID       float64 `json:"id"`
			Status   float64 `json:"status"`
			Type     float64 `json:"type"`
		} `json:"comm_msg_info"`
	} `json:"list"`
}

type ArticleInfo struct {
	A8scene float64 `json:"a8scene"`
	Bizinfo struct {
		Alias          string  `json:"alias"`
		Headimg        string  `json:"headimg"`
		IsBanned       float64 `json:"is_banned"`
		IsOriginalAcct float64 `json:"is_original_acct"`
		IsScanAcct     float64 `json:"is_scan_acct"`
		IsSubscribed   float64 `json:"is_subscribed"`
		IsVerify       float64 `json:"is_verify"`
		Nickname       string  `json:"nickname"`
		Signature      string  `json:"signature"`
		Username       string  `json:"username"`
		VerifySource   struct {
			Description   string  `json:"description"`
			Introurl      string  `json:"introurl"`
			Name          string  `json:"name"`
			Type          float64 `json:"type"`
			Verifybiztype float64 `json:"verifybiztype"`
		} `json:"verify_source"`
	} `json:"bizinfo"`
	CanMsgContinue  float64       `json:"can_msg_continue"`
	Errmsg          string        `json:"errmsg"`
	GeneralMsgList  string        `json:"general_msg_list"`
	HomePageList    []interface{} `json:"home_page_list"`
	HomePageListLen float64       `json:"home_page_list_len"`
	JssdkInfo       struct {
		Appid     string  `json:"appid"`
		Noncestr  string  `json:"noncestr"`
		Signature string  `json:"signature"`
		Timestamp float64 `json:"timestamp"`
	} `json:"jssdk_info"`
	MsgCount    float64 `json:"msg_count"`
	NextOffset  float64 `json:"next_offset"`
	RealType    float64 `json:"real_type"`
	Ret         float64 `json:"ret"`
	Scene       float64 `json:"scene"`
	UseDemo     float64 `json:"use_demo"`
	UseVideoTab float64 `json:"use_video_tab"`
	VideoCount  float64 `json:"video_count"`
}
