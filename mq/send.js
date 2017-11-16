#!/usr/bin/env node

"use strict";

const amqp = require('amqplib');
const qName = "testQ";
const mqAddr = process.env.MQADDR || "localhost:5672";
const mqURL = `amqp://${mqAddr}`;

(async function() {
    console.log("connecting to %s", mqURL);
    let connection = await amqp.connect(mqURL)
    let channel = await connection.createChannel();
    let qConf = await channel.assertQueue(qName, {durable: false});
    setInterval(() => {
        let msg = "message " + new Date().toLocaleTimeString();
        channel.sendToQueue(qName, Buffer.from(msg));
    }, 1000);
})();