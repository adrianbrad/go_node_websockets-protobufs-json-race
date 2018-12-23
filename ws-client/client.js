'use strict';
const Message = require('./message_pb');
const WebSocket = require('ws');
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

const wsProto = new WebSocket('ws://localhost:8080/proto', {
    origin: 'http://localhost:8080'
});
// wsProto.binaryType = "arraybuffer"

let wsJson;

wsProto.on('open', function open() {
    console.log('connected proto');
});

wsProto.on('close', function close() {
    console.log('disconnected proto')

    wsJson = new WebSocket('ws://localhost:8080/json', {
        origin: 'http://localhost:8080'
    });

    wsJson.on('close', function close() {
        console.log('disconnected json');
    });

    wsJson.on('message', function incoming(message) {
        // console.log('json message received');
        // console.log(message);
        wsJson.send(message)
    })

    wsJson.on('open', function open() {
        console.log('connected json')
    });
});


wsProto.on('message', function incoming(data) {
    // console.log('proto message received');
    // var bytes = Array.prototype.slice.call(data, 0);
    // var message = proto.message.Message.deserializeBinary(bytes);

    wsProto.send(data);
});


