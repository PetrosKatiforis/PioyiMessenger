<script>
  import IoIosSend from "svelte-icons/io/IoIosSend.svelte"
  
  export let onSubmit

  // Stores input state
  let message = ""

  function submit() {
    // If the message is empty, don't bother sending it
    if (!message) return

    onSubmit(message)  

    // Clear the input value
    message = ""
  }
</script>

<form on:submit|preventDefault={submit}>
  <input
    placeholder="Broadcast a message to all users"
    
    bind:value={message}
  />

  <button type="submit">
    <IoIosSend />
  </button>
</form>

<style>
  form {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  input {
    flex: 1;
    border: none;
    padding: 0.5em;

    background: none;
    outline: none;
  }

  button {
    border: none;
    background: var(--bg-light);

    border-radius: 50%;
    padding: 0.5em;

    display: flex;
    align-items: center;
    justify-content: center;

    cursor: pointer;
  }

  button > :global(*) {
    width: 1.5rem;
  }
</style>
