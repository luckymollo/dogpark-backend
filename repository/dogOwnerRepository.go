package repository

import (
	"log"
)

func FindDogOwnerById(id int64) DogOwner {
	findByIdStmt, err := db.Prepare("select name, age, email from dog_owner where id = ?1 ")

	if err != nil {
		log.Panic(err)
	}
	defer findByIdStmt.Close()
	var name, email string
	var age int
	findByIdStmt.QueryRow(id).Scan(&name, &age, &email)

	return DogOwner{Id: id, Name: name, Age: age, Email: email}

}

func SaveDownOwner(owner DogOwner) int64 {
	insertStmt, err := db.Prepare("insert into dog_owner (name, age, email) values ($1, $2, $3 )")
	if err != nil {
		log.Panic(err)
	}

	res, err := insertStmt.Exec(owner.Name, owner.Age, owner.Email)

	if err != nil {
		log.Panic(err)
	}

	id, err := res.LastInsertId()
	return id

}

func UpdateDogOwner(owner DogOwner) {
	updateStmt, err := db.Prepare("update dog_owner set name = $1, age = $2, email = $3 where id = $4")

	if err != nil {
		log.Panic(err)
	}

	updateStmt.Exec(owner.Name, owner.Age, owner.Email, owner.Id)
}
