var https = require("https");
var fs = require('fs');

const n = process.argv[2];

const token = fs.readFileSync('session.env', 'utf8');



https.get({
    host: 'adventofcode.com',
    port: 443,
    path: `/2022/day/${n}/input`,
    method: 'GET',
    headers: {
      cookie:token
    }
  }, (res) => {
    var str = '';

    //another chunk of data has been received, so append it to `str`
    res.on('data', function (chunk) {
      str += chunk;
    });
  
    //the whole response has been received, so we just print it out here
    res.on('end', function () {
      console.log(str);
    });
  });