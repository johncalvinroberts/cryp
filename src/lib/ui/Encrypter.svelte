<script>
  import Dropzone from "./Dropzone.svelte";
  import EncryptedPreview from "./EncryptedPreview.svelte";
  import { encrypter, State } from "../stores/encrypter";
  import FilesPreview from "./FilesPreview.svelte";
  import Done from "./Done.svelte";
  import Failure from "./Failure.svelte";
  import Processing from "./Processing.svelte";

  const elements = {
    [State.INITIAL]: Dropzone,
    [State.SHOULD_DECRYPT]: EncryptedPreview,
    [State.SHOULD_ENCRYPT]: FilesPreview,
    [State.PROCESSING]: Processing,
    [State.FAILURE]: Failure,
    [State.DONE]: Done,
  };
  $: console.log($encrypter.state);
</script>

<div class="wrapper vertical-center">
  <div class="inner">
    <svelte:component this={elements[$encrypter.state]} />
  </div>
</div>

<style>
  .wrapper {
    width: 100%;
  }
  .inner {
    width: 100%;
    max-width: 400px;
  }
</style>
