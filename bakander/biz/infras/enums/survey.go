package enums

type QuestionType string

const (
	QuestionTypeRating   QuestionType = "rating"   // 评分
	QuestionTypeSingle   QuestionType = "single"   //单选
	QuestionTypeMultiple QuestionType = "multiple" //多选
	QuestionTypeText     QuestionType = "text"     // 文本
	QuestionTypeImage    QuestionType = "image"    // 图片
	QuestionTypeVideo    QuestionType = "video"    // 视频
	QuestionTypeAudio    QuestionType = "audio"    // 音频
	QuestionTypeFile     QuestionType = "file"     // 文件
	QuestionTypeDateTime QuestionType = "datetime" // 日期时间
)
