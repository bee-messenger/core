package repo

import (
	"github.com/bee-messenger/core/store"
	"github.com/bee-messenger/core/entity"
)




type Repo struct {
	store store.Store
}

func NewChatRepo(store store.Store) (*Repo, error) {
	return &Repo{
		store: store,
	}, nil
}

func (c Repo) GetAllChat() ([]entity.ChatInfo, error) {
	chl, err := c.store.ChatList()
	ci := make([]entity.ChatInfo,0)
	if err != nil {
		return nil, err
	}
	for _, val := range chl {
		members := make([]entity.Contact, 0)
		m, _ := c.store.ContactByIDs(val.Members)
		for _, me := range m {
			members = append(members, entity.Contact{
				ID: me.ID,
				Name: me.Name,
			})
		}
		ci = append(ci, entity.ChatInfo{
			ID: val.ID,
			Name: val.Name,
			Members: members,
		})
	}
	return ci, nil
}

func (c Repo) GetByIDChat(id string) (*entity.Chat, error) {
	ct, err := c.store.ChatByID(id)
	if err != nil {
		return nil, err
	}
	members := make([]entity.Contact, 0)
	m, _ := c.store.ContactByIDs(ct.Members)
	for _, me := range m {
		members = append(members, entity.Contact{
			ID: me.ID,
			Name: me.Name,
		})
	}
	messages := make([]entity.Message, 0)
	bhm, err := c.store.ChatMessages(id)
	if err != nil {
		return nil, err
	}
	for _, m := range bhm {
		messages = append(messages, entity.Message{
			ID: m.ID,
			CreatedAt: m.CreatedAt,
			Text: m.Text,
			Status: entity.Status(m.Status),
			Author: entity.Contact{
				ID: m.Author.ID,
				Name: m.Author.Name,
			},
		})
	}
	return &entity.Chat{
		Info: entity.ChatInfo{
			ID: ct.ID,
			Name: ct.Name,
			Members: members,
		},
		Messages: messages,
	}, nil
}

func(c Repo) GetChatInfoByID(id string) (*entity.ChatInfo, error) {
	ct, err := c.store.ChatByID(id)
	if err != nil {
		return nil, err
	}
	members := make([]entity.Contact, 0)
	m, _ := c.store.ContactByIDs(ct.Members)
	for _, me := range m {
		members = append(members, entity.Contact{
			ID: me.ID,
			Name: me.Name,
		})
	}
	return &entity.ChatInfo{
			ID: ct.ID,
			Name: ct.Name,
			Members: members,
		
	}, nil
}

func (c Repo) AddMessage(chatId string, msg entity.Message) error {
	m := store.BHTextMessage{
		ID: msg.ID,
		ChatID: chatId,
		CreatedAt: msg.CreatedAt,
		Text: msg.Text,
		Status: store.Status(msg.Status),
		Author: store.BHContact(msg.Author),
	}
	err := c.store.InsertTextMessage(m)
	if err != nil {
		return err
	}
	return nil
}

func (c Repo) CreateChat(chat entity.ChatInfo) error {
	m := []string{}
	for _, val := range chat.Members {
		m = append(m, val.ID)
	}
	ci := store.BHChat{
		ID: chat.ID,
		Name: chat.Name,
		Members: m,
	}
	err := c.store.InsertChat(ci)
	if err != nil {
		return err
	}
	return nil
}


func (c Repo) GetAllContact() ([]entity.Contact, error) {
	cons := make([]entity.Contact, 0)
	bhcl, err := c.store.AllContacts()
	if err != nil {
		return nil, err
	}
	for _, val := range bhcl {
		cons = append(cons, entity.Contact{
			Name: val.Name,
			ID:   val.ID,
		})
	}
	return cons, nil
}

func (c Repo) GetByIDContact(id string) (*entity.Contact, error) {
	con, err := c.store.ContactByID(id)
	if err != nil {
		return nil, err
	}
	return &entity.Contact{
		ID:   con.ID,
		Name: con.Name,
	}, nil
}

func (c Repo) AddContact(con entity.Contact) error {
	err := c.store.InsertContact(store.BHContact{
		Name: con.Name,
		ID:   con.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
