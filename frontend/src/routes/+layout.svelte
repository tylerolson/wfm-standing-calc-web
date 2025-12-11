<script lang="ts">
  import "./layout.css";
  import favicon from "$lib/assets/icons8-warframe.svg";

  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>

<nav
  class="fixed top-0 flex h-18 w-full items-center justify-between bg-[#1F2937] px-4 text-xl text-white shadow-lg"
>
  <p class=" text-2xl">Warframe Platinum Efficiency Calculator</p>
  <div>
    <a href="/">
      <button type="button" class="button">HOME</button>
    </a>
    <a href="/about">
      <button type="button" class="button">ABOUT</button>
    </a>
    <a href="/docs">
      <button type="button" class="button">API</button>
    </a>
  </div>
</nav>
<div class="pb-18"></div>

{@render children()}

<style lang="postcss">
  @reference "tailwindcss";
  .button {
    @apply ml-2 inline-flex cursor-pointer items-center justify-center rounded-lg px-4 py-1 outline-2 outline-gray-700 transition-all hover:bg-gray-700;
  }
</style>
