var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Github Action Badge Generator', url: 'https://github.com/abtris/ga-badge/actions/workflows/node.js.yml' });
});

module.exports = router;
