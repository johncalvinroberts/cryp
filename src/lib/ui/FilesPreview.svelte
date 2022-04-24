<script lang="ts">
  import Empty from "./Empty.svelte";
  import { encrypter } from "../stores/encrypter";
  import FileSize from "./FileSize.svelte";
  import Eye from "../icons/eye.svelte";

  const { store, reset, handleEncrypt } = encrypter;
  const files = $store.filesToEncrypt;
  const totalFileBytes = files.accepted.reduce((memo, current) => {
    return memo + current.size;
  }, 0);
  let showPassword = false;
  let password = "";
  let hint = "";

  const handleSubmit = (e: SubmitEvent) => {
    e.preventDefault();
    handleEncrypt(password, hint);
  };
</script>

<div class="wrapper">
  {#if files.accepted.length > 1}
    <div class="title">
      <h6>
        {files.accepted.length} Files -
      </h6>
      <span class="vertical-center">
        <FileSize bytes={totalFileBytes} />
      </span>
    </div>
  {/if}
  <ul>
    {#each files.accepted as file}
      <li>
        <span class="file-name truncate">{file.name}</span>
        -
        <FileSize bytes={file.size} class="file-size" />
      </li>
    {/each}
  </ul>
  {#if !files.accepted?.length}
    <Empty />
  {/if}

  {#if files.rejected?.length}
    <h6>Rejected Files</h6>
    <ul>
      {#each files.rejected as rejected}
        <li>{rejected.file.name} - {rejected.error.message}</li>
      {/each}
    </ul>
  {/if}
  <form on:submit={handleSubmit} on:reset={reset} autocomplete="off">
    <div class="input-box">
      <input
        name="secret"
        placeholder="Password"
        type={showPassword ? "text" : "password"}
        value={password}
        autocomplete="off"
        spellcheck="false"
        on:input={(e) => (password = e.currentTarget.value)}
      />
      <button
        type="button"
        class="vertical-center"
        on:click={() => (showPassword = !showPassword)}
        title={showPassword ? "Hide" : "Show"}
      >
        <Eye strikethrough={!showPassword} />
      </button>
    </div>
    <div class="input-box">
      <input
        name="hint"
        placeholder="Hint (optional)"
        bind:value={hint}
        autocomplete="off"
        spellcheck="false"
      />
    </div>
    <div class="bottom-box">
      <button type="reset"> Cancel </button>
      <button type="submit" disabled={!password}> Encrypt </button>
    </div>
  </form>
</div>

<style>
  .wrapper {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
  }
  h6 {
    text-align: center;
    margin: 0;
  }
  li {
    display: flex;
    width: 100%;
    justify-content: flex-start;
    align-items: center;
  }
  .file-name {
    display: block;
    max-width: calc(70%);
  }
  .title {
    display: flex;
    width: 100%;
    justify-content: center;
    margin-top: 3rem;
  }
</style>
