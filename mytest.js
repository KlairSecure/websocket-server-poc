const WebSocket = require('ws');
const async = require('async');

const numClients = 500;
const duration = 5 * 60 * 1000; // 5 minutes in milliseconds

const clients = [];

function createWebSocketClient(clientIndex) {
    const ws = new WebSocket('ws://localhost:8080/ws');

    ws.on('open', () => {
        console.log(`Client ${clientIndex} connected`);
        setInterval(() => {
            // Send a message to the server every second
            ws.send(`${clientIndex}: Hello from client`);
          }, 1);
    });

    ws.on('message', (message) => {
        console.log(`${clientIndex} ${message}`)
        if (message === 'pong') {
            // You can handle the received "ok" messages here if needed
        }
    });

    ws.on('close', () => {
        console.log(`Client ${clientIndex} disconnected`);
    });

    return ws;
}

function closeWebSocketClient(clientIndex, ws) {
    ws.close();
    console.log(`Client ${clientIndex} disconnected`);
}

function startWebSocketClients() {
    for (let i = 0; i < numClients; i++) {
        const wsClient = createWebSocketClient(i);
        clients.push(wsClient);
    }
}

function stopWebSocketClients() {
    clients.forEach((ws, index) => {
        closeWebSocketClient(index, ws);
    });
}

startWebSocketClients();

setTimeout(() => {
    stopWebSocketClients();
    console.log('Test completed.');
}, duration);
