package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		//create cors header
		utils.SetCorsHeaders(&w)

		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_command := vars["id_command"]

		//delete command lines with id_command linked
		sqlQuery := ("DELETE FROM command_line WHERE id_command = " + id_command + ";")

		//execute the sql query
		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//create the sql query
		sqlQuery = ("DELETE FROM command WHERE id = " + id_command + ";")

		//execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.JsonResponseCommand{
			Type:    "success",
			Message: "Command deleted",
		})
	}
}
