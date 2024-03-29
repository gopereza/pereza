# Pereza [![Build Status](https://travis-ci.org/gopereza/pereza.svg?branch=master)](https://travis-ci.org/gopereza/pereza)[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/gopereza/pereza/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/gopereza/pereza/?branch=master)
experimental json marshaler

### JSON
Compare between **encoding/json** & **github.com/mailru/easyjson** & **github.com/gopereza/pereza**

##### Very simple structures
```golang
struct {
	State bool `json:"state"`
}
```
```text
BenchmarkBoolStateEncodingJSON   	 5000000	       306 ns/op	     177 B/op	       2 allocs/op
BenchmarkBoolStateEasyJSON       	20000000	       110 ns/op	     128 B/op	       1 allocs/op
BenchmarkBoolStatePerezaJSON     	200000000	         6.39 ns/op	       0 B/op	       0 allocs/op
```

```golang
struct {
	State int `json:"state"`
}
```
```text
BenchmarkIntStateEncodingJSON    	 5000000	       331 ns/op	     184 B/op	       2 allocs/op
BenchmarkIntStateEasyJSON        	10000000	       129 ns/op	     128 B/op	       1 allocs/op
BenchmarkIntStatePerezaJSON      	30000000	        58.1 ns/op	      32 B/op	       1 allocs/op
```

```golang
struct {
	State string `json:"state"`
}
```
```text
BenchmarkStringStateEncodingJSON   	 2000000	       761 ns/op	     480 B/op	       3 allocs/op
BenchmarkStringStateEasyJSON       	 2000000	       719 ns/op	     720 B/op	       4 allocs/op
BenchmarkStringStatePerezaJSON     	20000000	        69.6 ns/op	     144 B/op	       1 allocs/op
```

##### Bool structure with 16 properties
```golang
struct {
	A bool `json:"a"`
	B bool `json:"b"`
	C bool `json:"c"`
	D bool `json:"d"`
	E bool `json:"e"`
	F bool `json:"f"`
	G bool `json:"g"`
	H bool `json:"h"`
	I bool `json:"i"`
	J bool `json:"j"`
	K bool `json:"k"`
	L bool `json:"l"`
	M bool `json:"m"`
	N bool `json:"n"`
	O bool `json:"o"`
	P bool `json:"p"`
}
```
```text
BenchmarkHexaBoolStateEncodingJSON     	 1000000	      1549 ns/op	     624 B/op	       4 allocs/op
BenchmarkHexaBoolStateEasyJSON         	 2000000	       691 ns/op	     752 B/op	       4 allocs/op
BenchmarkHexaBoolStatePerezaJSON     	20000000	        85.3 ns/op	     176 B/op	       1 allocs/op
```

```text
BenchmarkAlphabetBoolStateEncodingJSON   	 1000000	      2076 ns/op	     640 B/op	       4 allocs/op
BenchmarkAlphabetBoolStateEasyJSON       	 2000000	       912 ns/op	     864 B/op	       4 allocs/op
BenchmarkAlphabetBoolStatePerezaJSON     	10000000	       127 ns/op	     288 B/op	       1 allocs/op
```

##### Multi type one level structure
```golang
struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	CreatedAt uint32 `json:"created_at"`
	UpdatedAt uint32 `json:"updated_at"`
	Enabled   bool   `json:"enabled"`
}
```

```text
BenchmarkShortUserEncodingJSON   	 1000000	      1383 ns/op	     688 B/op	       4 allocs/op
BenchmarkShortUserEasyJSON       	 2000000	       687 ns/op	     720 B/op	       4 allocs/op
BenchmarkShortUserPerezaJSON     	10000000	       158 ns/op	     144 B/op	       1 allocs/op
```

### BSON
Compare between **go.mongodb.org/mongo-driver/bson** & **github.com/gopereza/pereza**

##### Empty structure
```golang
struct {}
```

```text
BenchmarkEmptyStateMongoMarshalBSON    	 3000000	       481 ns/op	     288 B/op	       2 allocs/op
BenchmarkEmptyStatePerezaMarshalBSON   	2000000000	         0.30 ns/op	       0 B/op	       0 allocs/op
```

##### Very simple structures
```golang
struct {
	State bool `bson:"state"`
}
```

```text
BenchmarkBoolStateMongoMarshalBSON    	 3000000	       586 ns/op	     289 B/op	       3 allocs/op
BenchmarkBoolStatePerezaMarshalBSON   	200000000	         6.46 ns/op	       0 B/op	       0 allocs/op
```

##### Bool structure with 26 properties
```golang
struct {
	A bool `bson:"a"`
	B bool `bson:"b"`
	C bool `bson:"c"`
	D bool `bson:"d"`
	E bool `bson:"e"`
	F bool `bson:"f"`
	G bool `bson:"g"`
	H bool `bson:"h"`
	I bool `bson:"i"`
	J bool `bson:"j"`
	K bool `bson:"k"`
	L bool `bson:"l"`
	M bool `bson:"m"`
	N bool `bson:"n"`
	O bool `bson:"o"`
	P bool `bson:"p"`
	Q bool `bson:"q"`
	R bool `bson:"r"`
	S bool `bson:"s"`
	T bool `bson:"t"`
	U bool `bson:"u"`
	V bool `bson:"v"`
	W bool `bson:"w"`
	X bool `bson:"x"`
	Y bool `bson:"y"`
	Z bool `bson:"z"`
}
```

```text
BenchmarkAlphabetMongoMarshalBSON    	 1000000	      2127 ns/op	     320 B/op	       3 allocs/op
BenchmarkAlphabetPerezaMarshalBSON   	20000000	        74.9 ns/op	     112 B/op	       1 allocs/op
```