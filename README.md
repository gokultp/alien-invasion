# Alien Invasion

## Assumptions taken

1. Input should be of format
   `<City Name><Space><Direction>=<CityName><Space><Direction>=<CityName>...`
   - `Direction`: Should be one of   `west`, `east`, `south`, `north`

   - There should be no spaces in City names.

   Eg:
   ```
      Foo north=Bar west=Baz south=Qu-ux
      Bar south=Foo west=Bee

   ```
2. Every aliens will move to one of its neighbouring city in every iteration if there is one.


## How to build

### Prerequisits
   1. Golang 1.12.x
   2. Gnu Make (only for build purpose, can be build even without this)

### Building
   ```sh
      $ git clone <>
      $ cd alien-invasion
      $ make build
   ```
### Running Unit Tests

   ```
      $ make test
   ```
### Running the build

   ```
      # execute with default params
      $ cd build && ./alien-invasion

      # specifying number of aliens
      $ ./alien-invasion --aliens=10

      # get help on other available params
      $ ./alien-invasion --help

   ```

output of help will look like the following

```
Usage of ./build/alien-invantion:
  -aliens int
        --aliens=<value> to specify number of aliens (default 2)
  -input string
        use --input=<file path> to specify map input file (default "./input.txt")
  -output string
        use --output=<file path> to specify map output file (default "./out.txt")
```

here input and output files also can be set with params.


