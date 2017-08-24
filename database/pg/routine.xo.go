// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// RoutineTable is the database name for the table.
const RoutineTable = "information_schema.routines"

// Routine represents a row from 'information_schema.routines'.
type Routine struct {
	SpecificCatalog                 SQLIdentifier  `yaml:"specific_catalog,omitempty"`                    // specific_catalog
	SpecificSchema                  SQLIdentifier  `yaml:"specific_schema,omitempty"`                     // specific_schema
	SpecificName                    SQLIdentifier  `yaml:"specific_name,omitempty"`                       // specific_name
	RoutineCatalog                  SQLIdentifier  `yaml:"routine_catalog,omitempty"`                     // routine_catalog
	RoutineSchema                   SQLIdentifier  `yaml:"routine_schema,omitempty"`                      // routine_schema
	RoutineName                     SQLIdentifier  `yaml:"routine_name,omitempty"`                        // routine_name
	RoutineType                     CharacterData  `yaml:"routine_type,omitempty"`                        // routine_type
	ModuleCatalog                   SQLIdentifier  `yaml:"module_catalog,omitempty"`                      // module_catalog
	ModuleSchema                    SQLIdentifier  `yaml:"module_schema,omitempty"`                       // module_schema
	ModuleName                      SQLIdentifier  `yaml:"module_name,omitempty"`                         // module_name
	UdtCatalog                      SQLIdentifier  `yaml:"udt_catalog,omitempty"`                         // udt_catalog
	UdtSchema                       SQLIdentifier  `yaml:"udt_schema,omitempty"`                          // udt_schema
	UdtName                         SQLIdentifier  `yaml:"udt_name,omitempty"`                            // udt_name
	DataType                        CharacterData  `yaml:"data_type,omitempty"`                           // data_type
	CharacterMaximumLength          CardinalNumber `yaml:"character_maximum_length,omitempty"`            // character_maximum_length
	CharacterOctetLength            CardinalNumber `yaml:"character_octet_length,omitempty"`              // character_octet_length
	CharacterSetCatalog             SQLIdentifier  `yaml:"character_set_catalog,omitempty"`               // character_set_catalog
	CharacterSetSchema              SQLIdentifier  `yaml:"character_set_schema,omitempty"`                // character_set_schema
	CharacterSetName                SQLIdentifier  `yaml:"character_set_name,omitempty"`                  // character_set_name
	CollationCatalog                SQLIdentifier  `yaml:"collation_catalog,omitempty"`                   // collation_catalog
	CollationSchema                 SQLIdentifier  `yaml:"collation_schema,omitempty"`                    // collation_schema
	CollationName                   SQLIdentifier  `yaml:"collation_name,omitempty"`                      // collation_name
	NumericPrecision                CardinalNumber `yaml:"numeric_precision,omitempty"`                   // numeric_precision
	NumericPrecisionRadix           CardinalNumber `yaml:"numeric_precision_radix,omitempty"`             // numeric_precision_radix
	NumericScale                    CardinalNumber `yaml:"numeric_scale,omitempty"`                       // numeric_scale
	DatetimePrecision               CardinalNumber `yaml:"datetime_precision,omitempty"`                  // datetime_precision
	IntervalType                    CharacterData  `yaml:"interval_type,omitempty"`                       // interval_type
	IntervalPrecision               CardinalNumber `yaml:"interval_precision,omitempty"`                  // interval_precision
	TypeUdtCatalog                  SQLIdentifier  `yaml:"type_udt_catalog,omitempty"`                    // type_udt_catalog
	TypeUdtSchema                   SQLIdentifier  `yaml:"type_udt_schema,omitempty"`                     // type_udt_schema
	TypeUdtName                     SQLIdentifier  `yaml:"type_udt_name,omitempty"`                       // type_udt_name
	ScopeCatalog                    SQLIdentifier  `yaml:"scope_catalog,omitempty"`                       // scope_catalog
	ScopeSchema                     SQLIdentifier  `yaml:"scope_schema,omitempty"`                        // scope_schema
	ScopeName                       SQLIdentifier  `yaml:"scope_name,omitempty"`                          // scope_name
	MaximumCardinality              CardinalNumber `yaml:"maximum_cardinality,omitempty"`                 // maximum_cardinality
	DtdIdentifier                   SQLIdentifier  `yaml:"dtd_identifier,omitempty"`                      // dtd_identifier
	RoutineBody                     CharacterData  `yaml:"routine_body,omitempty"`                        // routine_body
	RoutineDefinition               CharacterData  `yaml:"routine_definition,omitempty"`                  // routine_definition
	ExternalName                    CharacterData  `yaml:"external_name,omitempty"`                       // external_name
	ExternalLanguage                CharacterData  `yaml:"external_language,omitempty"`                   // external_language
	ParameterStyle                  CharacterData  `yaml:"parameter_style,omitempty"`                     // parameter_style
	IsDeterministic                 YesOrNo        `yaml:"is_deterministic,omitempty"`                    // is_deterministic
	SQLDataAccess                   CharacterData  `yaml:"sql_data_access,omitempty"`                     // sql_data_access
	IsNullCall                      YesOrNo        `yaml:"is_null_call,omitempty"`                        // is_null_call
	SQLPath                         CharacterData  `yaml:"sql_path,omitempty"`                            // sql_path
	SchemaLevelRoutine              YesOrNo        `yaml:"schema_level_routine,omitempty"`                // schema_level_routine
	MaxDynamicResultSets            CardinalNumber `yaml:"max_dynamic_result_sets,omitempty"`             // max_dynamic_result_sets
	IsUserDefinedCast               YesOrNo        `yaml:"is_user_defined_cast,omitempty"`                // is_user_defined_cast
	IsImplicitlyInvocable           YesOrNo        `yaml:"is_implicitly_invocable,omitempty"`             // is_implicitly_invocable
	SecurityType                    CharacterData  `yaml:"security_type,omitempty"`                       // security_type
	ToSQLSpecificCatalog            SQLIdentifier  `yaml:"to_sql_specific_catalog,omitempty"`             // to_sql_specific_catalog
	ToSQLSpecificSchema             SQLIdentifier  `yaml:"to_sql_specific_schema,omitempty"`              // to_sql_specific_schema
	ToSQLSpecificName               SQLIdentifier  `yaml:"to_sql_specific_name,omitempty"`                // to_sql_specific_name
	AsLocator                       YesOrNo        `yaml:"as_locator,omitempty"`                          // as_locator
	Created                         TimeStamp      `yaml:"created,omitempty"`                             // created
	LastAltered                     TimeStamp      `yaml:"last_altered,omitempty"`                        // last_altered
	NewSavepointLevel               YesOrNo        `yaml:"new_savepoint_level,omitempty"`                 // new_savepoint_level
	IsUdtDependent                  YesOrNo        `yaml:"is_udt_dependent,omitempty"`                    // is_udt_dependent
	ResultCastFromDataType          CharacterData  `yaml:"result_cast_from_data_type,omitempty"`          // result_cast_from_data_type
	ResultCastAsLocator             YesOrNo        `yaml:"result_cast_as_locator,omitempty"`              // result_cast_as_locator
	ResultCastCharMaxLength         CardinalNumber `yaml:"result_cast_char_max_length,omitempty"`         // result_cast_char_max_length
	ResultCastCharOctetLength       CardinalNumber `yaml:"result_cast_char_octet_length,omitempty"`       // result_cast_char_octet_length
	ResultCastCharSetCatalog        SQLIdentifier  `yaml:"result_cast_char_set_catalog,omitempty"`        // result_cast_char_set_catalog
	ResultCastCharSetSchema         SQLIdentifier  `yaml:"result_cast_char_set_schema,omitempty"`         // result_cast_char_set_schema
	ResultCastCharSetName           SQLIdentifier  `yaml:"result_cast_char_set_name,omitempty"`           // result_cast_char_set_name
	ResultCastCollationCatalog      SQLIdentifier  `yaml:"result_cast_collation_catalog,omitempty"`       // result_cast_collation_catalog
	ResultCastCollationSchema       SQLIdentifier  `yaml:"result_cast_collation_schema,omitempty"`        // result_cast_collation_schema
	ResultCastCollationName         SQLIdentifier  `yaml:"result_cast_collation_name,omitempty"`          // result_cast_collation_name
	ResultCastNumericPrecision      CardinalNumber `yaml:"result_cast_numeric_precision,omitempty"`       // result_cast_numeric_precision
	ResultCastNumericPrecisionRadix CardinalNumber `yaml:"result_cast_numeric_precision_radix,omitempty"` // result_cast_numeric_precision_radix
	ResultCastNumericScale          CardinalNumber `yaml:"result_cast_numeric_scale,omitempty"`           // result_cast_numeric_scale
	ResultCastDatetimePrecision     CardinalNumber `yaml:"result_cast_datetime_precision,omitempty"`      // result_cast_datetime_precision
	ResultCastIntervalType          CharacterData  `yaml:"result_cast_interval_type,omitempty"`           // result_cast_interval_type
	ResultCastIntervalPrecision     CardinalNumber `yaml:"result_cast_interval_precision,omitempty"`      // result_cast_interval_precision
	ResultCastTypeUdtCatalog        SQLIdentifier  `yaml:"result_cast_type_udt_catalog,omitempty"`        // result_cast_type_udt_catalog
	ResultCastTypeUdtSchema         SQLIdentifier  `yaml:"result_cast_type_udt_schema,omitempty"`         // result_cast_type_udt_schema
	ResultCastTypeUdtName           SQLIdentifier  `yaml:"result_cast_type_udt_name,omitempty"`           // result_cast_type_udt_name
	ResultCastScopeCatalog          SQLIdentifier  `yaml:"result_cast_scope_catalog,omitempty"`           // result_cast_scope_catalog
	ResultCastScopeSchema           SQLIdentifier  `yaml:"result_cast_scope_schema,omitempty"`            // result_cast_scope_schema
	ResultCastScopeName             SQLIdentifier  `yaml:"result_cast_scope_name,omitempty"`              // result_cast_scope_name
	ResultCastMaximumCardinality    CardinalNumber `yaml:"result_cast_maximum_cardinality,omitempty"`     // result_cast_maximum_cardinality
	ResultCastDtdIdentifier         SQLIdentifier  `yaml:"result_cast_dtd_identifier,omitempty"`          // result_cast_dtd_identifier
}

