<script lang="ts">
  import BasicVendorPanel from "$lib/BasicVendorPanel.svelte";
  let { data } = $props();
</script>

<svelte:head>
  <title>WFM Calculator</title>
</svelte:head>

{#if data.error !== null}
  <div class="m-6">
    <p class="text-3xl text-gray-400">{data.error}</p>
  </div>
{:else}
  <div class="mt-1">
    <div class="flex justify-center">
      {#if data.updating}
        <p class="pb-3 text-gray-400">Updating...</p>
      {:else}
        <p class="pb-3 text-gray-400">Updated at: {data.updatedAt?.toLocaleString()}</p>
      {/if}
    </div>

    <div
      class="mx-auto grid w-full max-w-[1500px] auto-cols-fr grid-cols-1 gap-8 p-4 sm:grid-cols-2 xl:grid-cols-3"
    >
      {#each data.vendors as vendor (vendor.slug)}
        <BasicVendorPanel {vendor}></BasicVendorPanel>
      {/each}
    </div>
  </div>
{/if}
