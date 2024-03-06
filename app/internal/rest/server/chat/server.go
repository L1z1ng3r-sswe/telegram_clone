package chat_rest

import (
	"fmt"
	"net/http"
	"time"

	logger "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/lib/zerolog"
	models_rest "github.com/L1z1ng3r-sswe/telegram_clone/app/internal/rest/domain/models"
	"github.com/gorilla/websocket"
)

type service interface {
	CreateCommunity(comm models_rest.Community) (models_rest.Community, error, string, string, int, string)
	JoinCommunity(communityMember models_rest.CommunityMember) (models_rest.CommunityMember, error, string, string, int, string)
	CreateCommunityMessage(msg models_rest.Message) (models_rest.Message, error, string, string, int, string)
	CreateBIChat(chat models_rest.BIChat) (models_rest.BIChat, error, string, string, int, string)
	CreateBIMessage(msg models_rest.Message) (models_rest.Message, error, string, string, int, string)
	GetAllBIMessages(biChatId string) ([]models_rest.Message, error, string, string, int, string)
}

type server struct {
	service         service
	log             *logger.Logger
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
	secretKey       string
	upgrader        websocket.Upgrader
	Hub             Hub
}

type Hub struct {
	GracefulStop  chan struct{}
	log           *logger.Logger
	Clients       map[int64]*models_rest.Client
	Communities   map[int64]models_rest.Community
	BroadcastComm chan models_rest.Message
}

func (h *Hub) Run() {
	h.log.AppInf(fmt.Sprintf("Hub is running"))

	defer func() {
		close(h.BroadcastComm)
		close(h.GracefulStop)
	}()

	for {
		select {
		case msg := <-h.BroadcastComm:
			comm := h.Communities[msg.CommunityId]
			for _, cl := range comm.Clients {
				cl.CommMessages <- msg
			}

		case <-h.GracefulStop:
			h.log.AppInf("Hub shutdown completed")
			return
		}
	}
}

func New(log *logger.Logger, service service, secretKey string) *server {
	return &server{
		secretKey: secretKey,
		service:   service,
		log:       log,
		Hub: Hub{
			GracefulStop: make(chan struct{}, 1),
			log:          log,
			// for bi-chats
			Clients: make(map[int64]*models_rest.Client),

			// for communities
			BroadcastComm: make(chan models_rest.Message),
			Communities:   make(map[int64]models_rest.Community),
		},

		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
