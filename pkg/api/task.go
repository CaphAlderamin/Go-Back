package api

import (
	. "Rip/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type WsMessageTasks struct {
	Type  string `json:"type"`
	Tasks []Task `json:"tasks"`
}

//type WsMessageTask struct{
//	Type	string	`json:"type"`
//	Task	Task 	`json:"task"`
//}

type WsMessageTask struct {
	Type string `json:"type"`
	Task Task   `json:"data"`
}

type ConnectUser struct {
	WebSocket *websocket.Conn `json:"webSocket"`
	UserIp    string          `json:"userIp"`
}

var users []ConnectUser

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketMessageReceiver(c *gin.Context) {
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	log.Println("Client connected:", conn.RemoteAddr().String())
	var socketClient = ConnectUser{conn, conn.RemoteAddr().String()}
	users = append(users, socketClient)
	log.Println("Number of connected users: ", len(users))

	GetTasks(socketClient)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket '", socketClient.UserIp, "' is not answering. Disconnecting...\n", err.Error())
			for i, v := range users {
				if v.UserIp == socketClient.UserIp {
					users = RemoveUser(users, i)
				}
			}
			log.Println("New number of connected users: ", len(users))
			return
		}

		var msg *WsMessageTask
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Print("Decode error")
		}
		switch msg.Type {
		//--------------------------------------------------------//
		case "GetIndex":
			text, err := json.Marshal("It works")

			err = conn.WriteMessage(messageType, text)
			if err != nil {
				return
			}
			break
			//--------------------------------------------------------//
		case "GetTasks":
			task := Task{}
			tasks, err := task.GetTasks()
			if err != nil {
				log.Fatalln(err)
			}

			marshaledTasks, err := json.Marshal(tasks)
			if err != nil {
				log.Fatalln(err)
			}

			err = conn.WriteMessage(messageType, marshaledTasks)
			if err != nil {
				return
			}
			break
			//--------------------------------------------------------//
		case "PostTask":
			ra, err := msg.Task.AddTask()
			if err != nil {
				log.Fatalln(err)
			}
			msg.Task.Id = int(ra)

			var messageToUser = WsMessageTask{
				Type: "PostTask",
				Task: msg.Task,
			}

			msgToUser, err := json.Marshal(messageToUser)
			if err != nil {
				log.Fatalln(err)
			}

			WriteMessageToEveryUser(messageType, msgToUser)
			break
			//--------------------------------------------------------//
		case "ModTaskStatus":
			task, err := msg.Task.ModTaskStatus()
			if err != nil {
				log.Fatalln(err)
			}

			var messageToUser = WsMessageTask{
				Type: "ModTaskStatus",
				Task: task,
			}

			msgToUser, err := json.Marshal(messageToUser)
			if err != nil {
				log.Fatalln(err)
			}

			WriteMessageToEveryUser(messageType, msgToUser)
			break
			//--------------------------------------------------------//
		case "ModTask":
			task, err := msg.Task.ModTask()
			if err != nil {
				log.Fatalln(err)
			}

			var messageToUser = WsMessageTask{
				Type: "ModTask",
				Task: task,
			}

			msgToUser, err := json.Marshal(messageToUser)
			if err != nil {
				log.Fatalln(err)
			}

			WriteMessageToEveryUser(messageType, msgToUser)
			break
			//--------------------------------------------------------//
		case "DelTask":
			id, err := msg.Task.DelTask()
			if err != nil {
				log.Fatalln(err)
			}

			task := Task{Id: id}

			var messageToUser = WsMessageTask{
				Type: "DelTask",
				Task: task,
			}

			msgToUser, err := json.Marshal(messageToUser)
			if err != nil {
				log.Fatalln(err)
			}

			WriteMessageToEveryUser(messageType, msgToUser)
			break
			//--------------------------------------------------------//
		}
	}
}

func RemoveUser(s []ConnectUser, index int) []ConnectUser {
	return append(s[:index], s[index+1:]...)
}

func WriteMessageToEveryUser(messageToUser int, msgToUser []byte) {
	for _, user := range users {
		err := user.WebSocket.WriteMessage(messageToUser, msgToUser)
		if err != nil {
			log.Println("Failed to send message to ", user.UserIp, err.Error())
		}
	}
}

func GetTasks(c ConnectUser) {
	task := Task{}
	tasks, err := task.GetTasks()
	if err != nil {
		log.Fatalln(err)
	}

	var messageToUser = WsMessageTasks{
		Type:  "GetTasks",
		Tasks: tasks,
	}

	msgToUser, err := json.Marshal(messageToUser)

	err = c.WebSocket.WriteMessage(1, msgToUser)
}
