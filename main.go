package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Lecture struct {
	Id                    int       `json:"id"`
	Name                  string    `json:"name"`
	Instructions          string    `json:"instructions"`
	Notice                string    `json:"notice"`
	LiveType              int       `json:"liveType"`
	TeacherId             int       `json:"teacherId"`
	LiveChatRoom          int       `json:"-"`
	CourseWareId          int       `json:"coursewareId"`
	IsSeries              int       `json:"isBelongToSeries"`
	SeriesLectureId       int       `json:"seriesId"`
	Day                   string    `json:"-"`
	StartTime             time.Time `json:"-"` //提前15分钟
	StartTimeStr          string    `json:"-"`
	EndTime               time.Time `json:"-"` //延后半小时
	EndTimeStr            string    `json:"-"`
	StartTimeUnix         int64     `json:"stime"`
	EndTimeUnix           int64     `json:"etime"`
	GotoStartTime         string    `json:"-"`
	GotoEndTime           string    `json:"-"`
	GradeIds              string    `json:"grade_ids"`
	GradeIdsStr           string    `json:"gradeIds"`
	GradeIdList           []int     `json:"-"`
	SubjectIds            string    `json:"subject_ids"`
	SubjectIdsStr         string    `json:"subjectIds"`
	SubjectIdList         []int     `json:"-"`
	SubjectDigits         int       `json:"-"`
	Status                int       `json:"status"`        //返回给前端的状态
	OriginStatus          int       `json:"origin_status"` //数据库原生状态
	Recommend             int       `json:"-"`
	PlaybackRecommend     int       `json:"-"`
	LivePlaybackRecommend int       `json:"-"`
	OperationDepartment   int       `json:"-"`
	ImageUrl              string    `json:"imageUrl"`          //讲座学科小图
	BigImageUrl           string    `json:"bigImageUrl"`       //讲座学科大图
	IsBespoke             int       `json:"-"`                 //是否预约
	ScheduleTime          string    `json:"scheduleTime"`      //讲座时间 06月17日 10:30-12:00
	StudyreportStatus     int       `json:"studyreportStatus"` //学习报告
	StudyreportUrl        string    `json:"studyreportUrl"`
	TeacherName           string    `json:"teacherName"`      //教师名称
	TeacherImg            string    `json:"teacherImg"`       //教师头像
	TeacherSpell          string    `json:"teacherSpell"`     //教师拼音名
	LiveNum               int       `json:"liveNum"`          //在学人数
	ReservationNum        int       `json:"reservationNum"`   //预约人数
	VisitNum              int       `json:"visitNum"`         //学完人数
	PlaybackIsLogin       int       `json:"playback_islogin"` //观看回放是否需要登录
	AdvertiseIds          string    `json:"advertise_ids"`    //关联的广告ID
	ActivityUrl           string    `json:"-"`
	IsProgram             int       `json:"isProgram"`
	PlayProtocol          int       `json:"-"`
	StatusByTime          int       `json:"-"`      // unStart 2, start 3, end 4
	IsGray                int       `json:"isGray"` //是否灰度讲座
	TestIds               string    `json:"testIds"`
	H5SourceIds           string    `json:"h5SourceIds"` //关联的H5素材ID
	ProgrammingUrl        string    `json:"ProgrammingUrl"`
	IsGently              int       `json:"isGently"`             //是否为轻直播（0 否  1 是）
	TicketIds             string    `json:"ticketIds"`            //优惠券id list
	CourseIds             string    `json:"courseIds"`            //课程id list
	GentlyNotice          string    `json:"gentlyNotice"`         //轻直播公告
	IsNeedLogin           int       `json:"isNeedLogin"`          //是否需要登录
	OfficialAccount       string    `json:"officialAccount"`      //公众号
	TeachingVersionIds    string    `json:"teachingVersionIds"`   //教学版本ids
	LiveRoomType          int       `json:"liveRoomType"`         //直播间类型
	FutureCoursewareId    int       `json:"future_courseware_id"` //未来课件id
	CreatorId             int       `json:"creatorId"`            // 未来号
	LoginCheckType        int       `json:"loginCheckType"`       //登录验证方式
	GoingId               int       `json:"goingId"`              //新版学习中心---系列讲座中正在直播的场次id
	FristView             int       `json:"fristView"`            //是否上首页
	SortWeight            int       `json:"sortWeight"`           //首页权重
	CoverPicture          string    `json:"coverPicture"`         //首页封面图
	UvNum                 int64     `json:"UvNum"`
	StuNum                int64     `json:"StuNum"`
	PlaybackUvNum         int64     `json:"PlaybackUvNum"`
	PlaybackVisitNum      int64     `json:"PlaybackVisitNum"`
	WxType                int       `json:"wxType"` //微信按钮类型
	ExistsPlayback        int       `json:"existsPlayback"`
	HotId                 int       `json:"hotId"`           // 热点ID
	IsOpenGraffiti        int       `json:"-"`               // 是否开启涂鸦插件
	IsIpad                int       `json:"-"`               // 是否iPad直播
	PlaybackChatPic       string    `json:"playbackChatPic"` // 回放聊天区遮罩图片url
	AppIsHandstand        int       `json:"-"`               // 是否竖版直播
	LiveScene             int       `json:"-"`               // 直播场景
	CourseListImg         string    `json:"-"`               // 课表页展示图
	Cover                 string    `json:"cover"`
	CoverWidth            int       `json:"cover_width"`
	CoverHeight           int       `json:"cover_height"`
	Cover32               string    `json:"cover_32"`
	Cover32Width          int       `json:"cover_32_width"`
	Cover32Height         int       `json:"cover_32_height"`
	ProvinceIds           string    `json:"province_ids"`
	HomeShowType          string    `json:"home_show_type"`
	Portrait              string    `json:"portrait"`
	LectureAuthorType     int       `json:"lecture_author_type"`
	ExtSourceId           int       `json:"ext_source_id"`
	SourceType            int       `json:"source_type"`
	VersionIds            string    `json:"version_ids"`
	IsDisplay             int       `json:"is_display"`
	ContentPoolType       int       `json:"content_pool_type"`
	ProvinceIdsList       []int     `json:"-"`
	ClassId               int       `json:"class_id"`
	IsMoreMic             int       `json:"is_more_mic"`
	IsClassGray           int       `json:"is_class_gray"`
	ClassLiveRange        int       `json:"class_live_range"`
	ClassIds              string    `json:"class_ids"`
	ClassLiveType         int       `json:"class_live_type"`
}

