<script lang="ts">
  import { onMount } from "svelte";
  import VendorTable from "./lib/VendorTable.svelte";
  import type { Vendor } from "./types.ts";

  let loadStatus = $state("Loading...");
  let vendors = $state<Vendor[]>([]);

  let filterText = $state("");

  onMount(async () => {
    const response = await fetch("/api/vendors");

    if (!response.ok) {
      loadStatus = `HTTP error ${response.status} (${response.statusText}) The backend server may be down, try again later.`;
      return;
    }

    const data: Vendor[] = await response.json();

    vendors = data;
  });
</script>

<main>
  {#if vendors.length == 0}
    <div>
      <p class="text-3xl">{loadStatus}</p>
    </div>
  {/if}
  <div class="m-6">
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
</main>
