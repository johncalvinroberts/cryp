<script lang="ts">
  import Check from "../icons/check.svelte";
  import { encrypter, reset } from "../stores/encrypter";
  import { getEncryptedFilename, download } from "../utils";

  const type = $encrypter.decryptedFiles ? "decrypted" : "encrypted";

  let isFailToDownload = false;

  const handleDownload = () =>
    type === "encrypted" ? downloadEncrypted() : downloadDecrypted();

  const downloadEncrypted = () => {
    isFailToDownload = false;
    try {
      const fileName = getEncryptedFilename($encrypter.filesToEncrypt);
      const file = new File([$encrypter.crypString], fileName);
      download(file, fileName);
    } catch (error) {
      isFailToDownload = true;
      console.error(error);
    }
  };

  const downloadDecrypted = () => {
    isFailToDownload = false;
    try {
      for (const file of $encrypter.decryptedFiles) {
        download(file, file.name);
      }
    } catch (error) {
      isFailToDownload = true;
      console.error(error);
    }
  };
</script>

<div>
  <h3 class="vertical-center">
    <span class="icon vertical-center">
      <Check width={40} />
    </span>
    Successfully {type}
  </h3>
  {#if isFailToDownload}
    <div class="error">Failed to Download File.</div>
  {/if}
  <div class="bottom-box">
    <button class="vertical-center success-button" on:click={handleDownload}>
      Download
    </button>
    <button on:click={reset}> Start Over </button>
  </div>
</div>

<style>
  .bottom-box {
    margin: 0 auto;
  }

  .bottom-box button {
    max-width: 150px;
    margin: 0 auto;
  }

  .success-button {
    color: var(--light);
    background: -webkit-linear-gradient(var(--purple), var(--info));
    border: solid 1px var(--light);
    animation: floating 20s ease-in-out infinite;
    border-radius: 2px;
    transition: all 0.3s ease;
  }
  .success-button:hover {
    border: solid 2px var(--light);
  }
  .icon {
    margin-right: 0.5rem;
  }

  .error {
    color: var(--error);
    text-align: center;
    padding-top: 1rem;
  }
</style>
