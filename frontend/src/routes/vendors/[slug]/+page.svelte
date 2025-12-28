<script lang="ts">
  import VendorTable from "$lib/VendorTable.svelte";
  let { data } = $props();
  let filterText = $state("");
</script>

<svelte:head>
  {#if data.vendor === undefined}
    <title>WFM Calculator</title>
  {:else}
    <title>{data.vendor.name}</title>
  {/if}
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

    <div class="mx-auto grid w-full max-w-[1500px] auto-cols-fr grid-cols-1 gap-8 px-4 py-4">
      <!-- svelte-ignore a11y_autofocus -->
      <input
        placeholder="Search"
        class=" rounded-lg bg-gray-700 px-6 py-3 text-gray-100 shadow-xl focus:outline-gray-400"
        autofocus
        bind:value={filterText}
      />
      <VendorTable vendor={data.vendor} {filterText}></VendorTable>
    </div>
  </div>
{/if}
