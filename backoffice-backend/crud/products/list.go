package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"math"
	"strconv"

	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {

		if utils.Securized(w, r) {
			//global vars
			var products []structures.GetProducts

			//connect the database
			db := utils.DbConnect()

			//get filters values and update the sqlQuery
			orderBy, rangeBy := utils.GetAllParams(r, "product")
			sqlQuery := "SELECT product.id, product.name, product.price, product.description, product.pending_validation, product.is_alive, category.name, user.first_name, user.last_name, role.role, user.id AS id_user, AVG(rate.stars_number) AS rate FROM product INNER JOIN category ON category.id = product.id_category INNER JOIN user ON product.id_user = user.id INNER JOIN role ON role.id = user.id_role LEFT JOIN rate ON rate.id_product = product.id GROUP BY product.id " + orderBy + " " + rangeBy

			//execute the sql query and check errors
			rows, err := db.Query(sqlQuery)
			utils.CheckErr(err)

			//parse the query
			for rows.Next() {
				//global vars
				var name, description, category, firstname, lastname, role string
				var id, price, id_user int
				var rate float64
				var pending_validation, is_alive bool
				var s sql.NullString
				//retrieve the values and check errors
				err = rows.Scan(&id, &name, &price, &description, &pending_validation, &is_alive, &category, &firstname, &lastname, &role, &id_user, &s)
				utils.CheckErr(err)
				if s.Valid {
					rate, err = strconv.ParseFloat(s.String, 64)
					utils.CheckErr(err)
				} else {
					rate = 0
				}

				//round the rate
				rate = math.Round(rate*10) / 10

				//add the values to the response
				products = append(products, structures.GetProducts{
					Id:                 id,
					Name:               name,
					Price:              price,
					Description:        description,
					Pending_validation: pending_validation,
					Is_alive:           is_alive,
					Category:           category,
					Firstname:          firstname,
					Lastname:           lastname,
					Role:               role,
					Id_user:            id_user,
					Rate:               rate,
				})
			}
			//close the rows
			rows.Close()

			//close the database connection
			db.Close()

			//create the json response
			utils.SetXTotalCountHeader(&w, len(products))
			json.NewEncoder(w).Encode(products)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