// Constants defining each column in the table.
const (
	RoutineSpecificCatalogField                 = "specific_catalog"
	RoutineSpecificSchemaField                  = "specific_schema"
	RoutineSpecificNameField                    = "specific_name"
	RoutineRoutineCatalogField                  = "routine_catalog"
	RoutineRoutineSchemaField                   = "routine_schema"
	RoutineRoutineNameField                     = "routine_name"
	RoutineRoutineTypeField                     = "routine_type"
	RoutineModuleCatalogField                   = "module_catalog"
	RoutineModuleSchemaField                    = "module_schema"
	RoutineModuleNameField                      = "module_name"
	RoutineUdtCatalogField                      = "udt_catalog"
	RoutineUdtSchemaField                       = "udt_schema"
	RoutineUdtNameField                         = "udt_name"
	RoutineDataTypeField                        = "data_type"
	RoutineCharacterMaximumLengthField          = "character_maximum_length"
	RoutineCharacterOctetLengthField            = "character_octet_length"
	RoutineCharacterSetCatalogField             = "character_set_catalog"
	RoutineCharacterSetSchemaField              = "character_set_schema"
	RoutineCharacterSetNameField                = "character_set_name"
	RoutineCollationCatalogField                = "collation_catalog"
	RoutineCollationSchemaField                 = "collation_schema"
	RoutineCollationNameField                   = "collation_name"
	RoutineNumericPrecisionField                = "numeric_precision"
	RoutineNumericPrecisionRadixField           = "numeric_precision_radix"
	RoutineNumericScaleField                    = "numeric_scale"
	RoutineDatetimePrecisionField               = "datetime_precision"
	RoutineIntervalTypeField                    = "interval_type"
	RoutineIntervalPrecisionField               = "interval_precision"
	RoutineTypeUdtCatalogField                  = "type_udt_catalog"
	RoutineTypeUdtSchemaField                   = "type_udt_schema"
	RoutineTypeUdtNameField                     = "type_udt_name"
	RoutineScopeCatalogField                    = "scope_catalog"
	RoutineScopeSchemaField                     = "scope_schema"
	RoutineScopeNameField                       = "scope_name"
	RoutineMaximumCardinalityField              = "maximum_cardinality"
	RoutineDtdIdentifierField                   = "dtd_identifier"
	RoutineRoutineBodyField                     = "routine_body"
	RoutineRoutineDefinitionField               = "routine_definition"
	RoutineExternalNameField                    = "external_name"
	RoutineExternalLanguageField                = "external_language"
	RoutineParameterStyleField                  = "parameter_style"
	RoutineIsDeterministicField                 = "is_deterministic"
	RoutineSQLDataAccessField                   = "sql_data_access"
	RoutineIsNullCallField                      = "is_null_call"
	RoutineSQLPathField                         = "sql_path"
	RoutineSchemaLevelRoutineField              = "schema_level_routine"
	RoutineMaxDynamicResultSetsField            = "max_dynamic_result_sets"
	RoutineIsUserDefinedCastField               = "is_user_defined_cast"
	RoutineIsImplicitlyInvocableField           = "is_implicitly_invocable"
	RoutineSecurityTypeField                    = "security_type"
	RoutineToSQLSpecificCatalogField            = "to_sql_specific_catalog"
	RoutineToSQLSpecificSchemaField             = "to_sql_specific_schema"
	RoutineToSQLSpecificNameField               = "to_sql_specific_name"
	RoutineAsLocatorField                       = "as_locator"
	RoutineCreatedField                         = "created"
	RoutineLastAlteredField                     = "last_altered"
	RoutineNewSavepointLevelField               = "new_savepoint_level"
	RoutineIsUdtDependentField                  = "is_udt_dependent"
	RoutineResultCastFromDataTypeField          = "result_cast_from_data_type"
	RoutineResultCastAsLocatorField             = "result_cast_as_locator"
	RoutineResultCastCharMaxLengthField         = "result_cast_char_max_length"
	RoutineResultCastCharOctetLengthField       = "result_cast_char_octet_length"
	RoutineResultCastCharSetCatalogField        = "result_cast_char_set_catalog"
	RoutineResultCastCharSetSchemaField         = "result_cast_char_set_schema"
	RoutineResultCastCharSetNameField           = "result_cast_char_set_name"
	RoutineResultCastCollationCatalogField      = "result_cast_collation_catalog"
	RoutineResultCastCollationSchemaField       = "result_cast_collation_schema"
	RoutineResultCastCollationNameField         = "result_cast_collation_name"
	RoutineResultCastNumericPrecisionField      = "result_cast_numeric_precision"
	RoutineResultCastNumericPrecisionRadixField = "result_cast_numeric_precision_radix"
	RoutineResultCastNumericScaleField          = "result_cast_numeric_scale"
	RoutineResultCastDatetimePrecisionField     = "result_cast_datetime_precision"
	RoutineResultCastIntervalTypeField          = "result_cast_interval_type"
	RoutineResultCastIntervalPrecisionField     = "result_cast_interval_precision"
	RoutineResultCastTypeUdtCatalogField        = "result_cast_type_udt_catalog"
	RoutineResultCastTypeUdtSchemaField         = "result_cast_type_udt_schema"
	RoutineResultCastTypeUdtNameField           = "result_cast_type_udt_name"
	RoutineResultCastScopeCatalogField          = "result_cast_scope_catalog"
	RoutineResultCastScopeSchemaField           = "result_cast_scope_schema"
	RoutineResultCastScopeNameField             = "result_cast_scope_name"
	RoutineResultCastMaximumCardinalityField    = "result_cast_maximum_cardinality"
	RoutineResultCastDtdIdentifierField         = "result_cast_dtd_identifier"
)

