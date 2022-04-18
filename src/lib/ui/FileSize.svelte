<script lang="ts">
  import fileSize from "filesize";
  export let bytes: number | undefined = undefined;
  let className = "";
  export { className as class };
</script>

<span class={className}>
  {fileSize(bytes)}
</span>

<style>
  span {
    opacity: 0.5;
    font-size: 0.8rem;
    font-weight: normal;
    margin-left: 0.2rem;
  }
</style>
