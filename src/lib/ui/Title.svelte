<script lang="ts">
  import { encrypter } from "../stores/encrypter";
  import { STATE } from "../constants";
  const { store } = encrypter;
</script>

<div class="container" class:active={$store.state !== STATE.INITIAL}>
  <h1>cryp.</h1>
  <h1 class="gradient">cryp.</h1>
  <small>Encrypt and decrypt any files</small>
</div>

<style>
  h1 {
    letter-spacing: 1rem;
    width: 100%;
    text-align: center;
    margin-right: -35px;
    transition: all 0.3s ease;
    z-index: 10;
    position: relative;
    font-weight: 900 !important;
  }
  h1.gradient {
    background: -webkit-linear-gradient(var(--purple), var(--info));
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
    z-index: 9;
    position: absolute;
    right: 0;
    left: 0;
    top: 0;
    animation: floating 20s ease-in-out infinite;
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

  .container {
    position: relative;
  }
</style>
