<script>
  import { scale } from "svelte/transition"
  import usernameStore from "$lib/stores/username.js"

  export let author
  export let content

  // Check if the current user is the owner of the message
  let isOwner = author == $usernameStore

  // Useful for styling the server message differently
  let isServer = author == "Server"
</script>

<div 
  class="message__container"
  class:active={isOwner}
  class:server={isServer}
>
  <!-- Get the first character of the client's username for the icon -->
  <div class="message__author">
    { author.slice(0, 3) }
  </div>

  {#key author + content}
    <div 
      class="message__bubble"
      transition:scale
     >
       <p>{ content }</p>
    </div>
  {/key}
</div>

<style>
  .message__container {
    display: flex;
    align-items: flex-start;

    gap: 1em;
    margin-bottom: 1em;

    font-size: 0.85rem;
  }

  .message__container.active {
    flex-direction: row-reverse;
  }

  .message__container.server .message__author {
    display: none;
  }

  .message__container.server {
    justify-content: center;
    text-align: center;

    margin-bottom: 0em;
  }

  .message__container.server .message__bubble {
    background: none;
    opacity: 50%;

    margin: 0 auto;
  }

  :global(.message__container:not(.server) + .message__container.server),
  :global(.message__container.server + .message__container:not(.server)) {
    margin-top: 3em;
  }

  .message__bubble {
    background: var(--bg-light);
    max-width: var(--size-bubble);
    padding: 0.75em;

    border-radius: 5px;
    font-size: 0.8rem;
  }

  .message__container.active .message__bubble {
    background: var(--bg-primary);
    color: white;
  }

  .message__author {
    min-width: 35px;
    min-height: 35px;

    max-width: 35px;
    max-height: 35px;

    border-radius: 50%;
    background: var(--bg-light);

    display: flex;
    align-items: center;
    justify-content: center;

    font-size: 0.8rem;
    word-break: keep-all;

    overflow: hidden;
  }
</style>