func main() {
	redisReply := "\xde\x001\xa2Id\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa4Name\xa0\xacInstructions\xa0\xa6Notice\xa0\xa8LiveType\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa9TeacherId\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacLiveChatRoom\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacCourseWareId\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa8IsSeries\xd3\x00\x00\x00\x00\x00\x00\x00\x01\xafSeriesLectureId\xd3\x00\x00\x00\x00\x00\x00\x01\xb6\xa3Day\xa0\xa9StartTime\xc7\x0c\xff\x00\x00\x00\x00\xff\xff\xff\xf1\x88n\t\x00\xacStartTimeStr\xa0\xa7EndTime\xc7\x0c\xff\x00\x00\x00\x00\xff\xff\xff\xf1\x88n\t\x00\xaaEndTimeStr\xa0\xadStartTimeUnix\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xabEndTimeUnix\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa8GradeIds\xa0\xabGradeIdsStr\xa0\xabGradeIdList\xc0\xaaSubjectIds\xa0\xadSubjectIdsStr\xa0\xadSubjectIdList\xc0\xadSubjectDigits\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa6Status\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacOriginStatus\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa9Recommend\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xb1PlaybackRecommend\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xb5LivePlaybackRecommend\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xb3OperationDepartment\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa8ImageUrl\xa0\xabBigImageUrl\xa0\xa9IsBespoke\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacScheduleTime\xa0\xb1StudyreportStatus\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xaeStudyreportUrl\xa0\xabTeacherName\xa0\xaaTeacherImg\xa0\xacTeacherSpell\xa0\xa7LiveNum\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xaeReservationNum\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xa8VisitNum\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xafPlaybackIsLogin\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacAdvertiseIds\xa0\xabH5Sourceids\xa0\xabActivityUrl\xa0\xa9IsProgram\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacPlayProtocol\xd3\x00\x00\x00\x00\x00\x00\x00\x00\xacStatusByTime\xd3\x00\x00\x00\x00\x00\x00\x00\x00"
	bytes := []byte(redisReply)
	ins := &Lecture{}
	if err := json.Unmarshal(bytes, ins); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(redisReply)
}
