// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// ForeignServerTable is the database name for the table.
const ForeignServerTable = "information_schema.foreign_servers"

// ForeignServer represents a row from 'information_schema.foreign_servers'.
type ForeignServer struct {
	ForeignServerCatalog      SQLIdentifier `yaml:"foreign_server_catalog,omitempty"`       // foreign_server_catalog
	ForeignServerName         SQLIdentifier `yaml:"foreign_server_name,omitempty"`          // foreign_server_name
	ForeignDataWrapperCatalog SQLIdentifier `yaml:"foreign_data_wrapper_catalog,omitempty"` // foreign_data_wrapper_catalog
	ForeignDataWrapperName    SQLIdentifier `yaml:"foreign_data_wrapper_name,omitempty"`    // foreign_data_wrapper_name
	ForeignServerType         CharacterData `yaml:"foreign_server_type,omitempty"`          // foreign_server_type
	ForeignServerVersion      CharacterData `yaml:"foreign_server_version,omitempty"`       // foreign_server_version
	AuthorizationIdentifier   SQLIdentifier `yaml:"authorization_identifier,omitempty"`     // authorization_identifier
}

// Constants defining each column in the table.
const (
	ForeignServerForeignServerCatalogField      = "foreign_server_catalog"
	ForeignServerForeignServerNameField         = "foreign_server_name"
	ForeignServerForeignDataWrapperCatalogField = "foreign_data_wrapper_catalog"
	ForeignServerForeignDataWrapperNameField    = "foreign_data_wrapper_name"
	ForeignServerForeignServerTypeField         = "foreign_server_type"
	ForeignServerForeignServerVersionField      = "foreign_server_version"
	ForeignServerAuthorizationIdentifierField   = "authorization_identifier"
)

// WhereClauses for every type in ForeignServer.
var (
	ForeignServerForeignServerCatalogWhere      SQLIdentifierField = "foreign_server_catalog"
	ForeignServerForeignServerNameWhere         SQLIdentifierField = "foreign_server_name"
	ForeignServerForeignDataWrapperCatalogWhere SQLIdentifierField = "foreign_data_wrapper_catalog"
	ForeignServerForeignDataWrapperNameWhere    SQLIdentifierField = "foreign_data_wrapper_name"
	ForeignServerForeignServerTypeWhere         CharacterDataField = "foreign_server_type"
	ForeignServerForeignServerVersionWhere      CharacterDataField = "foreign_server_version"
	ForeignServerAuthorizationIdentifierWhere   SQLIdentifierField = "authorization_identifier"
)

// QueryOneForeignServer retrieves a row from 'information_schema.foreign_servers' as a ForeignServer.
func QueryOneForeignServer(db XODB, where WhereClause, order OrderBy) (*ForeignServer, error) {
	const origsqlstr = `SELECT ` +
		`foreign_server_catalog, foreign_server_name, foreign_data_wrapper_catalog, foreign_data_wrapper_name, foreign_server_type, foreign_server_version, authorization_identifier ` +
		`FROM information_schema.foreign_servers WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	fs := &ForeignServer{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&fs.ForeignServerCatalog, &fs.ForeignServerName, &fs.ForeignDataWrapperCatalog, &fs.ForeignDataWrapperName, &fs.ForeignServerType, &fs.ForeignServerVersion, &fs.AuthorizationIdentifier)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return fs, nil
}

// QueryForeignServer retrieves rows from 'information_schema.foreign_servers' as a slice of ForeignServer.
func QueryForeignServer(db XODB, where WhereClause, order OrderBy) ([]*ForeignServer, error) {
	const origsqlstr = `SELECT ` +
		`foreign_server_catalog, foreign_server_name, foreign_data_wrapper_catalog, foreign_data_wrapper_name, foreign_server_type, foreign_server_version, authorization_identifier ` +
		`FROM information_schema.foreign_servers WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ForeignServer
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		fs := ForeignServer{}

		err = q.Scan(&fs.ForeignServerCatalog, &fs.ForeignServerName, &fs.ForeignDataWrapperCatalog, &fs.ForeignDataWrapperName, &fs.ForeignServerType, &fs.ForeignServerVersion, &fs.AuthorizationIdentifier)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &fs)
	}
	return vals, nil
}

// AllForeignServer retrieves all rows from 'information_schema.foreign_servers' as a slice of ForeignServer.
func AllForeignServer(db XODB, order OrderBy) ([]*ForeignServer, error) {
	const origsqlstr = `SELECT ` +
		`foreign_server_catalog, foreign_server_name, foreign_data_wrapper_catalog, foreign_data_wrapper_name, foreign_server_type, foreign_server_version, authorization_identifier ` +
		`FROM information_schema.foreign_servers`

	sqlstr := origsqlstr + order.String()

	var vals []*ForeignServer
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		fs := ForeignServer{}

		err = q.Scan(&fs.ForeignServerCatalog, &fs.ForeignServerName, &fs.ForeignDataWrapperCatalog, &fs.ForeignDataWrapperName, &fs.ForeignServerType, &fs.ForeignServerVersion, &fs.AuthorizationIdentifier)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &fs)
	}
	return vals, nil
}
