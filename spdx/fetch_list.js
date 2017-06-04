const fs = require('fs');
const spdxLicenseList = require('spdx-license-list/full');

fs.writeFileSync("./json/licenses.json", JSON.stringify(spdxLicenseList));
console.log('put ./json/licenses.json')
