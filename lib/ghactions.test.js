const getBadge = require('./ghactions');

test('basic fail test', () => {
  expect(getBadge('xxxx')).toBe('Wrong URL, can\'t generate badge');
});

test('basic ok test', () => {
  expect(getBadge('https://github.com/abtris/ga-badge/actions/workflows/node.js.yml')).toBe('Wrong URL, can\'t generate badge');
});
