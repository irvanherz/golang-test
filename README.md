# README

## Instalation
Setup this program using below simple steps:
1. Clone this repo to your server.<br/>
`git clone https://github.com/irvanherz/golang-test`<br/>
then `cd` to the cloned folder<br/>
`cd golang-test`
2. Build docker container.<br/>
`docker build . -t go-dock`
3. Run container.<br/>
`docker run -p 5000:5000 go-dock`

## APIs
1. Get list of palindromes between to numbers (`from` and `to`).<br/>
Sample usage:<br/>
`www.example.com/go/soal0?from=1&to=100`<br/>
Output:
```JSON
{
  "success":true,
  "data":{
    "from":1,
    "to":100,
    "result":[
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      11,
      22,
      33,
      44,
      55,
      66,
      77,
      88,
      99
    ],
    "count":18
  }
}
```
2. Sort books to the correct order.<br/>
Sample usage<br/>
`www.example.com/go/soal1?input=3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14`<br/>
Output:
```JSON
{
  "success":true,
  "data":{
    "raw_input":"3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14",
    "input":[
      "3A13",
      "5X19",
      "9Y20",
      "2C18",
      "1N20",
      "3N20",
      "1M21",
      "1F14",
      "9A21",
      "3N21",
      "0E13",
      "5G14",
      "8A23",
      "9E22",
      "3N14"
    ],
    "raw_result":"0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3N14 3A13",
    "result":[
      "0E13",
      "9E22",
      "9A21",
      "9Y20",
      "8A23",
      "1M21",
      "1N20",
      "1F14",
      "2C18",
      "5X19",
      "5G14",
      "3N21",
      "3N20",
      "3N14",
      "3A13"
    ],
    "count":15
  }
}
```
3. Find missing numbers inside number sequences.<br/>
Sample usage:<br/>
`www.example.com/go/soal2?input=123457123459123460123462`<br/>
Output:
```JSON
{
  "success":true,
  "data":{
    "input":"123457123459123460123462",
    "result":[
      123458,
      123461
    ],
    "count":2
  }
}
```

