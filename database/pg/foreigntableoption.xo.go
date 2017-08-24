// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// ForeignTableOptionTable is the database name for the table.
const ForeignTableOptionTable = "information_schema.foreign_table_options"

// ForeignTableOption represents a row from 'information_schema.foreign_table_options'.
type ForeignTableOption struct {
	ForeignTableCatalog SQLIdentifier `yaml:"foreign_table_catalog,omitempty"` // foreign_table_catalog
	ForeignTableSchema  SQLIdentifier `yaml:"foreign_table_schema,omitempty"`  // foreign_table_schema
	ForeignTableName    SQLIdentifier `yaml:"foreign_table_name,omitempty"`    // foreign_table_name
	OptionName          SQLIdentifier `yaml:"option_name,omitempty"`           // option_name
	OptionValue         CharacterData `yaml:"option_value,omitempty"`          // option_value
}

// Constants defining each column in the table.
const (
	ForeignTableOptionForeignTableCatalogField = "foreign_table_catalog"
	ForeignTableOptionForeignTableSchemaField  = "foreign_table_schema"
	ForeignTableOptionForeignTableNameField    = "foreign_table_name"
	ForeignTableOptionOptionNameField          = "option_name"
	ForeignTableOptionOptionValueField         = "option_value"
)

// WhereClauses for every type in ForeignTableOption.
var (
	ForeignTableOptionForeignTableCatalogWhere SQLIdentifierField = "foreign_table_catalog"
	ForeignTableOptionForeignTableSchemaWhere  SQLIdentifierField = "foreign_table_schema"
	ForeignTableOptionForeignTableNameWhere    SQLIdentifierField = "foreign_table_name"
	ForeignTableOptionOptionNameWhere          SQLIdentifierField = "option_name"
	ForeignTableOptionOptionValueWhere         CharacterDataField = "option_value"
)

// QueryOneForeignTableOption retrieves a row from 'information_schema.foreign_table_options' as a ForeignTableOption.
func QueryOneForeignTableOption(db XODB, where WhereClause, order OrderBy) (*ForeignTableOption, error) {
	const origsqlstr = `SELECT ` +
		`foreign_table_catalog, foreign_table_schema, foreign_table_name, option_name, option_value ` +
		`FROM information_schema.foreign_table_options WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	fto := &ForeignTableOption{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&fto.ForeignTableCatalog, &fto.ForeignTableSchema, &fto.ForeignTableName, &fto.OptionName, &fto.OptionValue)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return fto, nil
}

// QueryForeignTableOption retrieves rows from 'information_schema.foreign_table_options' as a slice of ForeignTableOption.
func QueryForeignTableOption(db XODB, where WhereClause, order OrderBy) ([]*ForeignTableOption, error) {
	const origsqlstr = `SELECT ` +
		`foreign_table_catalog, foreign_table_schema, foreign_table_name, option_name, option_value ` +
		`FROM information_schema.foreign_table_options WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ForeignTableOption
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		fto := ForeignTableOption{}

		err = q.Scan(&fto.ForeignTableCatalog, &fto.ForeignTableSchema, &fto.ForeignTableName, &fto.OptionName, &fto.OptionValue)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &fto)
	}
	return vals, nil
}

// AllForeignTableOption retrieves all rows from 'information_schema.foreign_table_options' as a slice of ForeignTableOption.
func AllForeignTableOption(db XODB, order OrderBy) ([]*ForeignTableOption, error) {
	const origsqlstr = `SELECT ` +
		`foreign_table_catalog, foreign_table_schema, foreign_table_name, option_name, option_value ` +
		`FROM information_schema.foreign_table_options`

	sqlstr := origsqlstr + order.String()

	var vals []*ForeignTableOption
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		fto := ForeignTableOption{}

		err = q.Scan(&fto.ForeignTableCatalog, &fto.ForeignTableSchema, &fto.ForeignTableName, &fto.OptionName, &fto.OptionValue)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &fto)
	}
	return vals, nil
}
