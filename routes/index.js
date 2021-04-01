var express = require('express');
var router = express.Router();
var getBadge = require('../lib/ghactions.js')


router.get('/', function(req, res, next) {
  res.render('index', { title: 'Badge Generator', url: 'https://github.com/abtris/ga-badge/actions/workflows/node.js.yml', label: 'Build Status' });
});

router.post('/', function (req, res, next) {
  output = getBadge(req.body.gh_url, { title: req.body.title })
  res.render('badge', { title: 'Badge Generator', snippet: output, url: req.body.gh_url, label: req.body.title });
});

module.exports = router;
