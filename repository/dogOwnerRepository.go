package repository

import "log"

func FindById(id int) DogOwner {
	findByIdStmt, err := db.Prepare("select name, age, email from dog_owner where id = $1")
	if err != nil {
		log.Panic(err)
	}
	var name, email string
	var age int
	findByIdStmt.QueryRow(id).Scan(&name, &age, &email)

	return DogOwner{Name: name, Age: age, Email: email}

}
