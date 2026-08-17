package interviewer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/pkg/errors"

	db "kcers-survey/biz/dal/db"
	"kcers-survey/biz/dal/db/ent"
	logs2 "kcers-survey/biz/dal/db/ent/logs"
	"kcers-survey/biz/dal/db/ent/predicate"
	"kcers-survey/biz/dal/db/ent/role"
	user2 "kcers-survey/biz/dal/db/ent/user"
	surveyresponse2 "kcers-survey/biz/dal/db/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/ent/surveyresponseanswers"
	survey2 "kcers-survey/biz/dal/db/ent/survey"
	smslog "kcers-survey/biz/dal/db/ent/smslog"
	"kcers-survey/biz/dal/config"
	"kcers-survey/biz/dal/sms"
	"kcers-survey/biz/infras/enums"
	"kcers-survey/biz/infras/service"
	"kcers-survey/biz/pkg/encrypt"
	"kcers-survey/idl_gen/model/pub"

)

type Interviewer struct {
	ctx context.Context
	c   *app.RequestContext
	db  *ent.Client
}

type LoginResp struct {
	UserId     int64    `json:"userId"`
	Mobile     string   `json:"mobile"`
	Name       string   `json:"name"`
	NeedSetup  bool     `json:"needSetup"`
	RoleIDs    []int64  `json:"roleIds,omitempty"`
	RoleValues []string `json:"roleValues,omitempty"`
}

type InterviewerSurveyListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type InterviewerUpdateAnswerReq struct {
	SurveyID   int64    `json:"surveyId"`
	Sn         string   `json:"sn"`
	QuestionID int64    `json:"questionId"`
	Answer     []string `json:"answer"`
	AnswerText string   `json:"answerText"`
	Type       string   `json:"type"`
}

const InterviewerRoleValue = "interviewer"

func (s *Interviewer) Login(mobile, password string) (*LoginResp, error) {
	u, err := s.db.User.Query().
		Where(user2.UsernameEQ(mobile), user2.Status(1)).
		WithRoles().
		First(s.ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "query user failed")
	}

	if u == nil {
		// 验证该手机号在调查问卷中是否有记录
		researcherName := s.GetResearcherName(mobile)
		if researcherName == "" {
			return nil, errors.New("该手机号未关联任何调查问卷，无法登录")
		}
		return &LoginResp{
			Mobile:    mobile,
			Name:      researcherName,
			NeedSetup: true,
		}, nil
	} else {
		if ok := encrypt.VerifyPassword(password, u.Password); !ok {
			return nil, errors.New("密码错误")
		}
		hasRole := false
		for _, r := range u.Edges.Roles {
			if r.Value == InterviewerRoleValue {
				hasRole = true
				break
			}
		}
		if !hasRole {
			roleID := s.getInterviewerRoleID()
			if roleID == 0 {
				return nil, errors.New("调查员角色未配置，请联系管理员")
			}
			_, err = u.Update().AddRoleIDs(roleID).Save(s.ctx)
			if err != nil {
				return nil, errors.Wrap(err, "assign interviewer role failed")
			}
			u, err = s.db.User.Query().
				Where(user2.IDEQ(u.ID)).
				WithRoles().
				First(s.ctx)
			if err != nil {
				return nil, errors.Wrap(err, "reload user failed")
			}
		}
	}

	return &LoginResp{
		UserId: u.ID,
		Mobile: u.Mobile,
		Name:   u.Name,
	}, nil
}

func (s *Interviewer) GetResearcherName(mobile string) string {
	first, err := s.db.SurveyResponse.Query().
		Where(surveyresponse2.ResearcherPhone(mobile), surveyresponse2.ResearcherNEQ("")).
		Order(ent.Desc(surveyresponse2.FieldCreatedAt)).
		First(s.ctx)
	if err != nil || first == nil {
		return ""
	}
	return first.Researcher
}

