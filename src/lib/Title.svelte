<script lang="ts">
  import { encrypter } from "./stores/encrypter";
</script>

<div class:active={$encrypter.state !== "INITIAL"}>
  <h1>cryp.</h1>
  <small>Encrypt The Files You Send</small>
</div>

<style>
  h1 {
    letter-spacing: 1rem;
    width: 100%;
    text-align: center;
    margin-right: -35px;
    transition: all 0.3s ease;
  }
  small {
    text-align: center;
    width: 100%;
    display: block;
    max-height: 100px;
  }
  div {
    margin-bottom: 5rem;
    transition: all 0.3s ease;
  }
  .active h1 {
    font-size: 3rem;
    margin-bottom: 0;
  }
  .active {
    margin-bottom: 1rem;
  }

  .active small {
    max-height: 0;
    overflow: hidden;
  }
</style>
