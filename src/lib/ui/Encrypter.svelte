<script>
  import Guu from "guu";
  import Dropzone from "./Dropzone.svelte";
  import EncryptedPreview from "./EncryptedPreview.svelte";
  import { encrypter } from "../stores/encrypter";
  import FilesPreview from "./FilesPreview.svelte";
  import Done from "./Done.svelte";
  import Failure from "./Failure.svelte";
  import Processing from "./Processing.svelte";
  import { STATE } from "../constants";
  const log = new Guu("Encrypter.svelte", "goldenrod");
  const { store } = encrypter;
  const elements = {
    [STATE.INITIAL]: Dropzone,
    [STATE.SHOULD_DECRYPT]: EncryptedPreview,
    [STATE.SHOULD_ENCRYPT]: FilesPreview,
    [STATE.PROCESSING]: Processing,
    [STATE.FAILURE]: Failure,
    [STATE.DONE]: Done,
  };
  $: log.info($store.state);
</script>

<div class="wrapper vertical-center">
  <div class="inner">
    <svelte:component this={elements[$store.state]} />
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
