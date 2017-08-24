// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// SequenceTable is the database name for the table.
const SequenceTable = "information_schema.sequences"

// Sequence represents a row from 'information_schema.sequences'.
type Sequence struct {
	SequenceCatalog       SQLIdentifier  `yaml:"sequence_catalog,omitempty"`        // sequence_catalog
	SequenceSchema        SQLIdentifier  `yaml:"sequence_schema,omitempty"`         // sequence_schema
	SequenceName          SQLIdentifier  `yaml:"sequence_name,omitempty"`           // sequence_name
	DataType              CharacterData  `yaml:"data_type,omitempty"`               // data_type
	NumericPrecision      CardinalNumber `yaml:"numeric_precision,omitempty"`       // numeric_precision
	NumericPrecisionRadix CardinalNumber `yaml:"numeric_precision_radix,omitempty"` // numeric_precision_radix
	NumericScale          CardinalNumber `yaml:"numeric_scale,omitempty"`           // numeric_scale
	StartValue            CharacterData  `yaml:"start_value,omitempty"`             // start_value
	MinimumValue          CharacterData  `yaml:"minimum_value,omitempty"`           // minimum_value
	MaximumValue          CharacterData  `yaml:"maximum_value,omitempty"`           // maximum_value
	Increment             CharacterData  `yaml:"increment,omitempty"`               // increment
	CycleOption           YesOrNo        `yaml:"cycle_option,omitempty"`            // cycle_option
}

// Constants defining each column in the table.
const (
	SequenceSequenceCatalogField       = "sequence_catalog"
	SequenceSequenceSchemaField        = "sequence_schema"
	SequenceSequenceNameField          = "sequence_name"
	SequenceDataTypeField              = "data_type"
	SequenceNumericPrecisionField      = "numeric_precision"
	SequenceNumericPrecisionRadixField = "numeric_precision_radix"
	SequenceNumericScaleField          = "numeric_scale"
	SequenceStartValueField            = "start_value"
	SequenceMinimumValueField          = "minimum_value"
	SequenceMaximumValueField          = "maximum_value"
	SequenceIncrementField             = "increment"
	SequenceCycleOptionField           = "cycle_option"
)

// WhereClauses for every type in Sequence.
var (
	SequenceSequenceCatalogWhere       SQLIdentifierField  = "sequence_catalog"
	SequenceSequenceSchemaWhere        SQLIdentifierField  = "sequence_schema"
	SequenceSequenceNameWhere          SQLIdentifierField  = "sequence_name"
	SequenceDataTypeWhere              CharacterDataField  = "data_type"
	SequenceNumericPrecisionWhere      CardinalNumberField = "numeric_precision"
	SequenceNumericPrecisionRadixWhere CardinalNumberField = "numeric_precision_radix"
	SequenceNumericScaleWhere          CardinalNumberField = "numeric_scale"
	SequenceStartValueWhere            CharacterDataField  = "start_value"
	SequenceMinimumValueWhere          CharacterDataField  = "minimum_value"
	SequenceMaximumValueWhere          CharacterDataField  = "maximum_value"
	SequenceIncrementWhere             CharacterDataField  = "increment"
	SequenceCycleOptionWhere           YesOrNoField        = "cycle_option"
)

// QueryOneSequence retrieves a row from 'information_schema.sequences' as a Sequence.
func QueryOneSequence(db XODB, where WhereClause, order OrderBy) (*Sequence, error) {
	const origsqlstr = `SELECT ` +
		`sequence_catalog, sequence_schema, sequence_name, data_type, numeric_precision, numeric_precision_radix, numeric_scale, start_value, minimum_value, maximum_value, increment, cycle_option ` +
		`FROM information_schema.sequences WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	s := &Sequence{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&s.SequenceCatalog, &s.SequenceSchema, &s.SequenceName, &s.DataType, &s.NumericPrecision, &s.NumericPrecisionRadix, &s.NumericScale, &s.StartValue, &s.MinimumValue, &s.MaximumValue, &s.Increment, &s.CycleOption)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return s, nil
}

// QuerySequence retrieves rows from 'information_schema.sequences' as a slice of Sequence.
func QuerySequence(db XODB, where WhereClause, order OrderBy) ([]*Sequence, error) {
	const origsqlstr = `SELECT ` +
		`sequence_catalog, sequence_schema, sequence_name, data_type, numeric_precision, numeric_precision_radix, numeric_scale, start_value, minimum_value, maximum_value, increment, cycle_option ` +
		`FROM information_schema.sequences WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*Sequence
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		s := Sequence{}

		err = q.Scan(&s.SequenceCatalog, &s.SequenceSchema, &s.SequenceName, &s.DataType, &s.NumericPrecision, &s.NumericPrecisionRadix, &s.NumericScale, &s.StartValue, &s.MinimumValue, &s.MaximumValue, &s.Increment, &s.CycleOption)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &s)
	}
	return vals, nil
}

// AllSequence retrieves all rows from 'information_schema.sequences' as a slice of Sequence.
func AllSequence(db XODB, order OrderBy) ([]*Sequence, error) {
	const origsqlstr = `SELECT ` +
		`sequence_catalog, sequence_schema, sequence_name, data_type, numeric_precision, numeric_precision_radix, numeric_scale, start_value, minimum_value, maximum_value, increment, cycle_option ` +
		`FROM information_schema.sequences`

	sqlstr := origsqlstr + order.String()

	var vals []*Sequence
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		s := Sequence{}

		err = q.Scan(&s.SequenceCatalog, &s.SequenceSchema, &s.SequenceName, &s.DataType, &s.NumericPrecision, &s.NumericPrecisionRadix, &s.NumericScale, &s.StartValue, &s.MinimumValue, &s.MaximumValue, &s.Increment, &s.CycleOption)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &s)
	}
	return vals, nil
}
