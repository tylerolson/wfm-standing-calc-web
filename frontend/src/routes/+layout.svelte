<script lang="ts">
  import "./layout.css";
  import favicon from "$lib/assets/favicon.svg";

  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>

<nav class="flex h-16 items-center justify-center bg-[#1F2937] text-3xl text-white">
  <p>Warframe Platinum Efficiency Calculator</p>
</nav>

{@render children()}
