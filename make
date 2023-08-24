docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

#docker exec -it 8da bash
#rabbitmqadmin publish exchange=amq.default routing_key="TestQueue" payload="Hello Ali"
