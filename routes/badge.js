var express = require('express');
var router = express.Router()

var getBadge = require('../lib/ghactions.js')


router.post('/', function (req, res, next) {
  output = getBadge(req.body.gh_url)
  res.render('badge', { title: 'Github Action Badge Generator', snippet: output, url: req.body.gh_url });
});

router.get('/', function (req, res, next) {
  res.send('respond with a resource badge');
});

module.exports = router;
