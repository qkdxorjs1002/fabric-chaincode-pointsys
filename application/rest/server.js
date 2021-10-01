const express = require('express');
const app = express();
let path = require('path');
let sdk = require('./sdk');

const PORT = 8001;
const HOST = '0.0.0.0';
app.use(express.json());
app.use(express.urlencoded({ extended: true }))

app.get('/member/add', function (req, res) {
   let args = [req.query.name];
   sdk.send(false, 'AddMember', args, res);
});

app.get('/member/query', function (req, res) {
   let args = [req.query.name];
   sdk.send(true, 'QueryMember', args, res);
});

app.get('/member/delete', function (req, res) {
   let args = [req.query.name];
   sdk.send(true, 'DeleteMember', args, res);
});

app.get('/point/update', function (req, res) {
   let args = [req.query.name, req.query.point];
   sdk.send(true, 'UpdateMemberPoint', args, res);
});

app.use(express.static(path.join(__dirname, '../client')));
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
