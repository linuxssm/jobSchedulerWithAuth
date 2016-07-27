package conf


var DB_AUTH string = "httpauth"
var DB_JOB string = "jobScheduler"
var DB_PASS string = "passwd"
var DB_USER string = "root"
var DB_HOST string = "localhost"
var DB_PORT string = "3306"

var DB_AUTH_OPEN_INFO string = DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_AUTH
var DB_JOB_OPEN_INFO string = DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_JOB