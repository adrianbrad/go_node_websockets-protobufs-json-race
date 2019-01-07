'use strict';
const Message = require('./message_pb');
const WebSocket = require('ws');
const host = "ws://localhost:8080"

class WebsocketConnection {

    constructor(host, endpoint, closeCb) {
        this._endpoint = endpoint;
        this._host = host;
        this._closeCb = closeCb;
    }

    connect() {
        this._ws = new WebSocket(this._host + "/" + this._endpoint, {
            origin: this._host
        });
        this._ws.on('open', () => this.onOpen());
        this._ws.on('message', (data) => this.incomingMessage(data));
        this._ws.on('close', () => this.onClose());
    }

    onOpen() {
        console.log('connected ' + this._endpoint);
    }

    incomingMessage(data) {
        // console.log(this._endpoint +' message received');
        // var bytes = Array.prototype.slice.call(data, 0);
        // var message = proto.message.Message.deserializeBinary(bytes);
        // console.log(message);
        this._ws.send(data)
    }

    onClose() {
        console.log('disconnected ' + this._endpoint);
        if (this._closeCb) this._closeCb();
    }
}


var wsProto = new WebsocketConnection(host, "proto", startJs);
wsProto.connect();

function startJs() {
    const js = new WebsocketConnection(host, "json");
    js.connect();
}


