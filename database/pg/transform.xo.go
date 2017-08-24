// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// TransformTable is the database name for the table.
const TransformTable = "information_schema.transforms"

// Transform represents a row from 'information_schema.transforms'.
type Transform struct {
	UdtCatalog      SQLIdentifier `yaml:"udt_catalog,omitempty"`      // udt_catalog
	UdtSchema       SQLIdentifier `yaml:"udt_schema,omitempty"`       // udt_schema
	UdtName         SQLIdentifier `yaml:"udt_name,omitempty"`         // udt_name
	SpecificCatalog SQLIdentifier `yaml:"specific_catalog,omitempty"` // specific_catalog
	SpecificSchema  SQLIdentifier `yaml:"specific_schema,omitempty"`  // specific_schema
	SpecificName    SQLIdentifier `yaml:"specific_name,omitempty"`    // specific_name
	GroupName       SQLIdentifier `yaml:"group_name,omitempty"`       // group_name
	TransformType   CharacterData `yaml:"transform_type,omitempty"`   // transform_type
}

// Constants defining each column in the table.
const (
	TransformUdtCatalogField      = "udt_catalog"
	TransformUdtSchemaField       = "udt_schema"
	TransformUdtNameField         = "udt_name"
	TransformSpecificCatalogField = "specific_catalog"
	TransformSpecificSchemaField  = "specific_schema"
	TransformSpecificNameField    = "specific_name"
	TransformGroupNameField       = "group_name"
	TransformTransformTypeField   = "transform_type"
)

// WhereClauses for every type in Transform.
var (
	TransformUdtCatalogWhere      SQLIdentifierField = "udt_catalog"
	TransformUdtSchemaWhere       SQLIdentifierField = "udt_schema"
	TransformUdtNameWhere         SQLIdentifierField = "udt_name"
	TransformSpecificCatalogWhere SQLIdentifierField = "specific_catalog"
	TransformSpecificSchemaWhere  SQLIdentifierField = "specific_schema"
	TransformSpecificNameWhere    SQLIdentifierField = "specific_name"
	TransformGroupNameWhere       SQLIdentifierField = "group_name"
	TransformTransformTypeWhere   CharacterDataField = "transform_type"
)

// QueryOneTransform retrieves a row from 'information_schema.transforms' as a Transform.
func QueryOneTransform(db XODB, where WhereClause, order OrderBy) (*Transform, error) {
	const origsqlstr = `SELECT ` +
		`udt_catalog, udt_schema, udt_name, specific_catalog, specific_schema, specific_name, group_name, transform_type ` +
		`FROM information_schema.transforms WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	t := &Transform{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&t.UdtCatalog, &t.UdtSchema, &t.UdtName, &t.SpecificCatalog, &t.SpecificSchema, &t.SpecificName, &t.GroupName, &t.TransformType)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

// QueryTransform retrieves rows from 'information_schema.transforms' as a slice of Transform.
func QueryTransform(db XODB, where WhereClause, order OrderBy) ([]*Transform, error) {
	const origsqlstr = `SELECT ` +
		`udt_catalog, udt_schema, udt_name, specific_catalog, specific_schema, specific_name, group_name, transform_type ` +
		`FROM information_schema.transforms WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*Transform
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		t := Transform{}

		err = q.Scan(&t.UdtCatalog, &t.UdtSchema, &t.UdtName, &t.SpecificCatalog, &t.SpecificSchema, &t.SpecificName, &t.GroupName, &t.TransformType)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &t)
	}
	return vals, nil
}

// AllTransform retrieves all rows from 'information_schema.transforms' as a slice of Transform.
func AllTransform(db XODB, order OrderBy) ([]*Transform, error) {
	const origsqlstr = `SELECT ` +
		`udt_catalog, udt_schema, udt_name, specific_catalog, specific_schema, specific_name, group_name, transform_type ` +
		`FROM information_schema.transforms`

	sqlstr := origsqlstr + order.String()

	var vals []*Transform
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		t := Transform{}

		err = q.Scan(&t.UdtCatalog, &t.UdtSchema, &t.UdtName, &t.SpecificCatalog, &t.SpecificSchema, &t.SpecificName, &t.GroupName, &t.TransformType)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &t)
	}
	return vals, nil
}
