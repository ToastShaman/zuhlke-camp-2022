install: node_modules

node_modules: package.json
	npm install -g aws-cdk && npm install
