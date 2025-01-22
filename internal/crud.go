package internal

import (
	"baseCrud/infrastructure/connectors"
	"baseCrud/internal/entity"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var database *sql.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {
	database := connectors.Connection()
	rows, err := database.Query("select * from base_crud_bd.users")

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	users := []entity.User{}

	for rows.Next() {
		p := entity.User{}
		err := rows.Scan(&p.Id, &p.Email, &p.UserName, &p.Password, &p.PhoneNumber, &p.DateOfBirth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	tmpl, _ := template.ParseFiles("../templates/users_table.html")
	tmpl.Execute(w, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	database := connectors.Connection()
	params := mux.Vars(r)

	rows, err := database.Query("select * from base_crud_bd.users where id = ?;", params["id"])
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	users := []entity.User{}

	for rows.Next() {
		p := entity.User{}
		err := rows.Scan(&p.Id, &p.Email, &p.UserName, &p.Password, &p.PhoneNumber, &p.DateOfBirth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	tmpl, _ := template.ParseFiles("../templates/users_table.html")
	tmpl.Execute(w, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	database := connectors.Connection()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user := entity.User{
			Email:       r.FormValue("email"),
			UserName:    r.FormValue("username"),
			Password:    r.FormValue("password"),
			PhoneNumber: r.FormValue("phonenumber"),
			DateOfBirth: r.FormValue("dateofbirth"),
		}

		_, err = database.Exec("insert into base_crud_bd.users (email, user_name, password, phone_number, date_of_birth)  values (?, ?, ?, ?, ?)",
			user.Email,
			user.UserName,
			user.Password,
			user.PhoneNumber,
			user.DateOfBirth)

		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/users", 301)

	} else {
		http.ServeFile(w, r, "../templates/create_users.html")
	}

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	database := connectors.Connection()
	params := mux.Vars(r)

	_, err := database.Exec("delete from base_crud_bd.users where id = ?;", params["id"])
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/users", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	database := connectors.Connection()
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.QueryRow("select * from base_crud_bd.users where id = ?", id)
	user := entity.User{}
	err := row.Scan(&user.Id, &user.Email, &user.UserName, &user.Password, &user.PhoneNumber, &user.DateOfBirth)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("../templates/update_userInfo.html")
		tmpl.Execute(w, user)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	database := connectors.Connection()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		params := mux.Vars(r)

		newPhoneNumber := r.FormValue("phonenumber")
		_, err = database.Exec("update base_crud_bd.users set phone_number = ? where id = ?;", newPhoneNumber, params["id"])

		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/users", 301)

	}
}
