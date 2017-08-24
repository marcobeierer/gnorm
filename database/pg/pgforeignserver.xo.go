// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/pkg/errors"
)

// PgForeignServerTable is the database name for the table.
const PgForeignServerTable = "information_schema._pg_foreign_servers"

// PgForeignServer represents a row from 'information_schema._pg_foreign_servers'.
type PgForeignServer struct {
	Oid                       Oid              `yaml:"oid,omitempty"`                          // oid
	Srvoptions                []sql.NullString `yaml:"srvoptions,omitempty"`                   // srvoptions
	ForeignServerCatalog      SQLIdentifier    `yaml:"foreign_server_catalog,omitempty"`       // foreign_server_catalog
	ForeignServerName         SQLIdentifier    `yaml:"foreign_server_name,omitempty"`          // foreign_server_name
	ForeignDataWrapperCatalog SQLIdentifier    `yaml:"foreign_data_wrapper_catalog,omitempty"` // foreign_data_wrapper_catalog
	ForeignDataWrapperName    SQLIdentifier    `yaml:"foreign_data_wrapper_name,omitempty"`    // foreign_data_wrapper_name
	ForeignServerType         CharacterData    `yaml:"foreign_server_type,omitempty"`          // foreign_server_type
	ForeignServerVersion      CharacterData    `yaml:"foreign_server_version,omitempty"`       // foreign_server_version
	AuthorizationIdentifier   SQLIdentifier    `yaml:"authorization_identifier,omitempty"`     // authorization_identifier
}

// Constants defining each column in the table.
const (
	PgForeignServerOidField                       = "oid"
	PgForeignServerSrvoptionsField                = "srvoptions"
	PgForeignServerForeignServerCatalogField      = "foreign_server_catalog"
	PgForeignServerForeignServerNameField         = "foreign_server_name"
	PgForeignServerForeignDataWrapperCatalogField = "foreign_data_wrapper_catalog"
	PgForeignServerForeignDataWrapperNameField    = "foreign_data_wrapper_name"
	PgForeignServerForeignServerTypeField         = "foreign_server_type"
	PgForeignServerForeignServerVersionField      = "foreign_server_version"
	PgForeignServerAuthorizationIdentifierField   = "authorization_identifier"
)

// WhereClauses for every type in PgForeignServer.
var (
	PgForeignServerOidWhere                       OidField           = "oid"
	PgForeignServerSrvoptionsWhere                NullStringField    = "srvoptions"
	PgForeignServerForeignServerCatalogWhere      SQLIdentifierField = "foreign_server_catalog"
	PgForeignServerForeignServerNameWhere         SQLIdentifierField = "foreign_server_name"
	PgForeignServerForeignDataWrapperCatalogWhere SQLIdentifierField = "foreign_data_wrapper_catalog"
	PgForeignServerForeignDataWrapperNameWhere    SQLIdentifierField = "foreign_data_wrapper_name"
	PgForeignServerForeignServerTypeWhere         CharacterDataField = "foreign_server_type"
	PgForeignServerForeignServerVersionWhere      CharacterDataField = "foreign_server_version"
	PgForeignServerAuthorizationIdentifierWhere   SQLIdentifierField = "authorization_identifier"
)

// QueryOnePgForeignServer retrieves a row from 'information_schema._pg_foreign_servers' as a PgForeignServer.
func QueryOnePgForeignServer(db XODB, where WhereClause, order OrderBy) (*PgForeignServer, error) {
	const origsqlstr = `SELECT ` +
		`oid, srvoptions, foreign_server_catalog, foreign_server_name, foreign_data_wrapper_catalog, foreign_data_wrapper_name, foreign_server_type, foreign_server_version, authorization_identifier ` +
		`FROM information_schema._pg_foreign_servers WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	pfs := &PgForeignServer{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&pfs.Oid, &pfs.Srvoptions, &pfs.ForeignServerCatalog, &pfs.ForeignServerName, &pfs.ForeignDataWrapperCatalog, &pfs.ForeignDataWrapperName, &pfs.ForeignServerType, &pfs.ForeignServerVersion, &pfs.AuthorizationIdentifier)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return pfs, nil
}

// QueryPgForeignServer retrieves rows from 'information_schema._pg_foreign_servers' as a slice of PgForeignServer.
func QueryPgForeignServer(db XODB, where WhereClause, order OrderBy) ([]*PgForeignServer, error) {
	const origsqlstr = `SELECT ` +
		`oid, srvoptions, foreign_server_catalog, foreign_server_name, foreign_data_wrapper_catalog, foreign_data_wrapper_name, foreign_server_type, foreign_server_version, authorization_identifier ` +
		`FROM information_schema._pg_foreign_servers WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*PgForeignServer
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		pfs := PgForeignServer{}

		err = q.Scan(&pfs.Oid, &pfs.Srvoptions, &pfs.ForeignServerCatalog, &pfs.ForeignServerName, &pfs.ForeignDataWrapperCatalog, &pfs.ForeignDataWrapperName, &pfs.ForeignServerType, &pfs.ForeignServerVersion, &pfs.AuthorizationIdentifier)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &pfs)
	}
	return vals, nil
}

// AllPgForeignServer retrieves all rows from 'information_schema._pg_foreign_servers' as a slice of PgForeignServer.
func AllPgForeignServer(db XODB, order OrderBy) ([]*PgForeignServer, error) {
	const origsqlstr = `SELECT ` +
		`oid, srvoptions, foreign_server_catalog, foreign_server_name, foreign_data_wrapper_catalog, foreign_data_wrapper_name, foreign_server_type, foreign_server_version, authorization_identifier ` +
		`FROM information_schema._pg_foreign_servers`

	sqlstr := origsqlstr + order.String()

	var vals []*PgForeignServer
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		pfs := PgForeignServer{}

		err = q.Scan(&pfs.Oid, &pfs.Srvoptions, &pfs.ForeignServerCatalog, &pfs.ForeignServerName, &pfs.ForeignDataWrapperCatalog, &pfs.ForeignDataWrapperName, &pfs.ForeignServerType, &pfs.ForeignServerVersion, &pfs.AuthorizationIdentifier)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &pfs)
	}
	return vals, nil
}
