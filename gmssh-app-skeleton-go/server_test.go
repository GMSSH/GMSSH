package simplejrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/DemonZack/simplejrpc-go/core"
	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/net/gsock"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

type CustomMiddleware struct{}

func (c *CustomMiddleware) ProcessRequest(req *gsock.Request) {
	log.Println("[*] ProcessRequest. ", req)
}

func (c *CustomMiddleware) ProcessResponse(resp any) (any, error) {
	log.Println("[*] ProcessResponse. ", resp)
	return resp, nil
}

func TestJsonDefaultRpcServer(t *testing.T) {
	mockSockPath := "zack.sock"

	ds := NewDefaultServer(
		gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
		gsock.WithJsonRpcSimpleServiceMiddlewares([]gsock.RPCMiddleware{
			&CustomMiddleware{},
		}...),
	)

	err := ds.StartServer(mockSockPath)
	if err != nil {
		t.Fatalf("start server failed :%v", err)
	}
	t.Log("start server successfully")
}

type CustomHandler struct{}

func (c *CustomHandler) Hello(req *gsock.Request) (any, error) {
	fmt.Println("[*] request >>> ", req)
	resp := "Hello World"
	return resp, nil
}

func (c *CustomHandler) ValidPerson(req *gsock.Request) (any, error) {
	fmt.Println("[*] request >>> ", req.RawRequest().Params)

	var person ExampleUser

	err := json.Unmarshal(*req.RawRequest().Params, &person)
	if err != nil {
		return nil, errors.New("test unmarshal error")
	}

	err = core.Container.Valid().Walk(&person)
	if err != nil {
		return nil, err
	}
	// resp := "Hello World"
	return person, nil
}

func (c *CustomHandler) ProcessRequest(req *gsock.Request) {
	log.Println("[*] ProcessRequest. ", req)
}

func (c *CustomHandler) ProcessResponse(resp any) (any, error) {
	log.Println("[*] ProcessResponse. ", resp)
	return resp, nil
}

type ExampleUser struct {
	Username string `json:"username" validate:"min_length:6#The length is too small"`
	Age      any    `json:"age" validate:"required#Required parameters are missing Age|range:18,100|int#Test verification error return"`
}

func TestJsonDefaultValidateRpcServer(t *testing.T) {
	env := "test"
	gpath.GmCfgPath = filepath.Join(filepath.Dir(""), "..")
	core.InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	mockSockPath := "zack.sock"

	ds := NewDefaultServer(
		gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
		gsock.WithJsonRpcSimpleServiceMiddlewares([]gsock.RPCMiddleware{
			&CustomMiddleware{},
		}...),
	)

	hand := &CustomHandler{}
	ds.RegisterHandle("ValidPerson", hand.ValidPerson, []gsock.RPCMiddleware{hand}...)
	err := ds.StartServer(mockSockPath)
	if err != nil {
		t.Fatalf("start server failed :%v", err)
	}
	t.Log("start server successfully")
}

func TestJsonDefaultRpcServerWithHandleMiddleware(t *testing.T) {
	mockSockPath := "zack.sock"

	ds := NewDefaultServer(
		gsock.WithJsonRpcSimpleServiceHandler(gsock.NewJsonRpcSimpleServiceHandler()),
		gsock.WithJsonRpcSimpleServiceMiddlewares([]gsock.RPCMiddleware{
			&CustomMiddleware{},
		}...),
	)

	hand := &CustomHandler{}
	ds.RegisterHandle("hello", hand.Hello, []gsock.RPCMiddleware{hand}...)
	// ds.RegisterHandle("raisehello", hand.RaiseHello, []gsock.RPCMiddleware{hand}...)
	err := ds.StartServer(mockSockPath)
	if err != nil {
		t.Fatalf("start server failed :%v", err)
	}
	t.Log("start server successfully")
}
