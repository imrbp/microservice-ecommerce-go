package init

import (

)

func main() {

  // if err := RabbitMQ.Channel.ExchangeDeclare(
  //   "proceed_payment",   // name
  //   "direct", // type
  //   true,     // durable
  //   false,    // auto-deleted
  //   false,    // internal
  //   false,    // no-wait
  //   nil,      // arguments
  //   ) ; err != nil {
  //   panic(err)
  // }

  // if err := RabbitMQ.Channel.ExchangeDeclare(
  //   "payment_status",   // name
  //   "fanout", // type
  //   true,     // durable
  //   false,    // auto-deleted
  //   false,    // internal
  //   false,    // no-wait
  //   nil,      // arguments
  //   ) ; err != nil {
  //   panic(err)
  // }

  // if err := RabbitMQ.Channel.ExchangeDeclare(
  //   "shipping_status",   // name
  //   "fanout", // type
  //   true,     // durable
  //   false,    // auto-deleted
  //   false,    // internal
  //   false,    // no-wait
  //   nil,      // arguments
  //   ) ; err != nil {
  //   panic(err)
  // }

  // _, err := RabbitMQ.Channel.QueueDeclare(
  //   "order_service",    // name
  //   false, // durable
  //   false, // delete when unused
  //   true,  // exclusive
  //   false, // no-wait
  //   nil,   // arguments
  //   )


  // if err != nil {
  //   panic(err)
  // }


  // _, err = RabbitMQ.Channel.QueueDeclare(
  //   "payment_service",    // name
  //   false, // durable
  //   false, // delete when unused
  //   true,  // exclusive
  //   false, // no-wait
  //   nil,   // arguments
  //   )

  // if err != nil {
  //   panic(err)
  // }

  // _, err = RabbitMQ.Channel.QueueDeclare(
  //   "warehouse_service",    // name
  //   false, // durable
  //   false, // delete when unused
  //   true,  // exclusive
  //   false, // no-wait
  //   nil,   // arguments
  //   )

  // if err != nil {
  //   panic(err)
  // }

  // if err := RabbitMQ.Channel.QueueBind(
  //   "payment_service", // queue name
  //   "",     // routing key
  //   "proceed_payment", // exchange
  //   false,
  //   nil,
  //   ); err != nil {
  //   panic(err)
  // }
  // 
  // if err := RabbitMQ.Channel.QueueBind(
  //   "order_service", // queue name
  //   "",     // routing key
  //   "payment_status", // exchange
  //   false,
  //   nil,
  //   ); err != nil {
  //   panic(err)
  // }

  // if err := RabbitMQ.Channel.QueueBind(
  //   "warehouse_service", // queue name
  //   "",     // routing key
  //   "payment_status", // exchange
  //   false,
  //   nil,
  //   ); err != nil {
  //   panic(err)
  // }
  // 
  // if err := RabbitMQ.Channel.QueueBind(
  //   "order_service", // queue name
  //   "",     // routing key
  //   "shipping_status", // exchange
  //   false,
  //   nil,
  //   ); err != nil {
  //   panic(err)
  // }

}
