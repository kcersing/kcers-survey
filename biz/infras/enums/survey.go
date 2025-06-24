package enums

type QuestionType string

const (
	QuestionTypeRating   QuestionType = "rating"          // 评分
	QuestionTypeSingle   QuestionType = "single_choice"   //单选
	QuestionTypeMultiple QuestionType = "multiple_choice" //多选
	QuestionTypeText     QuestionType = "text"            // 文本
	QuestionTypeNumber   QuestionType = "number"          // 数字
	QuestionTypeImage    QuestionType = "image"           // 图片
	QuestionTypeVideo    QuestionType = "video"           // 视频
	QuestionTypeAudio    QuestionType = "audio"           // 音频
	QuestionTypeFile     QuestionType = "file"            // 文件
	QuestionTypeDate     QuestionType = "date"            // 日期
	QuestionTypeDateTime QuestionType = "datetime"        // 日期时间
	QuestionTypeH2       QuestionType = "h2"              // 标题
	QuestionTypePage     QuestionType = "page"            // 单页

)

type JumpRulesOperators string

const (
	JumpRulesOperatorsEquals      JumpRulesOperators = "equals"
	JumpRulesOperatorsNotEquals   JumpRulesOperators = "notEquals"
	JumpRulesOperatorsContains    JumpRulesOperators = "contains"
	JumpRulesOperatorsNotContains JumpRulesOperators = "notContains"
)
