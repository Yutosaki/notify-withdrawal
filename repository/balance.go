package repository

import (
	"database/sql"
	"fmt"
)

func ChangeBalance(db *sql.DB) int {
	var balance int
	fmt.Print("貯金残高を入力してください: ")
	fmt.Scan(&balance)

	_, err := db.Exec("INSERT INTO balance (amount) VALUES ($1)", balance)
	if err != nil {
		fmt.Println("残高の保存に失敗しました:", err)
	}
	return balance
}

func GetBalance(db *sql.DB) int {
	var balance int
	err := db.QueryRow("SELECT amount FROM balance ORDER BY id DESC LIMIT 1").Scan(&balance)
	if err != nil {
		fmt.Println("残高の取得に失敗しました:", err)
	}
	return balance
}

func UpdateBalance(db *sql.DB, newBalance int) {
	_, err := db.Exec("INSERT INTO balance (amount) VALUES ($1)", newBalance)
	if err != nil {
		fmt.Println("残高の更新に失敗しました:", err)
	}
}
