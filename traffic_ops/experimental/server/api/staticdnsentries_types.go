
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type StaticdnsentriesTypes struct {
	Name        string                     `db:"name" json:"name"`
	Description string                     `db:"description" json:"description"`
	CreatedAt   time.Time                  `db:"created_at" json:"createdAt"`
	Links       StaticdnsentriesTypesLinks `json:"_links" db:-`
}

type StaticdnsentriesTypesLinks struct {
	Self string `db:"self" json:"_self"`
}

type StaticdnsentriesTypesLink struct {
	ID  string `db:"staticdnsentries_type" json:"name"`
	Ref string `db:"staticdnsentries_types_name_ref" json:"_ref"`
}

// @Title getStaticdnsentriesTypesById
// @Description retrieves the staticdnsentries_types information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    StaticdnsentriesTypes
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentries_types/{id} [get]
func getStaticdnsentriesType(name string, db *sqlx.DB) (interface{}, error) {
	ret := []StaticdnsentriesTypes{}
	arg := StaticdnsentriesTypes{}
	arg.Name = name
	queryStr := "select *, concat('" + API_PATH + "staticdnsentries_types/', name) as self"
	queryStr += " from staticdnsentries_types WHERE name=:name"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getStaticdnsentriesTypess
// @Description retrieves the staticdnsentries_types
// @Accept  application/json
// @Success 200 {array}    StaticdnsentriesTypes
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentries_types [get]
func getStaticdnsentriesTypes(db *sqlx.DB) (interface{}, error) {
	ret := []StaticdnsentriesTypes{}
	queryStr := "select *, concat('" + API_PATH + "staticdnsentries_types/', name) as self"
	queryStr += " from staticdnsentries_types"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postStaticdnsentriesTypes
// @Description enter a new staticdnsentries_types
// @Accept  application/json
// @Param                 Body body     StaticdnsentriesTypes   true "StaticdnsentriesTypes object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentries_types [post]
func postStaticdnsentriesType(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v StaticdnsentriesTypes
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO staticdnsentries_types("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ",created_at"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ",:created_at"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putStaticdnsentriesTypes
// @Description modify an existing staticdnsentries_typesentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     StaticdnsentriesTypes   true "StaticdnsentriesTypes object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentries_types/{id}  [put]
func putStaticdnsentriesType(name string, payload []byte, db *sqlx.DB) (interface{}, error) {
	var arg StaticdnsentriesTypes
	err := json.Unmarshal(payload, &arg)
	arg.Name = name
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE staticdnsentries_types SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",created_at = :created_at"
	sqlString += " WHERE name=:name"
	result, err := db.NamedExec(sqlString, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delStaticdnsentriesTypesById
// @Description deletes staticdnsentries_types information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    StaticdnsentriesTypes
// @Resource /api/2.0
// @Router /api/2.0/staticdnsentries_types/{id} [delete]
func delStaticdnsentriesType(name string, db *sqlx.DB) (interface{}, error) {
	arg := StaticdnsentriesTypes{}
	arg.Name = name
	result, err := db.NamedExec("DELETE FROM staticdnsentries_types WHERE name=:name", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
