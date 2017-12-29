package <%- packagename%>

import (
	"github.com/jmoiron/sqlx"
)

//Data ...
type repository struct {
	DB *sqlx.DB
}

// Create Repository
func NewRepository(db *sqlx.DB) Repository {
	return repository{DB: db}
}

func (r repository) Get(ID string) (<%- packagename%>Response, error) {

	//you can add your logic for database
	//r.DB.Get(&your_db_object,"your_sql_query",your_params1, ...)

	return <%- packagename%>Response{ ID }, nil
}
