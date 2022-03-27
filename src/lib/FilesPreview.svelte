<script lang="ts">
  import Empty from "./Empty.svelte";
  import { encrypter, reset, handleEncrypt } from "./stores/encrypter";
  import FileSize from "./FileSize.svelte";
  import Eye from "./icons/eye.svelte";
  const files = $encrypter.files;
  const totalFileBytes = files.accepted.reduce((memo, current) => {
    return memo + current.size;
  }, 0);
  let showPassword = false;
  let password = "";
  let hint = "";
  const handleSubmit = async (e: SubmitEvent) => {
    e.preventDefault();
    await handleEncrypt({ password, hint });
  };
</script>

<div class="wrapper">
  {#if files.accepted.length > 1}
    <div class="title">
      <h6>
        {files.accepted.length} Files -
      </h6>
      <FileSize bytes={totalFileBytes} />
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
  <form on:submit={handleSubmit} on:reset={reset}>
    <div class="input-box">
      <input
        name="password"
        placeholder="Password"
        type={showPassword ? "text" : "password"}
        value={password}
        autocomplete="off"
        spellcheck="false"
        on:input={(e) => (password = e.currentTarget.value)}
      />
      <button
        type="button"
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
  }
</style>
