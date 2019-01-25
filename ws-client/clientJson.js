'use strict';

const WebSocket = require('ws');
const host = "ws://localhost:8080"

class WebsocketConnection {

    constructor(host, endpoint) {
        this._endpoint = endpoint;
        this._host = host;
    }

    connect() {
        this._ws = new WebSocket(this._host + "/" + this._endpoint, {
            origin: this._host
        });
        this._ws.on('message', (data) => this.incomingMessage(data));
    }

    onOpen() {
        console.log('connected ' + this._endpoint);
    }

    incomingMessage(data) {
        this._ws.send(data)
    }

    onClose(event) {
        console.log('disconnected ' + this._endpoint + " " + event);
    }
}

var wsProto = new WebsocketConnection(host, "json");
wsProto.connect();