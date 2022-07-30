package main

import(
	"github.com/aws/aws-lambda-go/lambda"
	"TestLambda/greeting"
)

func excuteFunction(){
	greeting.SayHello()
} 

func main()  {
	lambda.Start(excuteFunction)
}