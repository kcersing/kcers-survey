package survey

import (
	"context"
	"math"

	db "kcers-survey/biz/dal/db"
	"kcers-survey/biz/dal/db/ent"
	surveyquestion2 "kcers-survey/biz/dal/db/ent/surveyquestion"
	surveyresponse2 "kcers-survey/biz/dal/db/ent/surveyresponse"
	surveyresponseanswers2 "kcers-survey/biz/dal/db/ent/surveyresponseanswers"
	service2 "kcers-survey/biz/infras/service"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"testing"
)

func answerCount9(sq *ent.SurveyQuestion, ids []int64, db *ent.Client, ctx context.Context) (resp *StatisticsBasic) {
	// 1. 查询所有答案
	answers, err := db.SurveyResponseAnswers.Query().
		Where(
			surveyresponseanswers2.SurveyQuestionID(sq.ID),
			surveyresponseanswers2.Delete(0),
			surveyresponseanswers2.SurveyResponseIDIn(ids...),
		).
		WithResponse(func(q *ent.SurveyResponseQuery) {
			q.Select(surveyresponse2.FieldArea)
		}).
		Order(ent.Asc(surveyresponseanswers2.FieldSurveyResponseID)).
		Order(ent.Asc(surveyresponseanswers2.FieldCreatedAt)). // 按时间排序
		All(ctx)
	if err != nil {
		hlog.Error(err)
		return nil
	}

	if len(answers) == 0 {
		return nil
	}

	// 2. 去重：同一份答卷对同一道题只取最新的一条
	latestAnswers := make(map[int64]*ent.SurveyResponseAnswers)
	for _, ans := range answers {
		key := ans.SurveyResponseID
		existing, exists := latestAnswers[key]
		if !exists || ans.CreatedAt.After(existing.CreatedAt) {
			latestAnswers[key] = ans
		}
	}

	// 3. 使用去重后的数据
	uniqueAnswers := make([]*ent.SurveyResponseAnswers, 0, len(latestAnswers))
	for _, ans := range latestAnswers {
		uniqueAnswers = append(uniqueAnswers, ans)
	}

	// 4. 统计回答总数（按答卷去重）
	totalAnswers := int64(len(uniqueAnswers))

	// 5. 初始化选项统计
	optionMap := make(map[string]*Basic)
	for _, o := range sq.Options {
		optionMap[o.Content] = &Basic{
			Name:       o.Content,
			Value:      0,
			RegionData: make(map[string]int),
		}
	}

	// 6. 统计选项频次（含地区）
	var allOptions []string
	for _, ans := range uniqueAnswers {
		if len(ans.Answer) > 0 {
			for _, an := range ans.Answer {
				allOptions = append(allOptions, an)
				if stat, exists := optionMap[an]; exists {
					stat.Value++
					if ans.Edges.Response != nil && ans.Edges.Response.Area != "" {
						stat.RegionData[ans.Edges.Response.Area]++
					}
				}
			}
		}
	}

	// 6.1 地区(省)维度分母：各省回答本题的答卷数(已按答卷去重)、各省选项总数(选项个数)
	regionRespCount := make(map[string]int64)
	regionOptionCount := make(map[string]int64)
	for _, ans := range uniqueAnswers {
		if ans.Edges.Response == nil || ans.Edges.Response.Area == "" {
			continue
		}
		area := ans.Edges.Response.Area
		regionRespCount[area]++
		regionOptionCount[area] += int64(len(ans.Answer))
	}

	// 7. 构建结果
	resp = &StatisticsBasic{
		Count:               totalAnswers,
		Count2:              len(allOptions),
		Name:                SName(sq.ParentID, db) + sq.Content,
		RegionResponseCount: regionRespCount,
		RegionOptionCount:   regionOptionCount,
	}

	for _, stat := range optionMap {
		resp.Data = append(resp.Data, *stat)
	}

	// 8. 计算占比
	for i, v := range resp.Data {
		resp.Data[i].Proportion = math.Round(float64(v.Value)/float64(resp.Count)*10000) / 100
		resp.Data[i].Proportion2 = math.Round(float64(v.Value)/float64(resp.Count2)*10000) / 100
	}

	return
}

// regionPercent 计算某省某选项的百分比:省频次 / 省回答总数,保留 2 位小数(分母为 0 时返回 0)
func regionPercent(part, total int64) float64 {
	if total == 0 {
		return 0
	}
	return math.Round(float64(part)/float64(total)*10000) / 100
}

