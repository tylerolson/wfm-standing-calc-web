<script lang="ts">
  import { onMount } from "svelte";
  import VendorTable from "./lib/VendorTable.svelte";
  import type { Vendor } from "./types.ts";

  let vendors = $state<Vendor[]>([]);

  onMount(async () => {
    const response = await fetch("/api/vendors");
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data: Vendor[] = await response.json();

    vendors = data;
  });
</script>

<main>
  {#if vendors.length == 0}
    <div>
      <p>Loading</p>
    </div>
  {/if}
  <div>
    {#each vendors as vendor: Vendor (vendor.name)}
      <VendorTable {vendor}></VendorTable>
    {/each}
  </div>
</main>
