<script>
  import usernameStore from "$lib/stores/username.js"

  // Storing the input username
  let username = ""

  let playErrorAnimation = false

  function handleSubmit() {
    // Check if the username is valud
    if (username.length < 5 || username === "Server") {
      playErrorAnimation = true
      
      return
    }

    // Upgate username store, will automatically redirect to chat
    usernameStore.set(username)
  }
</script>

<div class="form__container">
  <form on:submit|preventDefault={handleSubmit}>
    <h1>Welcome to Pioyi Messenger</h1>

    <p
      class:error={playErrorAnimation}
      on:animationend={() => { playErrorAnimation = false }}
    >
      Please enter a username to access the global chat server.
      You must make use of at least 5 characters!
    </p>

    <div class="form__group">
      <input
        placeholder="Please enter a username"
        bind:value={username}
      />

      <button type="submit">Enter Chat</button>
    </div>
  </form>
</div>

<style>
  .form__container {
    margin: 2em;
  }

  form {
    display: block;
    background: var(--bg-light);
    border-radius: 5px;

    padding: 1.5em;
    
    max-width: 400px;
    margin: 2em auto;
  }

  h1 {
    font-size: 1.3rem;
    margin-bottom: 0.7em;
  }

  p {
    margin-bottom: 2em;

    font-size: 0.8rem;
  }

  .form__group {
    width: 100%;
    display: flex;

    align-items: center;
    justify-content: space-between;

    flex-wrap: wrap;
    gap: 0.5em;
  }

  input {
    flex: 1;
    border: none;

    padding: 0.1em 0.5em;
    outline: none;
  }

  button {
    border: none;
    padding: 0.3em 1em;

    cursor: pointer;
  }

  @keyframes errorAnimation {
    from {
      color: red;
      scale: 90%;
    }

    to {
      color: black;
      scale: 100%;
    }
  }

  p.error {
    animation-name: errorAnimation;
    animation-duration: 500ms;
  }
</style>