func (s *Interviewer) SetupPassword(mobile, name, password string) (*LoginResp, error) {
	// 再次验证手机号在调查问卷中有记录
	researcherName := s.GetResearcherName(mobile)
	if researcherName == "" {
		return nil, errors.New("该手机号未关联任何调查问卷，无法注册")
	}

	u, err := s.createInterviewer(mobile, researcherName, password)
	if err != nil {
		return nil, err
	}

	roleIDs, _ := s.GetRoleIDs(u.ID)
	roleValues, _ := s.GetRoleValues(u.ID)

	return &LoginResp{
		UserId:    u.ID,
		Mobile:    u.Mobile,
		Name:      u.Name,
		NeedSetup: false,
		RoleIDs:   roleIDs,
		RoleValues: roleValues,
	}, nil
}

func (s *Interviewer) createInterviewer(mobile, name, password string) (*ent.User, error) {
	roleID := s.getInterviewerRoleID()
	if roleID == 0 {
		return nil, errors.New("调查员角色未配置，请联系管理员")
	}

	hashedPassword, err := encrypt.Crypt(password)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt password failed")
	}

	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "starting transaction")
	}

	gender := enums.ReturnMemberGenderKey("未知")
	u, err := tx.User.Create().
		SetUsername(mobile).
		SetMobile(mobile).
		SetName(name).
		SetPassword(hashedPassword).
		SetGender(gender).
		AddRoleIDs(roleID).
		SetStatus(1).
		Save(s.ctx)
	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create interviewer failed"))
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "commit transaction failed")
	}

	return u, nil
}

func (s *Interviewer) getInterviewerRoleID() int64 {
	r, err := s.db.Role.Query().Where(role.ValueEQ(InterviewerRoleValue)).First(s.ctx)
	if err != nil {
		return 0
	}
	return r.ID
}

func (s *Interviewer) GetRoleIDs(userID int64) ([]int64, error) {
	u, err := s.db.User.Query().Where(user2.IDEQ(userID)).WithRoles().First(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query user failed")
	}
	ids := make([]int64, 0)
	for _, r := range u.Edges.Roles {
		ids = append(ids, r.ID)
	}
	return ids, nil
}

func (s *Interviewer) GetRoleValues(userID int64) ([]string, error) {
	u, err := s.db.User.Query().Where(user2.IDEQ(userID)).WithRoles().First(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query user failed")
	}
	values := make([]string, 0)
	for _, r := range u.Edges.Roles {
		values = append(values, r.Value)
	}
	return values, nil
}

func (s *Interviewer) GetSurveys(mobile string, page, pageSize int64) (resp []*pub.QueryByPhoneResp, total int, err error) {
	all, err := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.ResearcherPhone(mobile),
			surveyresponse2.SnNEQ(""),
			surveyresponse2.AnswersCountNEQ(0),
			surveyresponse2.Delete(0),
		).
		Order(ent.Desc(surveyresponse2.FieldCreatedAt)).
		Offset(int(page-1) * int(pageSize)).
		Limit(int(pageSize)).
		All(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	surveyIDs := make([]int64, 0, len(all))
	for _, v := range all {
		surveyIDs = append(surveyIDs, v.SurveyID)
	}

	surveyMap := make(map[int64]string)
	if len(surveyIDs) > 0 {
		surveys, err := s.db.Survey.Query().
			Where(survey2.IDIn(surveyIDs...)).
			All(s.ctx)
		if err != nil {
			return nil, 0, err
		}
		for _, sv := range surveys {
			surveyMap[sv.ID] = sv.Title
		}
	}

	for _, v := range all {
		resp = append(resp, &pub.QueryByPhoneResp{
			SurveyID:        v.SurveyID,
			SurveyTitle:     surveyMap[v.SurveyID],
			Sn:              v.Sn,
			Respondent:      v.Respondent,
			RespondentPhone: v.RespondentPhone,
			Address:         v.Address,
			CreatedAt:       v.CreatedAt.Add(8 * time.Hour).Format(time.DateTime),
			AnswersCount:    v.AnswersCount,
		})
	}

	total, err = s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.ResearcherPhone(mobile),
			surveyresponse2.SnNEQ(""),
			surveyresponse2.AnswersCountNEQ(0),
			surveyresponse2.Delete(0),
		).
		Count(s.ctx)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (s *Interviewer) UpdateAnswer(req *InterviewerUpdateAnswerReq) error {
	resp, err := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.SurveyID(req.SurveyID),
			surveyresponse2.Sn(req.Sn),
			surveyresponse2.Delete(0),
		).
		First(s.ctx)
	if err != nil {
		return errors.Wrap(err, "query survey response failed")
	}

	ra, err := s.db.SurveyResponseAnswers.Query().
		Where(
			surveyresponseanswers2.SurveyID(req.SurveyID),
			surveyresponseanswers2.SurveyResponseID(resp.ID),
			surveyresponseanswers2.SurveyQuestionID(req.QuestionID),
		).
		First(s.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return errors.Wrap(err, "query answer failed")
	}

	if ra == nil {
		return errors.New("答案记录不存在")
	}

	rau := ra.Update()
	if req.Type == "input" || req.Type == "text" {
		rau.SetAnswerText(req.AnswerText)
	} else {
		rau.SetAnswer(req.Answer)
	}
	_, err = rau.Save(s.ctx)
	if err != nil {
		return errors.Wrap(err, "update answer failed")
	}

	return nil
}

