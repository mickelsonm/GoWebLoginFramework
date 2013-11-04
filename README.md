# CURT Admin

GoAdmin is the second iteration of the CURT eCommerce Management Portal. This version was built using the MySQL infrastructure that CURT is moving toward and uses a more scalable model.

## Setup Process

1. Install Capistrano

	```bash
	gem install capistrano
	```

2. Install JSON Gem && JSON-Pure Gem

	```bash
	gem install json
	gem install json_pure
	```

3. Install GruntJS

	```bash
	npm install -g grunt-cli
	```

## Build Process

Grunt

This task will run jshint against all javascript, run the compass compiler against all our compass styling, it will run uglify to minify the javascript and add a banner; while also watching for changes and re-running the validation on save.

```bash
grunt
```

## Deployment Process

Capistrano Deployment

Capistrano will pull from the Github repository using remote_cache, this makes things deploy a little quicker. It will run `go get` against all needed go dependencies. Then it will configure the database settings and email settings through either a defined deploy_settings.json file or using prompts.

After all configuration is finished, it will minify javascript and stylesheets using the yui-compressor jar file on the remote server.

Once minification is finished, it will compile an executable, kill the running process, and start a new process. Pretty cool, eh?

```bash
cap deploy
```

Contributors
-----------

**Alex Ninneman**

+ http://twitter.com/ninnemana
+ http://github.com/ninnemana