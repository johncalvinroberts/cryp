<script lang="ts">
  import Check from "./icons/check.svelte";
  import streamSaver from "streamsaver";
  import { encrypter, reset } from "./stores/encrypter";

  let isFailToDownload = false;
  const handleDownload = () => {
    isFailToDownload = false;
    try {
      const fileStream = new streamSaver("encrypted.txt");
      new Response($encrypter.ciphertext).body.pipeTo(fileStream).then(
        () => null,
        () => (isFailToDownload = true)
      );
    } catch (error) {
      isFailToDownload = true;
      console.error(error);
    }
  };
</script>

<div>
  <div class="vertical-center button-wrapper">
    <button class="vertical-center" on:click={handleDownload}>
      <span class="icon vertical-center">
        <Check width={40} />
      </span>
      Success
      <small>Click to download</small>
    </button>
  </div>
  {#if isFailToDownload}
    <div class="error">Failed to Download File.</div>
  {/if}
  <div class="bottom-box">
    <button on:click={reset}> Start Over </button>
  </div>
</div>

<style>
  .button-wrapper {
    width: 100%;
    padding: 0 4rem;
  }

  .bottom-box {
    margin: 0 auto;
  }

  .bottom-box button {
    max-width: 150px;
    margin: 0 auto;
  }

  .button-wrapper button {
    color: var(--light);
    background: -webkit-linear-gradient(var(--purple), var(--info));
    /* background-color: var(--purple); */
    font-size: 2rem;
    border: solid 1px transparent;
    padding: 2rem;
    animation: floating 20s ease-in-out infinite;
    border-radius: 2px;
    transition: all 0.3s ease;
  }
  .button-wrapper button:hover {
    border: solid 2px var(--light);
  }
  .icon {
    margin-right: 0.5rem;
  }

  small {
    margin-top: 1rem;
    font-size: 1rem;
  }
  .error {
    color: var(--error);
    text-align: center;
    padding-top: 1rem;
  }
</style>
