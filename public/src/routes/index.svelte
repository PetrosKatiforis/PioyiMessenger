<script>
  import usernameStore from "$lib/stores/username.js"
  
  import Chat from "$lib/components/Chat.svelte"
  import LoginPage from "$lib/components/LoginPage.svelte"
</script>

<svelte:head>
  <title>Pioyi Messenger Online</title>  
</svelte:head>

{#if $usernameStore}
  <Chat />

{:else}
  <LoginPage />
{/if}
