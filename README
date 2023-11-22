Chat bot server demo using Go + Websocket

## Overview

The chat application is built using Go and uses WebSockets for real-time communication.


#### 👨‍💻 Full list what has been used(currently):
[PostgeSQL](https://github.com/jackc/pgx) For user information<br/>
[Gorilla Websocket](https://github.com/gorilla/websocket) Handle websocket connection<br/>
[Gin](https://github.com/gin-gonic/gin) Web framework<br/>

## Structure

The main components of the application are:

- `Client`: Represents a single chat client. A client has a connection and can send and receive messages.
- `ChatRoom`: Represents a chat room. A chat room has multiple clients and broadcasts messages to all clients.
- `Chat`: Represents a chat message. A message has a sender and a content.
- `User`: Represents a user. A user has a username, email, and password.

## Setup

To run the application, you need to have Go installed on your machine. Then, you can clone this repository and run the main file.

```bash
git clone https://github.com/yourusername/chat-application.git
cd chat-application
go build -v -o chat-server .
./chat-server
```