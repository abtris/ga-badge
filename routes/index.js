var express = require('express');
var router = express.Router();
var getBadge = require('../lib/ghactions.js')


const defaultSnippet = '[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)'

router.get('/', function(req, res, next) {
  res.render('index', { title: 'Badge Generator', snippet: defaultSnippet, url: 'https://github.com/abtris/ga-badge/actions/workflows/node.js.yml', label: 'Build Status', branch: 'default' });
});

router.post('/', function (req, res, next) {
  output = getBadge(req.body.gh_url, { title: req.body.title, branch: req.body.branch })
  res.render('index', { title: 'Badge Generator', snippet: output, url: req.body.gh_url, label: req.body.title, branch: req.body.branch });
});

module.exports = router;
