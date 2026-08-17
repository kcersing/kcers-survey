package pub

import (
	"context"
	"strconv"
	"time"

	"kcers-survey/biz/infras/service/interviewer"
	surveyService "kcers-survey/biz/infras/service/survey"
	"kcers-survey/biz/mw"
	"kcers-survey/biz/pkg/errno"
	"kcers-survey/biz/pkg/utils"
	pub "kcers-survey/idl_gen/model/pub"

	"github.com/cloudwego/hertz/pkg/app"
)

// QueryByPhone .
// @router /service/pub/query-by-phone [POST]
func QueryByPhone(ctx context.Context, c *app.RequestContext) {
	var req pub.QueryByPhoneReq
	err := c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	list, total, err := surveyService.NewSurvey(ctx, c).QueryByPhone(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
}

// InterviewerLogin .
// @router /service/pub/interviewer/login [POST]
func InterviewerLogin(ctx context.Context, c *app.RequestContext) {
	var req pub.InterviewerLoginReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	s := interviewer.NewInterviewer(ctx, c)
	resp, err := s.Login(req.Mobile, req.Password)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	if resp.NeedSetup {
		utils.SendResponse(c, errno.Success, map[string]interface{}{
			"needSetup": true,
			"mobile":    resp.Mobile,
			"name":      resp.Name,
		}, 0, "")
		return
	}

	roleIDs, _ := s.GetRoleIDs(resp.UserId)
	roleValues, _ := s.GetRoleValues(resp.UserId)
	payloadMap := map[string]interface{}{
		"userId":     strconv.Itoa(int(resp.UserId)),
		"mobile":     resp.Mobile,
		"roleIds":    roleIDs,
		"roleValues": roleValues,
	}
	token, expire, err := mw.GetInterviewerJWTMw().TokenGenerator(payloadMap)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, map[string]interface{}{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
		"name":   resp.Name,
		"mobile": resp.Mobile,
	}, 0, "")
}

// InterviewerSendSms .
// @router /service/pub/interviewer/send-sms [POST]
func InterviewerSendSms(ctx context.Context, c *app.RequestContext) {
	var req pub.InterviewerSendSmsReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	s := interviewer.NewInterviewer(ctx, c)
	if s.GetResearcherName(req.Mobile) == "" {
		utils.SendResponse(c, errno.NewErrNo(10003, "该手机号未关联任何调查问卷"), nil, 0, "")
		return
	}

	if err := s.SendSmsCaptcha(req.Mobile); err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
}

// InterviewerSetupPassword .
// @router /service/pub/interviewer/setup-password [POST]
func InterviewerSetupPassword(ctx context.Context, c *app.RequestContext) {
	var req pub.InterviewerSetupReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	if req.Password == "" || len(req.Password) < 6 {
		utils.SendResponse(c, errno.NewErrNo(10003, "密码至少6位"), nil, 0, "")
		return
	}

	s := interviewer.NewInterviewer(ctx, c)
	if !s.VerifySmsCaptcha(req.Mobile, req.SmsCode) {
		utils.SendResponse(c, errno.NewErrNo(10003, "验证码错误"), nil, 0, "")
		return
	}

	resp, err := s.SetupPassword(req.Mobile, req.Name, req.Password)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	payloadMap := map[string]interface{}{
		"userId":     strconv.Itoa(int(resp.UserId)),
		"mobile":     resp.Mobile,
		"roleIds":    resp.RoleIDs,
		"roleValues": resp.RoleValues,
	}
	token, expire, err := mw.GetInterviewerJWTMw().TokenGenerator(payloadMap)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, map[string]interface{}{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
		"name":   resp.Name,
		"mobile": resp.Mobile,
	}, 0, "")
}
