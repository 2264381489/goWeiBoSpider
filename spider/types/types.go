package types

type GetIndexResp struct {
	Ok   int  `json:"ok"`
	Data Data `json:"data"`
}

type Data struct {
	CardlistInfo CardlistInfo `json:"cardlistInfo"`
	Cards        []Cards      `json:"cards"`
	Scheme       string       `json:"scheme"`
	ShowAppTips  int          `json:"showAppTips"`
}

type CardlistInfo struct {
	Containerid       string `json:"containerid"`
	VP                int    `json:"v_p"`
	ShowStyle         int    `json:"show_style"`
	Total             int    `json:"total"`
	AutoLoadMoreIndex int    `json:"autoLoadMoreIndex"`
	SinceID           int64  `json:"since_id"`
}

type Visible struct {
	Type   int `json:"type"`
	ListID int `json:"list_id"`
}
type Badge struct {
	Dzwbqlx2016          int `json:"dzwbqlx_2016"`
	FollowWhitelistVideo int `json:"follow_whitelist_video"`
	UserNameCertificate  int `json:"user_name_certificate"`
	Suishoupai2018       int `json:"suishoupai_2018"`
	Memorial2018         int `json:"memorial_2018"`
	PcNew                int `json:"pc_new"`
	PartyCardidState     int `json:"party_cardid_state"`
	UserIdentityAuth     int `json:"user_identity_auth"`
	UserRealityAuth      int `json:"user_reality_auth"`
}
type User struct {
	ID                int64  `json:"id"`
	ScreenName        string `json:"screen_name"`
	ProfileImageURL   string `json:"profile_image_url"`
	ProfileURL        string `json:"profile_url"`
	StatusesCount     int    `json:"statuses_count"`
	Verified          bool   `json:"verified"`
	VerifiedType      int    `json:"verified_type"`
	VerifiedTypeExt   int    `json:"verified_type_ext"`
	VerifiedReason    string `json:"verified_reason"`
	CloseBlueV        bool   `json:"close_blue_v"`
	Description       string `json:"description"`
	Gender            string `json:"gender"`
	Mbtype            int    `json:"mbtype"`
	Svip              int    `json:"svip"`
	Urank             int    `json:"urank"`
	Mbrank            int    `json:"mbrank"`
	FollowMe          bool   `json:"follow_me"`
	Following         bool   `json:"following"`
	FollowCount       int    `json:"follow_count"`
	FollowersCount    string `json:"followers_count"`
	FollowersCountStr string `json:"followers_count_str"`
	CoverImagePhone   string `json:"cover_image_phone"`
	AvatarHd          string `json:"avatar_hd"`
	Like              bool   `json:"like"`
	LikeMe            bool   `json:"like_me"`
	Badge             Badge  `json:"badge"`
	SpecialFollow     bool   `json:"special_follow"`
}
type NumberDisplayStrategy struct {
	ApplyScenarioFlag    int    `json:"apply_scenario_flag"`
	DisplayTextMinNumber int    `json:"display_text_min_number"`
	DisplayText          string `json:"display_text"`
}
type CommentManageInfo struct {
	CommentPermissionType int `json:"comment_permission_type"`
	ApprovalCommentType   int `json:"approval_comment_type"`
	CommentSortType       int `json:"comment_sort_type"`
}
type HotPage struct {
	Fid            string `json:"fid"`
	FeedDetailType int    `json:"feed_detail_type"`
}
type EditConfig struct {
	Edited bool `json:"edited"`
}
type Geo struct {
	Width  interface{} `json:"width"`
	Height interface{} `json:"height"`
	Croped bool        `json:"croped"`
}
type Large struct {
	Size string `json:"size"`
	URL  string `json:"url"`
	Geo  Geo    `json:"geo"`
}
type Pics struct {
	Pid   string `json:"pid"`
	URL   string `json:"url"`
	Size  string `json:"size"`
	Geo   Geo    `json:"geo"`
	Large Large  `json:"large"`
}
type Title struct {
	Text      string `json:"text"`
	BaseColor int    `json:"base_color"`
}

type MediaInfo struct {
	StreamURL   string  `json:"stream_url"`
	StreamURLHd string  `json:"stream_url_hd"`
	Duration    float64 `json:"duration"`
}
type Urls struct {
	Mp4720PMp4 string `json:"mp4_720p_mp4"`
	Mp4HdMp4   string `json:"mp4_hd_mp4"`
	Mp4LdMp4   string `json:"mp4_ld_mp4"`
}
type PageInfo struct {
	Type             string    `json:"type"`
	ObjectType       int       `json:"object_type"`
	URLOri           string    `json:"url_ori"`
	PagePic          PagePic   `json:"page_pic"`
	PageURL          string    `json:"page_url"`
	ObjectID         string    `json:"object_id"`
	PageTitle        string    `json:"page_title"`
	Title            string    `json:"title"`
	Content1         string    `json:"content1"`
	Content2         string    `json:"content2"`
	VideoOrientation string    `json:"video_orientation"`
	PlayCount        string    `json:"play_count"`
	MediaInfo        MediaInfo `json:"media_info"`
	Urls             Urls      `json:"urls"`
}

type PagePic struct {
	URL string `json:"url"`
}

