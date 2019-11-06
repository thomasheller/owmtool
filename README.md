# owmtool

## Build

- Install Go - try gimme: https://github.com/travis-ci/gimme
- Clone repo to `~/go/src/github.com/thomasheller/owmtool`
- `cd ~/go/src/github.com/thomasheller/owmtool && go build`

## Usage

```bash
$ cat positions.json
{
  "positions": [
    { "lat": 52.524702, "lon": 13.367882 },
    { "lat": 50.036789, "lon": 8.561119 },
    { "lat": 46.961532, "lon": 7.503579 }
  ]
}
$ OPENWEATHERMAP_API_KEY=... ./owmtool < positions.json
2019/11/06 22:25:25 http://api.openweathermap.org/data/2.5/weather?lat=52.524702&lon=13.367882&units=metric&lang=de&appid=...
2019/11/06 22:25:26 http://api.openweathermap.org/data/2.5/weather?lat=50.036789&lon=8.561119&units=metric&lang=de&appid=...
2019/11/06 22:25:27 http://api.openweathermap.org/data/2.5/weather?lat=46.961532&lon=7.503579&units=metric&lang=de&appid=...
{
    "results": [
        {
            "input": {
                "lat": 52.524702,
                "lon": 13.367882
            },
            "output": {
                "lat": 52.52,
                "lon": 13.37
            },
            "distance_meters": 542.1200806736914,
            "location_name": "Tiergarten"
        },
        {
            "input": {
                "lat": 50.036789,
                "lon": 8.561119
            },
            "output": {
                "lat": 50.04,
                "lon": 8.56
            },
            "distance_meters": 365.8812841677877,
            "location_name": "Frankfurt Main Flughafen"
        },
        {
            "input": {
                "lat": 46.961532,
                "lon": 7.503579
            },
            "output": {
                "lat": 46.96,
                "lon": 7.5
            },
            "distance_meters": 320.61242379105903,
            "location_name": "Ostermundigen"
        }
    ],
    "stats": {
        "avg_distance_meters": 409.5379295441794,
        "max_distance_meters": 542.1200806736914,
        "min_distance_meters": 320.61242379105903
    },
    "errors": []
}
```
