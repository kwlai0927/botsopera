package opera

import (
	"errors"
)

type Status int64

const (
	Rest Status = iota
	Operate
	Error
)

type dramaMeta struct {
	id    DramaID
	drama Drama
	chS   chan *Message
	chR   chan *Message
}

type botMeta struct {
	id  BotID
	bot Bot
	chS chan *BotMessage
	chR chan *BotMessage
}

type RoomID string

type room struct {
	id        RoomID
	botID     BotID
	botMeta   *botMeta
	dramaID   DramaID
	dramaMeta *dramaMeta
}

// TODO: get from diff gorountine protect
type Opera struct {
	rooms  map[RoomID]*room
	dramas map[DramaID]*dramaMeta
	bots   map[BotID]*botMeta
}

// -- Bot

func (o *Opera) RegisterBot(id BotID, bot Bot) error {

	_, ok := o.bots[id]
	if ok {
		return errors.New("Drama id already used")
	}

	chS := make(chan *BotMessage)
	chR := make(chan *BotMessage)

	o.bots[id] = &botMeta{
		id:  id,
		bot: bot,
		chS: chS,
		chR: chR,
	}

	go func() {
		for msg := range chR {
			room, ok := o.rooms[msg.RoomID]
			if ok {
				room.dramaMeta.chS <- msg.Message
			}
		}
	}()
	go bot.Register(chR, chS)

	return nil
}

func (o *Opera) UnregesterBot(id BotID) error {

	meta, ok := o.bots[id]
	if !ok {
		return errors.New("drama doesn't exist")
	}

	roomIDs := []RoomID{}
	for roomID, room := range o.rooms {
		if room.botID == id {
			roomIDs = append(roomIDs, roomID)
		}
	}

	for _, roomID := range roomIDs {
		delete(o.rooms, roomID)
	}

	close(meta.chR)
	close(meta.chS)
	delete(o.bots, id)

	return nil
}

// -- Drama

func (o *Opera) RegisterDrama(id DramaID, drama Drama) error {

	_, ok := o.dramas[id]
	if ok {
		return errors.New("Drama id already used")
	}

	chS := make(chan *Message)
	chR := make(chan *Message)

	o.dramas[id] = &dramaMeta{
		id:    id,
		drama: drama,
		chS:   chS,
		chR:   chR,
	}

	go func(dramaID DramaID) {
		for msg := range chR {
			for id, room := range o.rooms {
				if room.dramaMeta.id == dramaID {
					room.botMeta.chS <- &BotMessage{
						RoomID:  id,
						Message: msg,
					}
				}
			}
		}
	}(id)
	go drama.Register(chR, chS)

	return nil
}

func (o *Opera) UnregisterDrama(id DramaID) error {

	meta, ok := o.dramas[id]
	if !ok {
		return errors.New("drama doesn't exist")
	}

	for _, room := range o.rooms {
		if room.dramaID == id {
			room.dramaID = ""
			room.dramaMeta = nil
		}
	}

	close(meta.chR)
	close(meta.chS)
	delete(o.dramas, id)

	return nil
}

// -- Room

func (o *Opera) CreateRoom(id RoomID, byBotID BotID) error {

	_, ok := o.rooms[id]
	if ok {
		return errors.New("room id already used")
	}

	o.rooms[id] = &room{
		id:      id,
		botID:   byBotID,
		botMeta: o.bots[byBotID],
	}

	return nil
}

func (o *Opera) DeleteRoom(id RoomID) error {

	_, ok := o.rooms[id]
	if !ok {
		return errors.New("room id doesn't exist")
	}

	delete(o.rooms, id)

	return nil
}

func (o *Opera) PlayDramaForRooms(id DramaID, roomIDs []RoomID) error {

	meta := o.dramas[id]

	for _, roomID := range roomIDs {
		if room, ok := o.rooms[roomID]; ok {
			room.dramaID = id
			room.dramaMeta = meta
		}
	}

	return nil
}

func (o *Opera) CancelDramaForRooms(id DramaID, roomIDs []RoomID) {

	for _, roomID := range roomIDs {
		if room, ok := o.rooms[roomID]; ok && room.dramaID == id {
			room.dramaID = ""
			room.dramaMeta = nil
		}
	}

}