type RetweetedStatus struct {
	Visible                  Visible               `json:"visible"`
	CreatedAt                string                `json:"created_at"`
	ID                       string                `json:"id"`
	Mid                      string                `json:"mid"`
	CanEdit                  bool                  `json:"can_edit"`
	ShowAdditionalIndication int                   `json:"show_additional_indication"`
	Text                     string                `json:"text"`
	TextLength               int                   `json:"textLength"`
	Source                   string                `json:"source"`
	Favorited                bool                  `json:"favorited"`
	PicIds                   []string              `json:"pic_ids"`
	ThumbnailPic             string                `json:"thumbnail_pic"`
	BmiddlePic               string                `json:"bmiddle_pic"`
	OriginalPic              string                `json:"original_pic"`
	IsPaid                   bool                  `json:"is_paid"`
	MblogVipType             int                   `json:"mblog_vip_type"`
	User                     User                  `json:"user"`
	PicStatus                string                `json:"picStatus"`
	RepostsCount             int                   `json:"reposts_count"`
	CommentsCount            int                   `json:"comments_count"`
	ReprintCmtCount          int                   `json:"reprint_cmt_count"`
	AttitudesCount           int                   `json:"attitudes_count"`
	PendingApprovalCount     int                   `json:"pending_approval_count"`
	IsLongText               bool                  `json:"isLongText"`
	ShowMlevel               int                   `json:"show_mlevel"`
	DarwinTags               []interface{}         `json:"darwin_tags"`
	AdMarked                 bool                  `json:"ad_marked"`
	Mblogtype                int                   `json:"mblogtype"`
	ItemCategory             string                `json:"item_category"`
	Rid                      string                `json:"rid"`
	NumberDisplayStrategy    NumberDisplayStrategy `json:"number_display_strategy"`
	ContentAuth              int                   `json:"content_auth"`
	CommentManageInfo        CommentManageInfo     `json:"comment_manage_info"`
	PicNum                   int                   `json:"pic_num"`
	HotPage                  HotPage               `json:"hot_page"`
	NewCommentStyle          int                   `json:"new_comment_style"`
	Mlevel                   int                   `json:"mlevel"`
	EditConfig               EditConfig            `json:"edit_config"`
	PageInfo                 PageInfo              `json:"page_info"`
	Pics                     []Pics                `json:"pics"`
	Bid                      string                `json:"bid"`
}

type Mblog struct {
	Visible                  Visible               `json:"visible"`
	CreatedAt                string                `json:"created_at"`
	ID                       string                `json:"id"`
	Mid                      string                `json:"mid"`
	CanEdit                  bool                  `json:"can_edit"`
	ShowAdditionalIndication int                   `json:"show_additional_indication"`
	Text                     string                `json:"text"`
	Source                   string                `json:"source"`
	Favorited                bool                  `json:"favorited"`
	PicIds                   []interface{}         `json:"pic_ids"`
	IsPaid                   bool                  `json:"is_paid"`
	MblogVipType             int                   `json:"mblog_vip_type"`
	User                     User                  `json:"user"`
	RetweetedStatus          RetweetedStatus       `json:"retweeted_status"`
	RepostsCount             int64                 `json:"reposts_count"`
	CommentsCount            int64                 `json:"comments_count"`
	ReprintCmtCount          int                   `json:"reprint_cmt_count"`
	AttitudesCount           int64                 `json:"attitudes_count"`
	PendingApprovalCount     int                   `json:"pending_approval_count"`
	IsLongText               bool                  `json:"isLongText"`
	ShowMlevel               int                   `json:"show_mlevel"`
	DarwinTags               []interface{}         `json:"darwin_tags"`
	AdMarked                 bool                  `json:"ad_marked"`
	Mblogtype                int                   `json:"mblogtype"`
	ItemCategory             string                `json:"item_category"`
	Rid                      string                `json:"rid"`
	ExternSafe               int                   `json:"extern_safe"`
	NumberDisplayStrategy    NumberDisplayStrategy `json:"number_display_strategy"`
	ContentAuth              int                   `json:"content_auth"`
	CommentManageInfo        CommentManageInfo     `json:"comment_manage_info"`
	HideHotFlow              int                   `json:"hide_hot_flow"`
	RepostType               int                   `json:"repost_type"`
	PicNum                   int                   `json:"pic_num"`
	HotPage                  HotPage               `json:"hot_page"`
	NewCommentStyle          int                   `json:"new_comment_style"`
	AbSwitcher               int                   `json:"ab_switcher"`
	Mlevel                   int                   `json:"mlevel"`
	RegionName               string                `json:"region_name"`
	RegionOpt                int                   `json:"region_opt"`
	AnalysisExtra            string                `json:"analysis_extra"`
	MblogMenuNewStyle        int                   `json:"mblog_menu_new_style"`
	EditConfig               EditConfig            `json:"edit_config"`
	RawText                  string                `json:"raw_text"`
	Bid                      string                `json:"bid"`
	Pics                     []Pics                `json:"pics"`
	PageInfo                 PageInfo              `json:"page_info"`
}
type Cards struct {
	CardType      int           `json:"card_type"`
	CommendInfo   []CommendInfo `json:"commend_info,omitempty"`
	ProfileTypeID string        `json:"profile_type_id,omitempty"`
	Itemid        string        `json:"itemid,omitempty"`
	Scheme        string        `json:"scheme,omitempty"`
	Mblog         Mblog         `json:"mblog,omitempty"`
}

type CommendInfo struct {
	Icon        string    `json:"icon"`
	Text        string    `json:"text"`
	AccessRight int       `json:"access_right"`
	Scheme      string    `json:"scheme"`
	ActionLog   ActionLog `json:"action_log"`
}
type ActionLog struct {
	ActCode int    `json:"act_code"`
	Oid     string `json:"oid"`
	Ext     string `json:"ext"`
	Fid     string `json:"fid"`
}
