var express = require('express');
var router = express.Router()

var getBadge = require('../lib/ghactions.js')


router.post('/', function (req, res, next) {
  res.send(getBadge(req.body.gh_url));
});

router.get('/', function (req, res, next) {
  res.send('respond with a resource badge');
});

module.exports = router;
