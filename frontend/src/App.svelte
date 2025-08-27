<script lang="ts">
  import { onMount } from "svelte";
  import VendorTable from "./lib/VendorTable.svelte";
  import type { Vendor, VendorsRespone } from "./types.ts";

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

    const data: VendorsRespone = await response.json();

    vendors = data.vendors;
    updatedAt = new Date(data.updatedAt);
    updating = data.updating;
  });
</script>

<main>
  {#if vendors.length == 0}
    <div class="m-6">
      <p class="text-gray-100 text-3xl">{loadStatus}</p>
    </div>
  {:else}
    <div class="m-6">
      {#if updating}
        <p class="pb-3 text-gray-100">Updating...</p>
      {:else}
        <p class="text-gray-100 pb-3">Updated at: {updatedAt?.toLocaleString()}</p>
      {/if}

      <!-- svelte-ignore a11y_autofocus -->
      <input
        placeholder="Search"
        class="rounded-md shadow-xl w-full py-3 px-6 bg-gray-700 focus:outline-gray-400"
        autofocus
        bind:value={filterText}
      />
      <div class="flex flex-wrap justify-around content-evenly">
        {#each vendors as vendor: Vendor (vendor.name)}
          <VendorTable {vendor} {filterText}></VendorTable>
        {/each}
      </div>
    </div>
  {/if}
</main>