// TestSurvey3 导出「问题 × 省份」交叉统计报表,一次运行导出 6 种范围:
// 问卷 1/2/3 × 全量 / 已完成。
// 导出范围定义(surveyScopes)与答卷筛选条件(surveyResponsePredicates)在 survey5_test.go 中,同包共用。
func TestSurvey3(t *testing.T) {
	dbs := db.InItDB("user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey port=5432 sslmode=disable TimeZone=Asia/Shanghai", true)

	ctx := context.Background()

	for _, sc := range surveyScopes {
		if err := exportSurveyCross(dbs, ctx, sc); err != nil {
			t.Errorf("%s 导出失败: %v", sc.Name, err)
		}
	}
}

// exportSurveyCross 导出某份问卷的「问题 × 省份」交叉统计报表
func exportSurveyCross(dbs *ent.Client, ctx context.Context, sc surveyScope) error {
	// 1. 获取符合条件的答卷ID
	sr, err := dbs.SurveyResponse.Query().
		Where(surveyResponsePredicates(sc)...).
		Order(ent.Asc(surveyresponse2.FieldID)).
		IDs(ctx)
	if err != nil {
		return err
	}

	if len(sr) == 0 {
		hlog.Warnf("%s: 没有符合条件的答卷", sc.Name)
		return nil
	}

	// 2. 获取所有题目
	sqarr, err := dbs.SurveyQuestion.Query().
		Where(
			surveyquestion2.SurveyID(sc.SurveyID),
			surveyquestion2.Delete(0),
			surveyquestion2.TypeIn("single_choice", "multiple_choice"),
		).
		Order(ent.Asc(surveyquestion2.FieldID, surveyquestion2.FieldParentID, surveyquestion2.FieldSort)).
		All(ctx)
	if err != nil {
		return err
	}

	// 3. 获取所有地区列表（用于动态列）
	regions, err := dbs.SurveyResponse.Query().
		Where(
			surveyresponse2.IDIn(sr...),
			surveyresponse2.Delete(0),
			surveyresponse2.AreaNEQ(""),
		).
		GroupBy(surveyresponse2.FieldArea).
		Strings(ctx)
	if err != nil {
		hlog.Warnf("获取地区列表失败: %v", err)
		regions = []string{}
	}

	// 4. 构建表头：全局 4 列 + 每个省 4 列(频次 / 回答总数 / 选项总数 / 百分比)
	tale := []interface{}{"问题", "频次", "回答总数", "选项总数"}
	for _, region := range regions {
		tale = append(tale,
			region+"-频次",
			region+"-回答总数",
			region+"-选项总数",
			region+"-百分比",
		)
	}

	// 5. 构建数据行
	var list []map[int]interface{}

	for _, sq := range sqarr {
		// 统计当前题目的数据
		resp := answerCount9(sq, sr, dbs, ctx)
		if resp == nil {
			continue
		}

		// 添加题目标题行
		li := map[int]interface{}{}
		li[1] = resp.Name
		li[2] = "频次"
		li[3] = "合计"
		li[4] = "合计2"
		// 地区列留空
		for idx := 5; idx <= len(tale); idx++ {
			li[idx] = ""
		}
		list = append(list, li)

		// 添加选项行
		for _, v := range resp.Data {
			li := map[int]interface{}{}
			li[1] = v.Name      // 选项名称
			li[2] = v.Value     // 频次
			li[3] = resp.Count  // 回答总数
			li[4] = resp.Count2 // 选项总数

			// 各省数据：每个省 4 列(频次 / 回答总数 / 选项总数 / 百分比)
			colIdx := 5
			for _, region := range regions {
				regionCount := resp.RegionResponseCount[region]
				li[colIdx] = v.RegionData[region]                                      // 省-频次
				li[colIdx+1] = regionCount                                             // 省-回答总数
				li[colIdx+2] = resp.RegionOptionCount[region]                          // 省-选项总数
				li[colIdx+3] = regionPercent(int64(v.RegionData[region]), regionCount) // 省-百分比
				colIdx += 4
			}
			list = append(list, li)
		}

		// 添加空行分隔
		list = append(list, map[int]interface{}{
			1: "",
		})
	}

	// 6. 导出（使用您原有的导出方法）
	domain, err := service2.Export(tale, list, sc.Name)
	if err != nil {
		return err
	}
	hlog.Infof("%s 导出成功: %s", sc.Name, domain)
	return nil
}
