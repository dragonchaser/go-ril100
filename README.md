# go-ril100

A quick and dirty hacked tool to query RIL100/DS100 abbreviations from the DB
(Deutsche Bahn) in Germany.

## usage

### text output

```
$> go-ril100 DFL
```

#### result

```
Name     : Flöha
Type     : station
Address  :
	Street : Bahnhofstr. 2a
	City   : 09557 Flöha
Location :
	Latitude  :50.854355
	Longitude :13.075152
```

### json output

```
$> go-ril100 DFL --json |jq
```


**Note:** jq is optional, just beautifies the output in our case.

#### result

```
{
  "type": "station",
  "name": "Flöha",
  "address": {
    "city": "Flöha",
    "zipcode": "09557",
    "street": "Bahnhofstr. 2a"
  },
  "location": {
    "latitude": 50.854355,
    "longitude": 13.075152
  }
}
```

## building

just run `$> make`
