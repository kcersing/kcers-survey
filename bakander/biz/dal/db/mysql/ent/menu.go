// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/menu"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
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
	// parent menu ID | 父菜单ID
	ParentID int64 `json:"parent_id,omitempty"`
	// index path | 菜单路由路径
	Path string `json:"path,omitempty"`
	// index name | 菜单名称
	Name string `json:"name,omitempty"`
	// sorting numbers | 排序编号
	OrderNo int64 `json:"order_no,omitempty"`
	// disable status | 是否停用
	Disabled int64 `json:"disabled,omitempty"`
	// 当前路由是否渲染菜单项，为 true 的话不会在菜单中显示，但可通过路由地址访问
	Ignore bool `json:"ignore,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Menu `json:"children,omitempty"`
	// Params holds the value of the params edge.
	Params []*MenuParam `json:"params,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ChildrenOrErr() ([]*Menu, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParamsOrErr returns the Params value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ParamsOrErr() ([]*MenuParam, error) {
	if e.loadedTypes[3] {
		return e.Params, nil
	}
	return nil, &NotLoadedError{edge: "params"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldIgnore:
			values[i] = new(sql.NullBool)
		case menu.FieldID, menu.FieldDelete, menu.FieldCreatedID, menu.FieldParentID, menu.FieldOrderNo, menu.FieldDisabled:
			values[i] = new(sql.NullInt64)
		case menu.FieldPath, menu.FieldName:
			values[i] = new(sql.NullString)
		case menu.FieldCreatedAt, menu.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case menu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case menu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case menu.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				m.Delete = value.Int64
			}
		case menu.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				m.CreatedID = value.Int64
			}
		case menu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				m.ParentID = value.Int64
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				m.OrderNo = value.Int64
			}
		case menu.FieldDisabled:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				m.Disabled = value.Int64
			}
		case menu.FieldIgnore:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ignore", values[i])
			} else if value.Valid {
				m.Ignore = value.Bool
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Menu entity.
func (m *Menu) QueryRoles() *RoleQuery {
	return NewMenuClient(m.config).QueryRoles(m)
}

// QueryParent queries the "parent" edge of the Menu entity.
func (m *Menu) QueryParent() *MenuQuery {
	return NewMenuClient(m.config).QueryParent(m)
}

// QueryChildren queries the "children" edge of the Menu entity.
func (m *Menu) QueryChildren() *MenuQuery {
	return NewMenuClient(m.config).QueryChildren(m)
}

// QueryParams queries the "params" edge of the Menu entity.
func (m *Menu) QueryParams() *MenuParamQuery {
	return NewMenuClient(m.config).QueryParams(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", m.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", m.ParentID))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(m.Path)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(fmt.Sprintf("%v", m.OrderNo))
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", m.Disabled))
	builder.WriteString(", ")
	builder.WriteString("ignore=")
	builder.WriteString(fmt.Sprintf("%v", m.Ignore))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
