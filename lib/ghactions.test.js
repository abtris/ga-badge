const getBadge = require('./ghactions');

test('basic test', () => {
  expect(getBadge('xxxx')).toBe('[![Build Status](https://github.com/abtris/sinopia-htaccess-gpg-email/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/sinopia-htaccess-gpg-email/actions)');
});
