// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/dictionary"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Dictionary is the model entity for the Dictionary schema.
type Dictionary struct {
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
	// the title shown in the ui | 展示名称 （建议配合i18n）
	Title string `json:"title,omitempty"`
	// the name of dictionary for search | 字典搜索名称
	Name string `json:"name,omitempty"`
	// the description of dictionary | 字典描述
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DictionaryQuery when eager-loading is set.
	Edges        DictionaryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DictionaryEdges holds the relations/edges for other nodes in the graph.
type DictionaryEdges struct {
	// DictionaryDetails holds the value of the dictionary_details edge.
	DictionaryDetails []*DictionaryDetail `json:"dictionary_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DictionaryDetailsOrErr returns the DictionaryDetails value or an error if the edge
// was not loaded in eager-loading.
func (e DictionaryEdges) DictionaryDetailsOrErr() ([]*DictionaryDetail, error) {
	if e.loadedTypes[0] {
		return e.DictionaryDetails, nil
	}
	return nil, &NotLoadedError{edge: "dictionary_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dictionary) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dictionary.FieldID, dictionary.FieldDelete, dictionary.FieldCreatedID, dictionary.FieldStatus:
			values[i] = new(sql.NullInt64)
		case dictionary.FieldTitle, dictionary.FieldName, dictionary.FieldDescription:
			values[i] = new(sql.NullString)
		case dictionary.FieldCreatedAt, dictionary.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dictionary fields.
func (d *Dictionary) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dictionary.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int64(value.Int64)
		case dictionary.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case dictionary.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case dictionary.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				d.Delete = value.Int64
			}
		case dictionary.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				d.CreatedID = value.Int64
			}
		case dictionary.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = value.Int64
			}
		case dictionary.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				d.Title = value.String
			}
		case dictionary.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case dictionary.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dictionary.
// This includes values selected through modifiers, order, etc.
func (d *Dictionary) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryDictionaryDetails queries the "dictionary_details" edge of the Dictionary entity.
func (d *Dictionary) QueryDictionaryDetails() *DictionaryDetailQuery {
	return NewDictionaryClient(d.config).QueryDictionaryDetails(d)
}

// Update returns a builder for updating this Dictionary.
// Note that you need to call Dictionary.Unwrap() before calling this method if this Dictionary
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dictionary) Update() *DictionaryUpdateOne {
	return NewDictionaryClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dictionary entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dictionary) Unwrap() *Dictionary {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dictionary is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dictionary) String() string {
	var builder strings.Builder
	builder.WriteString("Dictionary(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", d.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(d.Title)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Dictionaries is a parsable slice of Dictionary.
type Dictionaries []*Dictionary
