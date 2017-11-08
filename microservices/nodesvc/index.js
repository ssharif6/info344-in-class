// @ts-check
"use strict";

// Load express and morgan modules
const express = require('express');
const morgan = require('morgan');

const addr = process.env.ADDR || ":80";
// destructure assignment
const [host, port] = addr.split(":");
const portNum = parseInt(port);
const handlers = require('./handlers');


// value of a require(x) is the public api
const app = express();
app.use(morgan(process.env.LOG_FORMAT || "dev"));

app.use(handlers({}));

app.get("/", (req, res) => {
    res.set("Content-Type", "text/plain");
    res.send("Hello, Node.js!");
});

app.get("/v1/users/me/hello", (req, res) => {
    // req.get gets header values
    let userJSON = req.get("X-User");
    if (!userJSON) {
        // handle errors locally but don't expose to browser
        throw new Error("no X-User header provided");
    }

    let user = JSON.parse(userJSON);
    res.json({
        message: `Hello ${user.firstName} ${user.lastName}`
    });
});

// Error handler invoked when error or exception thrown
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.set("Content-Type", "text/plain");
    res.send(err.message);
});

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}...`);
});