<script lang="ts">
  import { onMount } from "svelte";
  import VendorTable from "./lib/VendorTable.svelte";
  import type { Vendor } from "./types.ts";

  let loadStatus = $state("Loading...");

  let vendors = $state<Vendor[]>([]);

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
  <div class="flex flex-wrap justify-evenly content-evenly">
    {#each vendors as vendor: Vendor (vendor.name)}
      <VendorTable {vendor}></VendorTable>
    {/each}
  </div>
</main>
