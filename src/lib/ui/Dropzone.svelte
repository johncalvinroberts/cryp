<script lang="ts">
  import FileDrop from "filedrop-svelte";
  import Plus from "../icons/plus.svelte";
  import { encrypter } from "../stores/encrypter";
  const { handleFiles } = encrypter;
  let isDraggingOver = false;
</script>

<FileDrop
  on:filedrop={(e) => handleFiles(e.detail.files)}
  on:dragover={() => (isDraggingOver = true)}
  on:dragleave={() => (isDraggingOver = false)}
>
  <div
    class="dropzone"
    class:is-dragging-over={isDraggingOver}
    on:mouseover={() => (isDraggingOver = true)}
    on:mouseleave={() => (isDraggingOver = false)}
    on:focus={() => (isDraggingOver = true)}
  >
    <Plus {isDraggingOver} />
  </div>
</FileDrop>
<small> Drop Files or Folders to Get Started </small>

<style>
  .dropzone {
    text-align: center;
    color: var(--dark);
    padding: 2rem 1rem;
    width: 100%;
    background-color: var(--yellow);
    margin-bottom: 5rem;
    transition: padding cubic-bezier(0.075, 0.82, 0.165, 1) 0.2s;
    cursor: pointer;
  }
  .dropzone:hover {
    padding: 3rem 1rem;
  }
  .is-dragging-over {
    padding: 3rem 1rem;
  }
  small {
    text-align: center;
    display: block;
  }
</style>
