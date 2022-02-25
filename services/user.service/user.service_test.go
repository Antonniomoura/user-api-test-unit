package user_service_test

import (
	"awesomeProject/models"
	"awesomeProject/repositories/user.repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	userId = oid.Hex()
	user := models.User{
		ID:       oid,
		Name:     "Antonio",
		Email:    "antonio@gmail.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := user_repository.Create(user)
	if err != nil {
		t.Errorf("error create")
		t.Fail()
		fmt.Printf(err.Error())
		return
	}
	t.Log("Finalizou com sucesso")
}

func TestRead(t *testing.T) {
	users, err := user_repository.Read()
	if err != nil {
		t.Error("Error")
		t.Fail()
		return
	}
	if len(users) == 0 {
		t.Error("Sem data")
		t.Fail()
		return
	}

	t.Log("sucesso")
	fmt.Println(users)
}

func TestReadById(t *testing.T) {
	user, err := user_repository.ReadById("62181fda334eb171461bb6f0")
	if err != nil {
		t.Error("Error one")
		t.Fail()
		return
	}

	t.Log("sucesso one")
	fmt.Println(user)
}

func TestDelete(t *testing.T) {

	err := user_repository.Delete(userId)

	if err != nil {
		t.Error("Error no delete")
		t.Fail()
		return
	}
	t.Log("Sucessp")
}

func TestUpdate(t *testing.T) {
	user := models.User{
		Name:  "Nopo",
		Email: "email.cqw@as.com",
	}
	err := user_repository.Update(user, userId)
	if err != nil {
		t.Error("Error ao tratar")
		t.Fail()
	}
	t.Log("Sucesso")
}