// WhereClauses for every type in Routine.
var (
	RoutineSpecificCatalogWhere                 SQLIdentifierField  = "specific_catalog"
	RoutineSpecificSchemaWhere                  SQLIdentifierField  = "specific_schema"
	RoutineSpecificNameWhere                    SQLIdentifierField  = "specific_name"
	RoutineRoutineCatalogWhere                  SQLIdentifierField  = "routine_catalog"
	RoutineRoutineSchemaWhere                   SQLIdentifierField  = "routine_schema"
	RoutineRoutineNameWhere                     SQLIdentifierField  = "routine_name"
	RoutineRoutineTypeWhere                     CharacterDataField  = "routine_type"
	RoutineModuleCatalogWhere                   SQLIdentifierField  = "module_catalog"
	RoutineModuleSchemaWhere                    SQLIdentifierField  = "module_schema"
	RoutineModuleNameWhere                      SQLIdentifierField  = "module_name"
	RoutineUdtCatalogWhere                      SQLIdentifierField  = "udt_catalog"
	RoutineUdtSchemaWhere                       SQLIdentifierField  = "udt_schema"
	RoutineUdtNameWhere                         SQLIdentifierField  = "udt_name"
	RoutineDataTypeWhere                        CharacterDataField  = "data_type"
	RoutineCharacterMaximumLengthWhere          CardinalNumberField = "character_maximum_length"
	RoutineCharacterOctetLengthWhere            CardinalNumberField = "character_octet_length"
	RoutineCharacterSetCatalogWhere             SQLIdentifierField  = "character_set_catalog"
	RoutineCharacterSetSchemaWhere              SQLIdentifierField  = "character_set_schema"
	RoutineCharacterSetNameWhere                SQLIdentifierField  = "character_set_name"
	RoutineCollationCatalogWhere                SQLIdentifierField  = "collation_catalog"
	RoutineCollationSchemaWhere                 SQLIdentifierField  = "collation_schema"
	RoutineCollationNameWhere                   SQLIdentifierField  = "collation_name"
	RoutineNumericPrecisionWhere                CardinalNumberField = "numeric_precision"
	RoutineNumericPrecisionRadixWhere           CardinalNumberField = "numeric_precision_radix"
	RoutineNumericScaleWhere                    CardinalNumberField = "numeric_scale"
	RoutineDatetimePrecisionWhere               CardinalNumberField = "datetime_precision"
	RoutineIntervalTypeWhere                    CharacterDataField  = "interval_type"
	RoutineIntervalPrecisionWhere               CardinalNumberField = "interval_precision"
	RoutineTypeUdtCatalogWhere                  SQLIdentifierField  = "type_udt_catalog"
	RoutineTypeUdtSchemaWhere                   SQLIdentifierField  = "type_udt_schema"
	RoutineTypeUdtNameWhere                     SQLIdentifierField  = "type_udt_name"
	RoutineScopeCatalogWhere                    SQLIdentifierField  = "scope_catalog"
	RoutineScopeSchemaWhere                     SQLIdentifierField  = "scope_schema"
	RoutineScopeNameWhere                       SQLIdentifierField  = "scope_name"
	RoutineMaximumCardinalityWhere              CardinalNumberField = "maximum_cardinality"
	RoutineDtdIdentifierWhere                   SQLIdentifierField  = "dtd_identifier"
	RoutineRoutineBodyWhere                     CharacterDataField  = "routine_body"
	RoutineRoutineDefinitionWhere               CharacterDataField  = "routine_definition"
	RoutineExternalNameWhere                    CharacterDataField  = "external_name"
	RoutineExternalLanguageWhere                CharacterDataField  = "external_language"
	RoutineParameterStyleWhere                  CharacterDataField  = "parameter_style"
	RoutineIsDeterministicWhere                 YesOrNoField        = "is_deterministic"
	RoutineSQLDataAccessWhere                   CharacterDataField  = "sql_data_access"
	RoutineIsNullCallWhere                      YesOrNoField        = "is_null_call"
	RoutineSQLPathWhere                         CharacterDataField  = "sql_path"
	RoutineSchemaLevelRoutineWhere              YesOrNoField        = "schema_level_routine"
	RoutineMaxDynamicResultSetsWhere            CardinalNumberField = "max_dynamic_result_sets"
	RoutineIsUserDefinedCastWhere               YesOrNoField        = "is_user_defined_cast"
	RoutineIsImplicitlyInvocableWhere           YesOrNoField        = "is_implicitly_invocable"
	RoutineSecurityTypeWhere                    CharacterDataField  = "security_type"
	RoutineToSQLSpecificCatalogWhere            SQLIdentifierField  = "to_sql_specific_catalog"
	RoutineToSQLSpecificSchemaWhere             SQLIdentifierField  = "to_sql_specific_schema"
	RoutineToSQLSpecificNameWhere               SQLIdentifierField  = "to_sql_specific_name"
	RoutineAsLocatorWhere                       YesOrNoField        = "as_locator"
	RoutineCreatedWhere                         TimeStampField      = "created"
	RoutineLastAlteredWhere                     TimeStampField      = "last_altered"
	RoutineNewSavepointLevelWhere               YesOrNoField        = "new_savepoint_level"
	RoutineIsUdtDependentWhere                  YesOrNoField        = "is_udt_dependent"
	RoutineResultCastFromDataTypeWhere          CharacterDataField  = "result_cast_from_data_type"
	RoutineResultCastAsLocatorWhere             YesOrNoField        = "result_cast_as_locator"
	RoutineResultCastCharMaxLengthWhere         CardinalNumberField = "result_cast_char_max_length"
	RoutineResultCastCharOctetLengthWhere       CardinalNumberField = "result_cast_char_octet_length"
	RoutineResultCastCharSetCatalogWhere        SQLIdentifierField  = "result_cast_char_set_catalog"
	RoutineResultCastCharSetSchemaWhere         SQLIdentifierField  = "result_cast_char_set_schema"
	RoutineResultCastCharSetNameWhere           SQLIdentifierField  = "result_cast_char_set_name"
	RoutineResultCastCollationCatalogWhere      SQLIdentifierField  = "result_cast_collation_catalog"
	RoutineResultCastCollationSchemaWhere       SQLIdentifierField  = "result_cast_collation_schema"
	RoutineResultCastCollationNameWhere         SQLIdentifierField  = "result_cast_collation_name"
	RoutineResultCastNumericPrecisionWhere      CardinalNumberField = "result_cast_numeric_precision"
	RoutineResultCastNumericPrecisionRadixWhere CardinalNumberField = "result_cast_numeric_precision_radix"
	RoutineResultCastNumericScaleWhere          CardinalNumberField = "result_cast_numeric_scale"
	RoutineResultCastDatetimePrecisionWhere     CardinalNumberField = "result_cast_datetime_precision"
	RoutineResultCastIntervalTypeWhere          CharacterDataField  = "result_cast_interval_type"
	RoutineResultCastIntervalPrecisionWhere     CardinalNumberField = "result_cast_interval_precision"
	RoutineResultCastTypeUdtCatalogWhere        SQLIdentifierField  = "result_cast_type_udt_catalog"
	RoutineResultCastTypeUdtSchemaWhere         SQLIdentifierField  = "result_cast_type_udt_schema"
	RoutineResultCastTypeUdtNameWhere           SQLIdentifierField  = "result_cast_type_udt_name"
	RoutineResultCastScopeCatalogWhere          SQLIdentifierField  = "result_cast_scope_catalog"
	RoutineResultCastScopeSchemaWhere           SQLIdentifierField  = "result_cast_scope_schema"
	RoutineResultCastScopeNameWhere             SQLIdentifierField  = "result_cast_scope_name"
	RoutineResultCastMaximumCardinalityWhere    CardinalNumberField = "result_cast_maximum_cardinality"
	RoutineResultCastDtdIdentifierWhere         SQLIdentifierField  = "result_cast_dtd_identifier"
)

