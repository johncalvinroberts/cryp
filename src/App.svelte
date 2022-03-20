<script lang="ts">
  import Title from "./lib/Title.svelte";
  import Encrypter from "./lib/Encrypter.svelte";
</script>

<main>
  <div class="inner">
    <Title />
    <Encrypter />
  </div>
</main>

<style>
  main {
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
  }
  .inner {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    width: 100%;
  }
</style>
