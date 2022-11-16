package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetBalance(userId string) *UserBalance {
	db, err := sql.Open("mysql", "root:1111@/balance_schema")
	balance := &UserBalance{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM balance where id=?", userId)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&balance.UserId, &balance.Balance)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return balance
}

func AddBalance(balance UserBalance) {

	db, err := sql.Open("mysql", "root:1111@/balance_schema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//здесь будет проверка на то, есть ли уже такой пользователь
	insert, err := db.Query(
		"INSERT INTO balance (user_id,balance) VALUES (?,?)",
		balance.UserId, balance.Balance)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