func (s *Interviewer) CheckOwnership(mobile string, surveyID int64, sn string) bool {
	exist, _ := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.SurveyID(surveyID),
			surveyresponse2.Sn(sn),
			surveyresponse2.ResearcherPhone(mobile),
			surveyresponse2.Delete(0),
		).
		Exist(s.ctx)
	return exist
}

func (s *Interviewer) ChangePassword(userID int64, oldPassword, newPassword string) error {
	u, err := s.db.User.Query().Where(user2.IDEQ(userID)).First(s.ctx)
	if err != nil {
		return errors.Wrap(err, "query user failed")
	}
	if ok := encrypt.VerifyPassword(oldPassword, u.Password); !ok {
		return errors.New("原密码错误")
	}
	hashed, err := encrypt.Crypt(newPassword)
	if err != nil {
		return errors.Wrap(err, "encrypt password failed")
	}
	_, err = u.Update().SetPassword(hashed).Save(s.ctx)
	return err
}

func (s *Interviewer) SendSmsCaptcha(mobile string) error {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	bizId, err := sms.SendVerifyCode(
		mobile,
		config.GlobalServerConfig.Aliyun.Sms.Captcha.SignName,
		config.GlobalServerConfig.Aliyun.Sms.Captcha.TemplateCode,
		code,
	)
	if err != nil {
		return err
	}

	_, err = s.db.SmsLog.Create().
		SetMobile(mobile).
		SetBizID(bizId).
		SetCode(code).
		SetContent("调查员注册验证码").
		SetTemplate(config.GlobalServerConfig.Aliyun.Sms.Captcha.TemplateCode).
		SetNotifyType(2). // 2=员工
		SetStatus(1).
		Save(s.ctx)
	return err
}

func (s *Interviewer) VerifySmsCaptcha(mobile, code string) bool {
	exist, _ := s.db.SmsLog.Query().
		Where(
			smslog.MobileEQ(mobile),
			smslog.CodeEQ(code),
			smslog.StatusEQ(1),
			smslog.CreatedAtGTE(time.Now().Add(-5*time.Minute)),
		).
		Exist(s.ctx)
	return exist
}

