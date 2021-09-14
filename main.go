package main

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yoshis2/smart-contract-beginner/contracts"
)

// ここの２つの変数を修正して実行する
const (
	// ganacheの起動したときのポートを指定 (8545 か 7545)
	GANACHE_PORT = "8545"
	// 先ほど作成したプログラムから取得した。　CONTRACT_ADDRESSを取得
	CONTRACT_ADDRESS = "0xEa2ff902dbeEECcc828757B881b343F9316752e5"
)

func main() {
	client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", GANACHE_PORT))
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		panic(err)
	}

	run(conn)
}

func run(conn *contracts.Contracts) error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("エラー : %s", err.Error()))
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("エラー : %s", err.Error()))
			return err
		}
		return c.JSON(http.StatusOK, reply) // Hello World
	})

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
