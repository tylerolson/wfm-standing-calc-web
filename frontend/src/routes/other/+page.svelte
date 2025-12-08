<script lang="ts">
  import { onMount } from "svelte";
  import VendorTable from "$lib/VendorTable.svelte";
  import type { Vendor, VendorsResponse } from "$lib/types.ts";

  let loadStatus = $state("Loading...");
  let vendors = $state<Vendor[]>([]);
  let updatedAt = $state<Date>();
  let updating = $state(false);

  let filterText = $state("");

  onMount(async () => {
    const response = await fetch("/api/vendors");

    if (!response.ok) {
      loadStatus = `HTTP error ${response.status} (${response.statusText}) The backend server may be down, try again later.`;
      return;
    }

    const data: VendorsResponse = await response.json();

    vendors = data.vendors;
    updatedAt = new Date(data.updatedAt);
    updating = data.updating;
  });
</script>

{#if vendors.length === 0}
  <div class="m-6">
    <p class="text-3xl text-gray-100">{loadStatus}</p>
  </div>
{:else}
  <div class="m-6">
    {#if updating}
      <p class="pb-3 text-gray-100">Updating...</p>
    {:else}
      <p class="pb-3 text-gray-100">Updated at: {updatedAt?.toLocaleString()}</p>
    {/if}

    <!-- svelte-ignore a11y_autofocus -->
    <input
      placeholder="Search"
      class="w-full rounded-md bg-gray-700 px-6 py-3 shadow-xl focus:outline-gray-400"
      autofocus
      bind:value={filterText}
    />
    <div class="flex flex-wrap content-evenly justify-around">
      {#each vendors as vendor: Vendor (vendor.name)}
        <VendorTable {vendor} {filterText}></VendorTable>
      {/each}
    </div>
  </div>
{/if}
