<script lang="ts">
  import Check from "./icons/check.svelte";
</script>

<div>
  <div class="icon vertical-center"><Check width={100} /></div>
  <h3>Success</h3>
</div>

<style>
  h3 {
    text-align: center;
  }
  .icon {
    padding-top: var(--small);
    width: 100%;
    color: var(--info);
  }
</style>
