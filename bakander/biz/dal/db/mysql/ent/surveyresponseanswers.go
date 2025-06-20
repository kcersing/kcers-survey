// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SurveyResponseAnswers is the model entity for the SurveyResponseAnswers schema.
type SurveyResponseAnswers struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// last delete  1:已删除
	Delete int64 `json:"delete,omitempty"`
	// created
	CreatedID int64 `json:"created_id,omitempty"`
	// 状态[0:禁用;1:正常]
	Status int64 `json:"status,omitempty"`
	// survey_id
	SurveyID int64 `json:"survey_id,omitempty"`
	// survey_response_id
	SurveyResponseID int64 `json:"survey_response_id,omitempty"`
	// survey_question_id
	SurveyQuestionID int64 `json:"survey_question_id,omitempty"`
	// 回答文本
	AnswerText string `json:"answer_text,omitempty"`
	// 回答数值
	AnswerValue  int64 `json:"answer_value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SurveyResponseAnswers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case surveyresponseanswers.FieldID, surveyresponseanswers.FieldDelete, surveyresponseanswers.FieldCreatedID, surveyresponseanswers.FieldStatus, surveyresponseanswers.FieldSurveyID, surveyresponseanswers.FieldSurveyResponseID, surveyresponseanswers.FieldSurveyQuestionID, surveyresponseanswers.FieldAnswerValue:
			values[i] = new(sql.NullInt64)
		case surveyresponseanswers.FieldAnswerText:
			values[i] = new(sql.NullString)
		case surveyresponseanswers.FieldCreatedAt, surveyresponseanswers.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SurveyResponseAnswers fields.
func (sra *SurveyResponseAnswers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case surveyresponseanswers.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sra.ID = int64(value.Int64)
		case surveyresponseanswers.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sra.CreatedAt = value.Time
			}
		case surveyresponseanswers.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sra.UpdatedAt = value.Time
			}
		case surveyresponseanswers.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				sra.Delete = value.Int64
			}
		case surveyresponseanswers.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				sra.CreatedID = value.Int64
			}
		case surveyresponseanswers.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sra.Status = value.Int64
			}
		case surveyresponseanswers.FieldSurveyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field survey_id", values[i])
			} else if value.Valid {
				sra.SurveyID = value.Int64
			}
		case surveyresponseanswers.FieldSurveyResponseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field survey_response_id", values[i])
			} else if value.Valid {
				sra.SurveyResponseID = value.Int64
			}
		case surveyresponseanswers.FieldSurveyQuestionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field survey_question_id", values[i])
			} else if value.Valid {
				sra.SurveyQuestionID = value.Int64
			}
		case surveyresponseanswers.FieldAnswerText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer_text", values[i])
			} else if value.Valid {
				sra.AnswerText = value.String
			}
		case surveyresponseanswers.FieldAnswerValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field answer_value", values[i])
			} else if value.Valid {
				sra.AnswerValue = value.Int64
			}
		default:
			sra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SurveyResponseAnswers.
// This includes values selected through modifiers, order, etc.
func (sra *SurveyResponseAnswers) Value(name string) (ent.Value, error) {
	return sra.selectValues.Get(name)
}

// Update returns a builder for updating this SurveyResponseAnswers.
// Note that you need to call SurveyResponseAnswers.Unwrap() before calling this method if this SurveyResponseAnswers
// was returned from a transaction, and the transaction was committed or rolled back.
func (sra *SurveyResponseAnswers) Update() *SurveyResponseAnswersUpdateOne {
	return NewSurveyResponseAnswersClient(sra.config).UpdateOne(sra)
}

// Unwrap unwraps the SurveyResponseAnswers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sra *SurveyResponseAnswers) Unwrap() *SurveyResponseAnswers {
	_tx, ok := sra.config.driver.(*txDriver)
	if !ok {
		panic("ent: SurveyResponseAnswers is not a transactional entity")
	}
	sra.config.driver = _tx.drv
	return sra
}

// String implements the fmt.Stringer.
func (sra *SurveyResponseAnswers) String() string {
	var builder strings.Builder
	builder.WriteString("SurveyResponseAnswers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sra.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sra.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sra.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", sra.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", sra.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sra.Status))
	builder.WriteString(", ")
	builder.WriteString("survey_id=")
	builder.WriteString(fmt.Sprintf("%v", sra.SurveyID))
	builder.WriteString(", ")
	builder.WriteString("survey_response_id=")
	builder.WriteString(fmt.Sprintf("%v", sra.SurveyResponseID))
	builder.WriteString(", ")
	builder.WriteString("survey_question_id=")
	builder.WriteString(fmt.Sprintf("%v", sra.SurveyQuestionID))
	builder.WriteString(", ")
	builder.WriteString("answer_text=")
	builder.WriteString(sra.AnswerText)
	builder.WriteString(", ")
	builder.WriteString("answer_value=")
	builder.WriteString(fmt.Sprintf("%v", sra.AnswerValue))
	builder.WriteByte(')')
	return builder.String()
}

// SurveyResponseAnswersSlice is a parsable slice of SurveyResponseAnswers.
type SurveyResponseAnswersSlice []*SurveyResponseAnswers
