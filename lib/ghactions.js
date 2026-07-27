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
function getBadge(input_url, reqOptions) {
  const defaultOptions = {};
  defaultOptions.schema = 'https';
  defaultOptions.title = 'Build Status';
  defaultOptions.hostname = 'github.com';
  console.log('reqoptions:', reqOptions)
  let options = { ...defaultOptions, ...reqOptions }

  let myURL = null
  try {
    myURL = new URL(input_url);
  } catch (error) {
    return "Wrong URL, can't generate badge"
  }
  const fullPathName = myURL.pathname.split('/')
  const repoOwner = fullPathName[1]
  const repoName = fullPathName[2]
  const actionsString = fullPathName[3]
  const workflowsString = fullPathName[4]
  const workflowFileName = fullPathName[5]

  if (actionsString != "actions" || workflowsString != "workflows") {
    return "Wrong URL, can't generate badge"
  }
  let branchOpt = ''
  console.log('options:', options)
  if (options.branch && options.branch != "default") {
    branchOpt = '?branch=' + options.branch
  }
  return '[![' + options.title + '](' + options.schema + '://' + options.hostname + '/' + repoOwner + '/' + repoName + '/actions/workflows/' + workflowFileName + '/badge.svg' + branchOpt + ')](' + options.schema + '://' + options.hostname + '/' + repoOwner + '/' + repoName + '/actions)'
}


module.exports = getBadge
