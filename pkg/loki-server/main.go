package loki_server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	fWebsocket "github.com/gofiber/websocket/v2"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/url"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{}

type StreamSelector struct {
	Key       string
	Value     string
	Operation string
}

type LokiClient interface {
	Tail(
		streamSelectors []StreamSelector,
		filter *string,
		start, end *int64,
		limit *int,
		connection *fWebsocket.Conn,
	) error
}

type lokiClient struct {
	url *url.URL
}

func (l *lokiClient) Tail(
	streamSelectors []StreamSelector,
	filter *string,
	start, end *int64,
	limit *int,
	connection *fWebsocket.Conn,
) error {
	streamSelectorSplits := make([]string, 0)
	for _, label := range streamSelectors {
		streamSelectorSplits = append(streamSelectorSplits, label.Key+label.Operation+fmt.Sprintf("\"%s\"", label.Value))
	}
	query := url.Values{}
	filterStr := ""
	if filter != nil {
		filterStr = *filter
	}
	query.Set("query", fmt.Sprintf("%v%v", fmt.Sprintf("{%v}", strings.Join(streamSelectorSplits, ",")), filterStr))
	if start != nil {
		query.Set("start", fmt.Sprintf("%v", start))
	} else {
		query.Set("start", fmt.Sprintf("%v", time.Now().Add(-time.Hour).UnixNano()))
	}
	if end != nil {
		query.Set("env", fmt.Sprintf("%v", end))
	}
	if limit != nil {
		query.Set("limit", fmt.Sprintf("%v", limit))
	} else {
		query.Set("limit", fmt.Sprintf("%v", 30))
	}
	u := url.URL{Scheme: "ws", Host: l.url.Host, Path: "/loki/api/v1/tail", RawQuery: query.Encode()}
	fmt.Println(u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return err
	}
	defer conn.Close()
	for {
		msgType, message, err := conn.ReadMessage()
		fmt.Println(message)
		connection.WriteMessage(msgType, message)
		if err != nil {
			connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, ""))
			return err
		}
		if msgType == websocket.CloseMessage {
			return nil
		}
	}
}

func NewLokiClient(serverUrl string) (LokiClient, error) {
	u, err := url.Parse(serverUrl)
	if err != nil {
		return nil, err
	}
	return &lokiClient{
		url: u,
	}, nil
}

type LogServer *fiber.App

type LokiClientOptions interface {
	GetLokiServerUrl() string
	GetLogServerPort() uint64
}

func NewLokiClientFx[T LokiClientOptions]() fx.Option {
	return fx.Module("loki-client",
		fx.Provide(func() LogServer {
			return fiber.New()
		}),
		fx.Invoke(func(o T, app LogServer, lifecycle fx.Lifecycle) {
			var a *fiber.App
			a = app
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go a.Listen(fmt.Sprintf(":%v", o.GetLogServerPort()))
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return a.Shutdown()
				},
			})
		}),
		fx.Provide(func(o T) (LokiClient, error) {
			return NewLokiClient(o.GetLokiServerUrl())
		}),
		fx.Invoke(func(app LogServer, lokiServer LokiClient) {
			var a *fiber.App
			a = app
			a.Use("/", func(c *fiber.Ctx) error {
				if fWebsocket.IsWebSocketUpgrade(c) {
					c.Locals("allowed", true)
					return c.Next()
				}
				return fiber.ErrUpgradeRequired
			})
		}))
}
