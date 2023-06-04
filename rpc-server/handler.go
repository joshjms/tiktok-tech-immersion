package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	_ "github.com/lib/pq"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

func ConnectDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println("Failed to connect to PostgreSQL database!")
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL database!")

	return db, nil
}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	resp := rpc.NewSendResponse()

	db, err := ConnectDatabase()
	if err != nil {
		resp.Code = 500
		return resp, err
	}

	chat := req.Message.Chat
	text := req.Message.Text
	sender := req.Message.Sender
	send_time := time.Now().Unix()

	_, err = db.Exec(fmt.Sprintf("INSERT INTO messages (chat, text, sender, send_time) VALUES ('%v', '%v', '%v', '%v')", chat, text, sender, send_time))
	if err != nil {
		resp.Code = 500
		return resp, err
	}

	return resp, err
}

func reverseMessages(messages []*rpc.Message) []*rpc.Message {
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	resp := rpc.NewPullResponse()

	db, err := ConnectDatabase()
	if err != nil {
		resp.Code = 500
		return resp, err
	}

	chat := req.Chat
	cursor := req.Cursor
	limit := req.Limit
	reverse := *req.Reverse

	query := fmt.Sprintf("SELECT * FROM messages WHERE chat = '%v' OFFSET %v LIMIT %v", chat, cursor, limit+1)

	rows, err := db.Query(query)
	if err != nil {
		resp.Code = 500
		return resp, err
	}

	defer rows.Close()

	var data []*rpc.Message
	for rows.Next() {
		var id int
		var chat string
		var text string
		var sender string
		var send_time int64

		err = rows.Scan(&id, &chat, &text, &sender, &send_time)
		if err != nil {
			resp.Code = 500
			return resp, err
		}

		data = append(data, &rpc.Message{
			Chat:     chat,
			Text:     text,
			Sender:   sender,
			SendTime: send_time,
		})
	}

	var nextCursor int64 = 0
	var hasMore bool = false

	if len(data) > int(limit) {
		nextCursor = cursor + int64(limit)
		hasMore = true
		data = data[:len(data)-1]
	}

	if reverse {
		data = reverseMessages(data)
	}

	resp.Code = 0
	resp.Messages = data
	resp.NextCursor = &nextCursor
	resp.HasMore = &hasMore

	return resp, nil
}
