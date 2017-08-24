// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// KeyColumnUsageTable is the database name for the table.
const KeyColumnUsageTable = "information_schema.key_column_usage"

// KeyColumnUsage represents a row from 'information_schema.key_column_usage'.
type KeyColumnUsage struct {
	ConstraintCatalog          SQLIdentifier  `yaml:"constraint_catalog,omitempty"`            // constraint_catalog
	ConstraintSchema           SQLIdentifier  `yaml:"constraint_schema,omitempty"`             // constraint_schema
	ConstraintName             SQLIdentifier  `yaml:"constraint_name,omitempty"`               // constraint_name
	TableCatalog               SQLIdentifier  `yaml:"table_catalog,omitempty"`                 // table_catalog
	TableSchema                SQLIdentifier  `yaml:"table_schema,omitempty"`                  // table_schema
	TableName                  SQLIdentifier  `yaml:"table_name,omitempty"`                    // table_name
	ColumnName                 SQLIdentifier  `yaml:"column_name,omitempty"`                   // column_name
	OrdinalPosition            CardinalNumber `yaml:"ordinal_position,omitempty"`              // ordinal_position
	PositionInUniqueConstraint CardinalNumber `yaml:"position_in_unique_constraint,omitempty"` // position_in_unique_constraint
}

// Constants defining each column in the table.
const (
	KeyColumnUsageConstraintCatalogField          = "constraint_catalog"
	KeyColumnUsageConstraintSchemaField           = "constraint_schema"
	KeyColumnUsageConstraintNameField             = "constraint_name"
	KeyColumnUsageTableCatalogField               = "table_catalog"
	KeyColumnUsageTableSchemaField                = "table_schema"
	KeyColumnUsageTableNameField                  = "table_name"
	KeyColumnUsageColumnNameField                 = "column_name"
	KeyColumnUsageOrdinalPositionField            = "ordinal_position"
	KeyColumnUsagePositionInUniqueConstraintField = "position_in_unique_constraint"
)

// WhereClauses for every type in KeyColumnUsage.
var (
	KeyColumnUsageConstraintCatalogWhere          SQLIdentifierField  = "constraint_catalog"
	KeyColumnUsageConstraintSchemaWhere           SQLIdentifierField  = "constraint_schema"
	KeyColumnUsageConstraintNameWhere             SQLIdentifierField  = "constraint_name"
	KeyColumnUsageTableCatalogWhere               SQLIdentifierField  = "table_catalog"
	KeyColumnUsageTableSchemaWhere                SQLIdentifierField  = "table_schema"
	KeyColumnUsageTableNameWhere                  SQLIdentifierField  = "table_name"
	KeyColumnUsageColumnNameWhere                 SQLIdentifierField  = "column_name"
	KeyColumnUsageOrdinalPositionWhere            CardinalNumberField = "ordinal_position"
	KeyColumnUsagePositionInUniqueConstraintWhere CardinalNumberField = "position_in_unique_constraint"
)

// QueryOneKeyColumnUsage retrieves a row from 'information_schema.key_column_usage' as a KeyColumnUsage.
func QueryOneKeyColumnUsage(db XODB, where WhereClause, order OrderBy) (*KeyColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, column_name, ordinal_position, position_in_unique_constraint ` +
		`FROM information_schema.key_column_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	kcu := &KeyColumnUsage{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&kcu.ConstraintCatalog, &kcu.ConstraintSchema, &kcu.ConstraintName, &kcu.TableCatalog, &kcu.TableSchema, &kcu.TableName, &kcu.ColumnName, &kcu.OrdinalPosition, &kcu.PositionInUniqueConstraint)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return kcu, nil
}

// QueryKeyColumnUsage retrieves rows from 'information_schema.key_column_usage' as a slice of KeyColumnUsage.
func QueryKeyColumnUsage(db XODB, where WhereClause, order OrderBy) ([]*KeyColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, column_name, ordinal_position, position_in_unique_constraint ` +
		`FROM information_schema.key_column_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*KeyColumnUsage
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		kcu := KeyColumnUsage{}

		err = q.Scan(&kcu.ConstraintCatalog, &kcu.ConstraintSchema, &kcu.ConstraintName, &kcu.TableCatalog, &kcu.TableSchema, &kcu.TableName, &kcu.ColumnName, &kcu.OrdinalPosition, &kcu.PositionInUniqueConstraint)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &kcu)
	}
	return vals, nil
}

// AllKeyColumnUsage retrieves all rows from 'information_schema.key_column_usage' as a slice of KeyColumnUsage.
func AllKeyColumnUsage(db XODB, order OrderBy) ([]*KeyColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, column_name, ordinal_position, position_in_unique_constraint ` +
		`FROM information_schema.key_column_usage`

	sqlstr := origsqlstr + order.String()

	var vals []*KeyColumnUsage
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		kcu := KeyColumnUsage{}

		err = q.Scan(&kcu.ConstraintCatalog, &kcu.ConstraintSchema, &kcu.ConstraintName, &kcu.TableCatalog, &kcu.TableSchema, &kcu.TableName, &kcu.ColumnName, &kcu.OrdinalPosition, &kcu.PositionInUniqueConstraint)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &kcu)
	}
	return vals, nil
}
