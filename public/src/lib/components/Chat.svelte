<script>
  import { onMount, beforeUpdate, afterUpdate } from "svelte"
  import usernameStore from "$lib/stores/username.js"

  import Message from "$lib/components/Message.svelte"
  import MessageForm from "$lib/components/MessageForm.svelte"

  // Storing the messages history from the server
  let messageHistory = []
  let socket

  // Storing a reference to the chat element for automatic scrolling when new messages are added
  let chatElement

  function sendMessage(value) {
    // Send message as plain text
    socket.send(value)
  }

  // Running websockets as client code
  onMount(() => {
    socket = new WebSocket(`ws://localhost:8080/server/socket?username=${$usernameStore}`)

    socket.addEventListener("message", message => {
      const data = JSON.parse(message.data)

      // Distinguish cases of different data being sent from the server
      // Once the connection is created, the server will reply with the message history (Array)
      if (data.length >= 0)
        messageHistory = data

      else
        messageHistory = [...messageHistory, data]
    })
  })

  let shouldScroll = false

  // Determine whether the chat should scroll at the bottom
  beforeUpdate(() => {
    shouldScroll = chatElement && (chatElement.scrollTop
      !== chatElement.scrollHeight)
  })

  // Scroll smoothly towards the bottom
  afterUpdate(() => {
    chatElement.scrollTo({
      top: chatElement.scrollHeight,
      left: 0,

      behavior: "smooth"
    })
  })
</script>

<div class="application">
  <div 
    class="messages"
    bind:this={chatElement}
  >
    {#each messageHistory as message}
      <Message {...message} />
    {/each}
  </div>

  <hr />
  <MessageForm onSubmit={sendMessage} />
</div>

<style>
  .application {
    max-width: var(--page-width);
    margin: 3em auto;

    padding: 2em;
  }

  .messages {
    margin-bottom: 3em;

    max-height: 60vh;
    padding: 0 2em;
    overflow-y: scroll;
  }

  @media (max-width: 500px) {
    .application {
      margin: 0;
      padding: 2em 1.5em;
    }

    .messages {
      padding: 0 1em;
    }
  }
</style>
