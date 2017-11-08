// @ts-check
"use strict";

const express = require('express');

module.exports = (mongoSession) => {
    if (!mongoSession) {
        throw new Error("Provide Mongo Session");
    }

    let router = express.Router();
    router.get("/v1/channels", (req, res) => {
        // query mongo using mongoSession
        res.json([{name: "general"}]);
    });

    return router;
}