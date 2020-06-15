# GoRocketChatRestSDK
## It's my implementation of RocketChat rest API. This is my first project on golang. coming soon...
A few facts about packet
 1. I was trying to tell methods same as official docs API
 2. All option parameters have names starts with the same method name, like this
 func (r *RocketChat) GroupCreate(channelName string, optional *GroupCreateSettings) string {
# rocketchat
--
    import "whatsapp/rocketchat"


## Usage

#### type GroupCreateSettings

```go
type GroupCreateSettings struct {
        Name     string   `json:"name"`     // The name of the new private group
        Members  []string `json:"members"`  // The users to add to the group when it is created.
        ReadOnly bool     `json:"readOnly"` //
}
```


#### type GroupDeleteSettings

```go
type GroupDeleteSettings struct {
        RoomId   string `json:"roomId"`
        RoomName string `json:"roomName"`
}
```


#### type Message

```go
type Message struct {
        ID string `json:"_id"`
        //rid   :       wojkjJkSQMFRCNngq
        Msg       string `json:"msg"`
        Ts        string `json:"ts"`
        ReplyID   string `json:"tmid"`
        User      User   `json:"u"`
        UpdatedAt string `json:"_updatedAt"`
}
```


#### type Messages

```go
type Messages struct {
        Messages []Message `json:"messages"`
}
```


#### type MethodResult

```go
type MethodResult struct {
        Success bool `json:"success"`
}
```


#### type MongoStore

```go
type MongoStore struct {
}
```


#### func  NewMongoStorage

```go
func NewMongoStorage(url string) *MongoStore
```

#### func (*MongoStore) Connect

```go
func (m *MongoStore) Connect()
```

#### func (*MongoStore) MarkMessageAsRead

```go
func (m *MongoStore) MarkMessageAsRead(messageId string)
```

#### func (*MongoStore) MessageIsRead

```go
func (m *MongoStore) MessageIsRead(messageId string) bool
```

#### type RocketChat

```go
type RocketChat struct {
        Cancel context.CancelFunc
}
```


#### func  CreateRocketChat

```go
func CreateRocketChat(host, email, password string) *RocketChat
```

#### func (*RocketChat) ChatSendMessage

```go
func (r *RocketChat) ChatSendMessage(to, text string)
```
ChatSendMessage - send message to one channel

#### func (*RocketChat) GroupCreate

```go
func (r *RocketChat) GroupCreate(options *GroupCreateSettings) *MethodResult
```
GroupCreate - Create a private channel.

#### func (*RocketChat) GroupDelete

```go
func (r *RocketChat) GroupDelete(options *GroupDeleteSettings) *MethodResult
```
GroupDelete Remove a private channel.

#### func (*RocketChat) GroupHistory

```go
func (r *RocketChat) GroupHistory() *Messages
```
GroupHistory Retrieves the messages from a private group, only if you're part of
the group.

#### func (*RocketChat) GroupList

```go
func (r *RocketChat) GroupList()
```
GroupList - Lists all of the private groups the calling user has joined.

#### type Storage

```go
type Storage interface {
        MarkMessageAsRead(messageId string)
        MessageIsRead(messageId string) bool
}
```


#### type User

```go
type User struct {
        ID       string `json:"userId"`
        Token    string `json:"authToken"`
        Username string `json:"username"`
}
```
