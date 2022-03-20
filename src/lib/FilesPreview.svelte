<script lang="ts">
  import Empty from "./Empty.svelte";
  import { encrypter, reset } from "./stores/encrypter";
  import FileSize from "./FileSize.svelte";
  import Eye from "./icons/eye.svelte";
  const files = $encrypter.files;
  const totalFileBytes = files.accepted.reduce((memo, current) => {
    return memo + current.size;
  }, 0);
  let showPassword = false;
  let password = "";
  let hint = "";
  const handleSubmit = (e: SubmitEvent) => {
    e.preventDefault();
    console.log({ password, hint });
  };
  console.log({ password, hint });
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
  .input-box {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--yellow);
    padding: 0.3rem;
    position: relative;
  }
  .input-box button {
    background-color: transparent;
    border: none;
    opacity: 0.5;
    transition: opacity 0.5 ease;
    position: absolute;
    right: 0;
  }
  .input-box button:hover {
    opacity: 1;
  }
  input {
    background-color: transparent;
    border: none;
    width: 100%;
    font-size: 0.9rem;
    text-align: center;
  }
  input:focus {
    outline: none;
  }
  .input-box:not(:last-child) {
    border-bottom: solid 1px var(--dark);
  }
  .input-box,
  .bottom-box {
    max-width: 300px;
  }
  form {
    flex: 0 0 300px;
  }

  .bottom-box {
    display: flex;
    width: 100%;
    justify-content: space-between;
    padding: 2rem 0;
    gap: 1rem;
  }

  .bottom-box button {
    background-color: var(--yellow);
    color: var(--dark);
    border-radius: 2px;
    border: none;
    padding: 0.2rem 1.5rem;
    flex-grow: 1;
    font-size: 0.9rem;
  }
  .bottom-box button:disabled {
    opacity: 0.5;
    background-color: var(--light);
    color: var(--dark);
  }
</style>
