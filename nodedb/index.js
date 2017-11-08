"use strict";
const mongodb = require('mongodb');
const MongoStore = require('./taskstore');
const express = require('express');
const app = express();

// middleware that will parse json and send to req.body

const addr = process.env.ADDR || "localhost:4000";
const [host, port] = addr.split(":");
console.log(host);
console.log(port);

// setup mongo
const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

mongodb.MongoClient.connect(mongoURL)
    .then(db => {
        // initialize mongostore
        let store = new MongoStore(db, "tasks");
        app.use(express.json());
        app.post("/v1/tasks", (req, res) => {
            // insert new task
            let task = {
                title: req.body.title,
                completed: false
            }

            store.insert(task)
                .then(tasks => {
                    res.json(tasks);
                })
                .catch(err => {
                    throw err;
                });
        });

        app.get("/v1/tasks", (req, res) => {
            // get all tasks
            console.log("hi");
            store.getAll(req.params.id)
            .then(tasks => {
                res.json(tasks);
            }).catch(err => {
                throw err;
            });
        });

        app.patch("/v1/tasks/:taskID", (req, res) => {
            console.log(req.params.taskID);
            store.update(req.params.taskID, {title: req.body.title})
            .then(updatedTask => {
                console.log(updatedTask);
                res.json(updatedTask);
            }).catch(err => {
                throw err;
            });
        });

        app.listen(4000, host, () => {
            console.log("Server is listening at http://" + addr)
        });
    }).catch(err => {
        throw err;
    })
