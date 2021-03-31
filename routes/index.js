var express = require('express');
var router = express.Router();
var getBadge = require('../lib/ghactions.js')


router.get('/', function(req, res, next) {
  res.render('index', { title: 'Github Action Badge Generator', url: 'https://github.com/abtris/ga-badge/actions/workflows/node.js.yml' });
});

router.post('/', function (req, res, next) {
  output = getBadge(req.body.gh_url)
  res.render('badge', { title: 'Github Action Badge Generator', snippet: output, url: req.body.gh_url });
});

module.exports = router;
