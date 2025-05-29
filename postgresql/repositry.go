package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func CheckUser(db *sql.DB, username string) bool {
	q := `
	SELECT *
	FROM users
	WHERE user_name = $1
	`
	var foundUsername string
	err := db.QueryRow(q, username).Scan(&foundUsername)
	if err != nil {
		if err == sql.ErrNoRows{
			return false
		}
		log.Printf("Error checking user %s: %v", username, err)
		return false
	}
	return true
}

func InsertUser(db *sql.DB, username string){
	// query
	q := `
	INSERT INTO users (
	user_name
	)
	VALUES (
	$1
	);
	`
	_, err := db.Exec(q, username)// 結果がいらないときはExec。返り値は影響の受けた情報
	if err != nil {
		// 強制終了
		log.Fatal(err)
	}
}

func DeleteUser(db *sql.DB, username string){
	q := `
	DELETE FROM users
	WHERE user_name = $1
	`
	_, err := db.Exec(q, username)
	if err != nil {
		// 強制終了
		log.Fatal(err)
	}
}

func UpsertDiary(db *sql.DB, user_name,date,content string){
	q := `
	INSERT INTO diary (
	user_name,
	diary_date,
	diary_content
	)
	VALUES (
	$1,
	$2,
	$3
	)
	ON CONFLICT (user_name, diary_date) DO UPDATE SET
	diary_content = EXCLUDED.diary_content;
	`
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("時間のレイアウトが間違っています")
		log.Fatal(err)
	}
	_, err = db.Exec(q,user_name,parsedTime,content)
	if err != nil {
		log.Fatal(err)
	}
}

func ShowDiary(db *sql.DB, user_name,input string){
	var q string
	date, ok := date(input)
	if input == "all" {
		// SELECT * にしてしまうとjunctionの列も取得してしまう
		q = `
		SELECT diary_date, diary_content
		FROM diary
		WHERE user_name = $1;
		`
		rows, err := db.Query(q,user_name)
		if err != nil {
			log.Fatal(err)
		}
		PrintShow(rows)
	}else if ok {
		// user_nameとdiary_dateが一致するdiaryを取得
		q = `
		SELECT diary_date, diary_content
		FROM diary
		WHERE user_name =$1
		AND diary_date = $2;
		`
		rows, err := db.Query(q,user_name,date)
		if err != nil {
			log.Fatal(err)
		}
		PrintShow(rows)
	}else {
		q = `
		SELECT diary_date, diary_content
		FROM diary
		WHERE diary_content LIKE $1
		AND user_name = $2;
		`
		rows, err := db.Query(q,"%"+input+"%",user_name)
		if err != nil {
			log.Fatal(err)
		}
		PrintShow(rows)
	}
}

func date(date string) (time.Time, bool) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, false
	}
		return parsedDate, true
}

func PrintShow(rows *sql.Rows){
	type diary struct{
		date time.Time
		content string
	}
	var diaries []*diary
	defer rows.Close()
	for rows.Next() {
		var date time.Time
		var content string
		if err := rows.Scan(&date, &content); err != nil{
			log.Fatal(err)
		}
		diaries = append(diaries, &diary{date, content})
	}
	if len(diaries) == 0 {
		fmt.Println("表示する日記がありません。") // データがない場合のメッセージ
		return
	}
	for _,v := range diaries{
		fmt.Println("date ", v.date.Format("2006-01-02"), "content ", v.content)
	}
}

func DeletePieceOfDiary(db *sql.DB, user_name,date string){
	q := `
	DELETE FROM diary
	WHERE diary_date = $1
	AND user_name = $2
	`
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)	
	if err != nil {
		log.Fatal(err)
	} 
	_, err = db.Exec(q, parsedDate, user_name)
	if err != nil {
		// 強制終了
		log.Fatal(err)
	}
}