// QueryOneRoutine retrieves a row from 'information_schema.routines' as a Routine.
func QueryOneRoutine(db XODB, where WhereClause, order OrderBy) (*Routine, error) {
	const origsqlstr = `SELECT ` +
		`specific_catalog, specific_schema, specific_name, routine_catalog, routine_schema, routine_name, routine_type, module_catalog, module_schema, module_name, udt_catalog, udt_schema, udt_name, data_type, character_maximum_length, character_octet_length, character_set_catalog, character_set_schema, character_set_name, collation_catalog, collation_schema, collation_name, numeric_precision, numeric_precision_radix, numeric_scale, datetime_precision, interval_type, interval_precision, type_udt_catalog, type_udt_schema, type_udt_name, scope_catalog, scope_schema, scope_name, maximum_cardinality, dtd_identifier, routine_body, routine_definition, external_name, external_language, parameter_style, is_deterministic, sql_data_access, is_null_call, sql_path, schema_level_routine, max_dynamic_result_sets, is_user_defined_cast, is_implicitly_invocable, security_type, to_sql_specific_catalog, to_sql_specific_schema, to_sql_specific_name, as_locator, created, last_altered, new_savepoint_level, is_udt_dependent, result_cast_from_data_type, result_cast_as_locator, result_cast_char_max_length, result_cast_char_octet_length, result_cast_char_set_catalog, result_cast_char_set_schema, result_cast_char_set_name, result_cast_collation_catalog, result_cast_collation_schema, result_cast_collation_name, result_cast_numeric_precision, result_cast_numeric_precision_radix, result_cast_numeric_scale, result_cast_datetime_precision, result_cast_interval_type, result_cast_interval_precision, result_cast_type_udt_catalog, result_cast_type_udt_schema, result_cast_type_udt_name, result_cast_scope_catalog, result_cast_scope_schema, result_cast_scope_name, result_cast_maximum_cardinality, result_cast_dtd_identifier ` +
		`FROM information_schema.routines WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	r := &Routine{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&r.SpecificCatalog, &r.SpecificSchema, &r.SpecificName, &r.RoutineCatalog, &r.RoutineSchema, &r.RoutineName, &r.RoutineType, &r.ModuleCatalog, &r.ModuleSchema, &r.ModuleName, &r.UdtCatalog, &r.UdtSchema, &r.UdtName, &r.DataType, &r.CharacterMaximumLength, &r.CharacterOctetLength, &r.CharacterSetCatalog, &r.CharacterSetSchema, &r.CharacterSetName, &r.CollationCatalog, &r.CollationSchema, &r.CollationName, &r.NumericPrecision, &r.NumericPrecisionRadix, &r.NumericScale, &r.DatetimePrecision, &r.IntervalType, &r.IntervalPrecision, &r.TypeUdtCatalog, &r.TypeUdtSchema, &r.TypeUdtName, &r.ScopeCatalog, &r.ScopeSchema, &r.ScopeName, &r.MaximumCardinality, &r.DtdIdentifier, &r.RoutineBody, &r.RoutineDefinition, &r.ExternalName, &r.ExternalLanguage, &r.ParameterStyle, &r.IsDeterministic, &r.SQLDataAccess, &r.IsNullCall, &r.SQLPath, &r.SchemaLevelRoutine, &r.MaxDynamicResultSets, &r.IsUserDefinedCast, &r.IsImplicitlyInvocable, &r.SecurityType, &r.ToSQLSpecificCatalog, &r.ToSQLSpecificSchema, &r.ToSQLSpecificName, &r.AsLocator, &r.Created, &r.LastAltered, &r.NewSavepointLevel, &r.IsUdtDependent, &r.ResultCastFromDataType, &r.ResultCastAsLocator, &r.ResultCastCharMaxLength, &r.ResultCastCharOctetLength, &r.ResultCastCharSetCatalog, &r.ResultCastCharSetSchema, &r.ResultCastCharSetName, &r.ResultCastCollationCatalog, &r.ResultCastCollationSchema, &r.ResultCastCollationName, &r.ResultCastNumericPrecision, &r.ResultCastNumericPrecisionRadix, &r.ResultCastNumericScale, &r.ResultCastDatetimePrecision, &r.ResultCastIntervalType, &r.ResultCastIntervalPrecision, &r.ResultCastTypeUdtCatalog, &r.ResultCastTypeUdtSchema, &r.ResultCastTypeUdtName, &r.ResultCastScopeCatalog, &r.ResultCastScopeSchema, &r.ResultCastScopeName, &r.ResultCastMaximumCardinality, &r.ResultCastDtdIdentifier)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return r, nil
}

// QueryRoutine retrieves rows from 'information_schema.routines' as a slice of Routine.
func QueryRoutine(db XODB, where WhereClause, order OrderBy) ([]*Routine, error) {
	const origsqlstr = `SELECT ` +
		`specific_catalog, specific_schema, specific_name, routine_catalog, routine_schema, routine_name, routine_type, module_catalog, module_schema, module_name, udt_catalog, udt_schema, udt_name, data_type, character_maximum_length, character_octet_length, character_set_catalog, character_set_schema, character_set_name, collation_catalog, collation_schema, collation_name, numeric_precision, numeric_precision_radix, numeric_scale, datetime_precision, interval_type, interval_precision, type_udt_catalog, type_udt_schema, type_udt_name, scope_catalog, scope_schema, scope_name, maximum_cardinality, dtd_identifier, routine_body, routine_definition, external_name, external_language, parameter_style, is_deterministic, sql_data_access, is_null_call, sql_path, schema_level_routine, max_dynamic_result_sets, is_user_defined_cast, is_implicitly_invocable, security_type, to_sql_specific_catalog, to_sql_specific_schema, to_sql_specific_name, as_locator, created, last_altered, new_savepoint_level, is_udt_dependent, result_cast_from_data_type, result_cast_as_locator, result_cast_char_max_length, result_cast_char_octet_length, result_cast_char_set_catalog, result_cast_char_set_schema, result_cast_char_set_name, result_cast_collation_catalog, result_cast_collation_schema, result_cast_collation_name, result_cast_numeric_precision, result_cast_numeric_precision_radix, result_cast_numeric_scale, result_cast_datetime_precision, result_cast_interval_type, result_cast_interval_precision, result_cast_type_udt_catalog, result_cast_type_udt_schema, result_cast_type_udt_name, result_cast_scope_catalog, result_cast_scope_schema, result_cast_scope_name, result_cast_maximum_cardinality, result_cast_dtd_identifier ` +
		`FROM information_schema.routines WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*Routine
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		r := Routine{}

		err = q.Scan(&r.SpecificCatalog, &r.SpecificSchema, &r.SpecificName, &r.RoutineCatalog, &r.RoutineSchema, &r.RoutineName, &r.RoutineType, &r.ModuleCatalog, &r.ModuleSchema, &r.ModuleName, &r.UdtCatalog, &r.UdtSchema, &r.UdtName, &r.DataType, &r.CharacterMaximumLength, &r.CharacterOctetLength, &r.CharacterSetCatalog, &r.CharacterSetSchema, &r.CharacterSetName, &r.CollationCatalog, &r.CollationSchema, &r.CollationName, &r.NumericPrecision, &r.NumericPrecisionRadix, &r.NumericScale, &r.DatetimePrecision, &r.IntervalType, &r.IntervalPrecision, &r.TypeUdtCatalog, &r.TypeUdtSchema, &r.TypeUdtName, &r.ScopeCatalog, &r.ScopeSchema, &r.ScopeName, &r.MaximumCardinality, &r.DtdIdentifier, &r.RoutineBody, &r.RoutineDefinition, &r.ExternalName, &r.ExternalLanguage, &r.ParameterStyle, &r.IsDeterministic, &r.SQLDataAccess, &r.IsNullCall, &r.SQLPath, &r.SchemaLevelRoutine, &r.MaxDynamicResultSets, &r.IsUserDefinedCast, &r.IsImplicitlyInvocable, &r.SecurityType, &r.ToSQLSpecificCatalog, &r.ToSQLSpecificSchema, &r.ToSQLSpecificName, &r.AsLocator, &r.Created, &r.LastAltered, &r.NewSavepointLevel, &r.IsUdtDependent, &r.ResultCastFromDataType, &r.ResultCastAsLocator, &r.ResultCastCharMaxLength, &r.ResultCastCharOctetLength, &r.ResultCastCharSetCatalog, &r.ResultCastCharSetSchema, &r.ResultCastCharSetName, &r.ResultCastCollationCatalog, &r.ResultCastCollationSchema, &r.ResultCastCollationName, &r.ResultCastNumericPrecision, &r.ResultCastNumericPrecisionRadix, &r.ResultCastNumericScale, &r.ResultCastDatetimePrecision, &r.ResultCastIntervalType, &r.ResultCastIntervalPrecision, &r.ResultCastTypeUdtCatalog, &r.ResultCastTypeUdtSchema, &r.ResultCastTypeUdtName, &r.ResultCastScopeCatalog, &r.ResultCastScopeSchema, &r.ResultCastScopeName, &r.ResultCastMaximumCardinality, &r.ResultCastDtdIdentifier)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &r)
	}
	return vals, nil
}

