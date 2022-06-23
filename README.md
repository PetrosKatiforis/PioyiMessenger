## Pioyi Messenger

Pioyi Messenger is a tiny real-time chat application. The server is build in Go using the gorilla/websocket package. The client is written in Svelte and is served statically.

The communication is made possible by websockets.

<br />

<p align="center">
    <img 
        width="400px"
        styles="display: block; margin: 0 auto;"
        src="https://i.imgur.com/Xuz4dWK.png" 
        alt="Pioyi Messenger Screenshot"
    />
</p>

<br />

### Building the project

The complilation of the project is done in two different stages

   - First, after having installed all the node dependencies `` cd public `` and build the client using `` npm run build ``. The frontend will now be localted inside the public/build folder.

   - Navigate to the project root and run `` go run . ``. The golang server will be started and the application will be served by default at localhost:8080
