<script lang="ts">
  import Title from "./lib/ui/Title.svelte";
  import Encrypter from "./lib/ui/Encrypter.svelte";
</script>

<main class="vertical-center">
  <div class="inner vertical-center">
    <Title />
    <Encrypter />
  </div>
</main>

<style>
  main {
    width: 100%;
    height: 100vh;
  }
  .inner {
    width: 100%;
  }
</style>
