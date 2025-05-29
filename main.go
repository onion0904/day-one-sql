package main

import (
	"bufio"
	"log"
	"strings"
	"os"
	"fmt"
	"postgresql/postgresql"
	"postgresql/postgresql/migrations"
)

func main() {
	db,err :=postgresql.Connect()
	if err != nil {
		log.Fatal(err)
	}
	migrations.MigrateUp(db)

	fmt.Print("input UserOperation: ")
	// 標準入力からの読み込み
	input := GetInput()
	input = strings.TrimSpace(input)
	UserOperation := strings.Fields(input)

	if UserOperation[0] == "delete" {
		postgresql.DeleteUser(db, UserOperation[1])
		return
	} else if !postgresql.CheckUser(db,UserOperation[0]) {
		postgresql.InsertUser(db,UserOperation[0])
	}

	fmt.Print("input operation: ")
	// 標準入力からの読み込み
	input = GetInput()
	operation := strings.TrimSpace(input)

	if operation == "upsert" {
		fmt.Print("input date and content: ")
		input = GetInput()
		input = strings.TrimSpace(input)
		parts := strings.SplitN(input, " ", 2)
		if len(parts) < 2 {
			fmt.Println("入力形式が正しくありません。")
			return
		}
		postgresql.UpsertDiary(db,UserOperation[0],parts[0],parts[1])
	} else if operation == "show" {
		fmt.Print("input ShowOperation: ")
		input = GetInput()
		input = strings.TrimSpace(input)
		postgresql.ShowDiary(db,UserOperation[0],input)
	} else if operation == "delete" {
		fmt.Print("input date: ")
		input = GetInput()
		input = strings.TrimSpace(input)
		postgresql.DeletePieceOfDiary(db,UserOperation[0],input)
	} else {
		fmt.Println("入力が正しくありません")
		return
	}
	return
}

func GetInput () string { 
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}