type LogSearchReq struct {
	Keyword  string `json:"keyword"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

type LogEntry struct {
	ID          int64  `json:"id"`
	API         string `json:"api"`
	Method      string `json:"method"`
	Success     bool   `json:"success"`
	ReqContent  string `json:"reqContent"`
	RespContent string `json:"respContent"`
	IP          string `json:"ip"`
	UserAgent   string `json:"userAgent"`
	Operatorsr  string `json:"operatorsr"`
	Time        int64  `json:"time"`
	CreatedAt   string `json:"createdAt"`
}

type LogSearchResp struct {
	List  []*LogEntry `json:"list"`
	Total int         `json:"total"`
}

type reqJSON struct {
	Sn string `json:"sn"`
}

func (s *Interviewer) SearchLogs(req *LogSearchReq) (*LogSearchResp, error) {
	kw := req.Keyword
	if kw == "" {
		return &LogSearchResp{List: []*LogEntry{}, Total: 0}, nil
	}

	snSet := make(map[string]bool)
	ipSet := make(map[string]bool)

	// Step 1a: 按关键字（手机号）搜索提交日志，提取 sn 和 ip
	firstLogs, err := s.db.Logs.Query().
		Where(
			logs2.APIEQ("/service/survey/response/create"),
			logs2.ReqContentContains(kw),
		).
		All(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query logs by keyword failed")
	}
	for _, l := range firstLogs {
		if l.IP != "" {
			ipSet[l.IP] = true
		}
		var rj reqJSON
		if err := json.Unmarshal([]byte(l.ReqContent), &rj); err == nil && rj.Sn != "" {
			snSet[rj.Sn] = true
		}
	}

	// Step 1b: 按 address 搜索 survey_response，提取 sn
	responses, err := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.AddressContains(kw),
			surveyresponse2.Delete(0),
		).
		All(s.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "query survey_response by address failed")
	}
	for _, r := range responses {
		if r.Sn != "" {
			snSet[r.Sn] = true
		}
		if r.IP != "" {
			ipSet[r.IP] = true
		}
	}

	// Step 1c: 按 researcher_phone 搜索 survey_response，提取 sn
	responsesByPhone, err := s.db.SurveyResponse.Query().
		Where(
			surveyresponse2.ResearcherPhoneContains(kw),
			surveyresponse2.Delete(0),
		).
		All(s.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "query survey_response by phone failed")
	}
	for _, r := range responsesByPhone {
		if r.Sn != "" {
			snSet[r.Sn] = true
		}
		if r.IP != "" {
			ipSet[r.IP] = true
		}
	}

	// Step 1d: 如果关键字本身是 IP，直接加入
	if isIP(kw) {
		ipSet[kw] = true
	}

	// Step 2: 根据提取的 sn 和 ip 构建查询
	var orPreds []predicate.Logs
	for sn := range snSet {
		orPreds = append(orPreds, logs2.ReqContentContains(sn))
	}
	ips := make([]string, 0, len(ipSet))
	for ip := range ipSet {
		ips = append(ips, ip)
	}
	if len(ips) > 0 {
		orPreds = append(orPreds, logs2.IPIn(ips...))
	}

	if len(orPreds) == 0 {
		return &LogSearchResp{List: []*LogEntry{}, Total: 0}, nil
	}

	finalPred := logs2.Or(orPreds...)

	allLogs, err := s.db.Logs.Query().
		Where(finalPred).
		Order(ent.Desc(logs2.FieldCreatedAt)).
		Offset(int((req.Page - 1) * req.PageSize)).
		Limit(int(req.PageSize)).
		All(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query related logs failed")
	}

	count, err := s.db.Logs.Query().
		Where(finalPred).
		Count(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "count related logs failed")
	}

	list := make([]*LogEntry, 0, len(allLogs))
	for _, l := range allLogs {
		list = append(list, &LogEntry{
			ID:          l.ID,
			API:         l.API,
			Method:      l.Method,
			Success:     l.Success,
			ReqContent:  l.ReqContent,
			RespContent: l.RespContent,
			IP:          l.IP,
			UserAgent:   l.UserAgent,
			Operatorsr:  l.Operatorsr,
			Time:        l.Time,
			CreatedAt:   l.CreatedAt.Format(time.DateTime),
		})
	}

	return &LogSearchResp{List: list, Total: count}, nil
}

func isIP(s string) bool {
	for _, c := range s {
		if c != '.' && c != ':' && (c < '0' || c > '9') {
			return false
		}
	}
	return len(s) >= 7
}

func NewInterviewer(ctx context.Context, c *app.RequestContext) *Interviewer {
	return &Interviewer{
		ctx: ctx,
		c:   c,
		db:  db.DB,
	}
}
