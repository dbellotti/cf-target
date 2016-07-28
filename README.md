# cf-target
This Go program will print out the currently targeted Cloudfoundry api, org and space.

The program was created with the intention of printing the Cloudfoundry target so that you may easily add
it to your terminal prompt.

## Usage
```
$ cf-target -h
Usage of cf-target:
  -api
      display api target
  -org
    	display org
  -space
    	display space
```

## Installation
```
$ go get github.com/dbellotti/cf-target
```

### Thanks
If you like this, it was __heavily__ influenced by Kevin Kelani's awesome [bosh-target](https://github.com/kkallday/bosh-target)

### Screenshot
![cf-target in action](https://raw.githubusercontent.com/dbellotti/cf-target/master/screenshot.png)
