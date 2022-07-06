package repo

// import (
// 	"GolangMongoDbTest/database"
// 	"GolangMongoDbTest/model"
// 	"GolangMongoDbTest/util"
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// type UserRepoInterface interface {
// 	UserRepoRegister(ctx context.Context, users *model.User) (*model.User, error)
// 	UserRepoUpdate(ctx context.Context, users *model.User, id string) (*model.User, error)
// 	UserRepoDelete(ctx context.Context, users *model.User, id string) (*model.User, error)
// 	UserRepoGetAll() ([]*model.UserGetAll, error)
// }

// type UserRepo struct {
// 	db *sql.DB
// }

// func NewUserRepo(db *sql.DB) UserRepoInterface {
// 	return &UserRepo{db: db}
// }

// func (u *UserRepo) UserRepoRegister(ctx context.Context, users *model.User) (*model.User, error) {
// 	pass, errHash := util.GenerateHashPassword(users.Password)
// 	if errHash != nil {
// 		fmt.Println("Error Hash : " + errHash.Error())
// 		return nil, errHash
// 	}
// 	users.Password = pass

// 	db, err := database.Connect()
// 	if err != nil {
// 		fmt.Println("connect db error :", err)
// 	}
// 	defer db.Close()

// 	sqlSt := `insert into users (username, email, password)values (?, ?, ?);`
// 	_, err = db.Exec(sqlSt, users.Username, users.Email, users.Password)
// 	if err != nil {
// 		fmt.Println("Query row error :", err)
// 	}
// 	defer db.Close()
// 	fmt.Println("repo user_id:", users.Idusers)

// 	return users, nil
// }

// func (u *UserRepo) UserRepoGetAll() ([]*model.UserGetAll, error) {
// 	var sliceUser []*model.UserGetAll
// 	db, err := database.Connect()
// 	if err != nil {
// 		fmt.Println("connect db error :", err)
// 	}
// 	defer db.Close()
// 	sqlSt := `Select id, username, email from users;`

// 	rows, err := db.Query(sqlSt)
// 	if err != nil {
// 		fmt.Println("query error :", err)
// 	}
// 	for rows.Next() {
// 		var user model.UserGetAll
// 		err = rows.Scan(&user.Idusers, &user.Username, &user.Email)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			sliceUser = append(sliceUser, &user)
// 		}
// 	}
// 	return sliceUser, nil
// }

// func (u *UserRepo) UserRepoUpdate(ctx context.Context, users *model.User, id string) (*model.User, error) {
// 	pass, errHash := util.GenerateHashPassword(users.Password)
// 	if errHash != nil {
// 		fmt.Println("Error Hash : " + errHash.Error())
// 		return nil, errHash
// 	}
// 	users.Password = pass

// 	db, err := database.Connect()
// 	if err != nil {
// 		fmt.Println("connect db error :", err)
// 	}
// 	defer db.Close()

// 	sqlSt := `update users set username = ?, email= ? where id = ?;`
// 	_, err = db.Exec(sqlSt, users.Username, users.Email, id)
// 	if err != nil {
// 		fmt.Println("Query row error :", err)
// 	}
// 	defer db.Close()

// 	sqlSt2 := `select username, email from users where id = ?;`
// 	err2 := db.QueryRow(sqlSt2, id).Scan(
// 		&users.Username,
// 		&users.Email,
// 	)
// 	if err2 != nil {
// 		fmt.Println("Query row error :", err)
// 	}
// 	defer db.Close()

// 	fmt.Println("repo user_id:", users.Idusers)

// 	return users, nil
// }

// func (u *UserRepo) UserRepoDelete(ctx context.Context, users *model.User, id string) (*model.User, error) {
// 	sqlSt := `delete from users where id = ?`
// 	db, err := database.Connect()
// 	if err != nil {
// 		fmt.Println("connect db error :", err)
// 	}
// 	defer db.Close()

// 	_, err = db.Exec(sqlSt, id)
// 	if err != nil {
// 		fmt.Println("query error :", err)
// 	}
// 	return users, nil
// }
