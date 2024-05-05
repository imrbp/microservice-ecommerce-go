package main

import (
	"order_service/internal"
	dependency "order_service/internal/depedency"

	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

	
func main() {
  log := dependency.NewLogger()
  db := dependency.NewDatabase()
  validate := dependency.NewCustomValidator(logrus.New())
  app := dependency.NewFiber() 
  conn, err := rabbitmq.NewConn(
    "amqp://guest:guest@localhost:5672",
    )

  if err != nil {
    log.Log.Panicf("Error Establish Connection : %v" , err)
  }

  internal.BootstrapWorker(&internal.BootstrapWorkerApp{
    DB :db,
    Log : log,
    Validation : validate,
    App : app,
    Conn : conn,
  })


}
