sql.Open("driver", "username:password@tcp(localhost:3306)/dbname")

```go
    func main(){

        fmt.Println("hello sql")

    }
```

db pooling == connection management

name = root
pass = 1234
dbname = my_database

sqlcmd: {

login: "mysql -u&name -p&pass",
showdb: "show databases",
showtable: "show tables"
create: "create database &dbname",
delete: "drop database &dbname",
use: "use &dbname"

}

datatype : {

TINYINT: {

bytes: 1,
minsign: -128,
maxsign: 127,
minunsign: 0,
maxunsign: 255,

},

SMALLINT: {

bytes: 2,
minsign: -32768,
maxsign: 32767,
minunsign: 0,
maxunsign: 65535,

},

MEDIUMINT: {

bytes: 3,
minsign: -8388608,
maxsign: 8388608,
minunsign: 0,
maxunsign: 16777215,

},

INT: {

bytes: 4,
minsign: -2147483648,
maxsign: 2147483647,
minunsign: 0,
maxunsign: 4294967295,

},

BIGINT: {

bytes: 8,
minsign: -9223372036854775808,
maxsign: 9223372036854775807,
minunsign: 0,
maxunsign: 18446744073709551615,

},

}
