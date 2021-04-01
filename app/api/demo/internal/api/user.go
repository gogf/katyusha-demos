package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/protobuf/demos"
	"github.com/gogf/katyusha/krpc"
	"golang.org/x/net/context"
)

var (
	userClient demos.UserClient
)

func init() {
	conn, err := krpc.Client.NewGrpcClientConn("demos")
	if err != nil {
		g.Log().Fatal(err)
	}
	userClient = demos.NewUserClient(conn)
}

func TestSignUp() {
	res, err := userClient.SignUp(context.Background(), &demos.SignUpReq{
		Passport:  "john",
		Password:  "123",
		Password2: "123",
		Nickname:  "JohnGuo",
	})
	g.Log().Print(res, err)
}

func TestCheckPassport() {
	res, err := userClient.CheckPassport(context.Background(), &demos.CheckPassportReq{
		Passport: "john",
	})
	g.Log().Print(res, err)
}

func TestCheckNickName() {
	res, err := userClient.CheckNickName(context.Background(), &demos.CheckNickNameReq{
		Nickname: "john",
	})
	g.Log().Print(res, err)
}

func TestGetUser() {
	res, err := userClient.GetUser(context.Background(), &demos.GetUserReq{
		UserId: 1,
	})
	g.Log().Print(res, err)
}

func TestSignIn() {
	res, err := userClient.SignIn(context.Background(), &demos.SignInReq{
		Passport: "john",
		Password: "123",
	})
	g.Log().Print(res, err)
}

func TestGetSession() {
	signInRes, err := userClient.SignIn(context.Background(), &demos.SignInReq{
		Passport: "john",
		Password: "123",
	})
	g.Log().Printf(`%+v, %v`, signInRes, err)
	getSessionRes, err := userClient.GetSession(context.Background(), &demos.GetSessionReq{
		Token: signInRes.Token,
	})
	g.Log().Printf(`%+v, %v`, getSessionRes, err)
}

func main() {
	//TestSignUp()
	//TestCheckPassport()
	//TestSignIn()
	//TestGetSession()

	res, err := userClient.SignUp(context.Background(), &demos.SignUpReq{
		Passport:  "john",
		Password:  "123",
		Password2: "1234",
		Nickname:  "JohnGuo",
	})
	g.Log().Print(res, err)
}
