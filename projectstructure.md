## Architecture

```mermaid
graph LR
    node1(Message Processor)
    node4(Observer Lobby Server)
    node5(wsclient1)
    node2(Message Reciever 1)
    node3(Message Sender 1)
    node6(Algorithm and room state watcher)
    node7(wsclient2)
    node8(Message Reciever 2)
    node9(Message Sender 2)
    node10(wsclient3)
    node11(Message Reciever 3)
    node12(Message Sender 3)
    node13(wsclient4)
    node14(Message Reciever 4)
    node15(Message Sender 4)
    node16(Message Processor arch)
    node17[clients list]
    node18(client1 message sender)
    node19(client2 message sender)
    node20(client3 message sender)

    node21[[Client connected]]
    node22[[Client disconnected]]
    node23[[Room created]]
    node24[[Room deleted]]
    node25[[Client Joined a Room]]
    node26[[Client Left a Room]]
    node27[[Event Triggers]]

    node28[(Clients Map <br>cid_key--struct)]
    node29[(Rooms Map <br>roomname_key--struct)]
    node30[(Room to client map <br>roomkey--map-clinetkeys)]

    node31[[Message Send in a processor <br> 1. Action to be done/message type <br> 2. Room address <br> 3. If not 2nd then client address <br> 4. Message body]]

node31-->node16
    node16-->node17
    node17-->node18
    node17-->node19
    node17-->node20
    node17-->node30


    node27-->node21
    node27-->node22
    node27-->node23
    node27-->node24
    node27-->node25
    node27-->node26
  
    node21-->node28
    node22-->node28

    node23-->node29
    node24-->node29
  
    node25-->node30
    node26-->node30
    node30-->node28
    node28-->node17

    node2-->node1-->node3
    node8-->node1-->node9
    node11-->node1-->node12
    node14-->node1-->node15

    node6-->node1
    node5-.->node2
    node5-.->node3

    node7-.->node8
    node7-.->node9

    node10-.->node11
    node10-.->node12

    node13-.->node14
    node13-.->node15


    node4-.->node5
    node4-.->node6
    node4-.->node1
    node4-.->node7
    node4-.->node10
    node4-.->node13
```
