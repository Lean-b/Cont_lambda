package main

import(
  "fmt",
  "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
   Name string '`json:"what is you name?"`'
   Email string '`json:"What is you email?"`'
   Comment string '`json:"send a message to the owner of this page "`'
}


type ReturnEvent struct {
  Message string '`json:"return email"`'
}



func HandleRequest(event MyEvent)(string error) {
  return Response{Message:
  fmt.Sprintf("Name: %s ,Email: %d,Comment: %g",
  event.name, event.email, event.comment)
},nil

}

func main()  {
  lambda.Start(HandleRequest)
}
