package pubsub

import (
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
)
var  (
	 PUBLISH = "publish"
	 SUBSCRIBE = "subscribe"
)

type PubSub struct {
	clients []Client
}

type Client struct {
	Id         string
	Connection *websocket.Conn
}

func (p *PubSub) AddClient(client *Client) *PubSub {
	p.clients = append(p.clients, *client)
	
	//server send message to client when adding to list
	msg := []byte("Hello client id: " + client.Id)
	client.Connection.WriteMessage(1, msg)
	//fmt.Println("adding new client to list", client.Id)
	return p
}

type Message struct {
	Action string `json:action`
	Topic string `json:topic`
	Message json.RawMessage `json:"message"`
}

func (p* PubSub) HandleReceiveMessage(client *Client, messageType int, payload []byte) (* PubSub) {
	m := &Message{}
	err := json.Unmarshal(payload, &m)
    if err != nil {
		fmt.Println("This is incorrect message payload.")
	    return p
	}

	switch m.Action {

	case PUBLISH:
		fmt.Println("This is a publish new message")
		break
	case SUBSCRIBE: 
    	fmt.Println("This is a subscribe new message")
	     break
	default:
		break

	}
	
	return p
}