# Social-Network
- [Audit Instructions](https://github.com/01-edu/public/tree/master/subjects/social-network/audit)

## Prequisites:
* Node Package Manager
* Golang version 1.22.5 or newer.
* Docker (Optional)

## Instructions
- Open the terminal and clone the repository using the following command:
```
git clone https://01.kood.tech/git/jvillem/social-network
```
- Then navigate to the root directory
```
cd social-network
```

## Running from source
Open up two terminal windows and navigate to the root directory of the repository
In one of the terminal windows navigate to the frontend directory and run the following command
```
npm install
```
After it finishes, type this command into the same terminal window

```
npm run dev
```

Expected output:
```
xaero@XAERO-X570-F:~/koodprojects/social-network/frontend$ npm run dev

> frontend@0.1.0 dev
> next dev

  ▲ Next.js 14.2.12
  - Local:        http://localhost:3000

 ✓ Starting...
 ✓ Ready in 1171ms
 ```

Now, leave that terminal window be and on your second terminal window navigate to the backend's cmd/api directory
```
cd backend/cmd/api
```
And then run the backend server
```
go run .
```
Expected output:
```
Foreign Key Status: 1
Starting Go server
```

And then feel free to navigate to http://localhost:3000 using a web browser of your choosing.

## Running via Docker
To build the docker container run the following command:
```docker compose build```
To run the container run:
```docker compose up``` 

## Troubleshooting
If you run the backend server and you encounter this error:
```
xaero@XAERO-X570-F:~/koodprojects/social-network/backend/cmd/api$ go run .
Failed to enable foreign keys: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
```
Within the same terminal install build-essential via this terminal command
```
sudo apt install build-essential
```

## Team
- [@jvillem](https://01.kood.tech/git/jvillem)
- [@mlutter](https://01.kood.tech/git/mlutter)
- [@mpuusaag](https://01.kood.tech/git/mpuusaag)
- [@skoppelm](https://01.kood.tech/git/skoppelm)
- [@ykaneko](https://01.kood.tech/git/ykaneko)