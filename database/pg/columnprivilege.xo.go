// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// ColumnPrivilegeTable is the database name for the table.
const ColumnPrivilegeTable = "information_schema.column_privileges"

// ColumnPrivilege represents a row from 'information_schema.column_privileges'.
type ColumnPrivilege struct {
	Grantor       SQLIdentifier `yaml:"grantor,omitempty"`        // grantor
	Grantee       SQLIdentifier `yaml:"grantee,omitempty"`        // grantee
	TableCatalog  SQLIdentifier `yaml:"table_catalog,omitempty"`  // table_catalog
	TableSchema   SQLIdentifier `yaml:"table_schema,omitempty"`   // table_schema
	TableName     SQLIdentifier `yaml:"table_name,omitempty"`     // table_name
	ColumnName    SQLIdentifier `yaml:"column_name,omitempty"`    // column_name
	PrivilegeType CharacterData `yaml:"privilege_type,omitempty"` // privilege_type
	IsGrantable   YesOrNo       `yaml:"is_grantable,omitempty"`   // is_grantable
}

// Constants defining each column in the table.
const (
	ColumnPrivilegeGrantorField       = "grantor"
	ColumnPrivilegeGranteeField       = "grantee"
	ColumnPrivilegeTableCatalogField  = "table_catalog"
	ColumnPrivilegeTableSchemaField   = "table_schema"
	ColumnPrivilegeTableNameField     = "table_name"
	ColumnPrivilegeColumnNameField    = "column_name"
	ColumnPrivilegePrivilegeTypeField = "privilege_type"
	ColumnPrivilegeIsGrantableField   = "is_grantable"
)

// WhereClauses for every type in ColumnPrivilege.
var (
	ColumnPrivilegeGrantorWhere       SQLIdentifierField = "grantor"
	ColumnPrivilegeGranteeWhere       SQLIdentifierField = "grantee"
	ColumnPrivilegeTableCatalogWhere  SQLIdentifierField = "table_catalog"
	ColumnPrivilegeTableSchemaWhere   SQLIdentifierField = "table_schema"
	ColumnPrivilegeTableNameWhere     SQLIdentifierField = "table_name"
	ColumnPrivilegeColumnNameWhere    SQLIdentifierField = "column_name"
	ColumnPrivilegePrivilegeTypeWhere CharacterDataField = "privilege_type"
	ColumnPrivilegeIsGrantableWhere   YesOrNoField       = "is_grantable"
)

// QueryOneColumnPrivilege retrieves a row from 'information_schema.column_privileges' as a ColumnPrivilege.
func QueryOneColumnPrivilege(db XODB, where WhereClause, order OrderBy) (*ColumnPrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.column_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	cp := &ColumnPrivilege{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&cp.Grantor, &cp.Grantee, &cp.TableCatalog, &cp.TableSchema, &cp.TableName, &cp.ColumnName, &cp.PrivilegeType, &cp.IsGrantable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cp, nil
}

// QueryColumnPrivilege retrieves rows from 'information_schema.column_privileges' as a slice of ColumnPrivilege.
func QueryColumnPrivilege(db XODB, where WhereClause, order OrderBy) ([]*ColumnPrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.column_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ColumnPrivilege
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		cp := ColumnPrivilege{}

		err = q.Scan(&cp.Grantor, &cp.Grantee, &cp.TableCatalog, &cp.TableSchema, &cp.TableName, &cp.ColumnName, &cp.PrivilegeType, &cp.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &cp)
	}
	return vals, nil
}

// AllColumnPrivilege retrieves all rows from 'information_schema.column_privileges' as a slice of ColumnPrivilege.
func AllColumnPrivilege(db XODB, order OrderBy) ([]*ColumnPrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.column_privileges`

	sqlstr := origsqlstr + order.String()

	var vals []*ColumnPrivilege
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		cp := ColumnPrivilege{}

		err = q.Scan(&cp.Grantor, &cp.Grantee, &cp.TableCatalog, &cp.TableSchema, &cp.TableName, &cp.ColumnName, &cp.PrivilegeType, &cp.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &cp)
	}
	return vals, nil
}
