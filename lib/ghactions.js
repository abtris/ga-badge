var url = require("url");

// ┌────────────────────────────────────────────────────────────────────────────────────────────────┐
// │                                              href                                              │
// ├──────────┬──┬─────────────────────┬────────────────────────┬───────────────────────────┬───────┤
// │ protocol │  │        auth         │          host          │           path            │ hash  │
// │          │  │                     ├─────────────────┬──────┼──────────┬────────────────┤       │
// │          │  │                     │    hostname     │ port │ pathname │     search     │       │
// │          │  │                     │                 │      │          ├─┬──────────────┤       │
// │          │  │                     │                 │      │          │ │    query     │       │
// "  https:   //    user   :   pass   @ sub.example.com : 8080   /p/a/t/h  ?  query=string   #hash "
// │          │  │          │          │    hostname     │ port │          │                │       │
// │          │  │          │          ├─────────────────┴──────┤          │                │       │
// │ protocol │  │ username │ password │          host          │          │                │       │
// ├──────────┴──┼──────────┴──────────┼────────────────────────┤          │                │       │
// │   origin    │                     │         origin         │ pathname │     search     │ hash  │
// ├─────────────┴─────────────────────┴────────────────────────┴──────────┴────────────────┴───────┤
// │                                              href                                              │
// └────────────────────────────────────────────────────────────────────────────────────────────────┘
// input: https://github.com/abtris/sinopia-htaccess-gpg-email/actions/workflows/node.js.yml
function getBadge(input_url, options) {
  const myURL = new URL(input_url);
  const fullPathName = myURL.pathname.split('/')
  repoOwner = fullPathName[1]
  repoName = fullPathName[2]
  actionsString = fullPathName[3]
  workflowsString = fullPathName[4]
  workflowFileName = fullPathName[5]

  if (actionsString != "actions" || workflowsString != "workflows") {
    return "Wrong URL, can't generate badge"
  }

  return '[![Build Status](https://github.com/' + repoOwner + '/' + repoName + '/actions/workflows/' + workflowFileName + '/badge.svg)](https://github.com/' + repoOwner + '/' + repoName + '/actions)'
}


module.exports = getBadge
