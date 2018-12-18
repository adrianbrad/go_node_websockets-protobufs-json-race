'use strict';
const Message = require('./message_pb');
const WebSocket = require('ws');
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

// startJson()
startBinary()
function startBinary() {

    let wsProto = new WebSocket('ws://localhost:8080/proto', {
        origin: 'http://localhost:8080'
    });

    wsProto.on('open', function open() {
        console.log('connected proto');
    });

    wsProto.on('message', function incoming(data) {
        // console.log('proto message received');
        // console.log(data);
        // var bytes = Array.prototype.slice.call(data, 0);
        // var message = proto.message.Message.deserializeBinary(bytes);
        // console.log(message)
        wsProto.send(data);
    });

    wsProto.on('close', function close() {
        console.log('disconnected proto')
        startJson()
    });
}

function startJson() {

    let wsJson = new WebSocket('ws://localhost:8080/json', {
        origin: 'http://localhost:8080'
    });

    wsJson.on('close', function close() {
        console.log('disconnected json');
        // startBinary()
    });

    wsJson.on('message', function incoming(message) {
        // console.log('json message received');
        // console.log(message);
        wsJson.send(message)
    })

    wsJson.on('open', function open() {
        console.log('connected json')
    });
}


