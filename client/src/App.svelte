<script lang="ts">
  import { onMount } from "svelte";
  import { get, getJSON, put, putJSON } from "./http";

  const url = `${window.location.protocol}//${window.location.hostname}:3000`;
  // const url = `http://localhost:3000`;

  let messages: string[] = [];

  function sendMessage(e: any): void {
    putJSON(url + "/mes", '"' + e.target.value + '"', (res) => {
      console.log(res);
    });
  }

  function getAllMessages(): void {
    getJSON(url + "/mes", (res) => {
      messages = JSON.parse(res);
    });
  }

  onMount(async () => {
    getAllMessages();
  });

  const interval = setInterval(() => {
    getAllMessages();
  }, 2000);
</script>

<div class="messages">
  {#each messages as m}
    <span>{m}</span>
  {/each}
</div>
<input type="text" on:change={(e) => sendMessage(e)} />

<style>
  .messages {
    display: flex;
    flex-direction: column;
  }
</style>