// AllRoutine retrieves all rows from 'information_schema.routines' as a slice of Routine.
func AllRoutine(db XODB, order OrderBy) ([]*Routine, error) {
	const origsqlstr = `SELECT ` +
		`specific_catalog, specific_schema, specific_name, routine_catalog, routine_schema, routine_name, routine_type, module_catalog, module_schema, module_name, udt_catalog, udt_schema, udt_name, data_type, character_maximum_length, character_octet_length, character_set_catalog, character_set_schema, character_set_name, collation_catalog, collation_schema, collation_name, numeric_precision, numeric_precision_radix, numeric_scale, datetime_precision, interval_type, interval_precision, type_udt_catalog, type_udt_schema, type_udt_name, scope_catalog, scope_schema, scope_name, maximum_cardinality, dtd_identifier, routine_body, routine_definition, external_name, external_language, parameter_style, is_deterministic, sql_data_access, is_null_call, sql_path, schema_level_routine, max_dynamic_result_sets, is_user_defined_cast, is_implicitly_invocable, security_type, to_sql_specific_catalog, to_sql_specific_schema, to_sql_specific_name, as_locator, created, last_altered, new_savepoint_level, is_udt_dependent, result_cast_from_data_type, result_cast_as_locator, result_cast_char_max_length, result_cast_char_octet_length, result_cast_char_set_catalog, result_cast_char_set_schema, result_cast_char_set_name, result_cast_collation_catalog, result_cast_collation_schema, result_cast_collation_name, result_cast_numeric_precision, result_cast_numeric_precision_radix, result_cast_numeric_scale, result_cast_datetime_precision, result_cast_interval_type, result_cast_interval_precision, result_cast_type_udt_catalog, result_cast_type_udt_schema, result_cast_type_udt_name, result_cast_scope_catalog, result_cast_scope_schema, result_cast_scope_name, result_cast_maximum_cardinality, result_cast_dtd_identifier ` +
		`FROM information_schema.routines`

	sqlstr := origsqlstr + order.String()

	var vals []*Routine
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		r := Routine{}

		err = q.Scan(&r.SpecificCatalog, &r.SpecificSchema, &r.SpecificName, &r.RoutineCatalog, &r.RoutineSchema, &r.RoutineName, &r.RoutineType, &r.ModuleCatalog, &r.ModuleSchema, &r.ModuleName, &r.UdtCatalog, &r.UdtSchema, &r.UdtName, &r.DataType, &r.CharacterMaximumLength, &r.CharacterOctetLength, &r.CharacterSetCatalog, &r.CharacterSetSchema, &r.CharacterSetName, &r.CollationCatalog, &r.CollationSchema, &r.CollationName, &r.NumericPrecision, &r.NumericPrecisionRadix, &r.NumericScale, &r.DatetimePrecision, &r.IntervalType, &r.IntervalPrecision, &r.TypeUdtCatalog, &r.TypeUdtSchema, &r.TypeUdtName, &r.ScopeCatalog, &r.ScopeSchema, &r.ScopeName, &r.MaximumCardinality, &r.DtdIdentifier, &r.RoutineBody, &r.RoutineDefinition, &r.ExternalName, &r.ExternalLanguage, &r.ParameterStyle, &r.IsDeterministic, &r.SQLDataAccess, &r.IsNullCall, &r.SQLPath, &r.SchemaLevelRoutine, &r.MaxDynamicResultSets, &r.IsUserDefinedCast, &r.IsImplicitlyInvocable, &r.SecurityType, &r.ToSQLSpecificCatalog, &r.ToSQLSpecificSchema, &r.ToSQLSpecificName, &r.AsLocator, &r.Created, &r.LastAltered, &r.NewSavepointLevel, &r.IsUdtDependent, &r.ResultCastFromDataType, &r.ResultCastAsLocator, &r.ResultCastCharMaxLength, &r.ResultCastCharOctetLength, &r.ResultCastCharSetCatalog, &r.ResultCastCharSetSchema, &r.ResultCastCharSetName, &r.ResultCastCollationCatalog, &r.ResultCastCollationSchema, &r.ResultCastCollationName, &r.ResultCastNumericPrecision, &r.ResultCastNumericPrecisionRadix, &r.ResultCastNumericScale, &r.ResultCastDatetimePrecision, &r.ResultCastIntervalType, &r.ResultCastIntervalPrecision, &r.ResultCastTypeUdtCatalog, &r.ResultCastTypeUdtSchema, &r.ResultCastTypeUdtName, &r.ResultCastScopeCatalog, &r.ResultCastScopeSchema, &r.ResultCastScopeName, &r.ResultCastMaximumCardinality, &r.ResultCastDtdIdentifier)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &r)
	}
	return vals, nil
}
