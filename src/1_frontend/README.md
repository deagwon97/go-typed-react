# Yarn, Typesctipt, React

### Dockerfile
```docker
RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    curl -sL https://deb.nodesource.com/setup_16.x | bash - &&\
    apt-get install nodejs -y &&\
    npm install -g yarn
```

### Create Project
```bash
yarn create react-app app_name --template typescript

yarn add scss

yarn add @types/scss
```

### package.json
```json
{
	...
	"homepage": "/", // 추가
	...
	"scripts": {
		...
    "build": "GENERATE_SOURCEMAP=false react-scripts build",// 변경
  },
}
```

### tsconfig.json

```json
{
	...
  "compilerOptions": {
    ...
    "baseUrl": "./src", // 추가
  },
  ...
}
```