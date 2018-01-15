'use strict';

const express = require('express');

// Constants
const PORT = 5002;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/headers', (req, res) => {
  res.send(JSON.stringify(req.headers));
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
