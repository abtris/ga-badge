const getBadge = require('./ghactions');

test('basic fail test', () => {
  expect(getBadge('xxxx')).toBe('Wrong URL, can\'t generate badge');
});

test('basic ok test', () => {
  expect(getBadge('https://github.com/abtris/ga-badge/actions/workflows/node.js.yml')).toBe('[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)');
});

test('basic option ok test', () => {
  expect(getBadge('https://github.com/abtris/ga-badge/actions/workflows/node.js.yml', { title: "test" })).toBe('[![test](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)');
});
