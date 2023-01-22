package websocket

import (
	"encoding/xml"
	"fmt"
	"github.com/plin2k/api-mocker/config"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const ProtocolName = "websocket"

type handler struct {
	api      *gin.Engine
	config   *config.Mocker
	source   *Source
	upgrader websocket.Upgrader
}

type connection struct {
	socket *websocket.Conn
	mu     sync.Mutex
}

func New(cfg *config.Mocker) *handler {
	gin.SetMode(gin.TestMode)
	return &handler{
		api:    gin.Default(),
		config: cfg,
		source: &Source{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

func (h *handler) Run() error {
	log.Printf("Starting Websocket server on port %d", h.config.Port)
	return h.api.Run(fmt.Sprintf(":%d", h.config.Port))
}

func (h *handler) Construct(body []byte) (err error) {
	err = xml.Unmarshal(body, h.source)
	if err != nil {
		return err
	}

	h.api.GET("/ws", h.wsEndpoint)

	return nil
}

func (h *handler) wsEndpoint(c *gin.Context) {
	wsConn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalln(err)
	}

	conn := &connection{
		socket: wsConn,
	}

	err = h.onOpen(conn)
	if err != nil {
		log.Fatalln(err)
	}

	go h.while(conn)

	h.reader(conn)
}

func (h *handler) reader(conn *connection) {
	for {
		messageType, msg, err := conn.socket.ReadMessage()
		if err != nil {
			log.Fatalln(err)
			return
		}

		switch messageType {
		case websocket.TextMessage, websocket.BinaryMessage:
			err = h.onMessage(conn, msg)
		case websocket.CloseMessage, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			websocket.CloseNormalClosure, websocket.CloseMandatoryExtension,
			websocket.CloseInvalidFramePayloadData, websocket.CloseMessageTooBig, websocket.CloseNoStatusReceived,
			websocket.ClosePolicyViolation, websocket.CloseUnsupportedData,
			websocket.CloseTryAgainLater, websocket.CloseServiceRestart, websocket.CloseTLSHandshake:
			err = h.onClose(conn, msg)
		case websocket.CloseProtocolError, websocket.CloseInternalServerErr:
			err = h.onError(conn, msg)
		case websocket.PingMessage:
			err = h.ping(conn, msg)
		case websocket.PongMessage:
			err = h.pong(conn, msg)
		default:
			err = h.response(conn, &Response{
				Value: "undefined message type",
				Type:  TypeText,
			})
		}

		if err != nil {
			log.Fatalln(err)
		}

	}
}

func (h *handler) while(conn *connection) {
	var threadFn = func(conn *connection, msg *Message) {
		if msg.Count == 0 {
			msg.Count = math.MaxUint
		}

		if msg.Delay < 1 {
			msg.Delay = 1
		}

		for i := uint(0); i < msg.Count; i++ {
			if err := h.response(conn, msg.Response); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Duration(msg.Delay) * time.Second)
		}
	}

	for i := range h.source.While.Messages {
		go threadFn(conn, h.source.While.Messages[i])
	}
}

func (h *handler) onOpen(conn *connection) (err error) {
	err = h.response(conn, h.source.OnOpen.Response)

	return err
}

func (h *handler) onClose(conn *connection, _ []byte) (err error) {
	err = h.response(conn, h.source.OnClose.Response)

	return err
}

func (h *handler) onMessage(conn *connection, _ []byte) (err error) {
	err = h.response(conn, h.source.OnMessage.Response)

	return err
}

func (h *handler) onError(conn *connection, _ []byte) (err error) {
	err = h.response(conn, h.source.OnError.Response)

	return err
}

func (h *handler) ping(conn *connection, _ []byte) (err error) {
	err = h.response(conn, h.source.Ping.Response)

	return err
}

func (h *handler) pong(conn *connection, _ []byte) (err error) {
	err = h.response(conn, h.source.Pong.Response)

	return err
}

func (h *handler) response(conn *connection, resp *Response) (err error) {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	err = conn.socket.WriteMessage(resp.GetType(), []byte(resp.Value))

	return err
}
