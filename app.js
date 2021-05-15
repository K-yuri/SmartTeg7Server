const express = require("express");
const app = express();

const port = 8715;

app.set("port", port);

app.get("/", (req, res) => {
    res.send("team7 smart-teg Server");
});

app.listen(port, () => console.log("Listening on", port));

module.exports = app;