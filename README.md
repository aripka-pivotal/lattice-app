# Lattice-App - a simple Go webapp for playing with Lattice

Lattice-App is packaged as a docker image at cloudfoundry/lattice-app

To push to [Lattice](https://github.com/cloudfoundry-incubator/lattice) using [ltc](https://github.com/cloudfoundry-incubator/lattice/ltc):

```bash
ltc create lattice-app cloudfoundry/lattice-app
```

###Displayed Application Name
This application will include the lattice/cloud foundry app name in the logging messages and landing page.  This value is set by reading the _PROC_GUID_ environment variable on lattice and the _"application name"_ value from the _VCAP_APPLICATION_ environment variable on cloud foundry.  Either of these values may be overridden by supplying an environment variable named _APP_NAME_.  When running locally, or where these environment variables are not set, this value defaults to "lattice-app"  

###Environment Variables to Customize Appearance
This application supports customization of the color scheme by supplying one of the following environment variables:
* _HEX_COLOR_ - hexidecimal representation of the RGB color value.  If an invalid value(including white) is supplied the default color will be used. 
* _COLOR_NUM_ - integer value representing one of 20 preset color values (indexes outside this range will use the default colors).

### Endpoints

`/`: a simple landing page displaying the index and uptime  
`/env`: displays environment variables  
`/exit`: instructs Lattice to exit with status code 1  

### To rebuild the dockerimage:

```bash
./build.sh
```

Assumes you have the go toolchain (with the ability to cross-compile to different platforms) and docker installed and pointing at your docker daemon.
